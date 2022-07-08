package sleeper

const sleeperBaseURL = "https://api.sleeper.app/v1"

// TrendingPlayerType is either add/drop for the GetTrendingPlayers API.
type TrendingPlayerType string

const (
	// TrendingPlayerTypeAdd is "add"
	TrendingPlayerTypeAdd TrendingPlayerType = "add"
	// TrendingPlayerTypeDrop is "drop"
	TrendingPlayerTypeDrop = "drop"
)
