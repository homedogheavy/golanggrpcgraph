package handler

import (
	"fmt"
	"net/http"

	"github.com/graphql-go/graphql"
	graphqlHandler "github.com/graphql-go/handler"
	"github.com/k0kubun/pp"
)

//type Task struct {
//	Id   int
//	Name string
//	Done bool
//}

var graphqlTasks = make([]*Task, 0)
var tasks = make([]Task, 0)

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

func NewGraphqlHandler() http.Handler {
	//query := fields.SetQuery()
	//mutateQuery := fields.SetMutation()
	//rootMutation := graphql.NewObject(graphql.ObjectConfig{Name: "RootMutation", Fields: mutateQuery})
	//rootQuery := graphql.NewObject(graphql.ObjectConfig{Name: "RootQuery", Fields: query})

	schema, err := graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    newQuery(),
			Mutation: newMutation(),
		},
	)
	if err != nil {

		fmt.Println("------------------------------------------------------------------------------------------------------------------------------------------------------")
		fmt.Println(err)
		return nil
	}

	return graphqlHandler.New(&graphqlHandler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})
	//return handler.New(&handler.Config{
	//	Schema:     &schema,
	//	Pretty:     true,
	//	GraphiQL:   false,
	//	Playground: true,
	//})
}

//func GraphqlSetting() http.Handler {
//	query := fields.SetQuery()
//	mutateQuery := fields.SetMutation()
//	rootMutation := graphql.NewObject(graphql.ObjectConfig{Name: "RootMutation", Fields: mutateQuery})
//	rootQuery := graphql.NewObject(graphql.ObjectConfig{Name: "RootQuery", Fields: query})
//
//	schema, err := graphql.NewSchema(graphql.SchemaConfig{
//		Query:    rootQuery,
//		Mutation: rootMutation,
//	})
//	if err != nil {
//		panic(err)
//	}
//	h := handler.New(&handler.Config{
//		Schema:     &schema,
//		Pretty:     true,
//		GraphiQL:   false,
//		Playground: true,
//	})
//	return h
//}

func SetQuery() graphql.Fields {
	//query := graphql.Fields{
	//	"SampleQuery": &graphql.Field{
	//		Type:    SampleQueryType(),
	//		Args:    SampleQueryArgs(),
	//		Resolve: SampleQueryResolve,
	//	},
	//}
	//return query

	query := graphql.Fields{
		"GetQuery": &graphql.Field{
			Type: SampleQueryType(),
			//Args:    SampleQueryArgs(),
			Resolve: SampleQueryResolve,
		},
	}
	return query
	//return graphql.NewObject(graphql.ObjectConfig{
	//	Name: "GetQuery",
	//	Fields: graphql.Fields{
	//		"tasks": NewUsers(),
	//	},
	//})
}
func SampleQueryType() *graphql.Object {
	sampleType := graphql.NewObject(graphql.ObjectConfig{
		Name: "SampleQuery",
		Fields: graphql.Fields{
			"tasks": NewUsers(),
			//"Message": &graphql.Field{
			//	Type: graphql.String,
			//},
		},
	})
	return sampleType
}

func SampleQueryArgs() map[string]*graphql.ArgumentConfig {
	sampleArgs := graphql.FieldConfigArgument{
		"ID": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"Password": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	}
	return sampleArgs
}

type SampleMutateResp struct {
	Message string
}

func SampleQueryResolve(params graphql.ResolveParams) (interface{}, error) {
	resp := &SampleMutateResp{
		Message: "Hello, This is SampleQuery",
	}
	return resp, nil
}

func newQuery() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name:   "RootQuery",
		Fields: SetQuery(),
	})
	//return graphql.NewObject(graphql.ObjectConfig{
	//	Name: "GetQuery",
	//	Fields: graphql.Fields{
	//		"tasks": NewUsers(),
	//	},
	//})
}
func newMutation() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "RootMutation",
		Fields: graphql.Fields{
			"PostTask": GraphqlPostTasks(),
		},
	})
	//return graphql.NewObject(graphql.ObjectConfig{
	//	Name: "postQuery",
	//	Fields: graphql.Fields{
	//		"postTask":   NewUsers(),
	//		"putTask":    NewUsers(),
	//		"deleteTask": NewUsers(),
	//	},
	//})
}

func SampleQueryArgs2() map[string]*graphql.ArgumentConfig {
	sampleArgs := graphql.FieldConfigArgument{
		"Name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		//"Password": &graphql.ArgumentConfig{
		//	Type: graphql.NewNonNull(graphql.String),
		//},
	}
	return sampleArgs
}
func GraphqlPostTasks() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(graphqlTask),
		Args: SampleQueryArgs2(),
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			fmt.Println("-----------------------------------------------------------------------------------------------")
			fmt.Println("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
			//var u []*model.User
			pp.Println(p.Args["Name"])
			//if err := db.Find(&u).Error; err != nil {
			//	// do something
			//}

			//req := &struct {
			//	Name string
			//}{}
			//if err := c.Bind(req); err != nil {
			//	return c.NoContent(http.StatusBadRequest)
			//}
			name, ok := p.Args["Name"].(string)
			if !ok {
				return graphqlTasks, nil
			}
			task := Task{
				Id:   len(tasks) + 1,
				Name: name,
				Done: false,
			}
			tasks = append(tasks, task)
			pp.Println(tasks)
			// return c.NoContent(http.StatusOK)
			return graphqlTasks, nil
		},
		Description: "graphqlTask",
	}

	//return func(c echo.Context) error {
	//	req := &struct {
	//		Name string
	//	}{}
	//	if err := c.Bind(req); err != nil {
	//		return c.NoContent(http.StatusBadRequest)
	//	}
	//	task := Task{
	//		Id:   len(Tasks) + 1,
	//		Name: req.Name,
	//		Done: false,
	//	}
	//	Tasks = append(Tasks, task)
	//	return c.NoContent(http.StatusOK)
	//}
}

func NewUsers() *graphql.Field {
	//return func(c echo.Context) error {
	//	return c.JSON(http.StatusOK, graphqlTasks)
	//}
	return &graphql.Field{
		Type: graphql.NewList(graphqlTask),
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			fmt.Println("-----------------------------------------------------------------------------------------------")
			fmt.Println("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
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
