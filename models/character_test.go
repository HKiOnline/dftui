package models

import (
	"testing"
)

func TestCharacterType(t *testing.T) {
	tests := []struct {
		name     string
		charType CharacterType
		expected string
	}{
		{"PC type", PC, "PC"},
		{"NPC type", NPC, "NPC"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if string(tt.charType) != tt.expected {
				t.Errorf("CharacterType = %v, want %v", tt.charType, tt.expected)
			}
		})
	}
}

func TestCharacterModel(t *testing.T) {
	char := Character{
		ID:   "test-1",
		Name: "Test Character",
		Type: PC,
	}

	if char.ID != "test-1" {
		t.Errorf("Character.ID = %v, want %v", char.ID, "test-1")
	}
	if char.Name != "Test Character" {
		t.Errorf("Character.Name = %v, want %v", char.Name, "Test Character")
	}
	if char.Type != PC {
		t.Errorf("Character.Type = %v, want %v", char.Type, PC)
	}
}

func TestCharacterCreation(t *testing.T) {
	tests := []struct {
		name         string
		character    Character
		expectedID   string
		expectedName string
		expectedType CharacterType
	}{
		{
			name: "PC character",
			character: Character{
				ID:   "pc-1",
				Name: "Hero",
				Type: PC,
			},
			expectedID:   "pc-1",
			expectedName: "Hero",
			expectedType: PC,
		},
		{
			name: "NPC character",
			character: Character{
				ID:   "npc-1",
				Name: "Villain",
				Type: NPC,
			},
			expectedID:   "npc-1",
			expectedName: "Villain",
			expectedType: NPC,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.character.ID != tt.expectedID {
				t.Errorf("Character.ID = %v, want %v", tt.character.ID, tt.expectedID)
			}
			if tt.character.Name != tt.expectedName {
				t.Errorf("Character.Name = %v, want %v", tt.character.Name, tt.expectedName)
			}
			if tt.character.Type != tt.expectedType {
				t.Errorf("Character.Type = %v, want %v", tt.character.Type, tt.expectedType)
			}
		})
	}
}
