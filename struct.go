package gqlstruct

import (
	"github.com/graphql-go/graphql"
	"reflect"
)

func objectConfig(t reflect.Type) (graphql.ObjectConfig, error) {
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
			return graphql.ObjectConfig{}, NewErrTypeNotRecognizedWithStruct(err, t, field.Type)
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
	}, nil
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

	objConfig, err := objectConfig(t)
	if err != nil {
		panic(err.Error())
	}

	return graphql.NewObject(objConfig)
}

func fromTypeOf(t reflect.Type) (graphql.Type, error) {
	objConfig, err := objectConfig(t)
	if err != nil {
		return nil, err
	}
	return graphql.NewObject(objConfig), nil
}
