package services

import (
	"fmt"

	"github.com/hkionline/dftui/dflib/dfdb"
	"github.com/hkionline/dftui/dflib/dfm"
)

// Backend provides access to Dark Fate RPG backend services
type Backend interface {
	// GetUserCharacters returns the characters visible to a user
	GetUserCharacters(username string) ([]dfm.Character, error)
}

// DFDBBackend implements the Backend interface using dfdb filesystem provider
type DFDBBackend struct {
	provider dfdb.Provider
}

// NewDFDBBackend creates a new backend service using dfdb
func NewDFDBBackend() (*DFDBBackend, error) {
	// Initialize dfdb with the db directory
	provider, err := dfdb.NewFsProvider("db/characters")
	if err != nil {
		return nil, fmt.Errorf("failed to initialize dfdb provider: %w", err)
	}

	return &DFDBBackend{provider: provider}, nil
}

// GetUserCharacters loads character data from db/characters directory using dfdb
// Returns PCs for the specified username and all NPCs
func (b *DFDBBackend) GetUserCharacters(username string) ([]dfm.Character, error) {
	// Create query to get PCs for this user
	pcQuery := dfm.CharacterQuery{
		Player: username,
		Group:  string(dfm.PC),
	}

	// Get user's PCs
	pcs, err := b.provider.List(pcQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to load PC characters: %w", err)
	}

	// Create query for all NPCs (no player filter)
	npcQuery := dfm.CharacterQuery{
		Group: string(dfm.NPC),
	}

	// Get all NPCs
	npcs, err := b.provider.List(npcQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to load NPC characters: %w", err)
	}

	// Combine results - PCs first, then NPCs
	characters := make([]dfm.Character, 0, len(pcs)+len(npcs))
	characters = append(characters, pcs...)
	characters = append(characters, npcs...)

	return characters, nil
}
