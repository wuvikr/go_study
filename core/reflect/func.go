package main

import (
	"fmt"
	"reflect"
)

func Add(a string, b int) int {
	return 1
}

func main() {
	typeFunc := reflect.TypeOf(Add)
	argInNum := typeFunc.NumIn()
	for i := 0; i < argInNum; i++ {
		argType := typeFunc.In(i)
		fmt.Printf("第%d个参数类型 argType: %v\n", i+1, argType)
	}
}
