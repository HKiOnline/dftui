package dfdb

import (
	"os"
	"path/filepath"
	"sync"
	"testing"

	"github.com/hkionline/dftui/dflib/dfm"
)

func TestNewFsProvider(t *testing.T) {
	// Create temp directory
	dir := t.TempDir()

	provider, err := NewFsProvider(dir)
	if err != nil {
		t.Fatalf("Failed to create provider: %v", err)
	}

	if provider == nil {
		t.Fatal("Provider should not be nil")
	}
}

func TestNewFsProviderCreatesDirectory(t *testing.T) {
	dir := filepath.Join(t.TempDir(), "new_dir", "characters")

	provider, err := NewFsProvider(dir)
	if err != nil {
		t.Fatalf("Failed to create provider: %v", err)
	}

	if provider == nil {
		t.Fatal("Provider should not be nil")
	}

	// Check directory exists
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		t.Error("Directory should have been created")
	}
}

func TestCreateAndRead(t *testing.T) {
	dir := t.TempDir()
	provider, _ := NewFsProvider(dir)

	char := dfm.Character{
		ID:       "550e8400-e29b-41d4-a716-446655440000",
		Name:     "Test Character",
		Spirit:   "human",
		Group:    "pc",
		Player:   "testuser",
		Category: "character",
	}

	// Create
	if err := provider.Create(char); err != nil {
		t.Fatalf("Failed to create character: %v", err)
	}

	// Read
	read, err := provider.Read(char.ID)
	if err != nil {
		t.Fatalf("Failed to read character: %v", err)
	}

	if read.Name != char.Name {
		t.Errorf("Name mismatch: got %s, want %s", read.Name, char.Name)
	}
	if read.Spirit != char.Spirit {
		t.Errorf("Spirit mismatch: got %s, want %s", read.Spirit, char.Spirit)
	}
}

func TestFileNamingConvention(t *testing.T) {
	dir := t.TempDir()
	provider, _ := NewFsProvider(dir)

	char := dfm.Character{
		ID:     "550e8400-e29b-41d4-a716-446655440000",
		Name:   "John Smith",
		Spirit: "human",
		Group:  "pc",
	}

	if err := provider.Create(char); err != nil {
		t.Fatalf("Failed to create character: %v", err)
	}

	// Check file was created with correct naming
	expectedFile := "john_smith_550e8400-e29b-41d4-a716-446655440000.json"
	path := filepath.Join(dir, expectedFile)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		t.Errorf("Expected file %s to exist", expectedFile)
	}
}

func TestUpdate(t *testing.T) {
	dir := t.TempDir()
	provider, _ := NewFsProvider(dir)

	char := dfm.Character{
		ID:     "550e8400-e29b-41d4-a716-446655440000",
		Name:   "Original Name",
		Spirit: "human",
		Group:  "pc",
	}

	provider.Create(char)

	// Update
	char.Name = "Updated Name"
	if err := provider.Update(char); err != nil {
		t.Fatalf("Failed to update character: %v", err)
	}

	// Verify update
	read, _ := provider.Read(char.ID)
	if read.Name != "Updated Name" {
		t.Errorf("Name not updated: got %s, want Updated Name", read.Name)
	}
}

func TestUpdateRenamesFile(t *testing.T) {
	dir := t.TempDir()
	provider, _ := NewFsProvider(dir)

	char := dfm.Character{
		ID:     "550e8400-e29b-41d4-a716-446655440000",
		Name:   "Old Name",
		Spirit: "human",
		Group:  "pc",
	}

	provider.Create(char)

	oldFile := filepath.Join(dir, "old_name_550e8400-e29b-41d4-a716-446655440000.json")
	if _, err := os.Stat(oldFile); os.IsNotExist(err) {
		t.Fatal("Old file should exist")
	}

	// Update with new name
	char.Name = "New Name"
	provider.Update(char)

	// Old file should be gone
	if _, err := os.Stat(oldFile); !os.IsNotExist(err) {
		t.Error("Old file should have been deleted")
	}

	// New file should exist
	newFile := filepath.Join(dir, "new_name_550e8400-e29b-41d4-a716-446655440000.json")
	if _, err := os.Stat(newFile); os.IsNotExist(err) {
		t.Error("New file should exist")
	}
}

