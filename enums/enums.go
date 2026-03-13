package enums

// Color represents a One Piece TCG card colour used in filter parameters.
type Color string

const (
	ColorRed    Color = "red"
	ColorGreen  Color = "green"
	ColorBlue   Color = "blue"
	ColorPurple Color = "purple"
	ColorBlack  Color = "black"
	ColorYellow Color = "yellow"
)

// CardType represents a One Piece TCG card type used in filter parameters.
type CardType string

const (
	CardTypeLeader    CardType = "Leader"
	CardTypeCharacter CardType = "Character"
	CardTypeEvent     CardType = "Event"
	CardTypeStage     CardType = "Stage"
)

// Rarity represents a One Piece TCG card rarity used in filter parameters.
type Rarity string

const (
	RarityCommon      Rarity = "Common"
	RarityUncommon    Rarity = "Uncommon"
	RarityRare        Rarity = "Rare"
	RaritySuperRare   Rarity = "Super Rare"
	RaritySecretRare  Rarity = "Secret Rare"
	RarityLeader      Rarity = "Leader"
	RarityPromo       Rarity = "Promo"
	RaritySpecial     Rarity = "Special"
	RarityManga       Rarity = "Manga"
	RarityTreasure    Rarity = "Treasure Rare"
	RarityAlternateArt Rarity = "Alternate Art"
)

// Attribute represents a One Piece TCG card attribute used in filter
// parameters.
type Attribute string

const (
	AttributeSlash   Attribute = "Slash"
	AttributeStrike  Attribute = "Strike"
	AttributeRanged  Attribute = "Ranged"
	AttributeWisdom  Attribute = "Wisdom"
	AttributeSpecial Attribute = "Special"
)
