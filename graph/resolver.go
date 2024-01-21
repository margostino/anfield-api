package graph

import (
	"context"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Fetcher func(ctx context.Context, entity string, year int, response any) (map[string]*float64, error)
}
