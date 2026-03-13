package optcgapi

import (
	"context"
	"net/http"
	"testing"

	"github.com/laiambryant/optcgapi-go-sdk/client"
	"github.com/laiambryant/optcgapi-go-sdk/query"
)

type fakeHTTPClient struct {
	fn func(req *http.Request) (*http.Response, error)
}

func (f *fakeHTTPClient) Do(req *http.Request) (*http.Response, error) { return f.fn(req) }

func TestNewDefaultsAndFields(t *testing.T) {
	sdk := New()
	if sdk.Client == nil {
		t.Fatalf("expected Client to be non-nil")
	}
	if sdk.Client.BaseURL != "https://www.optcgapi.com/api" {
		t.Fatalf("unexpected default BaseURL: %s", sdk.Client.BaseURL)
	}
	if sdk.sets == nil || sdk.decks == nil || sdk.cards == nil || sdk.cardsWithPrx == nil {
		t.Fatalf("expected all endpoints to be initialized")
	}
}

func TestNewWithOptionsOverrides(t *testing.T) {
	fake := &fakeHTTPClient{fn: func(req *http.Request) (*http.Response, error) {
		return client.NewMockResponse(200, `[]`), nil
	}}
	sdk := New(client.WithBaseURL("http://example"), client.WithUserAgent("custom-agent"), client.WithHTTPClient(fake))
	if sdk.Client.BaseURL != "http://example" {
		t.Fatalf("expected base url override, got %s", sdk.Client.BaseURL)
	}
	if sdk.Client.UserAgent != "custom-agent" {
		t.Fatalf("expected user agent override, got %s", sdk.Client.UserAgent)
	}
	if sdk.Client.HTTP != fake {
		t.Fatalf("expected provided HTTP client to be used")
	}
}

