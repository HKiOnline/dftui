package dfm

// Character represents a complete Dark Fate RPG character.
// Supports vampire, ghoul, and human spirit types.
type Character struct {
	// ID is the unique identifier (UUID v4)
	ID string `json:"id" yaml:"id"`
	// Player is the username of the player (for NPCs, uses gamemaster's username)
	Player string `json:"player" yaml:"player"`
	// Category is always "character"
	Category string `json:"category" yaml:"category"`
	// Spirit is the character type: "vampire", "ghoul", or "human"
	Spirit string `json:"spirit" yaml:"spirit"`
	// Group indicates PC ("pc") or NPC ("npc")
	Group string `json:"group" yaml:"group"`
	// Name is the full name of the character
	Name string `json:"name" yaml:"name"`
	// Gender is "male" or "female"
	Gender string `json:"gender" yaml:"gender"`
	// Aliases is a list of alternative names
	Aliases []string `json:"aliases" yaml:"aliases"`
	// Tags is a list of tags for the character
	Tags []string `json:"tags" yaml:"tags"`
	// Collectives is a list of collectives the character is affiliated with
	Collectives []string `json:"collectives" yaml:"collectives"`
	// EmbraceYear is the year of embrace (positive = AD, negative = BC)
	EmbraceYear int `json:"embrace_year" yaml:"embrace_year"`
	// SettingYear is the current year in the setting
	SettingYear int `json:"setting_year" yaml:"setting_year"`
	// Description is a character description
	Description string `json:"description" yaml:"description"`
	// Notes is free form notes about the character
	Notes string `json:"notes" yaml:"notes"`
	// Refresh is the fate point refresh value (default 3)
	Refresh int `json:"refresh" yaml:"refresh"`
	// FatePoint is the current number of fate points available (default 0)
	FatePoint int `json:"fatePoint" yaml:"fatePoint"`
	// BloodPotency is the potency of vampire's blood (vampire spirit only)
	BloodPotency int `json:"bloodPotency,omitempty" yaml:"bloodPotency,omitempty"`
	// Aspects is the list of character aspects
	Aspects []Aspect `json:"aspects" yaml:"aspects"`
	// Skills is the list of character skills
	Skills []Skill `json:"skills" yaml:"skills"`
	// Stunts is the list of character stunts
	Stunts []Stunt `json:"stunts" yaml:"stunts"`
	// Disciplines is the list of vampire disciplines (vampire spirit only)
	Disciplines []Discipline `json:"disciplines,omitempty" yaml:"disciplines,omitempty"`
	// Consequences is the list of consequence slots
	Consequences []Consequence `json:"consequences" yaml:"consequences"`
	// PhysicalStressLimit is the number of physical stress slots available
	PhysicalStressLimit int `json:"physicalStressLimit" yaml:"physicalStressLimit"`
	// PhysicalStressCurrent is the number of physical stress slots used
	PhysicalStressCurrent int `json:"physicalStressCurrent" yaml:"physicalStressCurrent"`
	// MentalStressLimit is the number of mental stress slots available
	MentalStressLimit int `json:"mentalStressLimit" yaml:"mentalStressLimit"`
	// MentalStressCurrent is the number of mental stress slots used
	MentalStressCurrent int `json:"mentalStressCurrent" yaml:"mentalStressCurrent"`
	// HungerStressLimit is the number of hunger stress slots (vampire spirit only)
	HungerStressLimit int `json:"hungerStressLimit,omitempty" yaml:"hungerStressLimit,omitempty"`
	// HungerStressCurrent is the number of hunger stress slots used (vampire spirit only)
	HungerStressCurrent int `json:"hungerStressCurrent,omitempty" yaml:"hungerStressCurrent,omitempty"`
}

// CharacterType represents whether a character is a PC or NPC
type CharacterType string

const (
	// PC represents a Player Character
	PC CharacterType = "pc"
	// NPC represents a Non-Player Character
	NPC CharacterType = "npc"
)

// SpiritType represents the type of character spirit
type SpiritType string

const (
	// SpiritVampire represents a vampire character
	SpiritVampire SpiritType = "vampire"
	// SpiritGhoul represents a ghoul character
	SpiritGhoul SpiritType = "ghoul"
	// SpiritHuman represents a human character
	SpiritHuman SpiritType = "human"
)
