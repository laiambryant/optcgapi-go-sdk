package optcgapi

import (
	"context"
	"net/url"

	"github.com/laiambryant/optcgapi-go-sdk/client"
	"github.com/laiambryant/optcgapi-go-sdk/endpoint"
	"github.com/laiambryant/optcgapi-go-sdk/models"
	"github.com/laiambryant/optcgapi-go-sdk/query"
)

// OPTCGAPI is the top-level SDK client. Create one with [New] and call its
// methods to interact with the One Piece TCG API.
type OPTCGAPI struct {
	// Client is the underlying HTTP client. It can be used to customise
	// transport behaviour directly, but in most cases the functional options
	// passed to [New] are sufficient.
	Client       *client.Client
	sets         *endpoint.Endpoint[[]models.Set]
	decks        *endpoint.Endpoint[[]models.StarterDeck]
	cards        *endpoint.Endpoint[[]models.Card]
	cardsWithPrx *endpoint.Endpoint[[]models.CardWithPricing]
}

// New creates an OPTCGAPI SDK instance. The default base URL is
// https://www.optcgapi.com/api. Pass [client.Option] values to override
// the base URL, user-agent, HTTP client, or enable response caching.
func New(opts ...client.Option) *OPTCGAPI {
	c := client.NewHTTPClient(nil, opts...)
	return &OPTCGAPI{
		Client:       c,
		sets:         endpoint.New[[]models.Set](c),
		decks:        endpoint.New[[]models.StarterDeck](c),
		cards:        endpoint.New[[]models.Card](c),
		cardsWithPrx: endpoint.New[[]models.CardWithPricing](c),
	}
}

// --- Set endpoints ---

// GetAllSets returns all booster set names and IDs from /allSets/.
func (o *OPTCGAPI) GetAllSets(ctx context.Context) ([]models.Set, error) {
	return o.sets.Fetch(ctx, "/allSets/")
}

// GetAllSetCards returns all cards from all booster sets from /allSetCards/.
func (o *OPTCGAPI) GetAllSetCards(ctx context.Context) ([]models.Card, error) {
	return o.cards.Fetch(ctx, "/allSetCards/")
}

// GetSetCards returns all cards in the given set from /sets/{setID}/.
func (o *OPTCGAPI) GetSetCards(ctx context.Context, setID string) ([]models.Card, error) {
	return o.cards.Fetch(ctx, "/sets/"+url.PathEscape(setID)+"/")
}

// GetSetCard returns the card(s) matching the given card set ID from
// /sets/card/{cardSetID}/. Multiple results may be returned for cards with
// parallel (alternate art) variants.
func (o *OPTCGAPI) GetSetCard(ctx context.Context, cardSetID string) ([]models.Card, error) {
	return o.cards.Fetch(ctx, "/sets/card/"+url.PathEscape(cardSetID)+"/")
}

// GetFilteredSetCards queries /sets/filtered/ with the given query parameters
// and returns the matching cards. Pass nil for an unfiltered search.
func (o *OPTCGAPI) GetFilteredSetCards(ctx context.Context, q *query.Query) ([]models.Card, error) {
	qs := ""
	if q != nil {
		qs = q.Build()
	}
	return o.cards.Fetch(ctx, "/sets/filtered/"+qs)
}

// GetSetCardTwoWeeks returns card(s) matching the given card set ID along with
// 14-day pricing history from /sets/card/twoweeks/{cardSetID}/.
func (o *OPTCGAPI) GetSetCardTwoWeeks(ctx context.Context, cardSetID string) ([]models.CardWithPricing, error) {
	return o.cardsWithPrx.Fetch(ctx, "/sets/card/twoweeks/"+url.PathEscape(cardSetID)+"/")
}

// --- Starter deck endpoints ---

// GetAllDecks returns all starter deck names and IDs from /allDecks/.
func (o *OPTCGAPI) GetAllDecks(ctx context.Context) ([]models.StarterDeck, error) {
	return o.decks.Fetch(ctx, "/allDecks/")
}

// GetAllStarterDeckCards returns all cards from all starter decks from
// /allSTCards/.
func (o *OPTCGAPI) GetAllStarterDeckCards(ctx context.Context) ([]models.Card, error) {
	return o.cards.Fetch(ctx, "/allSTCards/")
}

// GetDeckCards returns all cards in the given starter deck from
// /decks/{deckID}/.
func (o *OPTCGAPI) GetDeckCards(ctx context.Context, deckID string) ([]models.Card, error) {
	return o.cards.Fetch(ctx, "/decks/"+url.PathEscape(deckID)+"/")
}

// GetDeckCard returns the card(s) matching the given card set ID from
// /decks/card/{cardSetID}/.
func (o *OPTCGAPI) GetDeckCard(ctx context.Context, cardSetID string) ([]models.Card, error) {
	return o.cards.Fetch(ctx, "/decks/card/"+url.PathEscape(cardSetID)+"/")
}

// GetFilteredDeckCards queries /decks/filtered/ with the given query
// parameters and returns the matching starter deck cards. Pass nil for an
// unfiltered search.
func (o *OPTCGAPI) GetFilteredDeckCards(ctx context.Context, q *query.Query) ([]models.Card, error) {
	qs := ""
	if q != nil {
		qs = q.Build()
	}
	return o.cards.Fetch(ctx, "/decks/filtered/"+qs)
}

// GetDeckCardTwoWeeks returns starter deck card(s) matching the given card set
// ID along with 14-day pricing history from /decks/card/twoweeks/{cardSetID}/.
func (o *OPTCGAPI) GetDeckCardTwoWeeks(ctx context.Context, cardSetID string) ([]models.CardWithPricing, error) {
	return o.cardsWithPrx.Fetch(ctx, "/decks/card/twoweeks/"+url.PathEscape(cardSetID)+"/")
}

// --- Promo endpoints ---

// GetAllPromoCards returns all promo cards from /allPromoCards/.
func (o *OPTCGAPI) GetAllPromoCards(ctx context.Context) ([]models.Card, error) {
	return o.cards.Fetch(ctx, "/allPromoCards/")
}

// GetFilteredPromoCards queries /promos/filtered/ with the given query
// parameters and returns the matching promo cards. Pass nil for an unfiltered
// search.
func (o *OPTCGAPI) GetFilteredPromoCards(ctx context.Context, q *query.Query) ([]models.Card, error) {
	qs := ""
	if q != nil {
		qs = q.Build()
	}
	return o.cards.Fetch(ctx, "/promos/filtered/"+qs)
}

// GetPromoCard returns the promo card(s) matching the given card set ID from
// /promos/card/{cardSetID}/.
func (o *OPTCGAPI) GetPromoCard(ctx context.Context, cardSetID string) ([]models.Card, error) {
	return o.cards.Fetch(ctx, "/promos/card/"+url.PathEscape(cardSetID)+"/")
}