func TestGetAllSets(t *testing.T) {
	fake := &fakeHTTPClient{fn: func(req *http.Request) (*http.Response, error) {
		if req.URL.Path != "/allSets/" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		return client.NewMockResponse(200, `[{"set_name": "Romance Dawn", "set_id": "OP-01"}]`), nil
	}}
	sdk := New(client.WithBaseURL("http://example"), client.WithHTTPClient(fake))
	sets, err := sdk.GetAllSets(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(sets) != 1 || sets[0].SetName != "Romance Dawn" {
		t.Fatalf("unexpected sets: %#v", sets)
	}
}

func TestGetAllSetsError(t *testing.T) {
	fake := &fakeHTTPClient{fn: func(req *http.Request) (*http.Response, error) {
		return client.NewMockResponse(500, "error"), nil
	}}
	sdk := New(client.WithBaseURL("http://example"), client.WithHTTPClient(fake))
	_, err := sdk.GetAllSets(context.Background())
	if err == nil {
		t.Fatalf("expected error")
	}
}

func TestGetAllSetCards(t *testing.T) {
	fake := &fakeHTTPClient{fn: func(req *http.Request) (*http.Response, error) {
		if req.URL.Path != "/allSetCards/" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		return client.NewMockResponse(200, `[{"card_name": "Luffy", "card_set_id": "OP01-001", "set_id": "OP-01", "set_name": "Romance Dawn", "card_text": "", "rarity": "Leader", "card_color": "red", "card_type": "Leader", "sub_types": "", "attribute": "", "date_scraped": "", "card_image_id": "", "card_image": ""}]`), nil
	}}
	sdk := New(client.WithBaseURL("http://example"), client.WithHTTPClient(fake))
	cards, err := sdk.GetAllSetCards(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(cards) != 1 || cards[0].CardName != "Luffy" {
		t.Fatalf("unexpected cards: %#v", cards)
	}
}

func TestGetSetCards(t *testing.T) {
	fake := &fakeHTTPClient{fn: func(req *http.Request) (*http.Response, error) {
		if req.URL.Path != "/sets/OP-07/" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		return client.NewMockResponse(200, `[{"card_name": "Card1", "card_set_id": "OP07-001", "set_id": "OP-07", "set_name": "Set7", "card_text": "", "rarity": "", "card_color": "", "card_type": "", "sub_types": "", "attribute": "", "date_scraped": "", "card_image_id": "", "card_image": ""}]`), nil
	}}
	sdk := New(client.WithBaseURL("http://example"), client.WithHTTPClient(fake))
	cards, err := sdk.GetSetCards(context.Background(), "OP-07")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(cards) != 1 || cards[0].CardSetID != "OP07-001" {
		t.Fatalf("unexpected cards: %#v", cards)
	}
}

func TestGetSetCard(t *testing.T) {
	fake := &fakeHTTPClient{fn: func(req *http.Request) (*http.Response, error) {
		if req.URL.Path != "/sets/card/OP01-001/" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		return client.NewMockResponse(200, `[{"card_name": "Luffy", "card_set_id": "OP01-001", "set_id": "OP-01", "set_name": "Romance Dawn", "card_text": "", "rarity": "", "card_color": "", "card_type": "", "sub_types": "", "attribute": "", "date_scraped": "", "card_image_id": "", "card_image": ""}]`), nil
	}}
	sdk := New(client.WithBaseURL("http://example"), client.WithHTTPClient(fake))
	cards, err := sdk.GetSetCard(context.Background(), "OP01-001")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(cards) != 1 || cards[0].CardName != "Luffy" {
		t.Fatalf("unexpected cards: %#v", cards)
	}
}

func TestGetFilteredSetCards(t *testing.T) {
	fake := &fakeHTTPClient{fn: func(req *http.Request) (*http.Response, error) {
		if req.URL.Path != "/sets/filtered/" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		if req.URL.Query().Get("card_name") != "Zoro" {
			t.Fatalf("unexpected query: %s", req.URL.RawQuery)
		}
		return client.NewMockResponse(200, `[{"card_name": "Roronoa Zoro", "card_set_id": "OP01-025", "set_id": "OP-01", "set_name": "Romance Dawn", "card_text": "", "rarity": "", "card_color": "", "card_type": "", "sub_types": "", "attribute": "", "date_scraped": "", "card_image_id": "", "card_image": ""}]`), nil
	}}
	sdk := New(client.WithBaseURL("http://example"), client.WithHTTPClient(fake))
	q := query.New().CardName("Zoro")
	cards, err := sdk.GetFilteredSetCards(context.Background(), q)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(cards) != 1 || cards[0].CardName != "Roronoa Zoro" {
		t.Fatalf("unexpected cards: %#v", cards)
	}
}

func TestGetFilteredSetCardsNilQuery(t *testing.T) {
	fake := &fakeHTTPClient{fn: func(req *http.Request) (*http.Response, error) {
		if req.URL.Path != "/sets/filtered/" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		return client.NewMockResponse(200, `[]`), nil
	}}
	sdk := New(client.WithBaseURL("http://example"), client.WithHTTPClient(fake))
	cards, err := sdk.GetFilteredSetCards(context.Background(), nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(cards) != 0 {
		t.Fatalf("expected 0 cards, got %d", len(cards))
	}
}

func TestGetSetCardTwoWeeks(t *testing.T) {
	fake := &fakeHTTPClient{fn: func(req *http.Request) (*http.Response, error) {
		if req.URL.Path != "/sets/card/twoweeks/OP01-001/" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		body := `[{"card_name": "Luffy", "card_set_id": "OP01-001", "set_id": "OP-01", "set_name": "Romance Dawn", "card_text": "", "rarity": "", "card_color": "", "card_type": "", "sub_types": "", "attribute": "", "date_scraped": "", "card_image_id": "", "card_image": "", "Day1_Inventory_Price": 1.00, "Day1_Market_Price": 2.00}]`
		return client.NewMockResponse(200, body), nil
	}}
	sdk := New(client.WithBaseURL("http://example"), client.WithHTTPClient(fake))
	cards, err := sdk.GetSetCardTwoWeeks(context.Background(), "OP01-001")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(cards) != 1 || cards[0].CardName != "Luffy" {
		t.Fatalf("unexpected cards: %#v", cards)
	}
	if cards[0].Day1InventoryPrice == nil || *cards[0].Day1InventoryPrice != 1.00 {
		t.Fatalf("expected Day1 inventory price 1.00")
	}
}

func TestGetAllDecks(t *testing.T) {
	fake := &fakeHTTPClient{fn: func(req *http.Request) (*http.Response, error) {
		if req.URL.Path != "/allDecks/" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		return client.NewMockResponse(200, `[{"structure_deck_name": "Starter Deck 1: Straw Hat Crew", "structure_deck_id": "ST-01"}]`), nil
	}}
	sdk := New(client.WithBaseURL("http://example"), client.WithHTTPClient(fake))
	decks, err := sdk.GetAllDecks(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(decks) != 1 || decks[0].StructureDeckID != "ST-01" {
		t.Fatalf("unexpected decks: %#v", decks)
	}
}

func TestGetAllDecksError(t *testing.T) {
	fake := &fakeHTTPClient{fn: func(req *http.Request) (*http.Response, error) {
		return client.NewMockResponse(500, "error"), nil
	}}
	sdk := New(client.WithBaseURL("http://example"), client.WithHTTPClient(fake))
	_, err := sdk.GetAllDecks(context.Background())
	if err == nil {
		t.Fatalf("expected error")
	}
}

func TestGetAllStarterDeckCards(t *testing.T) {
	fake := &fakeHTTPClient{fn: func(req *http.Request) (*http.Response, error) {
		if req.URL.Path != "/allSTCards/" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		return client.NewMockResponse(200, `[]`), nil
	}}
	sdk := New(client.WithBaseURL("http://example"), client.WithHTTPClient(fake))
	cards, err := sdk.GetAllStarterDeckCards(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(cards) != 0 {
		t.Fatalf("expected 0 cards")
	}
}

func TestGetDeckCards(t *testing.T) {
	fake := &fakeHTTPClient{fn: func(req *http.Request) (*http.Response, error) {
		if req.URL.Path != "/decks/ST-01/" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		return client.NewMockResponse(200, `[{"card_name": "Luffy", "card_set_id": "ST01-001", "set_id": "ST-01", "set_name": "Straw Hat Crew", "card_text": "", "rarity": "", "card_color": "", "card_type": "", "sub_types": "", "attribute": "", "date_scraped": "", "card_image_id": "", "card_image": ""}]`), nil
	}}
	sdk := New(client.WithBaseURL("http://example"), client.WithHTTPClient(fake))
	cards, err := sdk.GetDeckCards(context.Background(), "ST-01")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(cards) != 1 || cards[0].CardSetID != "ST01-001" {
		t.Fatalf("unexpected cards: %#v", cards)
	}
}

func TestGetDeckCard(t *testing.T) {
	fake := &fakeHTTPClient{fn: func(req *http.Request) (*http.Response, error) {
		if req.URL.Path != "/decks/card/ST03-002/" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		return client.NewMockResponse(200, `[{"card_name": "Card", "card_set_id": "ST03-002", "set_id": "ST-03", "set_name": "Deck3", "card_text": "", "rarity": "", "card_color": "", "card_type": "", "sub_types": "", "attribute": "", "date_scraped": "", "card_image_id": "", "card_image": ""}]`), nil
	}}
	sdk := New(client.WithBaseURL("http://example"), client.WithHTTPClient(fake))
	cards, err := sdk.GetDeckCard(context.Background(), "ST03-002")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(cards) != 1 {
		t.Fatalf("expected 1 card")
	}
}

func TestGetFilteredDeckCards(t *testing.T) {
	fake := &fakeHTTPClient{fn: func(req *http.Request) (*http.Response, error) {
		if req.URL.Path != "/decks/filtered/" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		if req.URL.Query().Get("card_name") != "Zoro" || req.URL.Query().Get("color") != "red" {
			t.Fatalf("unexpected query: %s", req.URL.RawQuery)
		}
		return client.NewMockResponse(200, `[]`), nil
	}}
	sdk := New(client.WithBaseURL("http://example"), client.WithHTTPClient(fake))
	q := query.New().CardName("Zoro").Color("red")
	cards, err := sdk.GetFilteredDeckCards(context.Background(), q)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(cards) != 0 {
		t.Fatalf("expected 0 cards")
	}
}

func TestGetDeckCardTwoWeeks(t *testing.T) {
	fake := &fakeHTTPClient{fn: func(req *http.Request) (*http.Response, error) {
		if req.URL.Path != "/decks/card/twoweeks/ST01-001/" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		return client.NewMockResponse(200, `[{"card_name": "Luffy", "card_set_id": "ST01-001", "set_id": "ST-01", "set_name": "Straw Hat Crew", "card_text": "", "rarity": "", "card_color": "", "card_type": "", "sub_types": "", "attribute": "", "date_scraped": "", "card_image_id": "", "card_image": ""}]`), nil
	}}
	sdk := New(client.WithBaseURL("http://example"), client.WithHTTPClient(fake))
	cards, err := sdk.GetDeckCardTwoWeeks(context.Background(), "ST01-001")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(cards) != 1 {
		t.Fatalf("expected 1 card")
	}
}

func TestGetAllPromoCards(t *testing.T) {
	fake := &fakeHTTPClient{fn: func(req *http.Request) (*http.Response, error) {
		if req.URL.Path != "/allPromoCards/" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		return client.NewMockResponse(200, `[]`), nil
	}}
	sdk := New(client.WithBaseURL("http://example"), client.WithHTTPClient(fake))
	cards, err := sdk.GetAllPromoCards(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(cards) != 0 {
		t.Fatalf("expected 0 cards")
	}
}

func TestGetFilteredPromoCards(t *testing.T) {
	fake := &fakeHTTPClient{fn: func(req *http.Request) (*http.Response, error) {
		if req.URL.Path != "/promos/filtered/" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		return client.NewMockResponse(200, `[]`), nil
	}}
	sdk := New(client.WithBaseURL("http://example"), client.WithHTTPClient(fake))
	q := query.New().CardName("Luffy")
	cards, err := sdk.GetFilteredPromoCards(context.Background(), q)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(cards) != 0 {
		t.Fatalf("expected 0 cards")
	}
}

func TestGetPromoCard(t *testing.T) {
	fake := &fakeHTTPClient{fn: func(req *http.Request) (*http.Response, error) {
		if req.URL.Path != "/promos/card/P-001/" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		return client.NewMockResponse(200, `[{"card_name": "Promo Card", "card_set_id": "P-001", "set_id": "P", "set_name": "Promo", "card_text": "", "rarity": "Promo", "card_color": "", "card_type": "", "sub_types": "", "attribute": "", "date_scraped": "", "card_image_id": "", "card_image": ""}]`), nil
	}}
	sdk := New(client.WithBaseURL("http://example"), client.WithHTTPClient(fake))
	cards, err := sdk.GetPromoCard(context.Background(), "P-001")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(cards) != 1 || cards[0].CardName != "Promo Card" {
		t.Fatalf("unexpected cards: %#v", cards)
	}
}
