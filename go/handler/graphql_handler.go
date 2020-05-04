package handler

import (
	"net/http"

	"github.com/graphql-go/graphql"
	graphqlHandler "github.com/graphql-go/handler"
)

var tasks = make([]Task, 0)
var graphqlTask = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Task",
		Fields: graphql.Fields{
			"id":   &graphql.Field{Type: graphql.ID},
			"name": &graphql.Field{Type: graphql.String},
			"done": &graphql.Field{Type: graphql.Boolean},
		},
		Description: "tasks",
	},
)

func NewGraphqlHandler() http.Handler {
	schema, err := graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    newQuery(),
			Mutation: newMutation(),
		},
	)
	if err != nil {
		return nil
	}

	return graphqlHandler.New(&graphqlHandler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})
}

func newQuery() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name:   "RootQuery",
		Fields: GetQuery(),
	})
}

func newMutation() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "RootMutation",
		Fields: graphql.Fields{
			"PostTask":   GraphqlPostTasks(),
			"PutTask":    GraphqlPutTasks(),
			"DeleteTask": GraphqlDeleteTasks(),
		},
	})
}

func GetQuery() graphql.Fields {
	query := graphql.Fields{
		"GetQuery": &graphql.Field{
			Type:    GetQueryType(),
			Resolve: GetQueryResolve,
		},
	}
	return query
}

func GetQueryType() *graphql.Object {
	sampleType := graphql.NewObject(graphql.ObjectConfig{
		Name: "GetQuery",
		Fields: graphql.Fields{
			"tasks": GetUsers(),
		},
	})
	return sampleType
}

func GetUsers() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(graphqlTask),
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			return tasks, nil
		},
		Description: "graphqlTask",
	}
}

func GetQueryResolve(params graphql.ResolveParams) (interface{}, error) {
	return params, nil
}

func GraphqlPostTasks() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(graphqlTask),
		Args: PostQueryArgs(),
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			name, ok := p.Args["Name"].(string)
			if !ok {
				return tasks, nil
			}
			task := Task{
				Id:   len(tasks) + 1,
				Name: name,
				Done: false,
			}
			tasks = append(tasks, task)
			return tasks, nil
		},
		Description: "graphqlTask",
	}
}

func PostQueryArgs() map[string]*graphql.ArgumentConfig {
	return graphql.FieldConfigArgument{
		"Name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	}
}

func GraphqlPutTasks() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(graphqlTask),
		Args: PutQueryArgs(),
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			id, ok := p.Args["Id"].(int)
			if !ok {
				return tasks, nil
			}
			done, ok := p.Args["Done"].(bool)
			if !ok {
				return tasks, nil
			}
			for k, v := range tasks {
				if v.Id == id {
					tasks[k].Done = done
				}
			}
			return tasks, nil
		},
		Description: "graphqlTask",
	}
}

func PutQueryArgs() map[string]*graphql.ArgumentConfig {
	sampleArgs := graphql.FieldConfigArgument{
		"Id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"Done": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Boolean),
		},
	}
	return sampleArgs
}

func GraphqlDeleteTasks() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(graphqlTask),
		Args: DeleteQueryArgs(),
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			id, ok := p.Args["Id"].(int)
			if !ok {
				return tasks, nil
			}
			ttasks := make([]Task, 0)
			for _, v := range tasks {
				if v.Id != id {
					ttasks = append(ttasks, v)
				}
			}
			tasks = ttasks
			return tasks, nil
		},
		Description: "graphqlTask",
	}
}

func DeleteQueryArgs() map[string]*graphql.ArgumentConfig {
	sampleArgs := graphql.FieldConfigArgument{
		"Id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	}
	return sampleArgs
}
