package dfm

// Stunt represents a character stunt in the Fate Condensed system.
// Stunts are cool techniques, tricks, or bits of equipment that make characters unique.
type Stunt struct {
	// Title is the display name of the stunt
	Title string `json:"title" yaml:"title"`
	// Description explains what the stunt accomplishes
	Description string `json:"description" yaml:"description"`
}
