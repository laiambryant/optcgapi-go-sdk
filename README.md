# OPTCG API Go SDK

A Go SDK for the [One Piece TCG API](https://optcgapi.com/documentation).

[![Go Reference](https://pkg.go.dev/badge/github.com/laiambryant/optcgapi-sdk-go.svg)](https://pkg.go.dev/github.com/laiambryant/optcgapi-sdk-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/laiambryant/optcgapi-sdk-go)](https://goreportcard.com/report/github.com/laiambryant/optcgapi-sdk-go)
[![GitHub license](https://img.shields.io/github/license/laiambryant/optcgapi-sdk-go.svg)](https://github.com/laiambryant/optcgapi-sdk-go/blob/main/LICENSE)
[![GitHub issues](https://img.shields.io/github/issues/laiambryant/optcgapi-sdk-go.svg)](https://github.com/laiambryant/optcgapi-sdk-go/issues)
[![GitHub stars](https://img.shields.io/github/stars/laiambryant/optcgapi-sdk-go.svg)](https://github.com/laiambryant/optcgapi-sdk-go/stargazers)
[![Coverage Status](https://coveralls.io/repos/github/laiambryant/optcgapi-sdk-go/badge.svg?branch=main)](https://coveralls.io/github/laiambryant/optcgapi-sdk-go?branch=main)

## Installation

```bash
go get github.com/laiambryant/optcgapi-sdk-go
```

## Quick start

```go
import (
    optcgapi "github.com/laiambryant/optcgapi-sdk-go"
    "github.com/laiambryant/optcgapi-sdk-go/query"
    "github.com/laiambryant/optcgapi-sdk-go/enums"
)

sdk := optcgapi.New()

// All sets
sets, err := sdk.GetAllSets(ctx)

// All cards in a set
cards, err := sdk.GetSetCards(ctx, "OP-01")

// A specific card (returns all variants / parallel arts)
cards, err := sdk.GetSetCard(ctx, "OP01-001")

// Filtered search
q := query.New().
    CardName("Luffy").
    Color(enums.ColorRed).
    CardType(enums.CardTypeCharacter)
cards, err := sdk.GetFilteredSetCards(ctx, q)

// 14-day pricing history
cards, err := sdk.GetSetCardTwoWeeks(ctx, "OP01-001")
for i, day := range cards[0].PricingHistory() {
    fmt.Printf("Day %d: market $%.2f\n", i+1, *day.MarketPrice)
}
```

## Configuration

Pass functional options to `New` to override the defaults:

```go
sdk := optcgapi.New(
    client.WithBaseURL("https://www.optcgapi.com/api"),
    client.WithUserAgent("my-app/1.0"),
    client.WithCache(5 * time.Minute),
    client.WithHTTPClient(myCustomHTTPClient),
)
```

| Option | Default |
|--------|---------|
| `WithBaseURL` | `https://www.optcgapi.com/api` |
| `WithUserAgent` | `optcgapi-sdk-go` |
| `WithCache(ttl)` | disabled |
| `WithHTTPClient` | `http.DefaultClient` |

## API reference

### Sets

| Method | Endpoint |
|--------|----------|
| `GetAllSets(ctx)` | `GET /allSets/` |
| `GetAllSetCards(ctx)` | `GET /allSetCards/` |
| `GetSetCards(ctx, setID)` | `GET /sets/{setID}/` |
| `GetSetCard(ctx, cardSetID)` | `GET /sets/card/{cardSetID}/` |
| `GetFilteredSetCards(ctx, q)` | `GET /sets/filtered/` |
| `GetSetCardTwoWeeks(ctx, cardSetID)` | `GET /sets/card/twoweeks/{cardSetID}/` |

### Starter decks

| Method | Endpoint |
|--------|----------|
| `GetAllDecks(ctx)` | `GET /allDecks/` |
| `GetAllStarterDeckCards(ctx)` | `GET /allSTCards/` |
| `GetDeckCards(ctx, deckID)` | `GET /decks/{deckID}/` |
| `GetDeckCard(ctx, cardSetID)` | `GET /decks/card/{cardSetID}/` |
| `GetFilteredDeckCards(ctx, q)` | `GET /decks/filtered/` |
| `GetDeckCardTwoWeeks(ctx, cardSetID)` | `GET /decks/card/twoweeks/{cardSetID}/` |

### Promos

| Method | Endpoint |
|--------|----------|
| `GetAllPromoCards(ctx)` | `GET /allPromoCards/` |
| `GetFilteredPromoCards(ctx, q)` | `GET /promos/filtered/` |
| `GetPromoCard(ctx, cardSetID)` | `GET /promos/card/{cardSetID}/` |

## Query builder

Use `query.New()` to build filter parameters for the `GetFiltered*` methods:

```go
q := query.New().
    CardName("Zoro").
    Color(enums.ColorGreen).
    CardType(enums.CardTypeCharacter).
    Rarity(enums.RaritySuperRare).
    CardCost("4").
    CardPower("6000").
    Attribute(enums.AttributeSlash).
    SetID("OP-01")
```

Pass `nil` instead of a `*query.Query` for an unfiltered request.

### Enums

**`enums.Color`** — `ColorRed`, `ColorGreen`, `ColorBlue`, `ColorPurple`, `ColorBlack`, `ColorYellow`

**`enums.CardType`** — `CardTypeLeader`, `CardTypeCharacter`, `CardTypeEvent`, `CardTypeStage`

**`enums.Rarity`** — `RarityCommon`, `RarityUncommon`, `RarityRare`, `RaritySuperRare`, `RaritySecretRare`, `RarityLeader`, `RarityPromo`, `RaritySpecial`, `RarityManga`, `RarityTreasure`, `RarityAlternateArt`

**`enums.Attribute`** — `AttributeSlash`, `AttributeStrike`, `AttributeRanged`, `AttributeWisdom`, `AttributeSpecial`

## Error handling

```go
cards, err := sdk.GetSetCard(ctx, "OP01-001")
if errors.Is(err, client.ErrNotFound) {
    // 404 - card not found
}

var httpErr *client.HTTPError
if errors.As(err, &httpErr) {
    fmt.Printf("HTTP %d: %s\n", httpErr.Status, httpErr.Body)
}

var reqErr *client.RequestError
if errors.As(err, &reqErr) {
    fmt.Printf("request failed during %s: %v\n", reqErr.Op, reqErr.Err)
}
```

## Caching

Enable in-memory response caching to reduce API calls. The cache is keyed by full URL and is safe for concurrent use:

```go
sdk := optcgapi.New(client.WithCache(10 * time.Minute))
```

## Testing

```bash
go test ./...
go test ./... -coverprofile=coverage.out && go tool cover -func=coverage.out
```
