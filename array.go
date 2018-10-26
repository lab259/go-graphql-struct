package gqlstruct

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"reflect"
)

func Array(arr interface{}) graphql.Type {
	tArr := reflect.TypeOf(arr)
	if tArr.Kind() != reflect.Array && tArr.Kind() != reflect.Slice {
		panic(fmt.Sprintf("%s is not an array", tArr))
	}
	return graphql.NewList(graphql.NewObject(objectConfig(tArr.Elem())))
}
