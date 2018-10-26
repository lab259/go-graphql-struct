package gqlstruct

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"reflect"
)

func objectConfig(t reflect.Type) graphql.ObjectConfig {
	fields := graphql.Fields{}

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	// Goes field by field of the object.
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag, ok := field.Tag.Lookup("graphql")
		if !ok {
			// If the field is not tagged, ignore it.
			continue
		}

		objectType, err := buildFieldType(field.Type)
		if err != nil {
			panic(fmt.Sprintf("%s.%s:%s", objectType.Name(), field.Name, err.Error()))
		}

		// If the tag starts with "!" it is a NonNull type.
		if len(tag) > 0 && tag[0] == '!' {
			objectType = graphql.NewNonNull(objectType)
			tag = tag[1:]
		}

		resolve := fieldResolve(field)

		fields[tag] = &graphql.Field{
			Type:    objectType,
			Resolve: resolve,
		}
	}

	return graphql.ObjectConfig{
		Name:   t.Name(),
		Fields: fields,
	}
}

// Struct returns a `*graphql.Object` with the description extracted from the
// obj passed.
//
// This method extracts the information needed from the fields of the obj
// informed. All fields tagged with "graphql" are added.
//
// The "graphql" tag can be defined as:
//
// ```
// type T struct {
//     field string `graphql:"fieldname"`
// }
// ```
//
// * fieldname: The name of the field.
func Struct(obj interface{}) *graphql.Object {
	t := reflect.TypeOf(obj)
	return graphql.NewObject(objectConfig(t))
}
