package handler

import (
	"fmt"

	"github.com/graphql-go/graphql"
	graphqlHandler "github.com/graphql-go/handler"
)

type GraphqlTask struct {
	Id   int
	Name string
	Done bool
}

var graphqlTasks = make([]*GraphqlTask, 0)

var graphqlTask = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id":        &graphql.Field{Type: graphql.ID},
			"name":      &graphql.Field{Type: graphql.String},
			"age":       &graphql.Field{Type: graphql.String},
			"createdAt": &graphql.Field{Type: graphql.String},
			"updatedAt": &graphql.Field{Type: graphql.String},
			"deletedAt": &graphql.Field{Type: graphql.String},
		},
		Description: "Users data",
	},
)

func NewGraphqlHandler() (*graphqlHandler.Handler, error) {
	schema, err := graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    newQuery(),
			Mutation: newMutation(),
		},
	)
	if err != nil {
		return nil, err
	}

	return graphqlHandler.New(&graphqlHandler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	}), nil
}
func newQuery() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{

		Name: "Query",
		Fields: graphql.Fields{
			"tasks": NewUsers(),
		},
	})
}
func newMutation() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"postTask":   NewUsers(),
			"putTask":    NewUsers(),
			"deleteTask": NewUsers(),
		},
	})
}

func NewUsers() *graphql.Field {
	//return func(c echo.Context) error {
	//	return c.JSON(http.StatusOK, graphqlTasks)
	//}
	return &graphql.Field{
		Type: graphql.NewList(graphqlTask),
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			//var u []*model.User
			fmt.Println(p)
			//if err := db.Find(&u).Error; err != nil {
			//	// do something
			//}

			return graphqlTasks, nil
		},
		Description: "graphqlTask",
	}
}
