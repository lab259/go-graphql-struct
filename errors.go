package gqlstruct

import (
	"fmt"
	"reflect"
)

type typeNotRecognizedError struct {
	t reflect.Type
}

func (err *typeNotRecognizedError) Error() string {
	return fmt.Sprintf("'%s' not recognized", err.t)
}

func NewErrTypeNotRecognized(t reflect.Type) error {
	return &typeNotRecognizedError{
		t: t,
	}
}

type TypeNotRecognizedWithStructError struct {
	reason      error
	structType  reflect.Type
	fieldStruct reflect.StructField
}

func (err *TypeNotRecognizedWithStructError) Error() string {
	return fmt.Sprintf("%s.%s:%s", err.structType.Name(), err.fieldStruct.Name, err.reason.Error())
}

func NewErrTypeNotRecognizedWithStruct(reason error, structType reflect.Type, structField reflect.StructField) error {
	return &TypeNotRecognizedWithStructError{
		reason:      reason,
		structType:  structType,
		fieldStruct: structField,
	}
}
