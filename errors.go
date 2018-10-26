package gqlstruct

import (
	"fmt"
	"reflect"
)

type TypeNotRecognizedError struct {
	t reflect.Type
}

func (t *TypeNotRecognizedError) Error() string {
	return fmt.Sprintf("%s not recognized", t.t)
}

func NewErrTypeNotRecognized(t reflect.Type) error {
	return &TypeNotRecognizedError{
		t: t,
	}
}

type TypeNotRecognizedWithStructError struct {
	reason     error
	structType reflect.Type
	fieldType  reflect.Type
}

func (err *TypeNotRecognizedWithStructError) Error() string {
	return fmt.Sprintf("%s.%s:%s", err.structType.Name(), err.fieldType.Name, err.reason.Error())
}

func NewErrTypeNotRecognizedWithStruct(reason error, structType, fieldType reflect.Type) error {
	return &TypeNotRecognizedWithStructError{
		reason:     reason,
		structType: structType,
		fieldType:  fieldType,
	}
}
