package dfdb

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"unicode"

	"github.com/hkionline/dftui/dflib/dfm"
)

// UUID v4 pattern for file identification
var uuidV4Pattern = regexp.MustCompile(`_([0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12})\.json$`)

// Character name validation pattern: only alphanumeric and spaces allowed
var validNamePattern = regexp.MustCompile(`^[a-zA-Z0-9 ]+$`)

// ErrCharacterNotFound is returned when a character cannot be found
var ErrCharacterNotFound = errors.New("character not found")

// ErrInvalidCharacterName is returned when a character name contains invalid characters
var ErrInvalidCharacterName = errors.New("character name contains invalid characters: only alphanumeric characters and spaces are allowed")

// FsProvider implements the Provider interface using filesystem JSON storage.
type FsProvider struct {
	mu    sync.RWMutex
	cache map[string]dfm.Character // map of cached characters by ID
	files map[string]string        // map of filenames by character ID
	dir   string                   // directory where character files are stored
}

// NewFsProvider creates a new filesystem-based provider.
// If the directory does not exist, it will be created.
func NewFsProvider(dir string) (*FsProvider, error) {
	// Create directory if it doesn't exist
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create directory %s: %w", dir, err)
	}

	// Load existing characters into cache
	cache, files, err := loadCache(dir)
	if err != nil {
		return nil, fmt.Errorf("failed to load cache: %w", err)
	}

	return &FsProvider{
		cache: cache,
		files: files,
		dir:   dir,
	}, nil
}

// Create stores a new character.
func (f *FsProvider) Create(character dfm.Character) error {
	if err := validateCharacterName(character.Name); err != nil {
		return err
	}

	f.mu.Lock()
	defer f.mu.Unlock()

	// Generate filename
	filename := generateFilename(character.Name, character.ID)

	// Write to file
	if err := saveCharacter(character, filepath.Join(f.dir, filename)); err != nil {
		return err
	}

	// Update cache
	f.cache[character.ID] = character
	f.files[character.ID] = filename

	return nil
}

// Read retrieves a character by ID from the cache.
func (f *FsProvider) Read(characterID string) (dfm.Character, error) {
	f.mu.RLock()
	defer f.mu.RUnlock()

	if character, ok := f.cache[characterID]; ok {
		return character, nil
	}
	return dfm.Character{}, ErrCharacterNotFound
}

// Update modifies an existing character.
// If the character name has changed, the file will be renamed.
func (f *FsProvider) Update(character dfm.Character) error {
	if err := validateCharacterName(character.Name); err != nil {
		return err
	}

	f.mu.Lock()
	defer f.mu.Unlock()

	// Check if character exists
	oldFilename, ok := f.files[character.ID]
	if !ok {
		return ErrCharacterNotFound
	}

	// Generate new filename based on current name
	newFilename := generateFilename(character.Name, character.ID)

	// Write new file first (for safety)
	newPath := filepath.Join(f.dir, newFilename)
	if err := saveCharacter(character, newPath); err != nil {
		return err
	}

	// If filename changed, delete old file
	if newFilename != oldFilename {
		oldPath := filepath.Join(f.dir, oldFilename)
		if err := os.Remove(oldPath); err != nil && !os.IsNotExist(err) {
			// Log but don't fail - the new file is already written
			fmt.Fprintf(os.Stderr, "warning: failed to remove old file %s: %v\n", oldPath, err)
		}
	}

	// Update cache
	f.cache[character.ID] = character
	f.files[character.ID] = newFilename

	return nil
}

// Delete removes a character by ID.
func (f *FsProvider) Delete(characterID string) error {
	f.mu.Lock()
	defer f.mu.Unlock()

	filename, ok := f.files[characterID]
	if !ok {
		return ErrCharacterNotFound
	}

	// Remove file
	path := filepath.Join(f.dir, filename)
	if err := os.Remove(path); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("failed to delete file %s: %w", path, err)
	}

	// Remove from cache
	delete(f.cache, characterID)
	delete(f.files, characterID)

	return nil
}

// List returns characters matching the query filters.
func (f *FsProvider) List(query dfm.CharacterQuery) ([]dfm.Character, error) {
	f.mu.RLock()
	defer f.mu.RUnlock()

	var result []dfm.Character
	for _, character := range f.cache {
		if matchesQuery(character, query) {
			result = append(result, character)
		}
	}
	return result, nil
}

// validateCharacterName checks if a character name contains only valid characters.
func validateCharacterName(name string) error {
	if name == "" {
		return nil // Empty names are allowed
	}
	if !validNamePattern.MatchString(name) {
		return ErrInvalidCharacterName
	}
	return nil
}

// generateFilename creates a filename from character name and ID.
// Format: {name_lowercase_underscores}_{uuid}.json
func generateFilename(name, id string) string {
	// Convert name to lowercase and replace spaces with underscores
	safeName := strings.ToLower(name)
	safeName = strings.Map(func(r rune) rune {
		if r == ' ' {
			return '_'
		}
		if unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_' {
			return r
		}
		return -1 // Remove other characters
	}, safeName)

	if safeName == "" {
		safeName = "character"
	}

	return fmt.Sprintf("%s_%s.json", safeName, id)
}

// matchesQuery checks if a character matches the query filters.
func matchesQuery(character dfm.Character, query dfm.CharacterQuery) bool {
	if query.Spirit != "" && character.Spirit != query.Spirit {
		return false
	}
	if query.Player != "" && character.Player != query.Player {
		return false
	}
	if query.Group != "" && character.Group != query.Group {
		return false
	}
	return true
}

// loadCache loads all character files from the directory into memory.
func loadCache(dir string) (map[string]dfm.Character, map[string]string, error) {
	cache := make(map[string]dfm.Character)
	files := make(map[string]string)

	entries, err := os.ReadDir(dir)
	if err != nil {
		return cache, files, fmt.Errorf("failed to read directory: %w", err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		filename := entry.Name()

		// Extract UUID from filename
		matches := uuidV4Pattern.FindStringSubmatch(filename)
		if len(matches) < 2 {
			continue // Skip files that don't match the pattern
		}

		// Load the character
		path := filepath.Join(dir, filename)
		character, err := loadCharacter(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "warning: failed to load character from %s: %v\n", path, err)
			continue
		}

		cache[character.ID] = character
		files[character.ID] = filename
	}

	return cache, files, nil
}

// loadCharacter reads a character from a JSON file.
func loadCharacter(path string) (dfm.Character, error) {
	var character dfm.Character

	data, err := os.ReadFile(path)
	if err != nil {
		return character, err
	}

	if err := json.Unmarshal(data, &character); err != nil {
		return character, err
	}

	return character, nil
}

// saveCharacter writes a character to a JSON file.
func saveCharacter(character dfm.Character, path string) error {
	data, err := json.MarshalIndent(character, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal character: %w", err)
	}

	if err := os.WriteFile(path, data, 0644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}
