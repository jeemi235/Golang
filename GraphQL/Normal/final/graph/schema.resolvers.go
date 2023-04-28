package graph


			// This file will be automatically regenerated based on the schema, any resolver implementations
			// will be copied through when generating and any unknown code will be moved to the end.
			// Code generated by github.com/99designs/gqlgen version v0.17.30

import (
"context"
"fmt"
"io"
"strconv"
"time"
"sync"
"errors"
"bytes"
gqlparser "github.com/vektah/gqlparser/v2"
"github.com/vektah/gqlparser/v2/ast"
"github.com/99designs/gqlgen/graphql"
"github.com/99designs/gqlgen/graphql/introspection"
database "final/db"
"final/graph/model"
"log")


















// CreateTodo is the resolver for the createTodo field.
	func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) ( *model.Todo,  error){
		var Todo model.NewTodo

	Todo = model.NewTodo{
		ID:   input.ID,
		Name: input.Name,
	}

	db := database.Connect()
	if _, err := db.Exec(
		"INSERT INTO todo(id,name) values($1,$2)", Todo.ID, Todo.Name); err != nil {
		log.Println(err)
	}

	log.Println("Todo created succesfully")
	return (*model.Todo)(&Todo), nil
	}

// CreateTodoTwo is the resolver for the createTodoTwo field.
	func (r *mutationResolver) CreateTodoTwo(ctx context.Context, input model.NewTodo) ( *model.Todo,  error){
		panic(fmt.Errorf("not implemented: CreateTodoTwo - createTodoTwo"))
	}

// UpdateTodo is the resolver for the updateTodo field.
	func (r *mutationResolver) UpdateTodo(ctx context.Context, input model.NewTodo) ( *model.Todo,  error){
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
	func (r *mutationResolver) DeleteTodo(ctx context.Context, input model.ID) ( *model.Newid,  error){
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

// Todos is the resolver for the todos field.
This will print our all Todo's
	func (r *queryResolver) Todos(ctx context.Context) ( []*model.Todo,  error){
		db := database.Connect()
	rows, err := db.Query("select * from todo")
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	r.todos = nil

	for rows.Next() {
		var tod model.Todo
		if err := rows.Scan(&tod.ID, &tod.Name); err != nil {
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


type mutationResolver struct { *Resolver }
type queryResolver struct { *Resolver }



