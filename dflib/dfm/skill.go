package dfm

// Skill represents a character skill in the Fate Condensed system.
// Skills tell what characters can do.
type Skill struct {
	// Title is the skill name (e.g., "academics", "athletics", "larceny")
	Title string `json:"title" yaml:"title"`
	// Group categorizes the skill: "mental", "physical", or "social"
	Group string `json:"group" yaml:"group"`
	// Rating is the skill level (default 0)
	Rating int `json:"rating" yaml:"rating"`
}
