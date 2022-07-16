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

// LeagueInfoJSON is the return type of the league info API.
type LeagueInfoJSON struct {
	TotalRosters int    `json:"total_rosters"`
	Status       string `json:"status"`
	Sport        string `json:"sport"`
	Shard        int    `json:"shard"`
	Settings     struct {
		MaxKeepers           int `json:"max_keepers"`
		DraftRounds          int `json:"draft_rounds"`
		TradeReviewDays      int `json:"trade_review_days"`
		Squads               int `json:"squads"`
		ReserveAllowDnr      int `json:"reserve_allow_dnr"`
		CapacityOverride     int `json:"capacity_override"`
		PickTrading          int `json:"pick_trading"`
		DisableTrades        int `json:"disable_trades"`
		TaxiYears            int `json:"taxi_years"`
		TaxiAllowVets        int `json:"taxi_allow_vets"`
		BestBall             int `json:"best_ball"`
		LastReport           int `json:"last_report"`
		DisableAdds          int `json:"disable_adds"`
		WaiverType           int `json:"waiver_type"`
		BenchLock            int `json:"bench_lock"`
		ReserveAllowSus      int `json:"reserve_allow_sus"`
		Type                 int `json:"type"`
		ReserveAllowCov      int `json:"reserve_allow_cov"`
		WaiverClearDays      int `json:"waiver_clear_days"`
		DailyWaiversLastRan  int `json:"daily_waivers_last_ran"`
		WaiverDayOfWeek      int `json:"waiver_day_of_week"`
		StartWeek            int `json:"start_week"`
		PlayoffTeams         int `json:"playoff_teams"`
		NumTeams             int `json:"num_teams"`
		ReserveSlots         int `json:"reserve_slots"`
		PlayoffRoundType     int `json:"playoff_round_type"`
		DailyWaiversHour     int `json:"daily_waivers_hour"`
		WaiverBudget         int `json:"waiver_budget"`
		ReserveAllowOut      int `json:"reserve_allow_out"`
		OffseasonAdds        int `json:"offseason_adds"`
		PlayoffSeedType      int `json:"playoff_seed_type"`
		DailyWaivers         int `json:"daily_waivers"`
		PlayoffWeekStart     int `json:"playoff_week_start"`
		DailyWaiversDays     int `json:"daily_waivers_days"`
		LeagueAverageMatch   int `json:"league_average_match"`
		Leg                  int `json:"leg"`
		TradeDeadline        int `json:"trade_deadline"`
		ReserveAllowDoubtful int `json:"reserve_allow_doubtful"`
		TaxiDeadline         int `json:"taxi_deadline"`
		ReserveAllowNa       int `json:"reserve_allow_na"`
		TaxiSlots            int `json:"taxi_slots"`
		PlayoffType          int `json:"playoff_type"`
	} `json:"settings"`
	SeasonType       string             `json:"season_type"`
	Season           string             `json:"season"`
	ScoringSettings  map[string]float64 `json:"scoring_settings"`
	RosterPositions  []string           `json:"roster_positions"`
	PreviousLeagueID string             `json:"previous_league_id"`
	Name             string             `json:"name"`
	Metadata         struct {
		LatestLeagueWinnerRosterID string `json:"latest_league_winner_roster_id"`
	} `json:"metadata"`
	LoserBracketID        interface{} `json:"loser_bracket_id"`
	LeagueID              string      `json:"league_id"`
	LastReadID            interface{} `json:"last_read_id"`
	LastPinnedMessageID   interface{} `json:"last_pinned_message_id"`
	LastMessageTime       int64       `json:"last_message_time"`
	LastMessageTextMap    interface{} `json:"last_message_text_map"`
	LastMessageID         string      `json:"last_message_id"`
	LastMessageAttachment interface{} `json:"last_message_attachment"`
	LastAuthorIsBot       bool        `json:"last_author_is_bot"`
	LastAuthorID          string      `json:"last_author_id"`
	LastAuthorDisplayName string      `json:"last_author_display_name"`
	LastAuthorAvatar      interface{} `json:"last_author_avatar"`
	GroupID               interface{} `json:"group_id"`
	DraftID               string      `json:"draft_id"`
	CompanyID             interface{} `json:"company_id"`
	BracketID             interface{} `json:"bracket_id"`
	Avatar                interface{} `json:"avatar"`
}

// RostersJSON is the return type of the league rosters API.
type RostersJSON []RosterJSON

// RosterJSON is a single roster from the league rosters API.
type RosterJSON struct {
	Taxi     []string `json:"taxi"`
	Starters []string `json:"starters"`
	Settings struct {
		Wins             int `json:"wins"`
		WaiverPosition   int `json:"waiver_position"`
		WaiverBudgetUsed int `json:"waiver_budget_used"`
		TotalMoves       int `json:"total_moves"`
		Ties             int `json:"ties"`
		Losses           int `json:"losses"`
		Fpts             int `json:"fpts"`
	} `json:"settings"`
	RosterID  int         `json:"roster_id"`
	Reserve   []string    `json:"reserve"`
	Players   []string    `json:"players"`
	PlayerMap interface{} `json:"player_map"`
	OwnerID   string      `json:"owner_id"`
	LeagueID  string      `json:"league_id"`
	CoOwners  interface{} `json:"co_owners"`
	// known keys:
	// streak, record, p_nick_<player ID>, allow_pn_scoring, allow_pn_news
	Metadata interface{} `json:"metadata"`
}

