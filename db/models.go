package db

type ChipPlay struct {
	ChipName  string `json:"chipName"`
	NumPlayed int    `json:"numPlayed"`
}

type TopElementInfo struct {
	ID     int `json:"id"`
	Points int `json:"points"`
}

type TeamStat struct {
	Value   *int `json:"value,omitempty"`
	Element *int `json:"element,omitempty"`
}

type Stat struct {
	Identifier *string     `json:"identifier,omitempty"`
	TeamA      []*TeamStat `json:"teamA,omitempty"`
	TeamH      []*TeamStat `json:"teamH,omitempty"`
}

type Fixture struct {
	Code                 int     `json:"code"`
	Event                *int    `json:"event,omitempty"`
	Finished             *bool   `json:"finished,omitempty"`
	FinishedProvisional  *bool   `json:"finishedProvisional,omitempty"`
	ID                   int     `json:"id"`
	KickoffTime          *string `json:"kickoffTime,omitempty"`
	Minutes              *int    `json:"minutes,omitempty"`
	ProvisionalStartTime *bool   `json:"provisionalStartTime,omitempty"`
	Started              *bool   `json:"started,omitempty"`
	TeamA                *int    `json:"teamA,omitempty"`
	TeamAName            *string `json:"teamAName,omitempty"`
	TeamHName            *string `json:"teamHName,omitempty"`
	TeamAScore           *int    `json:"teamAScore,omitempty"`
	TeamH                *int    `json:"teamH,omitempty"`
	TeamHScore           *int    `json:"teamHScore,omitempty"`
	Stats                []*Stat `json:"stats,omitempty"`
	TeamHDifficulty      *int    `json:"teamHDifficulty,omitempty"`
	TeamADifficulty      *int    `json:"teamADifficulty,omitempty"`
	PulseID              *int    `json:"pulseId,omitempty"`
}

type Event struct {
	ID                     int             `json:"id"`
	Name                   string          `json:"name"`
	DeadlineTime           string          `json:"deadlineTime"`
	AverageEntryScore      int             `json:"averageEntryScore"`
	Finished               bool            `json:"finished"`
	DataChecked            bool            `json:"dataChecked"`
	HighestScoringEntry    int             `json:"highestScoringEntry"`
	DeadlineTimeEpoch      int             `json:"deadlineTimeEpoch"`
	DeadlineTimeGameOffset int             `json:"deadlineTimeGameOffset"`
	HighestScore           int             `json:"highestScore"`
	IsPrevious             bool            `json:"isPrevious"`
	IsCurrent              bool            `json:"isCurrent"`
	IsNext                 bool            `json:"isNext"`
	CupLeaguesCreated      bool            `json:"cupLeaguesCreated"`
	H2hKoMatchesCreated    bool            `json:"h2hKoMatchesCreated"`
	RankedCount            int             `json:"rankedCount"`
	ChipPlays              []*ChipPlay     `json:"chipPlays"`
	MostSelected           int             `json:"mostSelected"`
	MostTransferredIn      int             `json:"mostTransferredIn"`
	TopElement             int             `json:"topElement"`
	TopElementInfo         *TopElementInfo `json:"topElementInfo"`
	TransfersMade          int             `json:"transfersMade"`
	MostCaptained          int             `json:"mostCaptained"`
	MostViceCaptained      int             `json:"mostViceCaptained"`
	Fixtures               []*Fixture      `json:"fixtures"`
}

