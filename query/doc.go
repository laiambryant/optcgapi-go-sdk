// Package query provides a fluent builder for constructing query strings
// accepted by the filtered card endpoints of the OPTCG API.
//
// Create a new [Query] with [New], chain the desired filter methods, then pass
// it to the filtered card methods on [optcgapi.OPTCGAPI]:
//
//	q := query.New().
//	    CardName("Luffy").
//	    Color(enums.ColorRed).
//	    CardType(enums.CardTypeCharacter)
//
//	cards, err := sdk.GetFilteredSetCards(ctx, q)
//
// [Build] encodes all accumulated parameters as a URL query string (e.g.
// "?card_name=Luffy&color=red"). Values are percent-encoded via
// [net/url.QueryEscape].
package query