// UsersJSON is the return type of the league users API.
type UsersJSON []UserJSON

// UserJSON is a single user from the league users API.
type UserJSON struct {
	UserID   string      `json:"user_id"`
	Settings interface{} `json:"settings"`
	Metadata struct {
		TeamName                string `json:"team_name"`
		TeamNameUpdate          string `json:"team_name_update"`
		MentionPn               string `json:"mention_pn"`
		AllowPn                 string `json:"allow_pn"`
		TradeBlockPn            string `json:"trade_block_pn"`
		JoinVoicePn             string `json:"join_voice_pn"`
		UserMessagePn           string `json:"user_message_pn"`
		TransactionCommissioner string `json:"transaction_commissioner"`
		Archived                string `json:"archived"`
		TransactionFreeAgent    string `json:"transaction_free_agent"`
		TransactionTrade        string `json:"transaction_trade"`
		TransactionWaiver       string `json:"transaction_waiver"`
		PlayerLikePn            string `json:"player_like_pn"`
		MascotMessage           string `json:"mascot_message"`
		PlayerNicknameUpdate    string `json:"player_nickname_update"`

		// tf are these
		MascotItemTypeIDLeg1      string `json:"mascot_item_type_id_leg_1"`
		MascotItemTypeIDLeg2      string `json:"mascot_item_type_id_leg_2"`
		MascotItemTypeIDLeg3      string `json:"mascot_item_type_id_leg_3"`
		MascotItemTypeIDLeg4      string `json:"mascot_item_type_id_leg_4"`
		MascotItemTypeIDLeg5      string `json:"mascot_item_type_id_leg_5"`
		MascotItemTypeIDLeg6      string `json:"mascot_item_type_id_leg_6"`
		MascotItemTypeIDLeg7      string `json:"mascot_item_type_id_leg_7"`
		MascotItemTypeIDLeg8      string `json:"mascot_item_type_id_leg_8"`
		MascotItemTypeIDLeg9      string `json:"mascot_item_type_id_leg_9"`
		MascotItemTypeIDLeg10     string `json:"mascot_item_type_id_leg_10"`
		MascotItemTypeIDLeg11     string `json:"mascot_item_type_id_leg_11"`
		MascotItemTypeIDLeg12     string `json:"mascot_item_type_id_leg_12"`
		MascotItemTypeIDLeg13     string `json:"mascot_item_type_id_leg_13"`
		MascotItemTypeIDLeg14     string `json:"mascot_item_type_id_leg_14"`
		MascotItemTypeIDLeg15     string `json:"mascot_item_type_id_leg_15"`
		MascotItemTypeIDLeg16     string `json:"mascot_item_type_id_leg_16"`
		MascotItemTypeIDLeg17     string `json:"mascot_item_type_id_leg_17"`
		MascotMessageEmotionLeg1  string `json:"mascot_message_emotion_leg_1"`
		MascotMessageEmotionLeg2  string `json:"mascot_message_emotion_leg_2"`
		MascotMessageEmotionLeg3  string `json:"mascot_message_emotion_leg_3"`
		MascotMessageEmotionLeg4  string `json:"mascot_message_emotion_leg_4"`
		MascotMessageEmotionLeg5  string `json:"mascot_message_emotion_leg_5"`
		MascotMessageEmotionLeg6  string `json:"mascot_message_emotion_leg_6"`
		MascotMessageEmotionLeg7  string `json:"mascot_message_emotion_leg_7"`
		MascotMessageEmotionLeg8  string `json:"mascot_message_emotion_leg_8"`
		MascotMessageEmotionLeg9  string `json:"mascot_message_emotion_leg_9"`
		MascotMessageEmotionLeg10 string `json:"mascot_message_emotion_leg_10"`
		MascotMessageEmotionLeg11 string `json:"mascot_message_emotion_leg_11"`
		MascotMessageEmotionLeg12 string `json:"mascot_message_emotion_leg_12"`
	} `json:"metadata"`
	LeagueID    string `json:"league_id"`
	IsOwner     bool   `json:"is_owner"`
	IsBot       bool   `json:"is_bot"`
	DisplayName string `json:"display_name"`
	Avatar      string `json:"avatar"`
}

// MatchupsJSON is the return type of the league matchups API.
type MatchupsJSON []MatchupJSON

// MatchupJSON is a single matchup from the league matchups API.
// TODO: need to revisit once the season is running
type MatchupJSON struct {
	Starters     []string `json:"starters"`
	RosterID     int      `json:"roster_id"`
	Players      []string `json:"players"`
	MatchupID    int      `json:"matchup_id"`
	Points       float32  `json:"points"`
	CustomPoints float32  `json:"custom_points"`
}
