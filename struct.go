package gqlstruct

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"reflect"
	"time"
)

type GraphqlTyped interface {
	GraphqlType() graphql.Type
}

var GraphqlTypedType = reflect.TypeOf(new(GraphqlTyped)).Elem()

func nonNullable(t graphql.Type) graphql.Type {
	if x, ok := t.(*graphql.NonNull); ok {
		return x
	}
	return graphql.NewNonNull(t)
}

func nullable(t graphql.Type) graphql.Type {
	return t
}

var timeType = reflect.TypeOf(time.Time{})

func fieldType(field reflect.StructField, v reflect.Value) graphql.Type {
	t := field.Type

	fnTransformer := nullable

	if t.Kind() == reflect.Struct {
		vStruct := v
		tStruct := t
		if vStruct.CanAddr() {
			vStruct = vStruct.Addr()
			tStruct = reflect.PtrTo(t)
		}
		if tStruct.Implements(GraphqlTypedType) {
			return nonNullable(vStruct.Interface().(GraphqlTyped).GraphqlType())
		}
	}

	if t.Implements(GraphqlTypedType) {
		return v.Interface().(GraphqlTyped).GraphqlType()
	}

	// Check if it is a pointer or interface...
	if t.Kind() != reflect.Ptr && t.Kind() != reflect.Interface {
		// If it is not a pointer or a interface, applies a nonNullable
		// transformation in the result.
		fnTransformer = nonNullable
	} else {
		t = t.Elem()
	}

	if t == timeType {
		return fnTransformer(graphql.DateTime)
	}

	var result graphql.Type
	switch t.Kind() {
	case reflect.Bool:
		result = graphql.Boolean
	case reflect.String:
		result = graphql.String
	case
		reflect.Int, reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8,
		reflect.Uint, reflect.Uint64, reflect.Uint32, reflect.Uint16, reflect.Uint8:
		result = graphql.Int
	case
		reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128:
		result = graphql.Float
	default:
		panic(fmt.Sprintf("%s not recognized", t))
	}
	return fnTransformer(result)
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
		fields[tag] = &graphql.Field{
			Type: t,
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