func TestDelete(t *testing.T) {
	dir := t.TempDir()
	provider, _ := NewFsProvider(dir)

	char := dfm.Character{
		ID:     "550e8400-e29b-41d4-a716-446655440000",
		Name:   "To Delete",
		Spirit: "human",
		Group:  "pc",
	}

	provider.Create(char)

	// Delete
	if err := provider.Delete(char.ID); err != nil {
		t.Fatalf("Failed to delete character: %v", err)
	}

	// Verify deleted
	_, err := provider.Read(char.ID)
	if err != ErrCharacterNotFound {
		t.Error("Should return ErrCharacterNotFound after deletion")
	}
}

func TestReadNotFound(t *testing.T) {
	dir := t.TempDir()
	provider, _ := NewFsProvider(dir)

	_, err := provider.Read("non-existent-id")
	if err != ErrCharacterNotFound {
		t.Errorf("Expected ErrCharacterNotFound, got %v", err)
	}
}

func TestInvalidCharacterName(t *testing.T) {
	dir := t.TempDir()
	provider, _ := NewFsProvider(dir)

	tests := []struct {
		name     string
		charName string
		wantErr  bool
	}{
		{"Valid name", "John Smith", false},
		{"Valid with numbers", "John Smith 3rd", false},
		{"Invalid with dash", "John-Smith", true},
		{"Invalid with special char", "John@Smith", true},
		{"Invalid with slash", "John/Smith", true},
		{"Empty name allowed", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			char := dfm.Character{
				ID:     "550e8400-e29b-41d4-a716-446655440000",
				Name:   tt.charName,
				Spirit: "human",
				Group:  "pc",
			}

			err := provider.Create(char)
			if (err == ErrInvalidCharacterName) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}

			// Clean up for next test
			provider.Delete(char.ID)
		})
	}
}

func TestListAll(t *testing.T) {
	dir := t.TempDir()
	provider, _ := NewFsProvider(dir)

	chars := []dfm.Character{
		{ID: "id-1", Name: "Char One", Spirit: "human", Group: "pc"},
		{ID: "id-2", Name: "Char Two", Spirit: "vampire", Group: "npc"},
		{ID: "id-3", Name: "Char Three", Spirit: "ghoul", Group: "pc"},
	}

	for _, c := range chars {
		provider.Create(c)
	}

	// List all
	result, err := provider.List(dfm.CharacterQuery{})
	if err != nil {
		t.Fatalf("Failed to list: %v", err)
	}

	if len(result) != 3 {
		t.Errorf("Expected 3 characters, got %d", len(result))
	}
}

func TestListFilterBySpirit(t *testing.T) {
	dir := t.TempDir()
	provider, _ := NewFsProvider(dir)

	chars := []dfm.Character{
		{ID: "id-1", Name: "Human One", Spirit: "human", Group: "pc"},
		{ID: "id-2", Name: "Vampire One", Spirit: "vampire", Group: "pc"},
		{ID: "id-3", Name: "Human Two", Spirit: "human", Group: "npc"},
	}

	for _, c := range chars {
		provider.Create(c)
	}

	// Filter by spirit
	result, _ := provider.List(dfm.CharacterQuery{Spirit: "human"})
	if len(result) != 2 {
		t.Errorf("Expected 2 human characters, got %d", len(result))
	}
}

