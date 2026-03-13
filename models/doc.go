// Package models contains the data types returned by the OPTCG API.
//
// Each type corresponds directly to the JSON structure documented at
// https://optcgapi.com/documentation. Optional fields that may be absent or
// null in the API response are represented as pointer types so that callers can
// distinguish between a zero value and an absent value.
//
// Top-level response types:
//
//   - [Card] — a single card from any card endpoint
//   - [CardWithPricing] — a card with 14-day pricing history
//   - [Set] — a set entry from /allSets/
//   - [StarterDeck] — a starter deck entry from /allDecks/
package models
