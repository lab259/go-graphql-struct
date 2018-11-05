package gqlstruct

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"reflect"
)

type encoder struct {
	types map[string]graphql.Type
}

func NewEncoder() *encoder {
	return &encoder{
		types: make(map[string]graphql.Type),
	}
}

var defaultEncoder = NewEncoder()

func (enc *encoder) Struct(obj interface{}, options ...Option) (*graphql.Object, error) {
	t := reflect.TypeOf(obj)
	return enc.StructOf(t, options...)
}

func (enc *encoder) Args(obj interface{}) (graphql.FieldConfigArgument, error) {
	t := reflect.TypeOf(obj)
	return enc.ArgsOf(t)
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
func (enc *encoder) StructOf(t reflect.Type, options ...Option) (*graphql.Object, error) {
	if r, ok := enc.getType(t); ok {
		if d, ok := r.(*graphql.Object); ok {
			return d, nil
		}
		return nil, fmt.Errorf("%s is not an graphql.Object", r)
	}

	name := t.Name()
	if t.Kind() == reflect.Ptr {
		name = t.Elem().Name()
	}

	objCfg := graphql.ObjectConfig{
		Name:   name,
		Fields: graphql.Fields{},
	}

	// Apply options
	for _, opt := range options {
		err := opt.Apply(&objCfg)
		if err != nil {
			return nil, err
		}
	}

	r := graphql.NewObject(objCfg)
	enc.registerType(t, r)

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

		objectType, ok := enc.getType(field.Type)
		if !ok {
			ot, err := enc.buildFieldType(field.Type)
			if err != nil {
				return nil, NewErrTypeNotRecognizedWithStruct(err, t, field)
			}
			objectType = ot
			enc.registerType(field.Type, ot)
		}

		// If the tag starts with "!" it is a NonNull type.
		if len(tag) > 0 && tag[0] == '!' {
			objectType = graphql.NewNonNull(objectType)
			tag = tag[1:]
		}

		resolve := fieldResolve(field)

		gfield := &graphql.Field{
			Type:    objectType,
			Resolve: resolve,
		}
		r.AddFieldConfig(tag, gfield)
	}
	return r, nil
}

func (enc *encoder) FieldOf(t reflect.Type, options ...Option) (graphql.Field, error) {
	r := graphql.Field{}

	fieldType, err := enc.StructOf(t)
	if err != nil {
		return graphql.Field{}, err
	}
	r.Type = fieldType

	for _, option := range options {
		err = option.Apply(&r)
		if err != nil {
			return graphql.Field{}, err
		}
	}

	return r, nil
}

func (enc *encoder) Field(t interface{}, options ...Option) (graphql.Field, error) {
	return enc.FieldOf(reflect.TypeOf(t), options...)
}

func (enc *encoder) ArrayOf(t reflect.Type, options ...Option) (graphql.Type, error) {
	if t.Kind() == reflect.Ptr {
		// If pointer, get the Type of the pointer
		t = t.Elem()
	}
	var typeBuilt graphql.Type
	if cachedType, ok := enc.getType(t); ok {
		return graphql.NewList(cachedType), nil
	}
	if t.Kind() == reflect.Struct {
		bt, err := enc.StructOf(t, options...)
		if err != nil {
			return nil, err
		}
		typeBuilt = bt
	} else {
		ttt, err := enc.buildFieldType(t)
		if err != nil {
			return nil, err
		}
		typeBuilt = ttt
	}
	enc.registerType(t, typeBuilt)
	return graphql.NewList(typeBuilt), nil
}

func (enc *encoder) ArgsOf(t reflect.Type) (graphql.FieldConfigArgument, error) {
	r := graphql.FieldConfigArgument{}

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	if t.Kind() != reflect.Struct {
		return r, fmt.Errorf("cannot build args from a non struct")
	}

	// Goes field by field of the object.
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag, ok := field.Tag.Lookup("graphql")
		if !ok {
			// If the field is not tagged, ignore it.
			continue
		}

		objectType, ok := enc.getType(field.Type)
		if !ok {
			ot, err := enc.buildFieldType(field.Type)
			if err != nil {
				return nil, NewErrTypeNotRecognizedWithStruct(err, t, field)
			}
			objectType = ot
			enc.registerType(field.Type, ot)
		}

		// If the tag starts with "!" it is a NonNull type.
		if len(tag) > 0 && tag[0] == '!' {
			objectType = graphql.NewNonNull(objectType)
			tag = tag[1:]
		}

		graphQLArgument := &graphql.ArgumentConfig{
			Type: objectType,
		}
		r[tag] = graphQLArgument
	}

	return r, nil
}

func (enc *encoder) getType(t reflect.Type) (graphql.Type, bool) {
	name := t.Name()
	if t.Kind() == reflect.Ptr {
		name = t.Elem().Name()
	}
	gt, ok := enc.types[name]
	return gt, ok
}

func (enc *encoder) registerType(t reflect.Type, r graphql.Type) {
	name := t.Name()
	if t.Kind() == reflect.Ptr {
		name = t.Elem().Name()
	}
	enc.types[name] = r
}

func Struct(obj interface{}) *graphql.Object {
	r, err := defaultEncoder.Struct(obj)
	if err != nil {
		panic(err.Error())
	}
	return r
}

func ArrayOf(t reflect.Type) graphql.Type {
	r, err := defaultEncoder.ArrayOf(t)
	if err != nil {
		panic(err.Error())
	}
	return r
}

func ArgsOf(t reflect.Type) graphql.FieldConfigArgument {
	r, err := defaultEncoder.ArgsOf(t)
	if err != nil {
		panic(err.Error())
	}
	return r
}

func FieldOf(t reflect.Type) graphql.FieldConfigArgument {
	r, err := defaultEncoder.ArgsOf(t)
	if err != nil {
		panic(err.Error())
	}
	return r
}

func Field(t reflect.Type) graphql.Field {
	r, err := defaultEncoder.Field(t)
	if err != nil {
		panic(err.Error())
	}
	return r
}
