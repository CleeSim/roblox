package roblox

import (
	"net/http"
	"time"

	"github.com/cleesim/roblox/games"
	"github.com/cleesim/roblox/groups"
	"github.com/cleesim/roblox/users"
)

type Client struct {
	// http is the HTTP client used to make requests.
	http *http.Client

	// Games is the games service.
	Games *games.Service

	// Groups is the groups service.
	Groups *groups.Service

	// Users is the users service.
	Users *users.Service
}

var client *Client

// New creates a new Roblox API client.
func New() *Client {
	if client != nil {
		return client
	}

	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}

	client = &Client{
		http:   httpClient,
		Games:  games.New(httpClient),
		Groups: groups.New(httpClient),
		Users:  users.New(httpClient),
	}

	return client
}
