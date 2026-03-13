// Package enums defines named string types for the enumerated values accepted
// by the OPTCG API.
//
// Using these constants instead of bare strings makes queries self-documenting
// and protects against typos:
//
//	q := query.New().
//	    Color(enums.ColorRed).
//	    CardType(enums.CardTypeCharacter).
//	    Attribute(enums.AttributeSlash)
package enums
