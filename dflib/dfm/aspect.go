// Package dfm provides Dark Fate character models for use across dflib packages.
package dfm

// Aspect represents a character aspect in the Fate Condensed system.
// Aspects are short phrases that describe who a character is or what is important to them.
type Aspect struct {
	// Type indicates the aspect category: "high concept", "trouble", "relationship", "free", "clan", or "covenant"
	Type string `json:"type" yaml:"type"`
	// Title is the display name of the aspect
	Title string `json:"title" yaml:"title"`
	// Description is an optional explanation of the aspect's effect
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
}
