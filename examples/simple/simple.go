package main

import (
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

type Query struct {
	Hero   Person   `graphql:"!hero"`
	People []Person `graphql:"people"`
}

// main will initialize the schema and start a HTTP server on port 8080.
func main() {
	// It creates the schema
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: gqlstruct.Struct(Query{}), // Create the type based on the Query{} struct
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

	// Start the fasthttp server.
	fasthttp.ListenAndServe(":8080", h.ServeHTTP)
}
