package graph

import "github.com/margostino/anfield-api/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	TeamStore map[string]model.Team
}
