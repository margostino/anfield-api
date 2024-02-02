#  Anfield API

This project is a GraphQL API written in Golang hosted by Vercel to access to the Fantasy Premier League **public** information. The motivation of this project is to provide a simple, single and easy way to access to the Fantasy Premier League data, rather than using multiple endpoints separately.

## Endpoints

- Query: https://anfield-api-margostino.vercel.app/api/query

- Playground: https://anfield-api-margostino.vercel.app/api/playground

## Authentication

The API is using the Fantasy Premier League public API, but authentication is required simply to control some minimal rate limits. The API is using a simple token-based authentication, where the token is passed in the `Authorization` header. Yes, for free!

```http
Authorization : Bearer <api-token>
```

You can request a token by sending an email to me or creating an issue in this repository.

## Schema

You can find the GraphQL Schema [here](https://raw.githubusercontent.com/margostino/anfield-api/master/graph/schema.graphqls)

The API provides the following:
- Teams information (e.g. name, short name, code, strength attack, etc.)
- Players information (e.g. name, teams, points, cost, game stats, etc.)
- Events information (e.g. kick-off datetime, stats, etc.)
- Head to Head information with historical games between 2 teams (e.g. team names, team scores, etc.)

## Example of queries

### Get all players

```graphql
{
	players {
		webName
		team
		news
		position
		chanceOfPlayingNextRound
		expectedGoals
		form		
		nowCost
		pointsPerGame
		selectedByPercent
		selectedRank
		selectedRankType
		totalPoints
		minutes
		goalsScored
		assists
		cleanSheets
		goalsConceded
		ownGoals
		influence
		ownGoals
		ictIndex
		threat
		creativity
		bonus
		redCards
		yellowCards
	}
}
```

### Get player by name

```graphql
{
	player(webName: "Salah") {
		webName
		team
		news
		position
		chanceOfPlayingNextRound
	}
}
```

### Get all players of a team

```graphql
{
	players(teamShortName: "LIV") {
		webName
		team
		news
		position
		chanceOfPlayingNextRound
		expectedGoals
		form		
		nowCost
		pointsPerGame
		selectedByPercent
		selectedRank
		selectedRankType
		totalPoints
		minutes
		goalsScored
		assists
		cleanSheets
		goalsConceded
		ownGoals
		influence
		ownGoals
		ictIndex
		threat
		creativity
		bonus
		redCards
		yellowCards
	}
}
```

# Support
For support, issues, or contributions, please visit our GitHub repository.

# License

Anfield API is licensed under the [MIT license](https://github.com/margostino/anfield-api/blob/master/LICENSE)
