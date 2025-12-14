package services

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/hkionline/dftui/dflib/dfdb"
	"github.com/hkionline/dftui/dflib/dfm"
)

// TestGetUserCharacters tests the GetUserCharacters function with dfdb
func TestGetUserCharacters(t *testing.T) {
	testDir := "test_characters"
	defer os.RemoveAll(testDir)

	// Create test directory
	if err := os.MkdirAll(testDir, 0755); err != nil {
		t.Fatalf("Failed to create test directory: %v", err)
	}

	// Create test character files
	testChars := []dfm.Character{
		{
			ID:       "550e8400-e29b-41d4-a716-446655440000",
			Name:     "Test PC Character",
			Group:    string(dfm.PC),
			Spirit:   string(dfm.SpiritHuman),
			Player:   "testuser",
			Category: "character",
		},
		{
			ID:       "550e8400-e29b-41d4-a716-446655440001",
			Name:     "Another Test PC",
			Group:    string(dfm.PC),
			Spirit:   string(dfm.SpiritVampire),
			Player:   "testuser",
			Category: "character",
		},
		{
			ID:       "550e8400-e29b-41d4-a716-446655440002",
			Name:     "Test NPC Character",
			Group:    string(dfm.NPC),
			Spirit:   string(dfm.SpiritHuman),
			Player:   "gm",
			Category: "character",
		},
	}

	// Write test character files
	for _, char := range testChars {
		filename := filepath.Join(testDir, char.Name+"_"+char.ID+".json")
		data, err := json.MarshalIndent(char, "", "  ")
		if err != nil {
			t.Fatalf("Failed to marshal character: %v", err)
		}
		if err := os.WriteFile(filename, data, 0644); err != nil {
			t.Fatalf("Failed to write character file: %v", err)
		}
	}

	// Create backend with test directory
	backend, err := NewDFDBBackendForTest(testDir)
	if err != nil {
		t.Fatalf("Failed to create backend: %v", err)
	}

	// Test getting characters for testuser
	results, err := backend.GetUserCharacters("testuser")
	if err != nil {
		t.Fatalf("GetUserCharacters failed: %v", err)
	}

	// Should have 2 PCs and 1 NPC = 3 total
	if len(results) != 3 {
		t.Errorf("Expected 3 characters, got %d", len(results))
	}

	// First 2 should be PCs for testuser
	for i := 0; i < 2; i++ {
		if results[i].Group != string(dfm.PC) {
			t.Errorf("Expected PC at index %d, got %s", i, results[i].Group)
		}
		if results[i].Player != "testuser" {
			t.Errorf("Expected player 'testuser' at index %d, got '%s'", i, results[i].Player)
		}
	}

	// Last should be NPC
	if results[2].Group != string(dfm.NPC) {
		t.Errorf("Expected NPC at index 2, got %s", results[2].Group)
	}
}

// TestGetUserCharactersEmpty tests behavior with no characters
func TestGetUserCharactersEmpty(t *testing.T) {
	testDir := "test_empty"
	defer os.RemoveAll(testDir)

	if err := os.MkdirAll(testDir, 0755); err != nil {
		t.Fatalf("Failed to create test directory: %v", err)
	}

	backend, err := NewDFDBBackendForTest(testDir)
	if err != nil {
		t.Fatalf("Failed to create backend: %v", err)
	}

	results, err := backend.GetUserCharacters("nonexistent")
	if err != nil {
		t.Fatalf("GetUserCharacters failed: %v", err)
	}

	// Should return empty slice, not nil
	if results == nil {
		t.Errorf("Expected empty slice, got nil")
	}

	if len(results) != 0 {
		t.Errorf("Expected 0 characters, got %d", len(results))
	}
}

// NewDFDBBackendForTest creates a DFDBBackend for testing with a specific directory
func NewDFDBBackendForTest(dir string) (*DFDBBackend, error) {
	provider, err := dfdb.NewFsProvider(dir)
	if err != nil {
		return nil, err
	}
	return &DFDBBackend{provider: provider}, nil
}
