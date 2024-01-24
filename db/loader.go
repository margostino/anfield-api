package db

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/margostino/anfield-api/common"
	"github.com/margostino/anfield-api/fpl"
)

var Data, TeamIndex = load()

func get(url string, fplResponse interface{}) {
	response, err := http.Get(url)
	common.Check(err)
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	common.Check(err)

	err = json.Unmarshal(body, fplResponse)
	common.Check(err)
}

func load() (*Cache, *TeamIdx) {
	var fplStaticResponse fpl.StaticResponse
	var fplFixturesResponse fpl.FixturesResponse

	fplStaticUrl := os.Getenv("FPL_STATIC_URL")
	fplFixturesUrl := os.Getenv("FPL_FIXTURES_URL")

	get(fplStaticUrl, &fplStaticResponse)
	get(fplFixturesUrl, &fplFixturesResponse)

	responseFixtures, err := http.Get(fplFixturesUrl)
	common.Check(err)
	defer responseFixtures.Body.Close()

	response, err := http.Get(fplStaticUrl)
	common.Check(err)
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	common.Check(err)

	err = json.Unmarshal(body, &fplStaticResponse)
	common.Check(err)

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