func TestListFilterByGroup(t *testing.T) {
	dir := t.TempDir()
	provider, _ := NewFsProvider(dir)

	chars := []dfm.Character{
		{ID: "id-1", Name: "PC One", Spirit: "human", Group: "pc"},
		{ID: "id-2", Name: "NPC One", Spirit: "human", Group: "npc"},
		{ID: "id-3", Name: "PC Two", Spirit: "vampire", Group: "pc"},
	}

	for _, c := range chars {
		provider.Create(c)
	}

	// Filter by group
	result, _ := provider.List(dfm.CharacterQuery{Group: "pc"})
	if len(result) != 2 {
		t.Errorf("Expected 2 PC characters, got %d", len(result))
	}
}

func TestListFilterByPlayer(t *testing.T) {
	dir := t.TempDir()
	provider, _ := NewFsProvider(dir)

	chars := []dfm.Character{
		{ID: "id-1", Name: "Char One", Spirit: "human", Group: "pc", Player: "alice"},
		{ID: "id-2", Name: "Char Two", Spirit: "human", Group: "pc", Player: "bob"},
		{ID: "id-3", Name: "Char Three", Spirit: "human", Group: "pc", Player: "alice"},
	}

	for _, c := range chars {
		provider.Create(c)
	}

	// Filter by player
	result, _ := provider.List(dfm.CharacterQuery{Player: "alice"})
	if len(result) != 2 {
		t.Errorf("Expected 2 characters for alice, got %d", len(result))
	}
}

func TestListMultipleFilters(t *testing.T) {
	dir := t.TempDir()
	provider, _ := NewFsProvider(dir)

	chars := []dfm.Character{
		{ID: "id-1", Name: "Char One", Spirit: "vampire", Group: "pc", Player: "alice"},
		{ID: "id-2", Name: "Char Two", Spirit: "vampire", Group: "npc", Player: "alice"},
		{ID: "id-3", Name: "Char Three", Spirit: "human", Group: "pc", Player: "alice"},
		{ID: "id-4", Name: "Char Four", Spirit: "vampire", Group: "pc", Player: "bob"},
	}

	for _, c := range chars {
		provider.Create(c)
	}

	// Filter by spirit AND group AND player
	result, _ := provider.List(dfm.CharacterQuery{Spirit: "vampire", Group: "pc", Player: "alice"})
	if len(result) != 1 {
		t.Errorf("Expected 1 character matching all filters, got %d", len(result))
	}
	if len(result) > 0 && result[0].ID != "id-1" {
		t.Errorf("Expected id-1, got %s", result[0].ID)
	}
}

func TestConcurrentAccess(t *testing.T) {
	dir := t.TempDir()
	provider, _ := NewFsProvider(dir)

	// Create initial character
	char := dfm.Character{
		ID:     "550e8400-e29b-41d4-a716-446655440000",
		Name:   "Concurrent Test",
		Spirit: "human",
		Group:  "pc",
	}
	provider.Create(char)

	// Concurrent reads and writes
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(2)

		// Reader goroutine
		go func() {
			defer wg.Done()
			provider.Read(char.ID)
		}()

		// Writer goroutine
		go func() {
			defer wg.Done()
			provider.List(dfm.CharacterQuery{})
		}()
	}

	wg.Wait()
	// If we get here without deadlock or race, test passes
}

func TestLoadExistingFiles(t *testing.T) {
	dir := t.TempDir()

	// Create a provider and add a character
	provider1, _ := NewFsProvider(dir)
	char := dfm.Character{
		ID:     "550e8400-e29b-41d4-a716-446655440000",
		Name:   "Persistent Character",
		Spirit: "vampire",
		Group:  "pc",
	}
	provider1.Create(char)

	// Create a new provider pointing to the same directory
	provider2, err := NewFsProvider(dir)
	if err != nil {
		t.Fatalf("Failed to create second provider: %v", err)
	}

	// Should be able to read the character
	read, err := provider2.Read(char.ID)
	if err != nil {
		t.Fatalf("Failed to read character from new provider: %v", err)
	}

	if read.Name != char.Name {
		t.Errorf("Name mismatch: got %s, want %s", read.Name, char.Name)
	}
}
