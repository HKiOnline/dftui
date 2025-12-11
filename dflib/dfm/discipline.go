package dfm

// Discipline represents a vampire discipline in the Dark Fate system.
// Disciplines are unique to vampire spirit characters.
type Discipline struct {
	// Title is the discipline name (e.g., "animalism", "auspex", "celerity")
	Title string `json:"title" yaml:"title"`
	// Rating is the discipline level (default 0)
	Rating int `json:"rating" yaml:"rating"`
}
