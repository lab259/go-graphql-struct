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
