package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.43

import (
	"context"
	"fmt"
	"strings"

	"github.com/margostino/anfield-api/db"
	"github.com/margostino/anfield-api/graph/model"
)

// Team is the resolver for the team field.
func (r *queryResolver) Team(ctx context.Context, shortName string) (*model.Team, error) {
	var key = strings.ToLower(shortName)
	var data = db.Data.Teams[key]
	var response = toTeamGraph(data)
	return response, nil
}

// Teams is the resolver for the teams field.
func (r *queryResolver) Teams(ctx context.Context) ([]*model.Team, error) {
	panic(fmt.Errorf("not implemented: Player - player"))
}

// Player is the resolver for the player field.
func (r *queryResolver) Player(ctx context.Context, webName string) (*model.Player, error) {
	var key = strings.ToLower(webName)
	var data = db.Data.Players[key]
	var response = toPlayerGraph(data)
	return response, nil
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
