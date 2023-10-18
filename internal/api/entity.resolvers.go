package api

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.38

import (
	"context"
	"fmt"

	"go.infratographer.com/kubernetes-cluster-api/internal/ent/generated"
	"go.infratographer.com/x/gidx"
)

// FindClusterByID is the resolver for the findClusterByID field.
func (r *entityResolver) FindClusterByID(ctx context.Context, id gidx.PrefixedID) (*generated.Cluster, error) {
	panic(fmt.Errorf("not implemented: FindClusterByID - findClusterByID"))
}

// Entity returns EntityResolver implementation.
func (r *Resolver) Entity() EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }