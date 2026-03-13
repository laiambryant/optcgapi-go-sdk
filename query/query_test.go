package query

import (
	"fmt"
	"net/url"
	"strings"
	"testing"

	"github.com/laiambryant/optcgapi-go-sdk/enums"
)

func TestBuildEmpty(t *testing.T) {
	q := New()
	if got := q.Build(); got != "" {
		t.Fatalf("expected empty string for empty query, got %q", got)
	}
}

func TestChainingReturnsSamePointer(t *testing.T) {
	q := New()
	q2 := q.CardName("test")
	if q != q2 {
		t.Fatalf("expected methods to return same pointer receiver")
	}
}

func TestAllMethodsBuildsExpectedQuery(t *testing.T) {
	q := New()
	q.CardName("Luffy")
	q.Color(enums.ColorRed)
	q.SetID("OP-01")
	q.SetName("Romance Dawn")
	q.Rarity(enums.RaritySuperRare)
	q.CardType(enums.CardTypeCharacter)
	q.CardCost("5")
	q.CardPower("6000")
	q.Attribute(enums.AttributeSlash)
	q.CardImageID("OP01-001")
	pairs := [][2]string{
		{"card_name", "Luffy"},
		{"color", "red"},
		{"set_id", "OP-01"},
		{"set_name", "Romance Dawn"},
		{"rarity", "Super Rare"},
		{"card_type", "Character"},
		{"card_cost", "5"},
		{"card_power", "6000"},
		{"attribute", "Slash"},
		{"card_image_id", "OP01-001"},
	}
	var parts []string
	for _, p := range pairs {
		parts = append(parts, fmt.Sprintf("%s=%s", url.QueryEscape(p[0]), url.QueryEscape(p[1])))
	}
	expected := "?" + strings.Join(parts, "&")
	if got := q.Build(); got != expected {
		t.Fatalf("unexpected query string:\n got: %s\nwant: %s", got, expected)
	}
}

func TestQueryEscaping(t *testing.T) {
	q := New()
	q.CardName("a b&c=d")
	expected := "?" + fmt.Sprintf("%s=%s", url.QueryEscape("card_name"), url.QueryEscape("a b&c=d"))
	if got := q.Build(); got != expected {
		t.Fatalf("escaping mismatch: got %q want %q", got, expected)
	}
}

func TestParamCustom(t *testing.T) {
	q := New().Param("custom_key", "custom_value")
	expected := "?custom_key=custom_value"
	if got := q.Build(); got != expected {
		t.Fatalf("expected %q, got %q", expected, got)
	}
}
