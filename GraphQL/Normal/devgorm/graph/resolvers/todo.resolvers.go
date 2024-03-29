package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.28

import (
	"context"
	"devgorm/graph/customTypes"
	"devgorm/graph/generated"
	"fmt"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, text string) (*customTypes.Todo, error) {
	panic(fmt.Errorf("not implemented: CreateTodo - createTodo"))
}

// UpdateTodo is the resolver for the updateTodo field.
func (r *mutationResolver) UpdateTodo(ctx context.Context, input customTypes.TodoInput) (*customTypes.Todo, error) {
	panic(fmt.Errorf("not implemented: UpdateTodo - updateTodo"))
}

// DeleteTodo is the resolver for the deleteTodo field.
func (r *mutationResolver) DeleteTodo(ctx context.Context, todoID int) (*customTypes.Todo, error) {
	panic(fmt.Errorf("not implemented: DeleteTodo - deleteTodo"))
}

// GetTodos is the resolver for the getTodos field.
func (r *queryResolver) GetTodos(ctx context.Context) ([]*customTypes.Todo, error) {
	panic(fmt.Errorf("not implemented: GetTodos - getTodos"))
}

// GetTodo is the resolver for the getTodo field.
func (r *queryResolver) GetTodo(ctx context.Context, todoID int) (*customTypes.Todo, error) {
	panic(fmt.Errorf("not implemented: GetTodo - getTodo"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
