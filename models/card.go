package models

import (
	"encoding/json"
)

// Card represents a single One Piece TCG card as returned by the OPTCG API.
// Fields that may be null in the API response are pointer types so callers can
// distinguish an absent value from a zero value.
type Card struct {
	// InventoryPrice is the current inventory price.
	InventoryPrice *float64 `json:"inventory_price"`
	// MarketPrice is the current market price.
	MarketPrice *float64 `json:"market_price"`
	// CardName is the card name.
	CardName string `json:"card_name"`
	// SetName is the name of the set this card belongs to.
	SetName string `json:"set_name"`
	// CardText is the card effect or flavour text.
	CardText string `json:"card_text"`
	// SetID is the set identifier (e.g. "OP-01").
	SetID string `json:"set_id"`
	// Rarity is the card rarity (e.g. "Common", "Super Rare").
	Rarity string `json:"rarity"`
	// CardSetID is the card's unique set identifier (e.g. "OP01-001").
	CardSetID string `json:"card_set_id"`
	// CardColor is the card colour (e.g. "red", "blue").
	CardColor string `json:"card_color"`
	// CardType is the card type (e.g. "Leader", "Character", "Event", "Stage").
	CardType string `json:"card_type"`
	// Life is the leader's life value. Nil for non-Leader cards.
	Life *string `json:"life"`
	// CardCost is the card's cost to play. Nil for Leaders.
	CardCost *string `json:"card_cost"`
	// CardPower is the card's power value. Nil for Events and Stages.
	CardPower *string `json:"card_power"`
	// SubTypes lists the card's sub-types / traits (e.g. "Straw Hat Crew").
	SubTypes string `json:"sub_types"`
	// CounterAmount is the card's counter value. Nil if the card has no counter.
	CounterAmount *string `json:"counter_amount"`
	// Attribute is the card's attribute (e.g. "Slash", "Strike").
	Attribute string `json:"attribute"`
	// DateScraped is the timestamp when the data was last scraped.
	DateScraped string `json:"date_scraped"`
	// CardImageID is the identifier for the card image.
	CardImageID string `json:"card_image_id"`
	// CardImage is the URL of the card image.
	CardImage string `json:"card_image"`
}

// UnmarshalJSON handles counter_amount being sent as either a JSON string or a
// JSON number, depending on the API endpoint.
func (c *Card) UnmarshalJSON(data []byte) error {
	type noMethod Card
	var raw struct {
		noMethod
		CounterAmount json.RawMessage `json:"counter_amount"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	*c = Card(raw.noMethod)
	c.CounterAmount = parseFlexString(raw.CounterAmount)
	return nil
}

func parseFlexString(data json.RawMessage) *string {
	if len(data) == 0 || string(data) == "null" {
		return nil
	}
	var s string
	if err := json.Unmarshal(data, &s); err == nil {
		return &s
	}
	var n json.Number
	if err := json.Unmarshal(data, &n); err != nil {
		return nil
	}
	str := n.String()
	return &str
}
