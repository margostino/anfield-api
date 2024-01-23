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

	elementTypes := make(map[int]*ElementType)
	for _, elementType := range fplResponse.ElementTypes {
		elementTypes[elementType.ID] = &ElementType{
			ID:                elementType.ID,
			PluralName:        elementType.PluralName,
			PluralNameShort:   elementType.PluralNameShort,
			SingularName:      elementType.SingularName,
			SingularNameShort: elementType.SingularNameShort,
		}
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
			Position:                         elementTypes[player.ElementType].SingularNameShort,
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
			Teams:        teams,
			Players:      players,
			ElementTypes: elementTypes,
		}, &TeamIdx{
			Index: teamIndex,
		}
}
