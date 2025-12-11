package dfm

// Consequence represents a consequence aspect in the Fate Condensed system.
// Consequences are new aspects written when a character takes a hit,
// representing real harm and injury.
type Consequence struct {
	// Level indicates consequence severity: 2 (mild), 4 (moderate), or 6 (severe)
	Level int `json:"level" yaml:"level"`
	// IsActive indicates if this consequence is currently in play
	IsActive bool `json:"isActive" yaml:"isActive"`
	// Title is the consequence description (only filled when IsActive is true)
	Title string `json:"title" yaml:"title"`
}
