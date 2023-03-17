package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Gender byte   `json:"gender"`
}

func (User) F1()             {}
func (User) f2()             {}
func (*User) F3()            {}
func (User) F4(s string) int { return 1 }

func typeFieldInfo() {
	u := User{}
	t := reflect.TypeOf(u)
	for i := 0; i < t.NumField(); i++ {
		sf := t.Field(i)
		fmt.Printf("sf.Name: %v\n", sf.Name)
		fmt.Printf("sf.Tag.Get(\"json\"): %v\n", sf.Tag.Get("json"))
		fmt.Printf("sf.Offset: %v\n", sf.Offset)
	}
	fmt.Printf("t.Size(): %v\n", t.Size())
}

func typeMethodInfo() {
	typeUser := reflect.TypeOf(User{})
	methodCount := typeUser.NumMethod()
	for i := 0; i < methodCount; i++ {
		method := typeUser.Method(i)
		fmt.Printf("method.Name: %v\n", method.Name)
		fmt.Printf("method.Type: %v\n", method.Type)
	}
}

func main() {
	// typeFieldInfo()
	typeMethodInfo()
}
