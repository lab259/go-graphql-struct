[![CircleCI](https://circleci.com/gh/lab259/go-graphql-struct.svg?style=shield)](https://circleci.com/gh/lab259/go-graphql-struct)
[![codecov](https://codecov.io/gh/lab259/go-graphql-struct/branch/master/graph/badge.svg)](https://codecov.io/gh/lab259/go-graphql-struct)
[![Go Report Card](https://goreportcard.com/badge/github.com/lab259/go-graphql-struct)](https://goreportcard.com/report/github.com/lab259/go-graphql-struct)
[![Go Doc](https://img.shields.io/badge/godoc-reference-5272B4.svg?style=shield)](https://godoc.org/github.com/lab259/go-graphql-struct)

# go-graphql-struct (gqlstruct)

This library implements generating GraphQL Schema based on tagged
structs using the [github.com/graphql-go/graphql](https://github.com/graphql-go/graphql)
implementation.

Usually, building the schema is a one time task and it is done
statically. So, this library does not degrade the performance, not even
a little, but in that one-time initialization.

## Usage

TODO

## Custom Types

The default data types of the GraphQL can be count in one hand, which is
not a bad thing. However, that means that you may need to implement some
scalar types (or event complexes types) yourself.

In order to provide custom types for the fields the `GraphqlTyped`
interface was defined:

```go
type GraphqlTyped interface {
    GraphqlType() graphql.Type
}
```

An example:

```go
type TypeA string

func (*TypeA) GraphqlType() graphql.Type {
    return graphql.Int
}

```

Remember, this library is all about declaring the schema. If you need
marshalling/unmarshaling a custom type to another, use the implementation
of the [github.com/graphql-go/graphql](https://github.com/graphql-go/graphql)
library (check on the `graphql.NewScalar` and `graphql.ScalarConfig`).

## Resolver

To implement resolvers over a Custom Type, you will implement the
interface `GraphqlResolver`:

```go
type GraphqlResolver interface {
    GraphqlResolve(p graphql.ResolveParams) (interface{}, error)
}
```

**IMPORTANT**: Although the method `GraphqlResolve` is a member of a struct, it is
called statically. So, do not make any references of the struct itself,
inside of this method.

An example:

```go
type TypeA string

func (*TypeA) GraphqlType() graphql.Type {
    return graphql.Int
}
```

## Limitations

* This library do not deal with arrays yet.

## License

MIT