package db

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/margostino/anfield-api/common"
	"github.com/margostino/anfield-api/fpl"
)

type Player struct {
	ID                               string  `json:"id"`
	FirstName                        string  `json:"firstName"`
	SecondName                       string  `json:"secondName"`
	WebName                          string  `json:"webName"`
	News                             string  `json:"news"`
	NewsAdded                        string  `json:"newsAdded"`
	Team                             string  `json:"team"`
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
			ID:                               strconv.Itoa(player.ID),
			FirstName:                        player.FirstName,
			SecondName:                       player.SecondName,
			WebName:                          player.WebName,
			News:                             player.News,
			NewsAdded:                        player.NewsAdded.Format("2006-01-02"),
			Team:                             teamIndex[strconv.Itoa(player.TeamCode)].Name,
			ChanceOfPlayingNextRound:         player.ChanceOfPlayingNextRound,
			ChanceOfPlayingThisRound:         player.ChanceOfPlayingThisRound,
			CostChangeEvent:                  player.CostChangeEvent,
			CostChangeEventFall:              player.CostChangeEventFall,
			CostChangeStart:                  player.CostChangeStart,
			CostChangeStartFall:              player.CostChangeStartFall,
			DreamteamCount:                   player.DreamteamCount,
			ElementType:                      player.ElementType,
			EpNext:                           player.EpNext,
			EpThis:                           player.EpThis,
			EventPoints:                      player.EventPoints,
			Form:                             player.Form,
			InDreamteam:                      player.InDreamteam,
			NowCost:                          player.NowCost,
			Photo:                            player.Photo,
			PointsPerGame:                    player.PointsPerGame,
			SelectedByPercent:                player.SelectedByPercent,
			Special:                          player.Special,
			SquadNumber:                      player.SquadNumber,
			Status:                           player.Status,
			TeamCode:                         player.TeamCode,
			TotalPoints:                      player.TotalPoints,
			TransfersIn:                      player.TransfersIn,
			TransfersInEvent:                 player.TransfersInEvent,
			TransfersOut:                     player.TransfersOut,
			TransfersOutEvent:                player.TransfersOutEvent,
			ValueForm:                        player.ValueForm,
			ValueSeason:                      player.ValueSeason,
			Minutes:                          player.Minutes,
			GoalsScored:                      player.GoalsScored,
			Assists:                          player.Assists,
			CleanSheets:                      player.CleanSheets,
			GoalsConceded:                    player.GoalsConceded,
			OwnGoals:                         player.OwnGoals,
			PenaltiesSaved:                   player.PenaltiesSaved,
			PenaltiesMissed:                  player.PenaltiesMissed,
			YellowCards:                      player.YellowCards,
			RedCards:                         player.RedCards,
			Saves:                            player.Saves,
			Bonus:                            player.Bonus,
			Bps:                              player.Bps,
			Influence:                        player.Influence,
			Creativity:                       player.Creativity,
			Threat:                           player.Threat,
			IctIndex:                         player.IctIndex,
			Starts:                           player.Starts,
			ExpectedGoals:                    player.ExpectedGoals,
			ExpectedAssists:                  player.ExpectedAssists,
			ExpectedGoalInvolvements:         player.ExpectedGoalInvolvements,
			ExpectedGoalsConceded:            player.ExpectedGoalsConceded,
			InfluenceRank:                    player.InfluenceRank,
			InfluenceRankType:                player.InfluenceRankType,
			CreativityRank:                   player.CreativityRank,
			CreativityRankType:               player.CreativityRankType,
			ThreatRank:                       player.ThreatRank,
			ThreatRankType:                   player.ThreatRankType,
			IctIndexRank:                     player.IctIndexRank,
			IctIndexRankType:                 player.IctIndexRankType,
			CornersAndIndirectFreekicksOrder: player.CornersAndIndirectFreekicksOrder,
			CornersAndIndirectFreekicksText:  player.CornersAndIndirectFreekicksText,
			DirectFreekicksOrder:             player.DirectFreekicksOrder,
			DirectFreekicksText:              player.DirectFreekicksText,
			PenaltiesOrder:                   player.PenaltiesOrder,
			PenaltiesText:                    player.PenaltiesText,
			ExpectedGoalsPer90:               player.ExpectedGoalsPer90,
			SavesPer90:                       player.SavesPer90,
			ExpectedAssistsPer90:             player.ExpectedAssistsPer90,
			ExpectedGoalInvolvementsPer90:    player.ExpectedGoalInvolvementsPer90,
			ExpectedGoalsConcededPer90:       player.ExpectedGoalsConcededPer90,
			GoalsConcededPer90:               player.GoalsConcededPer90,
			NowCostRank:                      player.NowCostRank,
			NowCostRankType:                  player.NowCostRankType,
			FormRank:                         player.FormRank,
			FormRankType:                     player.FormRankType,
			PointsPerGameRank:                player.PointsPerGameRank,
			PointsPerGameRankType:            player.PointsPerGameRankType,
			SelectedRank:                     player.SelectedRank,
			SelectedRankType:                 player.SelectedRankType,
			StartsPer90:                      player.StartsPer90,
			CleanSheetsPer90:                 player.CleanSheetsPer90,
		}
	}

	return &Cache{
			Teams:   teams,
			Players: players,
		}, &TeamIdx{
			Index: teamIndex,
		}
}
