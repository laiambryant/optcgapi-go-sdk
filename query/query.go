package query

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/laiambryant/optcgapi-go-sdk/enums"
)

// Query is a fluent builder for filtered endpoint query parameters. Create one
// with [New], chain filter methods, then pass it to a filtered card method on
// [optcgapi.OPTCGAPI]. All methods return the same *Query pointer so calls can
// be chained.
type Query struct {
	params []param
}

type param struct {
	key   string
	value string
}

// New returns an empty Query.
func New() *Query {
	return &Query{}
}

func (q *Query) add(key, value string) {
	q.params = append(q.params, param{key, value})
}

// CardName filters cards whose name matches name (case-insensitive partial
// match).
func (q *Query) CardName(name string) *Query {
	q.add("card_name", name)
	return q
}

// Color filters cards by colour.
func (q *Query) Color(color enums.Color) *Query {
	q.add("color", string(color))
	return q
}

// SetID filters cards by set identifier (e.g. "OP-01").
func (q *Query) SetID(id string) *Query {
	q.add("set_id", id)
	return q
}

// SetName filters cards by set name.
func (q *Query) SetName(name string) *Query {
	q.add("set_name", name)
	return q
}

// Rarity filters cards by rarity.
func (q *Query) Rarity(rarity enums.Rarity) *Query {
	q.add("rarity", string(rarity))
	return q
}

// CardType filters cards by type (e.g. Leader, Character, Event, Stage).
func (q *Query) CardType(ct enums.CardType) *Query {
	q.add("card_type", string(ct))
	return q
}

// CardCost filters cards by cost.
func (q *Query) CardCost(cost string) *Query {
	q.add("card_cost", cost)
	return q
}

// CardPower filters cards by power value.
func (q *Query) CardPower(power string) *Query {
	q.add("card_power", power)
	return q
}

// Attribute filters cards by attribute.
func (q *Query) Attribute(attr enums.Attribute) *Query {
	q.add("attribute", string(attr))
	return q
}

// CardImageID filters cards by image identifier.
func (q *Query) CardImageID(id string) *Query {
	q.add("card_image_id", id)
	return q
}

// Param adds a custom key-value parameter to the query. This is useful for
// any filter parameters not covered by the typed methods.
func (q *Query) Param(key, value string) *Query {
	q.add(key, value)
	return q
}

// Build encodes all accumulated parameters as a URL query string beginning
// with "?". Returns an empty string if no parameters have been set. All keys
// and values are percent-encoded.
func (q *Query) Build() string {
	if len(q.params) == 0 {
		return ""
	}
	var parts []string
	for _, p := range q.params {
		k := url.QueryEscape(p.key)
		v := url.QueryEscape(p.value)
		parts = append(parts, fmt.Sprintf("%s=%s", k, v))
	}
	return "?" + strings.Join(parts, "&")
}
