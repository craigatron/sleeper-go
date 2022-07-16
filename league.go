package sleeper

// League is a Sleeper league.
type League struct {
	Client Client

	ID    string
	Token string

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
	c, err := NewClient()
	if err != nil {
		return l, err
	}

	l.Client = c

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
