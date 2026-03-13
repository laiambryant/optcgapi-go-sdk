// Package optcgapi provides a Go SDK for the One Piece TCG API.
//
// The SDK wraps the REST API at https://www.optcgapi.com/api/ and exposes
// typed methods for browsing sets, starter decks, promo cards, and querying
// cards with filters.
//
// # Quick start
//
//	sdk := optcgapi.New()
//
//	// All sets
//	sets, err := sdk.GetAllSets(ctx)
//
//	// Cards in a set
//	cards, err := sdk.GetSetCards(ctx, "OP-01")
//
//	// Filtered search
//	q := query.New().CardName("Luffy").Color("red")
//	cards, err := sdk.GetFilteredSetCards(ctx, q)
//
// # Configuration
//
// Pass functional options to New to override defaults:
//
//	sdk := optcgapi.New(
//	    client.WithBaseURL("https://www.optcgapi.com/api"),
//	    client.WithUserAgent("my-app/1.0"),
//	    client.WithCache(5*time.Minute),
//	)
//
// # Zero external dependencies
//
// The SDK uses only the Go standard library.
package optcgapi
