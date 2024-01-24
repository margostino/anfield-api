package fpl

import "time"

type StaticResponse struct {
	Teams        []*Team        `json:"teams"`
	Players      []*Player      `json:"elements"`
	ElementTypes []*ElementType `json:"element_types"`
	Events       []*Event       `json:"events"`
}

type FixturesResponse []*Fixture

type ChipPlay struct {
	ChipName  string `json:"chip_name"`
	NumPlayed int    `json:"num_played"`
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
	TeamA      []*TeamStat `json:"a,omitempty"`
	TeamH      []*TeamStat `json:"h,omitempty"`
}

type Fixture struct {
	Code                 int     `json:"code"`
	Event                *int    `json:"event,omitempty"`
	Finished             *bool   `json:"finished,omitempty"`
	FinishedProvisional  *bool   `json:"finished_provisional,omitempty"`
	ID                   int     `json:"id"`
	KickoffTime          *string `json:"kickoff_time,omitempty"`
	Minutes              *int    `json:"minutes,omitempty"`
	ProvisionalStartTime *bool   `json:"provisional_start_time,omitempty"`
	Started              *bool   `json:"started,omitempty"`
	TeamA                *int    `json:"team_a,omitempty"`
	TeamAScore           *int    `json:"team_a_score,omitempty"`
	TeamH                *int    `json:"team_h,omitempty"`
	TeamHScore           *int    `json:"team_h_score,omitempty"`
	Stats                []*Stat `json:"stats,omitempty"`
	TeamHDifficulty      *int    `json:"team_h_difficulty,omitempty"`
	TeamADifficulty      *int    `json:"team_a_difficulty,omitempty"`
	PulseID              *int    `json:"pulse_id,omitempty"`
}

type Event struct {
	ID                     int             `json:"id"`
	Name                   string          `json:"name"`
	DeadlineTime           time.Time       `json:"deadline_time"`
	AverageEntryScore      int             `json:"average_entry_score"`
	Finished               bool            `json:"finished"`
	DataChecked            bool            `json:"data_checked"`
	HighestScoringEntry    int             `json:"highest_scoring_entry"`
	DeadlineTimeEpoch      int             `json:"deadline_time_epoch"`
	DeadlineTimeGameOffset int             `json:"deadline_time_game_offset"`
	HighestScore           int             `json:"highest_score"`
	IsPrevious             bool            `json:"is_previous"`
	IsCurrent              bool            `json:"is_current"`
	IsNext                 bool            `json:"is_next"`
	CupLeaguesCreated      bool            `json:"cup_leagues_created"`
	H2hKoMatchesCreated    bool            `json:"h2h_ko_matches_created"`
	RankedCount            int             `json:"ranked_count"`
	ChipPlays              []*ChipPlay     `json:"chip_plays"`
	MostSelected           int             `json:"most_selected"`
	MostTransferredIn      int             `json:"most_transferred_in"`
	TopElement             int             `json:"top_element"`
	TopElementInfo         *TopElementInfo `json:"top_element_info"`
	TransfersMade          int             `json:"transfers_made"`
	MostCaptained          int             `json:"most_captained"`
	MostViceCaptained      int             `json:"most_vice_captained"`
	Fixtures               []*Fixture      `json:"fixtures"`
}

type ElementType struct {
	ID                int    `json:"id"`
	PluralName        string `json:"plural_name"`
	PluralNameShort   string `json:"plural_name_short"`
	SingularName      string `json:"singular_name"`
	SingularNameShort string `json:"singular_name_short"`
}

type Team struct {
	Code                int    `json:"code"`
	Name                string `json:"name"`
	ID                  int    `json:"id"`
	ShortName           string `json:"short_name"`
	Strength            int    `json:"strength"`
	StrengthOverallHome int    `json:"strength_overall_home"`
	StrengthOverallAway int    `json:"strength_overall_away"`
	StrengthAttackHome  int    `json:"strength_attack_home"`
	StrengthAttackAway  int    `json:"strength_attack_away"`
	StrengthDefenceHome int    `json:"strength_defence_home"`
	StrengthDefenceAway int    `json:"strength_defence_away"`
	PulseID             int    `json:"pulse_id"`
	Win                 int    `json:"win"`
	Draw                int    `json:"draw"`
	Loss                int    `json:"loss"`
	Played              int    `json:"played"`
	Points              int    `json:"points"`
	Position            int    `json:"position"`
	Unavailable         bool   `json:"unavailable"`
}

