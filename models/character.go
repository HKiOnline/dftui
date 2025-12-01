package models

// CharacterType represents whether a character is a PC or NPC
type CharacterType string

const (
	// PC represents a Player Character (user's character)
	PC CharacterType = "PC"
	// NPC represents a Non-Player Character (gamemaster's character)
	NPC CharacterType = "NPC"
)

// Character represents a character in the Dark Fate RPG system
type Character struct {
	// ID is the unique identifier for the character
	ID string
	// Name is the character's display name
	Name string
	// Type indicates whether this is a PC or NPC
	Type CharacterType
}
