package models

import "encoding/json"

// DayPrice holds a single day's inventory and market price from the 14-day
// pricing history returned by the twoweeks endpoints.
type DayPrice struct {
	// InventoryPrice is the inventory price for that day.
	InventoryPrice *float64
	// MarketPrice is the market price for that day.
	MarketPrice *float64
}

// CardWithPricing extends [Card] with a 14-day pricing history as returned by
// the /sets/card/twoweeks/ and /decks/card/twoweeks/ endpoints.
type CardWithPricing struct {
	Card

	// Day1 through Day13 hold the pricing history for the past 14 days
	// (Day1 being the most recent historical day).
	Day1InventoryPrice  *float64 `json:"Day1_Inventory_Price"`
	Day1MarketPrice     *float64 `json:"Day1_Market_Price"`
	Day2InventoryPrice  *float64 `json:"Day2_Inventory_Price"`
	Day2MarketPrice     *float64 `json:"Day2_Market_Price"`
	Day3InventoryPrice  *float64 `json:"Day3_Inventory_Price"`
	Day3MarketPrice     *float64 `json:"Day3_Market_Price"`
	Day4InventoryPrice  *float64 `json:"Day4_Inventory_Price"`
	Day4MarketPrice     *float64 `json:"Day4_Market_Price"`
	Day5InventoryPrice  *float64 `json:"Day5_Inventory_Price"`
	Day5MarketPrice     *float64 `json:"Day5_Market_Price"`
	Day6InventoryPrice  *float64 `json:"Day6_Inventory_Price"`
	Day6MarketPrice     *float64 `json:"Day6_Market_Price"`
	Day7InventoryPrice  *float64 `json:"Day7_Inventory_Price"`
	Day7MarketPrice     *float64 `json:"Day7_Market_Price"`
	Day8InventoryPrice  *float64 `json:"Day8_Inventory_Price"`
	Day8MarketPrice     *float64 `json:"Day8_Market_Price"`
	Day9InventoryPrice  *float64 `json:"Day9_Inventory_Price"`
	Day9MarketPrice     *float64 `json:"Day9_Market_Price"`
	Day10InventoryPrice *float64 `json:"Day10_Inventory_Price"`
	Day10MarketPrice    *float64 `json:"Day10_Market_Price"`
	Day11InventoryPrice *float64 `json:"Day11_Inventory_Price"`
	Day11MarketPrice    *float64 `json:"Day11_Market_Price"`
	Day12InventoryPrice *float64 `json:"Day12_Inventory_Price"`
	Day12MarketPrice    *float64 `json:"Day12_Market_Price"`
	Day13InventoryPrice *float64 `json:"Day13_Inventory_Price"`
	Day13MarketPrice    *float64 `json:"Day13_Market_Price"`
}

