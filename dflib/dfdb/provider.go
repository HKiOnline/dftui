// Package dfdb provides a database abstraction for storing Dark Fate characters.
package dfdb

import (
	"github.com/hkionline/dftui/dflib/dfm"
)

// Provider defines the interface for character storage backends.
type Provider interface {
	// Create stores a new character and returns an error if it fails.
	Create(character dfm.Character) error
	// Read retrieves a character by ID, returning an error if not found.
	Read(characterID string) (dfm.Character, error)
	// Update modifies an existing character, returning an error if it fails.
	Update(character dfm.Character) error
	// Delete removes a character by ID, returning an error if not found.
	Delete(characterID string) error
	// List returns characters matching the query filters.
	List(query dfm.CharacterQuery) ([]dfm.Character, error)
}

// ProviderConfiguration holds configuration for all provider types.
type ProviderConfiguration struct {
	// Provider is the type of provider: "filesystem"
	Provider string `yaml:"provider" json:"provider"`
	// Filesystem contains filesystem provider configuration
	Filesystem FsProviderConfiguration `yaml:"filesystem" json:"filesystem"`
}

// FsProviderConfiguration holds configuration for the filesystem provider.
type FsProviderConfiguration struct {
	// Directory is the path to store character JSON files
	Directory string `yaml:"directory" json:"directory"`
}
