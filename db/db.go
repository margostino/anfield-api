package db

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/margostino/anfield-api/common"
	"github.com/margostino/anfield-api/fpl"
)

type Player struct {
	ID         string `json:"id"`
	FirstName  string `json:"firstName"`
	SecondName string `json:"secondName"`
	WebName    string `json:"webName"`
	News       string `json:"news"`
	NewsAdded  string `json:"newsAdded"`
	Team       string `json:"team"`
}

type Team struct {
	ID                  string `json:"id"`
	Name                string `json:"name"`
	ShortName           string `json:"shortName"`
	StrengthOverallHome int    `json:"strengthOverallHome"`
	StrengthOverallAway int    `json:"strengthOverallAway"`
	StrengthAttackHome  int    `json:"strengthAttackHome"`
	StrengthAttackAway  int    `json:"strengthAttackAway"`
	StrengthDefenceHome int    `json:"strengthDefenceHome"`
	StrengthDefenceAway int    `json:"strengthDefenceAway"`
}

type Playerx struct {
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
	SquadNumber                      *int      `json:"squad_number"`
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
	CornersAndIndirectFreekicksOrder *int      `json:"corners_and_indirect_freekicks_order"`
	CornersAndIndirectFreekicksText  string    `json:"corners_and_indirect_freekicks_text"`
	DirectFreekicksOrder             *int      `json:"direct_freekicks_order"`
	DirectFreekicksText              string    `json:"direct_freekicks_text"`
	PenaltiesOrder                   *int      `json:"penalties_order"`
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

type Cache struct {
	Teams   map[string]*Team
	Players map[string]*Player
}

type TeamIdx struct {
	Index map[string]*Team
}

var Data, TeamIndex = load()

func load() (*Cache, *TeamIdx) {
	var fplResponse fpl.Response
	fplStaticUrl := os.Getenv("FPL_STATIC_URL")

	response, err := http.Get(fplStaticUrl)
	common.Check(err)
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	common.Check(err)

	err = json.Unmarshal(body, &fplResponse)
	common.Check(err)

	teams := make(map[string]*Team)
	teamIndex := make(map[string]*Team)
	for _, team := range fplResponse.Teams {
		key := strings.ToLower(team.ShortName)
		teamElement := &Team{
			ID:                  strconv.Itoa(team.Code),
			Name:                team.Name,
			ShortName:           team.ShortName,
			StrengthOverallHome: team.StrengthOverallHome,
			StrengthOverallAway: team.StrengthOverallAway,
			StrengthAttackHome:  team.StrengthAttackHome,
			StrengthAttackAway:  team.StrengthAttackAway,
			StrengthDefenceHome: team.StrengthDefenceHome,
			StrengthDefenceAway: team.StrengthDefenceAway,
		}
		teams[key] = teamElement
		teamIndex[strconv.Itoa(team.Code)] = teamElement

	}

	players := make(map[string]*Player)
	for _, player := range fplResponse.Elements {
		key := strings.ToLower(player.WebName)
		players[key] = &Player{
			ID:         strconv.Itoa(player.ID),
			FirstName:  player.FirstName,
			SecondName: player.SecondName,
			WebName:    player.WebName,
			News:       player.News,
			NewsAdded:  player.NewsAdded.Format("2006-01-02"),
			Team:       teamIndex[strconv.Itoa(player.TeamCode)].Name,
		}
	}

	return &Cache{
			Teams:   teams,
			Players: players,
		}, &TeamIdx{
			Index: teamIndex,
		}
}
