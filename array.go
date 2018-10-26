package gqlstruct

import (
	"github.com/graphql-go/graphql"
	"reflect"
)

func ArrayOf(t reflect.Type) graphql.Type {
	if t.Kind() == reflect.Ptr {
		// If pointer, get the Type of the pointer
		t = t.Elem()
	}
	var typeBuilt graphql.Type
	if t.Kind() == reflect.Struct {
		objConfig, err := objectConfig(t)
		if err != nil {
			panic(err.Error())
		}
		typeBuilt = graphql.NewObject(objConfig)
	} else {
		ttt, err := buildFieldType(t)
		if err != nil {
			panic(err)
		}
		typeBuilt = ttt
	}
	return graphql.NewList(typeBuilt)
}
