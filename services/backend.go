package services

import (
	"github.com/hkionline/dftui/models"
)

// Backend provides access to Dark Fate RPG backend services
// TODO: Implement actual backend integration with API client
type Backend interface {
	// GetUserCharacters returns the characters visible to a user
	GetUserCharacters(username string) ([]models.Character, error)
}

// StubBackend is a mock implementation for development
// TODO: Replace with actual backend client when API is ready
type StubBackend struct{}

// NewStubBackend creates a new stub backend service
func NewStubBackend() *StubBackend {
	return &StubBackend{}
}

// GetUserCharacters returns mock character data
// TODO: Implement actual backend integration
// This should call the Dark Fate RPG backend API to fetch:
// - User's player characters (PCs)
// - NPCs visible to the user based on their campaigns
func (b *StubBackend) GetUserCharacters(username string) ([]models.Character, error) {
	// Mock data for development
	// In a real implementation, this would:
	// 1. Authenticate/verify the username
	// 2. Query the backend API for user's characters
	// 3. Filter NPCs based on user's campaign visibility
	// 4. Return the combined list with proper error handling

	mockCharacters := []models.Character{
		{
			ID:   "char-1",
			Name: "Gandalf the Grey",
			Type: models.PC,
		},
		{
			ID:   "char-2",
			Name: "Shadowfax",
			Type: models.NPC,
		},
		{
			ID:   "char-3",
			Name: "Aragorn",
			Type: models.PC,
		},
		{
			ID:   "char-4",
			Name: "Saruman",
			Type: models.NPC,
		},
	}

	return mockCharacters, nil
}
