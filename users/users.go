package users

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

// Service is a service used to interact with Roblox users.
type Service struct {
	// http is the HTTP client used to make requests.
	http *http.Client
}

var service *Service

// New creates a new users service.
func New(httpClient *http.Client) *Service {
	if service != nil {
		return service
	}

	service = &Service{
		http: httpClient,
	}

	return service
}

// User represents a Roblox user.
type User struct {
	// ID is the user's ID.
	ID int64 `json:"id"`

	// Username is the user's username.
	Name string `json:"name"`

	// DisplayName is the user's display name.
	DisplayName string `json:"displayName"`

	// Description is the user's description.
	Description string `json:"description"`

	// IsBanned is whether the user is banned.
	IsBanned bool `json:"isBanned"`

	// Created is the date the user was created.
	Created string `json:"created"`

	// HasVerifiedBadge is whether the user has a verified badge.
	HasVerifiedBadge bool `json:"hasVerifiedBadge"`

	// ExternalAppDisplayName is the user's display name in an external app.
	ExternalAppDisplayName string `json:"externalAppDisplayName"`
}

// UserSearchRequest represents a request to search for Roblox users by usernames.
type UserSearchRequest struct {
	// Usernames is a list of usernames to search for.
	Usernames []string `json:"usernames"`
}

// UserSearchResponse represents a response to a user search request.
type UserSearchResponse struct {
	// Data is a list of user IDs and usernames.
	Data []struct {
		// ID is the user's ID.
		ID int64 `json:"id"`

		// Name is the user's username.
		Name string `json:"name"`
	} `json:"data"`
}

// Get fetches a user by their ID.
func (s *Service) Get(id int64) (*User, error) {
	resp, err := s.http.Get("https://users.roblox.com/v1/users/" + strconv.FormatInt(id, 10))
	if err != nil {
		return nil, fmt.Errorf("failed to fetch user data: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch user data: %w", errors.New(resp.Status))
	}

	var user User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	return &user, nil
}

// Search searches for users by their usernames.
func (s *Service) Search(usernames []string) (*UserSearchResponse, error) {
	reqBody := UserSearchRequest{
		Usernames: usernames,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequest("POST", "https://users.roblox.com/v1/usernames/users", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := s.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var result UserSearchResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &result, nil
}
