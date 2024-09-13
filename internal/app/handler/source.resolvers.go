package handler

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"
	"fmt"

	"github.com/simonkimi/minebangumi/internal/app/graph"
	"github.com/simonkimi/minebangumi/internal/app/graph/grmodel"
)

// ParseSource is the resolver for the parseSource field.
func (r *queryResolver) ParseSource(ctx context.Context, source *grmodel.ParseSourceInput) (*grmodel.ParseSourceResponse, error) {
	panic(fmt.Errorf("not implemented: ParseSource - parseSource"))
}

// Query returns graph.QueryResolver implementation.
func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
