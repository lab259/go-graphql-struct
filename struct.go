package gqlstruct

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"reflect"
	"time"
)

// GraphqlTyped is the interface implemented by types that will provide a
// special `graphql.Type`.
type GraphqlTyped interface {
	// GraphqlType returns the `graphql.Type` that represents the data type that
	// implements this interface.
	GraphqlType() graphql.Type
}

// GraphqlResolver is the interface implemented by types that will provide a
// a resolver.
type GraphqlResolver interface {
	// GraphqlResolve is the method will be set to the `graphql.Field` as the
	// resolver method.
	GraphqlResolve(p graphql.ResolveParams) (interface{}, error)
}

var (
	graphqlTypedType    = reflect.TypeOf(new(GraphqlTyped)).Elem()
	graphqlResolverType = reflect.TypeOf(new(GraphqlResolver)).Elem()
	timeType            = reflect.TypeOf(time.Time{})
)

func buildFieldType(fieldType reflect.Type) (graphql.Type, error) {
	if fieldType.Kind() == reflect.Struct && fieldType != timeType {
		// If the type is a struct, we need the a pointer to that struct to
		// check if it implements the interface.
		tStruct := reflect.PtrTo(fieldType)
		if tStruct.Implements(graphqlTypedType) {
			vStruct := reflect.New(fieldType)
			return vStruct.Interface().(GraphqlTyped).GraphqlType(), nil
		}
	}

	if fieldType.Implements(graphqlTypedType) {
		vStruct := reflect.New(fieldType.Elem())
		return vStruct.Interface().(GraphqlTyped).GraphqlType(), nil
	}

	// Check if it is a pointer or interface...
	if fieldType.Kind() == reflect.Ptr || fieldType.Kind() == reflect.Interface {
		// Updates the type with the type of the pointer
		fieldType = fieldType.Elem()
	}

	// Special case: If the type is the time.Time type.
	if fieldType == timeType {
		return graphql.DateTime, nil
	}

	switch fieldType.Kind() {
	case reflect.Struct:
		return graphql.NewObject(objectConfig(fieldType)), nil
	case reflect.Array, reflect.Slice:
		t, err := buildFieldType(fieldType.Elem())
		if err != nil {
			return nil, err
		}
		return graphql.NewList(t), nil
	case reflect.Bool:
		return graphql.Boolean, nil
	case reflect.String:
		return graphql.String, nil
	case
		reflect.Int, reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8,
		reflect.Uint, reflect.Uint64, reflect.Uint32, reflect.Uint16, reflect.Uint8:
		return graphql.Int, nil
	case
		reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128:
		return graphql.Float, nil
	}
	return nil, NewErrTypeNotRecognized(fieldType)
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

func Array(arr interface{}) graphql.Type {
	tArr := reflect.TypeOf(arr)
	if tArr.Kind() != reflect.Array && tArr.Kind() != reflect.Slice {
		panic(fmt.Sprintf("%s is not an array", tArr))
	}
	return graphql.NewList(graphql.NewObject(objectConfig(tArr.Elem())))
}
