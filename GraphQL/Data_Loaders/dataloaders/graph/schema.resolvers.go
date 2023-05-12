package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.28

import (
	"context"
	database "final/db"
	"final/graph/model"
	"fmt"
	"log"
	"strconv"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	var Todo model.NewTodo

	Todo = model.NewTodo{
		ID:   input.ID,
		Name: input.Name,
		User: input.User,
	}

	db := database.Connect()
	if _, err := db.Exec(
		"INSERT INTO todo(id,name,user_id) values($1,$2,$3)", Todo.ID, Todo.Name,Todo.User); err != nil {
		log.Println(err)
	}

	log.Println("Todo created succesfully")
	return (*model.Todo)(&Todo), nil
}

// UpdateTodo is the resolver for the updateTodo field.
func (r *mutationResolver) UpdateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	var Todo model.Todo

	Todo = model.Todo{
		ID:   input.ID,
		Name: input.Name,
	}

	db := database.Connect()
	if _, err := db.Exec(
		"UPDATE public.todo SET name=$1 WHERE id=$2;", Todo.Name, Todo.ID); err != nil {
		log.Println(err)
	}

	log.Println("Todo updated succesfully")
	return &Todo, nil
}

// DeleteTodo is the resolver for the deleteTodo field.
func (r *mutationResolver) DeleteTodo(ctx context.Context, input model.ID) (*model.Newid, error) {
	var Todo model.Newid

	Todo = model.Newid{
		ID: input.ID,
	}

	db := database.Connect()
	rows, err := db.Query("delete from todo where id=$1", Todo.ID)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	log.Println("Todo deleted succesfully")
	return &Todo, nil
}

// This will print our all Todos
func (r *queryResolver) Todos(ctx context.Context, id string) ([]*model.Todo, error) {
	db := database.Connect()
	x, _ := strconv.Atoi(id)
	fmt.Println(x)
	rows, err := db.Query("select * from todo where user_id=$1",x)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	r.todos=nil

	for rows.Next() {
		var tod model.Todo
		if err := rows.Scan(&tod.ID, &tod.Name, &tod.User); err != nil {
			log.Println(err)
		}
		r.todos = append(r.todos, &tod)
	}
	if err := rows.Err(); err != nil {
		log.Println(err)
	}
	// RETURN ALL THE TODOS
	log.Println("Showing list of all todo's")
	return r.todos, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