type Player struct {
	ChanceOfPlayingNextRound         int       `json:"chance_of_playing_next_round"`
	ChanceOfPlayingThisRound         int       `json:"chance_of_playing_this_round"`
	Code                             int       `json:"code"`
	CostChangeEvent                  int       `json:"cost_change_event"`
	CostChangeEventFall              int       `json:"cost_change_event_fall"`
	CostChangeStart                  int       `json:"cost_change_start"`
	CostChangeStartFall              int       `json:"cost_change_start_fall"`
	DreamteamCount                   int       `json:"dreamteam_count"`
	ElementType                      int       `json:"element_type"`
	EpNext                           string    `json:"ep_next"`
	EpThis                           string    `json:"ep_this"`
	EventPoints                      int       `json:"event_points"`
	FirstName                        string    `json:"first_name"`
	Form                             string    `json:"form"`
	ID                               int       `json:"id"`
	InDreamteam                      bool      `json:"in_dreamteam"`
	News                             string    `json:"news"`
	NewsAdded                        time.Time `json:"news_added"`
	NowCost                          int       `json:"now_cost"`
	Photo                            string    `json:"photo"`
	PointsPerGame                    string    `json:"points_per_game"`
	SecondName                       string    `json:"second_name"`
	SelectedByPercent                string    `json:"selected_by_percent"`
	Special                          bool      `json:"special"`
	SquadNumber                      int       `json:"squad_number"`
	Status                           string    `json:"status"`
	Team                             int       `json:"team"`
	TeamCode                         int       `json:"team_code"`
	TotalPoints                      int       `json:"total_points"`
	TransfersIn                      int       `json:"transfers_in"`
	TransfersInEvent                 int       `json:"transfers_in_event"`
	TransfersOut                     int       `json:"transfers_out"`
	TransfersOutEvent                int       `json:"transfers_out_event"`
	ValueForm                        string    `json:"value_form"`
	ValueSeason                      string    `json:"value_season"`
	WebName                          string    `json:"web_name"`
	Minutes                          int       `json:"minutes"`
	GoalsScored                      int       `json:"goals_scored"`
	Assists                          int       `json:"assists"`
	CleanSheets                      int       `json:"clean_sheets"`
	GoalsConceded                    int       `json:"goals_conceded"`
	OwnGoals                         int       `json:"own_goals"`
	PenaltiesSaved                   int       `json:"penalties_saved"`
	PenaltiesMissed                  int       `json:"penalties_missed"`
	YellowCards                      int       `json:"yellow_cards"`
	RedCards                         int       `json:"red_cards"`
	Saves                            int       `json:"saves"`
	Bonus                            int       `json:"bonus"`
	Bps                              int       `json:"bps"`
	Influence                        string    `json:"influence"`
	Creativity                       string    `json:"creativity"`
	Threat                           string    `json:"threat"`
	IctIndex                         string    `json:"ict_index"`
	Starts                           int       `json:"starts"`
	ExpectedGoals                    string    `json:"expected_goals"`
	ExpectedAssists                  string    `json:"expected_assists"`
	ExpectedGoalInvolvements         string    `json:"expected_goal_involvements"`
	ExpectedGoalsConceded            string    `json:"expected_goals_conceded"`
	InfluenceRank                    int       `json:"influence_rank"`
	InfluenceRankType                int       `json:"influence_rank_type"`
	CreativityRank                   int       `json:"creativity_rank"`
	CreativityRankType               int       `json:"creativity_rank_type"`
	ThreatRank                       int       `json:"threat_rank"`
	ThreatRankType                   int       `json:"threat_rank_type"`
	IctIndexRank                     int       `json:"ict_index_rank"`
	IctIndexRankType                 int       `json:"ict_index_rank_type"`
	CornersAndIndirectFreekicksOrder int       `json:"corners_and_indirect_freekicks_order"`
	CornersAndIndirectFreekicksText  string    `json:"corners_and_indirect_freekicks_text"`
	DirectFreekicksOrder             int       `json:"direct_freekicks_order"`
	DirectFreekicksText              string    `json:"direct_freekicks_text"`
	PenaltiesOrder                   int       `json:"penalties_order"`
	PenaltiesText                    string    `json:"penalties_text"`
	ExpectedGoalsPer90               float64   `json:"expected_goals_per_90"`
	SavesPer90                       float64   `json:"saves_per_90"`
	ExpectedAssistsPer90             float64   `json:"expected_assists_per_90"`
	ExpectedGoalInvolvementsPer90    float64   `json:"expected_goal_involvements_per_90"`
	ExpectedGoalsConcededPer90       float64   `json:"expected_goals_conceded_per_90"`
	GoalsConcededPer90               float64   `json:"goals_conceded_per_90"`
	NowCostRank                      int       `json:"now_cost_rank"`
	NowCostRankType                  int       `json:"now_cost_rank_type"`
	FormRank                         int       `json:"form_rank"`
	FormRankType                     int       `json:"form_rank_type"`
	PointsPerGameRank                int       `json:"points_per_game_rank"`
	PointsPerGameRankType            int       `json:"points_per_game_rank_type"`
	SelectedRank                     int       `json:"selected_rank"`
	SelectedRankType                 int       `json:"selected_rank_type"`
	StartsPer90                      float64   `json:"starts_per_90"`
	CleanSheetsPer90                 float64   `json:"clean_sheets_per_90"`
}
