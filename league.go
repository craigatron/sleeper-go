package sleeper

import (
	"math"
	"strconv"
	"strings"
)

// League is a Sleeper league.
type League struct {
	Client Client

	ID    string
	Token string

	Season string

	LeagueInfo LeagueInfoJSON
	Rosters    map[int]RosterJSON
	Users      map[string]UserJSON
	Matchups   []MatchupsJSON
}

// NewLeague creates a new Sleeper league with the given ID and token (for graphQL functionality).
func NewLeague(leagueID string, token string) (League, error) {
	l := League{
		ID:    leagueID,
		Token: token,
	}
	c, err := NewClientWithToken(token)
	if err != nil {
		return l, err
	}

	l.Client = c

	status, err := c.GetNflStatus()
	if err != nil {
		return l, err
	}
	l.Season = status.Season

	li, err := c.GetLeagueInfo(leagueID)
	if err != nil {
		return l, err
	}
	l.LeagueInfo = li

	rosters, err := c.GetLeagueRosters(leagueID)
	if err != nil {
		return l, err
	}
	rosterMap := make(map[int]RosterJSON)
	for _, r := range rosters {
		rosterMap[r.RosterID] = r
	}
	l.Rosters = rosterMap

	users, err := c.GetLeagueUsers(leagueID)
	if err != nil {
		return l, err
	}
	userMap := make(map[string]UserJSON)
	for _, u := range users {
		userMap[u.UserID] = u
	}
	l.Users = userMap

	matchupSlice := make([]MatchupsJSON, 0)
	// TODO: need to handle playoffs in a smart way
	for i := 1; i < li.Settings.PlayoffWeekStart; i++ {
		weekMatchups, err := c.GetLeagueMatchups(leagueID, i)
		if err != nil {
			return l, err
		}
		matchupSlice = append(matchupSlice, weekMatchups)
	}
	l.Matchups = matchupSlice

	return l, nil
}

// MatchupProjection is the projected score for a particular matchup.
type MatchupProjection struct {
	Matchup    *MatchupJSON
	Projection float64
}

// GetProjections returns the matchup projections for the current week.
func (l League) GetProjections() ([]MatchupProjection, error) {
	state, err := l.Client.GetNflStatus()
	if err != nil {
		return nil, err
	}

	currentWeek := state.Week

	matchups, err := l.Client.GetLeagueMatchups(l.ID, currentWeek)
	if err != nil {
		return nil, err
	}

	// build the list of every active player this week
	allActivePlayers := make([]string, 0)
	for _, m := range matchups {
		allActivePlayers = append(allActivePlayers, m.Starters...)
	}

	playerStats, err := l.Client.GetPlayerStats(allActivePlayers, currentWeek, l.Season)
	if err != nil {
		return nil, err
	}

	projectedStatsByPlayer := make(map[string]StatsJSON)
	actualStatsByPlayer := make(map[string]StatsJSON)
	for _, as := range playerStats.Data.Actual {
		actualStatsByPlayer[as.PlayerID] = as
	}
	for _, ps := range playerStats.Data.Projected {
		projectedStatsByPlayer[ps.PlayerID] = ps
	}

	gamesByID, err := l.getGamesByID(currentWeek)
	if err != nil {
		return nil, err
	}

	projections := make([]MatchupProjection, 0)
	for _, matchup := range matchups {
		projection := 0.0
		for _, startingPlayer := range matchup.Starters {
			if projectedStats, ok := projectedStatsByPlayer[startingPlayer]; ok {
				gameInfo := gamesByID[projectedStats.GameID]
				projection += l.calculatePlayerProjection(actualStatsByPlayer[startingPlayer], projectedStats, gameInfo)
			}
		}
		projections = append(projections, MatchupProjection{
			Matchup:    &matchup,
			Projection: projection,
		})
	}

	return projections, nil
}

type gameMetadata struct {
	status      string
	secondsLeft int
}

func (l League) getGamesByID(week int) (map[string]gameMetadata, error) {
	batchScores, err := l.Client.GetBatchScores(week, l.Season)
	if err != nil {
		return nil, err
	}
	gamesByID := make(map[string]gameMetadata)

	for _, game := range batchScores.Data.Scores {
		var secondsLeft int
		if game.Status == "complete" {
			secondsLeft = 0
		} else if game.Status == "pre_game" {
			secondsLeft = 3600
		} else {
			quarter := game.Metadata.QuarterNum
			quartersLeft := 4 - quarter
			// OT is represented as quarter 5. also there really isn't a math.max for integers in go???
			if quartersLeft < 0 {
				quartersLeft = 0
			}
			quarterTimeRemaining := strings.Split(game.Metadata.TimeRemaining, ":")
			quarterMinsRemaining, err := strconv.Atoi(quarterTimeRemaining[0])
			if err != nil {
				return nil, err
			}
			quarterSecsRemaining, err := strconv.Atoi(quarterTimeRemaining[1])
			if err != nil {
				return nil, err
			}
			secondsLeft = (quartersLeft * 15 * 60) + (quarterMinsRemaining * 60) + quarterSecsRemaining
		}
		gamesByID[game.GameID] = gameMetadata{
			status:      game.Status,
			secondsLeft: secondsLeft,
		}
	}
	return gamesByID, nil
}

func (l League) calculatePlayerProjection(actualStats StatsJSON, projectedStats StatsJSON, gameInfo gameMetadata) float64 {
	originalProjection := l.scoreStats(projectedStats.Stats)
	if gameInfo.status == "pre_game" {
		return originalProjection
	}
	var currentScore float64
	if actualStats.Stats == nil {
		currentScore = 0.0
	} else {
		currentScore = l.scoreStats(actualStats.Stats)
	}
	if gameInfo.status == "complete" {
		return currentScore
	}

	// else the game is in progress, let's calculate the projections on the fly
	fractionalGameLeft := float64(gameInfo.secondsLeft) / 3600.0
	minutesRemaining := float64(gameInfo.secondsLeft) / 60.0
	// TODO: gotta see how overtime is handled
	minutesPlayed := 60.0 - minutesRemaining

	// this is all from the sleeper client
	// TODO: think this is only the chunk that works for IDP, DEF roles had something else
	s := currentScore + (currentScore / math.Max(minutesPlayed, 1.0) * minutesRemaining * (minutesRemaining / 60.0))
	c := 0.2 * fractionalGameLeft * s
	i := (.35 + .65*(1-fractionalGameLeft)) * s
	u := .45 * fractionalGameLeft * s
	d := math.Max(c+i+u, currentScore)
	f := math.Max(originalProjection, currentScore)
	return f + (1-fractionalGameLeft)*(d-f)
}

func (l League) scoreStats(stats map[string]float64) float64 {
	total := 0.0
	for statKey, statValue := range stats {
		if statMultiplier, ok := l.LeagueInfo.ScoringSettings[statKey]; ok {
			total += statMultiplier * statValue
		}
	}
	return total
}
