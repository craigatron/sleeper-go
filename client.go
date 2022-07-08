package sleeper

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Client is a client for interacting with the read-only Sleeper API.
type Client struct {
	httpClient *http.Client

	sleeperURL string
}

// NewClient creates a new Sleeper Client.
func NewClient() Client {
	return Client{
		httpClient: &http.Client{
			Timeout: time.Minute,
		},
		sleeperURL: sleeperBaseURL,
	}
}

func (c *Client) sendRequest(path string, v interface{}) error {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s", c.sleeperURL, path), nil)
	if err != nil {
		return err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		return fmt.Errorf("unknown error, status code: %d", res.StatusCode)
	}

	return json.NewDecoder(res.Body).Decode(&v)
}

// GetAllPlayers returns information about every NFL player Sleeper knows about.
func (c Client) GetAllPlayers() (AllPlayersJSON, error) {
	res := AllPlayersJSON{}
	err := c.sendRequest("/players/nfl", &res)
	return res, err
}

// GetTrendingPlayers fetches the currently trending NFL players on Sleeper.
func (c Client) GetTrendingPlayers(trendType TrendingPlayerType) (TrendingPlayersJSON, error) {
	res := TrendingPlayersJSON{}
	err := c.sendRequest(fmt.Sprintf("/players/nfl/trending/%s", trendType), &res)
	return res, err
}

// GetNflStatus returns the current status of the NFL fantasy season on Sleeper.
func (c Client) GetNflStatus() (SportStatusJSON, error) {
	res := SportStatusJSON{}
	err := c.sendRequest("/state/nfl", &res)
	return res, err
}
