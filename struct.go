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

func fieldType(field reflect.StructField, v reflect.Value) graphql.Type {
	t := field.Type

	if t.Kind() == reflect.Struct {
		vStruct := v
		tStruct := t
		if vStruct.CanAddr() {
			vStruct = vStruct.Addr()
			tStruct = reflect.PtrTo(t)
		}
		if tStruct.Implements(graphqlTypedType) {
			return vStruct.Interface().(GraphqlTyped).GraphqlType()
		}
	}

	if t.Implements(graphqlTypedType) {
		return v.Interface().(GraphqlTyped).GraphqlType()
	}

	// Check if it is a pointer or interface...
	if t.Kind() == reflect.Ptr || t.Kind() == reflect.Interface {
		// Updates the type with the type of the pointer
		t = t.Elem()
	}

	if t == timeType {
		return graphql.DateTime
	}

	switch t.Kind() {
	case reflect.Bool:
		return graphql.Boolean
	case reflect.String:
		return graphql.String
	case
		reflect.Int, reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8,
		reflect.Uint, reflect.Uint64, reflect.Uint32, reflect.Uint16, reflect.Uint8:
		return graphql.Int
	case
		reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128:
		return graphql.Float
	}
	panic(fmt.Sprintf("%s not recognized", t))
}

func fieldResolve(field reflect.StructField, v reflect.Value) graphql.FieldResolveFn {
	t := field.Type

	if t.Kind() == reflect.Struct {
		vStruct := v
		tStruct := t
		if vStruct.CanAddr() {
			vStruct = vStruct.Addr()
			tStruct = reflect.PtrTo(t)
		}
		if tStruct.Implements(graphqlResolverType) {
			return vStruct.Interface().(GraphqlResolver).GraphqlResolve
		}
	}

	if t.Implements(graphqlResolverType) {
		return v.Interface().(GraphqlResolver).GraphqlResolve
	}

	return nil
}

func objectConfig(obj interface{}) graphql.ObjectConfig {
	fields := graphql.Fields{}

	val := reflect.ValueOf(obj).Elem()
	for i := 0; i < val.NumField(); i++ {
		fValue := val.Field(i)
		fType := val.Type().Field(i)
		tag, ok := fType.Tag.Lookup("graphql")
		if !ok {
			continue
		}

		t := fieldType(fType, fValue)
		if len(tag) > 0 && tag[0] == '!' {
			t = graphql.NewNonNull(t)
			tag = tag[1:]
		}

		resolve := fieldResolve(fType, fValue)

		fields[tag] = &graphql.Field{
			Type:    t,
			Resolve: resolve,
		}
	}

	return graphql.ObjectConfig{
		Name:   val.Type().Name(),
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
	return graphql.NewObject(objectConfig(obj))
}
