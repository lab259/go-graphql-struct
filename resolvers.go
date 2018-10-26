package gqlstruct

import (
	"github.com/graphql-go/graphql"
	"reflect"
)

// GraphqlResolver is the interface implemented by types that will provide a
// a resolver.
type GraphqlResolver interface {
	// GraphqlResolve is the method will be set to the `graphql.Field` as the
	// resolver method.
	GraphqlResolve(p graphql.ResolveParams) (interface{}, error)
}

func fieldResolve(field reflect.StructField) graphql.FieldResolveFn {
	t := field.Type

	if t.Kind() == reflect.Struct {
		// If the type is a struct, we need the a pointer to that struct to
		// check if it implements the interface.
		tStruct := reflect.PtrTo(t)
		if tStruct.Implements(graphqlResolverType) {
			vStruct := reflect.New(t)
			return vStruct.Interface().(GraphqlResolver).GraphqlResolve
		}
	}

	if t.Implements(graphqlResolverType) {
		vStruct := reflect.New(t).Elem()
		return vStruct.Interface().(GraphqlResolver).GraphqlResolve
	}

	return nil
}
