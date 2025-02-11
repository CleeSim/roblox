package games

// GameResponse represents a response from the games API.
type GameResponse struct {
	Data []Game `json:"data"`
}

// Game represents a Roblox game.
type Game struct {
	// ID is the game's ID.
	ID int64 `json:"id"`

	// RootPlaceID is the ID of the game's root place.
	RootPlaceID int64 `json:"rootPlaceId"`

	// Name is the game's name.
	Name string `json:"name"`

	// Description is the game's description.
	Description string `json:"description"`

	// SourceName is the game's source name.
	SourceName string `json:"sourceName"`

	// SourceDescription is the game's source description.
	SourceDescription string `json:"sourceDescription"`

	// Creator is the game's creator.
	Creator Creator `json:"creator"`

	// Price is the game's price.
	Price *int `json:"price"`

	// AllowedGearGenres is a list of allowed gear genres.
	AllowedGearGenres []string `json:"allowedGearGenres"`

	// AllowedGearCategories is a list of allowed gear categories.
	AllowedGearCategories []string `json:"allowedGearCategories"`

	// IsGenreEnforced is whether the genre is enforced.
	IsGenreEnforced bool `json:"isGenreEnforced"`

	// CopyingAllowed is whether copying is allowed.
	CopyingAllowed bool `json:"copyingAllowed"`

	// Playing is the number of players playing the game.
	Playing int `json:"playing"`

	// Visits is the number of visits the game has.
	Visits int `json:"visits"`

	// MaxPlayers is the maximum number of players allowed in the game.
	MaxPlayers int `json:"maxPlayers"`

	// Created is the date the game was created.
	Created string `json:"created"`

	// Updated is the date the game was last updated.
	Updated string `json:"updated"`

	// StudioAccessToApisAllowed is whether studio access to APIs is allowed.
	StudioAccessToApisAllowed bool `json:"studioAccessToApisAllowed"`

	// CreateVipServersAllowed is whether creating VIP servers is allowed.
	CreateVipServersAllowed bool `json:"createVipServersAllowed"`

	// UniverseAvatarType is the universe avatar type.
	UniverseAvatarType string `json:"universeAvatarType"`

	// Genre is the game's genre.
	Genre string `json:"genre"`

	// GenreL1 is the game's first genre.
	GenreL1 string `json:"genre_l1"`

	// GenreL2 is the game's second genre.
	GenreL2 string `json:"genre_l2"`

	// IsAllGenre is whether the game is all genre.
	IsAllGenre bool `json:"isAllGenre"`

	// IsFavoritedByUser is whether the game is favorited by the user.
	IsFavoritedByUser bool `json:"isFavoritedByUser"`

	// FavoritedCount is the number of times the game has been favorited.
	FavoritedCount int `json:"favoritedCount"`
}

// Creator represents a game's creator.
type Creator struct {
	// ID is the creator's ID.
	ID int64 `json:"id"`

	// Name is the creator's name.
	Name string `json:"name"`

	// Type is the creator's type.
	Type string `json:"type"`

	// IsRNVAccount is whether the creator is an RNV account.
	IsRNVAccount bool `json:"isRNVAccount"`

	// HasVerifiedBadge is whether the creator has a verified badge.
	HasVerifiedBadge bool `json:"hasVerifiedBadge"`
}

// RecommendedGamesResponse represents a response from the recommended games API.
type RecommendedGamesResponse struct {
	// Data is a list of recommended games.
	Data []RecommendedGame `json:"data"`
}

// RecommendedGame represents a recommended game.
type RecommendedGame struct {
	// CreatorID is the creator's ID.
	CreatorID int64 `json:"creatorId"`

	// CreatorName is the creator's name.
	CreatorName string `json:"creatorName"`

	// CreatorType is the creator's type.
	CreatorType string `json:"creatorType"`

	// CreatorHasVerifiedBadge is whether the creator has a verified badge.
	CreatorHasVerifiedBadge bool `json:"creatorHasVerifiedBadge"`

	// TotalUpVotes is the total number of upvotes.
	TotalUpVotes int `json:"totalUpVotes"`

	// TotalDownVotes is the total number of downvotes.
	TotalDownVotes int `json:"totalDownVotes"`

	// UniverseID is the games universe ID.
	UniverseID int64 `json:"universeId"`

	// Name is the game's name.
	Name string `json:"name"`

	// PlaceID is the game's place ID.
	PlaceID int64 `json:"placeId"`

	// PlayerCount is the number of players playing the game.
	PlayerCount int `json:"playerCount"`

	// ImageToken is the game's image token.
	ImageToken string `json:"imageToken"`

	// IsSponsored is whether the game is sponsored.
	IsSponsored bool `json:"isSponsored"`

	// NativeAdData is the game's native ad data.
	NativeAdData string `json:"nativeAdData"`

	// IsShowSponsoredLabel is whether the sponsored label is shown.
	IsShowSponsoredLabel bool `json:"isShowSponsoredLabel"`

	// Price is the game's price.
	Price *int `json:"price"`

	// AnalyticsIdentifier is the game's analytics identifier.
	AnalyticsIdentifier *string `json:"analyticsIdentifier"`

	// GameDescription is the game's description.
	GameDescription string `json:"gameDescription"`

	// Genre is the game's genre.
	Genre string `json:"genre"`

	// MinimumAge is the minimum age required to play the game.
	MinimumAge int `json:"minimumAge"`

	// AgeRecommendationDisplayName is the age recommendation display name.
	AgeRecommendationDisplayName string `json:"ageRecommendationDisplayName"`
}

// GameMediaResponse represents a response from the game media API.
type GameMediaResponse struct {
	// Data is a list of game media.
	Data []GameMedia `json:"data"`
}

// GameMedia represents a media asset associated with a game.
type GameMedia struct {
	// AssetTypeID is the type ID of the media asset
	AssetTypeID int `json:"assetTypeId"`

	// AssetType is the type of the media asset
	AssetType string `json:"assetType"`

	// ImageID is the ID of the image
	ImageID int64 `json:"imageId"`

	// VideoHash is the hash of the video if present
	VideoHash *string `json:"videoHash"`

	// VideoTitle is the title of the video if present
	VideoTitle *string `json:"videoTitle"`

	// Approved indicates if the media has been approved
	Approved bool `json:"approved"`

	// AltText is the alternative text for the media
	AltText *string `json:"altText"`
}
