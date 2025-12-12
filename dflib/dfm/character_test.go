package dfm

import (
	"encoding/json"
	"testing"
)

func TestCharacterJSONMarshal(t *testing.T) {
	char := Character{
		ID:       "550e8400-e29b-41d4-a716-446655440000",
		Player:   "testuser",
		Category: "character",
		Spirit:   "vampire",
		Group:    "pc",
		Name:     "Test Vampire",
		Gender:   "male",
		Aliases:  []string{"The Dark One"},
		Tags:     []string{"vampire", "elder"},
		Aspects: []Aspect{
			{Type: "high concept", Title: "Ancient Vampire Lord", Description: "A powerful elder vampire"},
			{Type: "trouble", Title: "Blood Hunger", Description: ""},
		},
		Skills: []Skill{
			{Title: "academics", Group: "mental", Rating: 2},
			{Title: "athletics", Group: "physical", Rating: 3},
		},
		Stunts: []Stunt{
			{Title: "Night Vision", Description: "Can see in complete darkness"},
		},
		Disciplines: []Discipline{
			{Title: "auspex", Rating: 3},
			{Title: "dominate", Rating: 2},
		},
		Consequences: []Consequence{
			{Level: 2, IsActive: false, Title: ""},
			{Level: 4, IsActive: false, Title: ""},
			{Level: 6, IsActive: false, Title: ""},
		},
		PhysicalStressLimit:   3,
		PhysicalStressCurrent: 0,
		MentalStressLimit:     3,
		MentalStressCurrent:   1,
		HungerStressLimit:     3,
		HungerStressCurrent:   0,
		Refresh:               3,
	}

	// Marshal to JSON
	data, err := json.MarshalIndent(char, "", "  ")
	if err != nil {
		t.Fatalf("Failed to marshal character: %v", err)
	}

	// Unmarshal back
	var parsed Character
	if err := json.Unmarshal(data, &parsed); err != nil {
		t.Fatalf("Failed to unmarshal character: %v", err)
	}

	// Verify fields
	if parsed.ID != char.ID {
		t.Errorf("ID mismatch: got %s, want %s", parsed.ID, char.ID)
	}
	if parsed.Name != char.Name {
		t.Errorf("Name mismatch: got %s, want %s", parsed.Name, char.Name)
	}
	if parsed.Spirit != char.Spirit {
		t.Errorf("Spirit mismatch: got %s, want %s", parsed.Spirit, char.Spirit)
	}
	if parsed.Group != char.Group {
		t.Errorf("Group mismatch: got %s, want %s", parsed.Group, char.Group)
	}
	if len(parsed.Aspects) != len(char.Aspects) {
		t.Errorf("Aspects length mismatch: got %d, want %d", len(parsed.Aspects), len(char.Aspects))
	}
	if len(parsed.Skills) != len(char.Skills) {
		t.Errorf("Skills length mismatch: got %d, want %d", len(parsed.Skills), len(char.Skills))
	}
	if len(parsed.Disciplines) != len(char.Disciplines) {
		t.Errorf("Disciplines length mismatch: got %d, want %d", len(parsed.Disciplines), len(char.Disciplines))
	}
	if parsed.HungerStressLimit != char.HungerStressLimit {
		t.Errorf("HungerStressLimit mismatch: got %d, want %d", parsed.HungerStressLimit, char.HungerStressLimit)
	}
}

func TestCharacterJSONUnmarshalFromSpec(t *testing.T) {
	// Test unmarshaling a JSON string matching the spec format
	jsonData := `{
		"id": "test-uuid",
		"player": "gamemaster",
		"category": "character",
		"spirit": "human",
		"group": "npc",
		"name": "John Smith",
		"gender": "male",
		"aliases": ["Johnny"],
		"tags": ["merchant"],
		"collectives": ["Guild of Merchants"],
		"embrace_year": 1982,
		"setting_year": 2024,
		"description": "A friendly merchant",
		"notes": "Met in session 1",
		"refresh": 3,
		"aspects": [
			{"type": "high concept", "title": "Clever Merchant"},
			{"type": "trouble", "title": "Greedy"}
		],
		"skills": [
			{"title": "rapport", "group": "social", "rating": 4}
		],
		"stunts": [
			{"title": "Silver Tongue", "description": "Bonus to rapport"}
		],
		"consequences": [
			{"level": 2, "isActive": false, "title": ""},
			{"level": 4, "isActive": false, "title": ""},
			{"level": 6, "isActive": false, "title": ""}
		],
		"physicalStressLimit": 3,
		"physicalStressCurrent": 0,
		"mentalStressLimit": 3,
		"mentalStressCurrent": 0
	}`

	var char Character
	if err := json.Unmarshal([]byte(jsonData), &char); err != nil {
		t.Fatalf("Failed to unmarshal character: %v", err)
	}

	if char.ID != "test-uuid" {
		t.Errorf("ID mismatch: got %s, want test-uuid", char.ID)
	}
	if char.Spirit != "human" {
		t.Errorf("Spirit mismatch: got %s, want human", char.Spirit)
	}
	if char.Group != "npc" {
		t.Errorf("Group mismatch: got %s, want npc", char.Group)
	}
	if char.EmbraceYear != 1982 {
		t.Errorf("EmbraceYear mismatch: got %d, want 1982", char.EmbraceYear)
	}
	if char.SettingYear != 2024 {
		t.Errorf("SettingYear mismatch: got %d, want 2024", char.SettingYear)
	}
	if len(char.Collectives) != 1 || char.Collectives[0] != "Guild of Merchants" {
		t.Errorf("Collectives mismatch: got %v", char.Collectives)
	}
}

func TestCharacterTypes(t *testing.T) {
	tests := []struct {
		name     string
		charType CharacterType
		expected string
	}{
		{"PC type", PC, "pc"},
		{"NPC type", NPC, "npc"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if string(tt.charType) != tt.expected {
				t.Errorf("CharacterType = %v, want %v", tt.charType, tt.expected)
			}
		})
	}
}

func TestSpiritTypes(t *testing.T) {
	tests := []struct {
		name     string
		spirit   SpiritType
		expected string
	}{
		{"Vampire", SpiritVampire, "vampire"},
		{"Ghoul", SpiritGhoul, "ghoul"},
		{"Human", SpiritHuman, "human"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if string(tt.spirit) != tt.expected {
				t.Errorf("SpiritType = %v, want %v", tt.spirit, tt.expected)
			}
		})
	}
}
