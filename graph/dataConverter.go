package graph

import (
	"fmt"
	"os"
	"strconv"

	"github.com/margostino/anfield-api/common"
	"github.com/margostino/anfield-api/db"
	"github.com/margostino/anfield-api/graph/model"
	"github.com/margostino/anfield-api/http"
	"github.com/margostino/anfield-api/pulse"
)

func toTeamGraph(team *db.Team) *model.Team {
	return &model.Team{
		ID:                  team.ID,
		Name:                team.Name,
		ShortName:           team.ShortName,
		StrengthOverallHome: team.StrengthOverallHome,
		StrengthOverallAway: team.StrengthOverallAway,
		StrengthAttackHome:  team.StrengthAttackHome,
		StrengthAttackAway:  team.StrengthAttackAway,
		StrengthDefenceHome: team.StrengthDefenceHome,
		StrengthDefenceAway: team.StrengthDefenceAway,
	}
}

func toPlayerGraph(player *db.Player) *model.Player {
	return &model.Player{
		ID:                               player.ID,
		FirstName:                        player.FirstName,
		SecondName:                       player.SecondName,
		WebName:                          player.WebName,
		News:                             player.News,
		NewsAdded:                        player.NewsAdded,
		Team:                             player.Team,
		Position:                         model.PlayerPosition(player.Position),
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

func toEventGraph(event db.Event) *model.Event {
	chipPlays := make([]*model.ChipPlay, 0)
	for _, chipPlay := range event.ChipPlays {
		if chipPlay != nil {
			chipPlays = append(chipPlays, &model.ChipPlay{
				ChipName:  chipPlay.ChipName,
				NumPlayed: chipPlay.NumPlayed,
			})
		}
	}

	fixtures := make([]*model.Fixture, 0)
	for _, fixture := range event.Fixtures {
		if fixture != nil {
			stats := make([]*model.Stat, 0)
			for _, stat := range fixture.Stats {
				if stat != nil {
					statsTeamA := make([]*model.TeamStat, 0)
					for _, teamA := range stat.TeamA {
						if teamA != nil {
							statTeamA := model.TeamStat{
								Value:   teamA.Value,
								Element: teamA.Element,
							}
							statsTeamA = append(statsTeamA, &statTeamA)
						}
					}
					statsTeamH := make([]*model.TeamStat, 0)
					for _, teamH := range stat.TeamH {
						if teamH != nil {
							statTeamH := model.TeamStat{
								Value:   teamH.Value,
								Element: teamH.Element,
							}
							statsTeamH = append(statsTeamH, &statTeamH)
						}
					}

					stats = append(stats, &model.Stat{
						Identifier: stat.Identifier,
						TeamA:      statsTeamA,
						TeamH:      statsTeamH,
					})
				}
			}

			fixtures = append(fixtures, &model.Fixture{
				ID:                   fixture.ID,
				Code:                 fixture.Code,
				TeamH:                fixture.TeamH,
				TeamA:                fixture.TeamA,
				TeamAName:            fixture.TeamAName,
				TeamHName:            fixture.TeamHName,
				TeamHScore:           fixture.TeamHScore,
				TeamAScore:           fixture.TeamAScore,
				Event:                fixture.Event,
				Finished:             fixture.Finished,
				Minutes:              fixture.Minutes,
				ProvisionalStartTime: fixture.ProvisionalStartTime,
				KickoffTime:          fixture.KickoffTime,
				Stats:                stats,
				PulseID:              fixture.PulseID,
				TeamHDifficulty:      fixture.TeamHDifficulty,
				TeamADifficulty:      fixture.TeamADifficulty,
			})
		}
	}
	return &model.Event{
		ID:                     event.ID,
		Name:                   event.Name,
		DeadlineTime:           event.DeadlineTime,
		AverageEntryScore:      event.AverageEntryScore,
		Finished:               event.Finished,
		DataChecked:            event.DataChecked,
		HighestScoringEntry:    event.HighestScoringEntry,
		DeadlineTimeEpoch:      event.DeadlineTimeEpoch,
		DeadlineTimeGameOffset: event.DeadlineTimeGameOffset,
		HighestScore:           event.HighestScore,
		IsPrevious:             event.IsPrevious,
		IsCurrent:              event.IsCurrent,
		IsNext:                 event.IsNext,
		ChipPlays:              chipPlays,
		Fixtures:               fixtures,
	}
}

func toH2HGraph(teamHData *db.Team, teamAData *db.Team) *model.H2h {
	pulseH2HUrl := os.Getenv("PULSE_H2H_URL")
	pulseHeaders := map[string]string{
		"Origin": "https://www.premierleague.com",
	}
	var pulseH2HResponse pulse.H2HResponse

	pulseH2HUrlWithParams := fmt.Sprintf("%s?teams=%d,%d&altIds=true&comps=1", pulseH2HUrl, teamHData.PulseID, teamAData.PulseID)
	http.Get(pulseH2HUrlWithParams, &pulseH2HResponse, nil, pulseHeaders)

	h2h := &model.H2h{
		StatsTeamA: make([]*model.StatsTeam, 0),
		StatsTeamH: make([]*model.StatsTeam, 0),
		Gameweeks:  make([]*model.Gameweek, 0),
	}

	for id, stats := range pulseH2HResponse.Stats {
		teamStats := make([]*model.StatsTeam, 0)
		for _, statsElement := range stats {
			teamStatsElement := &model.StatsTeam{
				Name:        statsElement.Name,
				Value:       statsElement.Value,
				Description: statsElement.Description,
			}
			teamStats = append(teamStats, teamStatsElement)
		}

		// convert string id to int
		idNumber, err := strconv.Atoi(id)
		common.Check(err)

		if idNumber == *&teamHData.PulseID {
			h2h.StatsTeamH = teamStats
		} else {
			h2h.StatsTeamA = teamStats
		}
	}

	for _, gameweek := range pulseH2HResponse.HeadToHeads {
		var teamAName, teamHName string
		var teamAScore, teamHScore int

		if gameweek.Teams[0].Team.Id == float32(teamHData.PulseID) {
			teamHName = gameweek.Teams[0].Team.Name
			teamAName = gameweek.Teams[1].Team.Name
			teamHScore = int(gameweek.Teams[0].Score)
			teamAScore = int(gameweek.Teams[1].Score)
		} else {
			teamHName = gameweek.Teams[1].Team.Name
			teamAName = gameweek.Teams[0].Team.Name
			teamHScore = int(gameweek.Teams[1].Score)
			teamAScore = int(gameweek.Teams[0].Score)
		}
		h2h.Gameweeks = append(h2h.Gameweeks, &model.Gameweek{
			Kickoff:    gameweek.Kickoff.Label,
			Stadium:    gameweek.Ground.Name,
			TeamAName:  teamAName,
			TeamHName:  teamHName,
			ScoreTeamA: teamAScore,
			ScoreTeamH: teamHScore,
		})
	}
	return h2h
}
