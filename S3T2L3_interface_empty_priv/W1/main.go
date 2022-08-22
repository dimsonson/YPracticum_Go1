package main

import (
	"fmt"
)

func Mul(a any, b int) any {
	switch a.(type) {
	case int:
		return a.(int) * b
	case string:
		var s string
		for i := 0; i < b; i++ {
			s += a.(string) + " "
		}
		return s
	default:
		return "Неверный тип параметра"
	}
}

func main() {
	fmt.Println(Mul(3, 5))
	fmt.Println(Mul("три", 5))
	fmt.Println(Mul(true, 5))
}

/* func Mul(a interface{}, b int) interface{} {
	switch va := a.(type) {
	case int:
		return va * b
	case string:
		return strings.Repeat(va, b)
	case fmt.Stringer:
		return strings.Repeat(va.String(), b)
	default:
		return nil
	}
}
*/
