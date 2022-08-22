package main

import (
	"fmt"
	"reflect"
)

type MyType struct{}

func NaiveIsNil(obj interface{}) bool {
	return obj == nil
}

func main() {
	var t *MyType
	fmt.Printf("Проверка типа (%v) на nil: %v\n", reflect.TypeOf(t), NaiveIsNil(t)) // TypeOf возвращает тип переданного объекта.
}