// UnmarshalJSON populates Card fields (including flex counter_amount) then
// populates the pricing day fields.
func (c *CardWithPricing) UnmarshalJSON(data []byte) error {
	if err := c.Card.UnmarshalJSON(data); err != nil {
		return err
	}
	type pricingFields struct {
		Day1InventoryPrice  *float64 `json:"Day1_Inventory_Price"`
		Day1MarketPrice     *float64 `json:"Day1_Market_Price"`
		Day2InventoryPrice  *float64 `json:"Day2_Inventory_Price"`
		Day2MarketPrice     *float64 `json:"Day2_Market_Price"`
		Day3InventoryPrice  *float64 `json:"Day3_Inventory_Price"`
		Day3MarketPrice     *float64 `json:"Day3_Market_Price"`
		Day4InventoryPrice  *float64 `json:"Day4_Inventory_Price"`
		Day4MarketPrice     *float64 `json:"Day4_Market_Price"`
		Day5InventoryPrice  *float64 `json:"Day5_Inventory_Price"`
		Day5MarketPrice     *float64 `json:"Day5_Market_Price"`
		Day6InventoryPrice  *float64 `json:"Day6_Inventory_Price"`
		Day6MarketPrice     *float64 `json:"Day6_Market_Price"`
		Day7InventoryPrice  *float64 `json:"Day7_Inventory_Price"`
		Day7MarketPrice     *float64 `json:"Day7_Market_Price"`
		Day8InventoryPrice  *float64 `json:"Day8_Inventory_Price"`
		Day8MarketPrice     *float64 `json:"Day8_Market_Price"`
		Day9InventoryPrice  *float64 `json:"Day9_Inventory_Price"`
		Day9MarketPrice     *float64 `json:"Day9_Market_Price"`
		Day10InventoryPrice *float64 `json:"Day10_Inventory_Price"`
		Day10MarketPrice    *float64 `json:"Day10_Market_Price"`
		Day11InventoryPrice *float64 `json:"Day11_Inventory_Price"`
		Day11MarketPrice    *float64 `json:"Day11_Market_Price"`
		Day12InventoryPrice *float64 `json:"Day12_Inventory_Price"`
		Day12MarketPrice    *float64 `json:"Day12_Market_Price"`
		Day13InventoryPrice *float64 `json:"Day13_Inventory_Price"`
		Day13MarketPrice    *float64 `json:"Day13_Market_Price"`
	}
	var p pricingFields
	if err := json.Unmarshal(data, &p); err != nil {
		return err
	}
	c.Day1InventoryPrice = p.Day1InventoryPrice
	c.Day1MarketPrice = p.Day1MarketPrice
	c.Day2InventoryPrice = p.Day2InventoryPrice
	c.Day2MarketPrice = p.Day2MarketPrice
	c.Day3InventoryPrice = p.Day3InventoryPrice
	c.Day3MarketPrice = p.Day3MarketPrice
	c.Day4InventoryPrice = p.Day4InventoryPrice
	c.Day4MarketPrice = p.Day4MarketPrice
	c.Day5InventoryPrice = p.Day5InventoryPrice
	c.Day5MarketPrice = p.Day5MarketPrice
	c.Day6InventoryPrice = p.Day6InventoryPrice
	c.Day6MarketPrice = p.Day6MarketPrice
	c.Day7InventoryPrice = p.Day7InventoryPrice
	c.Day7MarketPrice = p.Day7MarketPrice
	c.Day8InventoryPrice = p.Day8InventoryPrice
	c.Day8MarketPrice = p.Day8MarketPrice
	c.Day9InventoryPrice = p.Day9InventoryPrice
	c.Day9MarketPrice = p.Day9MarketPrice
	c.Day10InventoryPrice = p.Day10InventoryPrice
	c.Day10MarketPrice = p.Day10MarketPrice
	c.Day11InventoryPrice = p.Day11InventoryPrice
	c.Day11MarketPrice = p.Day11MarketPrice
	c.Day12InventoryPrice = p.Day12InventoryPrice
	c.Day12MarketPrice = p.Day12MarketPrice
	c.Day13InventoryPrice = p.Day13InventoryPrice
	c.Day13MarketPrice = p.Day13MarketPrice
	return nil
}

// PricingHistory returns the 14-day pricing history as a slice of [DayPrice]
// values, ordered from Day 1 (most recent) to Day 13 (oldest).
func (c *CardWithPricing) PricingHistory() []DayPrice {
	return []DayPrice{
		{InventoryPrice: c.Day1InventoryPrice, MarketPrice: c.Day1MarketPrice},
		{InventoryPrice: c.Day2InventoryPrice, MarketPrice: c.Day2MarketPrice},
		{InventoryPrice: c.Day3InventoryPrice, MarketPrice: c.Day3MarketPrice},
		{InventoryPrice: c.Day4InventoryPrice, MarketPrice: c.Day4MarketPrice},
		{InventoryPrice: c.Day5InventoryPrice, MarketPrice: c.Day5MarketPrice},
		{InventoryPrice: c.Day6InventoryPrice, MarketPrice: c.Day6MarketPrice},
		{InventoryPrice: c.Day7InventoryPrice, MarketPrice: c.Day7MarketPrice},
		{InventoryPrice: c.Day8InventoryPrice, MarketPrice: c.Day8MarketPrice},
		{InventoryPrice: c.Day9InventoryPrice, MarketPrice: c.Day9MarketPrice},
		{InventoryPrice: c.Day10InventoryPrice, MarketPrice: c.Day10MarketPrice},
		{InventoryPrice: c.Day11InventoryPrice, MarketPrice: c.Day11MarketPrice},
		{InventoryPrice: c.Day12InventoryPrice, MarketPrice: c.Day12MarketPrice},
		{InventoryPrice: c.Day13InventoryPrice, MarketPrice: c.Day13MarketPrice},
	}
}
