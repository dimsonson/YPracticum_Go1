package main

import "fmt"

type MyType struct {
	IntField int
	StrField string
	PtrField *float64 //false из за указателей, их адреса разные
}

func (mt MyType) IsEqual(mt2 MyType) bool {
	return mt == mt2
}

func main() {
	floatValue1, floatValue2 := 10.0, 10.0
	a := MyType{IntField: 1, StrField: "str", PtrField: &floatValue1}
	b := MyType{IntField: 1, StrField: "str", PtrField: &floatValue2}

	fmt.Printf("Равенство a и b: %v\n", a.IsEqual(b))
}
