package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.28

import (
	"context"
	"devgenius/graph/model"
	"fmt"
)

// RegisterNewNinja is the resolver for the registerNewNinja field.
func (r *mutationResolver) RegisterNewNinja(ctx context.Context, input *model.NewNinja) (*model.ResponseMessage, error) {
	//If there is a database, here you would insert the new data
	//as both models are similar, we type cast our input into the Ninja model
	newNinja := model.Ninja(*input)

	//we append into our array
	r.Ninjas = append(r.Ninjas, &newNinja)

	return &model.ResponseMessage{Message: "Ninja added"}, nil
}

// FindGenin is the resolver for the findGenin field.
func (r *queryResolver) FindGenin(ctx context.Context, name string) (*model.Ninja, error) {
	//This simulates a Find in a database
	for _, tempNinja := range r.Ninjas {

		//we look for ninjas that match the status of genin and the name
		if *tempNinja.Name == name && *tempNinja.Rank == "Genin" {
			return tempNinja, nil
		}
	}
	return nil, nil
}

// ReturnAllHokages is the resolver for the returnAllHokages field.
func (r *queryResolver) ReturnAllHokages(ctx context.Context) ([]*model.Ninja, error) {
	hokages := []*model.Ninja{}
	for _, tempNinja := range r.Ninjas {
		//we look for ninjas that match the status of hokage and the name
		if *tempNinja.Rank == "Hokage" {
			hokages = append(hokages, tempNinja)
		}
	}
	return hokages, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented: CreateTodo - createTodo"))
}
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	panic(fmt.Errorf("not implemented: Todos - todos"))
}