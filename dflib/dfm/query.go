package dfm

// CharacterQuery defines filters for listing characters.
// All filters are combined with AND logic.
// Empty string values mean "match all" for that field.
type CharacterQuery struct {
	// Spirit filters by character type: "vampire", "ghoul", "human", or empty for all
	Spirit string
	// Player filters by player username, or empty for all
	Player string
	// Group filters by character group: "pc", "npc", or empty for all
	Group string
}
