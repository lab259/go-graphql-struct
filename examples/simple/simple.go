package main

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/lab259/go-graphql-struct"
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

func main() {
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: gqlstruct.Struct(Query{}),
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(schema)
}
