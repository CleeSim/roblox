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
// Note that you need the Universe ID, not the Game ID.
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

// GetRecommended gets a list of recommended games.
// Note that you need the Universe ID, not the Game ID.
func (s *Service) GetRecommended(id int64, maxRows int8) ([]RecommendedGame, error) {
	if maxRows < 1 {
		return nil, errors.New("maxRows must be greater than 0")
	} else if maxRows > 6 {
		return nil, errors.New("maxRows must be less than or equal to 6")
	}

	resp, err := s.http.Get("https://games.roblox.com/v1/games/recommendations/game/" + strconv.FormatInt(id, 10) + "?maxRows=" + strconv.FormatInt(int64(maxRows), 10))
	if err != nil {
		return nil, err
	}

	var recommendedGames RecommendedGamesResponse
	if err := json.NewDecoder(resp.Body).Decode(&recommendedGames); err != nil {
		return nil, err
	}

	return recommendedGames.Data, nil
}

// GetMedia gets media associated with a game.
// Note that you need the Universe ID, not the Game ID.
func (s *Service) GetMedia(id int64) (*[]GameMedia, error) {
	resp, err := s.http.Get("https://games.roblox.com/v2/games/" + strconv.FormatInt(id, 10) + "/media")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var media GameMediaResponse
	if err := json.NewDecoder(resp.Body).Decode(&media); err != nil {
		return nil, err
	}

	return &media.Data, nil
}
