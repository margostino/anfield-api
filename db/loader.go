package db

import (
	"os"
	"sync"

	"github.com/margostino/anfield-api/fpl"
	"github.com/margostino/anfield-api/network"
)

var Data, TeamIndex = load()

func load() (*Cache, *TeamIdx) {
	var wg sync.WaitGroup
	var fplStaticResponse fpl.StaticResponse
	var fplFixturesResponse fpl.FixturesResponse

	fplStaticUrl := os.Getenv("FPL_STATIC_URL")
	fplFixturesUrl := os.Getenv("FPL_FIXTURES_URL")

	wg.Add(2)
	go network.Get(fplStaticUrl, &fplStaticResponse, &wg, nil)
	go network.Get(fplFixturesUrl, &fplFixturesResponse, &wg, nil)
	wg.Wait()

	teams, teamIndex := loadTeams(fplStaticResponse.Teams)
	events := loadEvents(fplStaticResponse.Events, fplFixturesResponse, teamIndex)
	elementTypes := loadTypes(fplStaticResponse.ElementTypes)
	players := loadPlayers(fplStaticResponse.Players, teamIndex, elementTypes)

	return &Cache{
			Teams:        teams,
			Players:      players,
			ElementTypes: elementTypes,
			Events:       events,
		}, &TeamIdx{
			Index: teamIndex,
		}
}