type Player struct {
	ID                               int     `json:"id"`
	FirstName                        string  `json:"firstName"`
	SecondName                       string  `json:"secondName"`
	WebName                          string  `json:"webName"`
	News                             string  `json:"news"`
	NewsAdded                        string  `json:"newsAdded"`
	Team                             string  `json:"team"`
	Position                         string  `json:"position"`
	ChanceOfPlayingNextRound         int     `json:"chance_of_playing_next_round"`
	ChanceOfPlayingThisRound         int     `json:"chance_of_playing_this_round"`
	CostChangeEvent                  int     `json:"cost_change_event"`
	CostChangeEventFall              int     `json:"cost_change_event_fall"`
	CostChangeStart                  int     `json:"cost_change_start"`
	CostChangeStartFall              int     `json:"cost_change_start_fall"`
	DreamteamCount                   int     `json:"dreamteam_count"`
	ElementType                      int     `json:"element_type"`
	EpNext                           string  `json:"ep_next"`
	EpThis                           string  `json:"ep_this"`
	EventPoints                      int     `json:"event_points"`
	Form                             string  `json:"form"`
	InDreamteam                      bool    `json:"in_dreamteam"`
	NowCost                          int     `json:"now_cost"`
	Photo                            string  `json:"photo"`
	PointsPerGame                    string  `json:"points_per_game"`
	SelectedByPercent                string  `json:"selected_by_percent"`
	Special                          bool    `json:"special"`
	SquadNumber                      int     `json:"squad_number"`
	Status                           string  `json:"status"`
	TeamCode                         int     `json:"team_code"`
	TotalPoints                      int     `json:"total_points"`
	TransfersIn                      int     `json:"transfers_in"`
	TransfersInEvent                 int     `json:"transfers_in_event"`
	TransfersOut                     int     `json:"transfers_out"`
	TransfersOutEvent                int     `json:"transfers_out_event"`
	ValueForm                        string  `json:"value_form"`
	ValueSeason                      string  `json:"value_season"`
	Minutes                          int     `json:"minutes"`
	GoalsScored                      int     `json:"goals_scored"`
	Assists                          int     `json:"assists"`
	CleanSheets                      int     `json:"clean_sheets"`
	GoalsConceded                    int     `json:"goals_conceded"`
	OwnGoals                         int     `json:"own_goals"`
	PenaltiesSaved                   int     `json:"penalties_saved"`
	PenaltiesMissed                  int     `json:"penalties_missed"`
	YellowCards                      int     `json:"yellow_cards"`
	RedCards                         int     `json:"red_cards"`
	Saves                            int     `json:"saves"`
	Bonus                            int     `json:"bonus"`
	Bps                              int     `json:"bps"`
	Influence                        string  `json:"influence"`
	Creativity                       string  `json:"creativity"`
	Threat                           string  `json:"threat"`
	IctIndex                         string  `json:"ict_index"`
	Starts                           int     `json:"starts"`
	ExpectedGoals                    string  `json:"expected_goals"`
	ExpectedAssists                  string  `json:"expected_assists"`
	ExpectedGoalInvolvements         string  `json:"expected_goal_involvements"`
	ExpectedGoalsConceded            string  `json:"expected_goals_conceded"`
	InfluenceRank                    int     `json:"influence_rank"`
	InfluenceRankType                int     `json:"influence_rank_type"`
	CreativityRank                   int     `json:"creativity_rank"`
	CreativityRankType               int     `json:"creativity_rank_type"`
	ThreatRank                       int     `json:"threat_rank"`
	ThreatRankType                   int     `json:"threat_rank_type"`
	IctIndexRank                     int     `json:"ict_index_rank"`
	IctIndexRankType                 int     `json:"ict_index_rank_type"`
	CornersAndIndirectFreekicksOrder int     `json:"corners_and_indirect_freekicks_order"`
	CornersAndIndirectFreekicksText  string  `json:"corners_and_indirect_freekicks_text"`
	DirectFreekicksOrder             int     `json:"direct_freekicks_order"`
	DirectFreekicksText              string  `json:"direct_freekicks_text"`
	PenaltiesOrder                   int     `json:"penalties_order"`
	PenaltiesText                    string  `json:"penalties_text"`
	ExpectedGoalsPer90               float64 `json:"expected_goals_per_90"`
	SavesPer90                       float64 `json:"saves_per_90"`
	ExpectedAssistsPer90             float64 `json:"expected_assists_per_90"`
	ExpectedGoalInvolvementsPer90    float64 `json:"expected_goal_involvements_per_90"`
	ExpectedGoalsConcededPer90       float64 `json:"expected_goals_conceded_per_90"`
	GoalsConcededPer90               float64 `json:"goals_conceded_per_90"`
	NowCostRank                      int     `json:"now_cost_rank"`
	NowCostRankType                  int     `json:"now_cost_rank_type"`
	FormRank                         int     `json:"form_rank"`
	FormRankType                     int     `json:"form_rank_type"`
	PointsPerGameRank                int     `json:"points_per_game_rank"`
	PointsPerGameRankType            int     `json:"points_per_game_rank_type"`
	SelectedRank                     int     `json:"selected_rank"`
	SelectedRankType                 int     `json:"selected_rank_type"`
	StartsPer90                      float64 `json:"starts_per_90"`
	CleanSheetsPer90                 float64 `json:"clean_sheets_per_90"`
}

type Team struct {
	ID                  int    `json:"id"`
	Name                string `json:"name"`
	ShortName           string `json:"shortName"`
	StrengthOverallHome int    `json:"strengthOverallHome"`
	StrengthOverallAway int    `json:"strengthOverallAway"`
	StrengthAttackHome  int    `json:"strengthAttackHome"`
	StrengthAttackAway  int    `json:"strengthAttackAway"`
	StrengthDefenceHome int    `json:"strengthDefenceHome"`
	StrengthDefenceAway int    `json:"strengthDefenceAway"`
}

type ElementType struct {
	ID                int    `json:"id"`
	PluralName        string `json:"plural_ame"`
	PluralNameShort   string `json:"plural_name_short"`
	SingularName      string `json:"singular_name"`
	SingularNameShort string `json:"singular_name_short"`
}

type Cache struct {
	Teams        map[string]*Team
	Players      map[string]*Player
	ElementTypes map[int]*ElementType
	Events       map[int]*Event
}

type TeamIdx struct {
	Index map[int]*Team
}
