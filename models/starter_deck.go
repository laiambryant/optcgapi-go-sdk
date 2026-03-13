package models

// StarterDeck represents a single One Piece TCG starter deck as returned by
// /allDecks/.
type StarterDeck struct {
	// StructureDeckName is the full name of the starter deck.
	StructureDeckName string `json:"structure_deck_name"`
	// StructureDeckID is the starter deck identifier (e.g. "ST-01").
	StructureDeckID string `json:"structure_deck_id"`
}
