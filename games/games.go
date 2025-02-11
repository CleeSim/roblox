package games

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

// Service is a service used to interact with Roblox games.
type Service struct {
	// http is the HTTP client used to make requests.
	http *http.Client
}

var service *Service

// New creates a new games service.
func New(httpClient *http.Client) *Service {
	if service != nil {
		return service
	}

	service = &Service{
		http: httpClient,
	}

	return service
}

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

// GetUniverseID gets the ID of the universe associated with the given game ID.
func (s *Service) GetUniverseID(id int64) (int64, error) {
	resp, err := s.http.Get("https://apis.roblox.com/universes/v1/places/" + strconv.FormatInt(id, 10) + "/universe")
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var universe struct {
		ID int64 `json:"universeId"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&universe); err != nil {
		return 0, err
	}

	return universe.ID, nil
}

// Get gets information about a game.
func (s *Service) Get(id int64) (*Game, error) {
	resp, err := s.http.Get("https://games.roblox.com/v1/games?universeIds=" + strconv.FormatInt(id, 10))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var GameResponse GameResponse
	if err := json.NewDecoder(resp.Body).Decode(&GameResponse); err != nil {
		return nil, err
	}

	if len(GameResponse.Data) == 0 {
		return nil, errors.New("game not found")
	}

	return &GameResponse.Data[0], nil
}
