package groups

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

// Service is a service used to interact with Roblox groups.
type Service struct {
	// http is the HTTP client used to make requests.
	http *http.Client
}

var service *Service

// New creates a new groups service.
func New(httpClient *http.Client) *Service {
	if service != nil {
		return service
	}

	service = &Service{
		http: httpClient,
	}

	return service
}

// Group represents a Roblox group.
type Group struct {
	// ID is the group's ID.
	ID int64 `json:"id"`

	// Name is the group's name.
	Name string `json:"name"`

	// Description is the group's description.
	Description string `json:"description"`

	// Owner represents the group owner's information.
	Owner struct {
		HasVerifiedBadge bool   `json:"hasVerifiedBadge"`
		UserID           int64  `json:"userId"`
		Username         string `json:"username"`
		DisplayName      string `json:"displayName"`
	} `json:"owner"`

	// Shout represents the current group shout.
	Shout struct {
		Body    string `json:"body"`
		Created string `json:"created"`
		Updated string `json:"updated"`
		Poster  struct {
			HasVerifiedBadge bool   `json:"hasVerifiedBadge"`
			UserID           int64  `json:"userId"`
			Username         string `json:"username"`
			DisplayName      string `json:"displayName"`
		} `json:"poster"`
	} `json:"shout"`

	// MemberCount is the number of members in the group.
	MemberCount int `json:"memberCount"`

	// IsBuildersClubOnly indicates if the group is restricted to Builders Club members.
	IsBuildersClubOnly bool `json:"isBuildersClubOnly"`

	// PublicEntryAllowed indicates if anyone can join the group.
	PublicEntryAllowed bool `json:"publicEntryAllowed"`

	// HasVerifiedBadge indicates if the group has a verified badge.
	HasVerifiedBadge bool `json:"hasVerifiedBadge"`
}

// Get fetches a group by its ID.
func (s *Service) Get(id int64) (*Group, error) {
	resp, err := s.http.Get("https://groups.roblox.com/v1/groups/" + strconv.FormatInt(id, 10))
	if err != nil {
		return nil, fmt.Errorf("failed to fetch group data: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch group data: %w", errors.New(resp.Status))
	}

	var group Group
	if err := json.NewDecoder(resp.Body).Decode(&group); err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	return &group, nil
}
