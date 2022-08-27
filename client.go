package sleeper

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

// Client is a client for interacting with the read-only Sleeper API.
type Client struct {
	httpClient *http.Client

	sleeperURL   string
	graphqlToken string

	NFLPlayers AllPlayersJSON
}

func NewClient() (Client, error) {
	return NewClientWithToken("")
}

// NewClientWithToken creates a new Sleeper Client with a GraphQL token.
func NewClientWithToken(graphqlToken string) (Client, error) {
	c := Client{
		httpClient: &http.Client{
			Timeout: time.Minute,
		},
		sleeperURL:   sleeperBaseURL,
		graphqlToken: graphqlToken,
	}
	players, err := c.GetAllPlayers()
	if err != nil {
		return c, err
	}
	c.NFLPlayers = players
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

func (c *Client) sendGraphqlRequest(op interface{}, v interface{}) error {
	if c.graphqlToken == "" {
		return errors.New("cannot send GraphQL requests without a Sleeper token")
	}
	jsonBody, err := json.Marshal(op)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", sleeperGraphqlURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("authorization", c.graphqlToken)

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

// GetBatchScores returns scores and game info for the given week's games.
func (c Client) GetBatchScores(week int, season string) (BatchScoresJSON, error) {
	res := BatchScoresJSON{}
	op := map[string]interface{}{
		"operationName": "batch_scores",
		"variables":     struct{}{},
		"query":         fmt.Sprintf("query batch_scores {scores: scores(sport: \"nfl\",season_type: \"regular\",season: \"%s\",week: %d){date game_id metadata season season_type sport status week start_time}}", season, week),
	}
	err := c.sendGraphqlRequest(op, &res)
	return res, err
}

// GetPlayerStats returns actual and projected stats for the given set of players.
func (c Client) GetPlayerStats(playerIds []string, week int, season string) (PlayerStatsJSON, error) {
	res := PlayerStatsJSON{}
	playersStr, err := json.Marshal(playerIds)
	if err != nil {
		return res, err
	}
	op := map[string]interface{}{
		"operationName": "get_player_score_and_projections_batch",
		"variables":     struct{}{},
		"query":         fmt.Sprintf("query get_player_score_and_projections_batch { actual: stats_for_players_in_week(sport: \"nfl\",season: \"%s\",category: \"stat\",season_type: \"regular\",week: %d,player_ids: %s){ game_id opponent player_id stats team week season } projected: stats_for_players_in_week(sport: \"nfl\",season: \"%s\",category: \"proj\",season_type: \"regular\",week: %d,player_ids: %s){ game_id opponent player_id stats team week season } }", season, week, playersStr, season, week, playersStr),
	}
	err = c.sendGraphqlRequest(op, &res)
	return res, err
}
