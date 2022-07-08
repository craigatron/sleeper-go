package sleeper

// AllPlayersJSON is the return type for the all players endpoint.
type AllPlayersJSON map[string]PlayerInfoJSON

// PlayerInfoJSON is information about a single player in the all players endpoint.
type PlayerInfoJSON struct {
	YahooID               int         `json:"yahoo_id"`
	SportradarID          string      `json:"sportradar_id"`
	PlayerID              string      `json:"player_id"`
	YearsExp              int         `json:"years_exp"`
	SwishID               interface{} `json:"swish_id"`
	RotoworldID           interface{} `json:"rotoworld_id"`
	FirstName             string      `json:"first_name"`
	Hashtag               string      `json:"hashtag"`
	Sport                 string      `json:"sport"`
	Status                string      `json:"status"`
	PandascoreID          interface{} `json:"pandascore_id"`
	BirthDate             string      `json:"birth_date"`
	BirthState            interface{} `json:"birth_state"`
	Position              string      `json:"position"`
	Metadata              interface{} `json:"metadata"`
	EspnID                int         `json:"espn_id"`
	SearchFullName        string      `json:"search_full_name"`
	DepthChartOrder       interface{} `json:"depth_chart_order"`
	SearchLastName        string      `json:"search_last_name"`
	Weight                string      `json:"weight"`
	LastName              string      `json:"last_name"`
	College               string      `json:"college"`
	Age                   int         `json:"age"`
	PracticeDescription   interface{} `json:"practice_description"`
	FantasyPositions      []string    `json:"fantasy_positions"`
	DepthChartPosition    interface{} `json:"depth_chart_position"`
	InjuryStartDate       interface{} `json:"injury_start_date"`
	Team                  interface{} `json:"team"`
	InjuryStatus          interface{} `json:"injury_status"`
	FullName              string      `json:"full_name"`
	FantasyDataID         int         `json:"fantasy_data_id"`
	BirthCountry          interface{} `json:"birth_country"`
	SearchFirstName       string      `json:"search_first_name"`
	GsisID                interface{} `json:"gsis_id"`
	StatsID               interface{} `json:"stats_id"`
	NewsUpdated           interface{} `json:"news_updated"`
	RotowireID            int         `json:"rotowire_id"`
	HighSchool            interface{} `json:"high_school"`
	Height                string      `json:"height"`
	InjuryNotes           interface{} `json:"injury_notes"`
	Number                int         `json:"number"`
	Active                bool        `json:"active"`
	InjuryBodyPart        interface{} `json:"injury_body_part"`
	SearchRank            int         `json:"search_rank"`
	BirthCity             interface{} `json:"birth_city"`
	PracticeParticipation interface{} `json:"practice_participation"`
}

// TrendingPlayersJSON is the return type of the trending players API.
type TrendingPlayersJSON []struct {
	PlayerID string `json:"player_id"`
	Count    int    `json:"count"`
}

// SportStatusJSON is the return type of the NFL status API.
type SportStatusJSON struct {
	Week               int    `json:"week"`
	SeasonType         string `json:"season_type"`
	SeasonStartDate    string `json:"season_start_date"`
	Season             string `json:"season"`
	PreviousSeason     string `json:"previous_season"`
	Leg                int    `json:"leg"`
	LeagueSeason       string `json:"league_season"`
	LeagueCreateSeason string `json:"league_create_season"`
	DisplayWeek        int    `json:"display_week"`
}
