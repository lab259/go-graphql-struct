package main

import (
	"fmt"
	"reflect"
)

type Marshaler interface {
	MarshalKV() (string, error)
}

type User struct {
	Email   string `kv:"email,omitempty"`
	Name    string `kv:"name,omitempty"`
	Github  string `kv:"github,omitempty"`
	private string
}

func (u User) MarshalKV() (string, error) {
	return fmt.Sprintf("name=%s,email=%s,github=%s", u.Name, u.Email, u.Github), nil
}

func main() {
	fmt.Println(encode(User{"boring", "Ariel", "a8m", ""}))
	fmt.Println(encode(&User{Github: "posener", Name: "Eyal", Email: "boring"}))
}

var marshalerType = reflect.TypeOf(new(Marshaler)).Elem()

func encode(i interface{}) (string, error) {
	t := reflect.TypeOf(i)
	if !t.Implements(marshalerType) {
		return "", fmt.Errorf("encode only supports structs that implement the Marshaler interface")
	}
	m, _ := reflect.ValueOf(i).Interface().(Marshaler)
	return m.MarshalKV()
}
