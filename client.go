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

	NFLPlayers AllPlayersJSON
}

// NewClient creates a new Sleeper Client.
func NewClient() (Client, error) {
	c := Client{
		httpClient: &http.Client{
			Timeout: time.Minute,
		},
		sleeperURL: sleeperBaseURL,
	}
	//players, err := c.GetAllPlayers()
	/*if err != nil {
		return c, err
	}*/
	//c.NFLPlayers = players
	return c, nil
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

// GetLeagueInfo returns info about the provided league.
func (c Client) GetLeagueInfo(leagueID string) (LeagueInfoJSON, error) {
	res := LeagueInfoJSON{}
	err := c.sendRequest(fmt.Sprintf("/league/%s", leagueID), &res)
	return res, err
}

// GetLeagueRosters returns the rosters currently active in the provided league.
func (c Client) GetLeagueRosters(leagueID string) (RostersJSON, error) {
	res := RostersJSON{}
	err := c.sendRequest(fmt.Sprintf("/league/%s/rosters", leagueID), &res)
	return res, err
}

// GetLeagueUsers returns the users currently active in the provided league.
func (c Client) GetLeagueUsers(leagueID string) (UsersJSON, error) {
	res := UsersJSON{}
	err := c.sendRequest(fmt.Sprintf("/league/%s/users", leagueID), &res)
	return res, err
}

// GetLeagueMatchups returns the matchups for the provided league and week.
func (c Client) GetLeagueMatchups(leagueID string, week int) (MatchupsJSON, error) {
	res := MatchupsJSON{}
	err := c.sendRequest(fmt.Sprintf("/league/%s/matchups/%d", leagueID, week), &res)
	return res, err
}
