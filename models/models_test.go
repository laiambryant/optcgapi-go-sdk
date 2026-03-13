package models

import (
	"encoding/json"
	"testing"
)

func TestCardUnmarshal(t *testing.T) {
	data := `{
		"inventory_price": 1.25,
		"market_price": 2.50,
		"card_name": "Monkey.D.Luffy",
		"set_name": "Romance Dawn",
		"card_text": "Your Turn Once Per Turn",
		"set_id": "OP-01",
		"rarity": "Leader",
		"card_set_id": "OP01-001",
		"card_color": "red",
		"card_type": "Leader",
		"life": "5",
		"card_cost": null,
		"card_power": "5000",
		"sub_types": "Straw Hat Crew",
		"counter_amount": null,
		"attribute": "Strike",
		"date_scraped": "2024-01-15",
		"card_image_id": "OP01-001",
		"card_image": "https://example.com/OP01-001.jpg"
	}`
	var c Card
	if err := json.Unmarshal([]byte(data), &c); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if c.CardName != "Monkey.D.Luffy" {
		t.Fatalf("expected card name Monkey.D.Luffy, got %s", c.CardName)
	}
	if c.CardSetID != "OP01-001" {
		t.Fatalf("expected card set ID OP01-001, got %s", c.CardSetID)
	}
	if c.InventoryPrice == nil || *c.InventoryPrice != 1.25 {
		t.Fatalf("expected inventory price 1.25")
	}
	if c.MarketPrice == nil || *c.MarketPrice != 2.50 {
		t.Fatalf("expected market price 2.50")
	}
	if c.Life == nil || *c.Life != "5" {
		t.Fatalf("expected life 5")
	}
	if c.CardCost != nil {
		t.Fatalf("expected nil card cost for leader")
	}
	if c.CounterAmount != nil {
		t.Fatalf("expected nil counter amount")
	}
	if c.CardColor != "red" {
		t.Fatalf("expected red, got %s", c.CardColor)
	}
	if c.CardType != "Leader" {
		t.Fatalf("expected Leader, got %s", c.CardType)
	}
}

func TestCardCharacterUnmarshal(t *testing.T) {
	data := `{
		"inventory_price": 0.50,
		"market_price": 0.75,
		"card_name": "Roronoa Zoro",
		"set_name": "Romance Dawn",
		"card_text": "Rush",
		"set_id": "OP-01",
		"rarity": "Super Rare",
		"card_set_id": "OP01-025",
		"card_color": "red",
		"card_type": "Character",
		"life": null,
		"card_cost": "3",
		"card_power": "5000",
		"sub_types": "Straw Hat Crew",
		"counter_amount": "1000",
		"attribute": "Slash",
		"date_scraped": "2024-01-15",
		"card_image_id": "OP01-025",
		"card_image": "https://example.com/OP01-025.jpg"
	}`
	var c Card
	if err := json.Unmarshal([]byte(data), &c); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if c.Life != nil {
		t.Fatalf("expected nil life for character")
	}
	if c.CardCost == nil || *c.CardCost != "3" {
		t.Fatalf("expected card cost 3")
	}
	if c.CounterAmount == nil || *c.CounterAmount != "1000" {
		t.Fatalf("expected counter 1000")
	}
	if c.CardPower == nil || *c.CardPower != "5000" {
		t.Fatalf("expected power 5000")
	}
}

func TestSetUnmarshal(t *testing.T) {
	data := `{"set_name": "Romance Dawn", "set_id": "OP-01"}`
	var s Set
	if err := json.Unmarshal([]byte(data), &s); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if s.SetName != "Romance Dawn" || s.SetID != "OP-01" {
		t.Fatalf("unexpected set: %#v", s)
	}
}

func TestStarterDeckUnmarshal(t *testing.T) {
	data := `{"structure_deck_name": "Starter Deck 1: Straw Hat Crew", "structure_deck_id": "ST-01"}`
	var d StarterDeck
	if err := json.Unmarshal([]byte(data), &d); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if d.StructureDeckName != "Starter Deck 1: Straw Hat Crew" || d.StructureDeckID != "ST-01" {
		t.Fatalf("unexpected deck: %#v", d)
	}
}

func TestCardWithPricingUnmarshal(t *testing.T) {
	data := `{
		"inventory_price": 1.00,
		"market_price": 2.00,
		"card_name": "Luffy",
		"set_name": "Romance Dawn",
		"card_text": "text",
		"set_id": "OP-01",
		"rarity": "Leader",
		"card_set_id": "OP01-001",
		"card_color": "red",
		"card_type": "Leader",
		"life": "5",
		"card_cost": null,
		"card_power": "5000",
		"sub_types": "Straw Hat Crew",
		"counter_amount": null,
		"attribute": "Strike",
		"date_scraped": "2024-01-15",
		"card_image_id": "OP01-001",
		"card_image": "https://example.com/OP01-001.jpg",
		"Day1_Inventory_Price": 0.90,
		"Day1_Market_Price": 1.80,
		"Day2_Inventory_Price": 0.85,
		"Day2_Market_Price": 1.75,
		"Day3_Inventory_Price": null,
		"Day3_Market_Price": null,
		"Day4_Inventory_Price": null,
		"Day4_Market_Price": null,
		"Day5_Inventory_Price": null,
		"Day5_Market_Price": null,
		"Day6_Inventory_Price": null,
		"Day6_Market_Price": null,
		"Day7_Inventory_Price": null,
		"Day7_Market_Price": null,
		"Day8_Inventory_Price": null,
		"Day8_Market_Price": null,
		"Day9_Inventory_Price": null,
		"Day9_Market_Price": null,
		"Day10_Inventory_Price": null,
		"Day10_Market_Price": null,
		"Day11_Inventory_Price": null,
		"Day11_Market_Price": null,
		"Day12_Inventory_Price": null,
		"Day12_Market_Price": null,
		"Day13_Inventory_Price": null,
		"Day13_Market_Price": null
	}`
	var c CardWithPricing
	if err := json.Unmarshal([]byte(data), &c); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if c.CardName != "Luffy" {
		t.Fatalf("expected Luffy, got %s", c.CardName)
	}
	if c.Day1InventoryPrice == nil || *c.Day1InventoryPrice != 0.90 {
		t.Fatalf("expected Day1 inventory 0.90")
	}
	if c.Day1MarketPrice == nil || *c.Day1MarketPrice != 1.80 {
		t.Fatalf("expected Day1 market 1.80")
	}
	if c.Day3InventoryPrice != nil {
		t.Fatalf("expected nil Day3 inventory")
	}

	history := c.PricingHistory()
	if len(history) != 13 {
		t.Fatalf("expected 13 day history, got %d", len(history))
	}
	if history[0].InventoryPrice == nil || *history[0].InventoryPrice != 0.90 {
		t.Fatalf("expected history[0] inventory 0.90")
	}
	if history[1].InventoryPrice == nil || *history[1].InventoryPrice != 0.85 {
		t.Fatalf("expected history[1] inventory 0.85")
	}
	if history[2].InventoryPrice != nil {
		t.Fatalf("expected nil history[2] inventory")
	}
}
