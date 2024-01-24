package db

import (
	"strings"

	"github.com/margostino/anfield-api/fpl"
)

func loadTeams(fplTeams []*fpl.Team) (map[string]*Team, map[int]*Team) {
	teams := make(map[string]*Team)
	teamIndex := make(map[int]*Team)
	for _, team := range fplTeams {
		key := strings.ToLower(team.ShortName)
		teamElement := &Team{
			ID:                  team.Code,
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
		teamIndex[team.ID] = teamElement
	}
	return teams, teamIndex
}
