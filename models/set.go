package models

// Set represents a single One Piece TCG booster set as returned by /allSets/.
type Set struct {
	// SetName is the full name of the set (e.g. "Romance Dawn").
	SetName string `json:"set_name"`
	// SetID is the set identifier (e.g. "OP-01").
	SetID string `json:"set_id"`
}
