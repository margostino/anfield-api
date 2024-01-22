package graph

import (
	"github.com/margostino/anfield-api/db"
	"github.com/margostino/anfield-api/graph/model"
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
		ID:         player.ID,
		FirstName:  player.FirstName,
		SecondName: player.SecondName,
		WebName:    player.WebName,
		News:       player.News,
		NewsAdded:  player.NewsAdded,
		Team:       player.Team,
	}
}
