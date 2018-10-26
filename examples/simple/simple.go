package main

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/lab259/go-graphql-struct"
	"github.com/lab259/graphql-fasthttp-handler"
	"github.com/valyala/fasthttp"
)

type Person struct {
	Name    string   `graphql:"!name"`
	Age     int      `graphql:"age"`
	Friends []Person `graphql:"friends"`
}

// main will initialize the schema and start a HTTP server on port 8080.
func main() {
	// It creates the schema
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"hero": &graphql.Field{
					Type: gqlstruct.Struct(&Person{}),
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						return &Person{
							Name: "Snake Eyes",
							Age:  1,
							Friends: []Person{
								{
									Name: "Scarlett",
								},
								{
									Name: "Duke",
								},
							},
						}, nil
					},
				},
			},
		}), // Create the type based on the Query{} struct
	})

	if err != nil {
		panic(err)
	}

	// Create the handler
	h := handler.New(&handler.Config{
		Schema:     &schema,
		Pretty:     true,
		GraphiQL:   false,
		Playground: true,
	})

	fmt.Println("Starting the server at 8080 ...")
	// Start the fasthttp server.
	err = fasthttp.ListenAndServe(":8080", h.ServeHTTP)
	if err != nil {
		panic(err)
	}
}
