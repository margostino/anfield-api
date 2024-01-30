package db

import (
	"github.com/margostino/anfield-api/fpl"
)

func loadFixtures(fplFixturesResponse fpl.FixturesResponse, teamIndex map[int]*Team) map[int][]*Fixture {
	fixtures := make(map[int][]*Fixture, len(fplFixturesResponse))
	for _, fixture := range fplFixturesResponse {
		stats := make([]*Stat, 0)
		for _, stat := range fixture.Stats {
			teamA := make([]*TeamStat, len(stat.TeamA))
			for _, teamAStat := range stat.TeamA {
				teamA = append(teamA, &TeamStat{
					Value:   teamAStat.Value,
					Element: teamAStat.Element,
				})
			}
			teamH := make([]*TeamStat, len(stat.TeamH))
			for _, teamHStat := range stat.TeamH {
				teamH = append(teamH, &TeamStat{
					Value:   teamHStat.Value,
					Element: teamHStat.Element,
				})
			}
			stats = append(stats, &Stat{
				Identifier: stat.Identifier,
				TeamA:      teamA,
				TeamH:      teamH,
			})
		}

		if fixture.Event != nil {
			if fixtures[*fixture.Event] == nil {
				fixtures[*fixture.Event] = make([]*Fixture, 0)
			}
			fixtures[*fixture.Event] = append(fixtures[*fixture.Event], &Fixture{
				Code:                 fixture.Code,
				Event:                fixture.Event,
				Finished:             fixture.Finished,
				FinishedProvisional:  fixture.FinishedProvisional,
				ID:                   fixture.ID,
				KickoffTime:          fixture.KickoffTime,
				Minutes:              fixture.Minutes,
				ProvisionalStartTime: fixture.ProvisionalStartTime,
				Started:              fixture.Started,
				TeamA:                fixture.TeamA,
				TeamAName:            &teamIndex[*fixture.TeamA].Name,
				TeamHName:            &teamIndex[*fixture.TeamH].Name,
				TeamAScore:           fixture.TeamAScore,
				TeamH:                fixture.TeamH,
				TeamHScore:           fixture.TeamHScore,
				Stats:                stats,
				TeamHDifficulty:      fixture.TeamHDifficulty,
				TeamADifficulty:      fixture.TeamADifficulty,
				PulseID:              fixture.PulseID,
				TeamAPulseID:         &teamIndex[*fixture.TeamA].PulseID,
				TeamHPulseID:         &teamIndex[*fixture.TeamH].PulseID,
			})
		}
	}
	return fixtures
}
