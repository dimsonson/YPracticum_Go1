package main
//тест использования функций в другом файле
import (
	"strconv"
	_"fmt"
)

type Stringer1 interface {
	String1() string
}

type myType1 int

// myType реализует интерфейс Stringer
func (t myType) String1() string { //(string, error) {
	//err := nil
	return strconv.Itoa(int(t)) //, err
	// представление типа myType в виде строки
}

/* func main() {
	a := myType(5)
	//var a myType = 5
	//a := new(myType) это инциализация структуры она здесь не работает
	//a = 5
	fmt.Printf("%T", a.String())
} */

/* type Stringer1 interface {
	String1() string
}

type myType struct {
	dig int
}

// myType реализует интерфейс Stringer
func (t myType) String1() string {
	return strconv.Itoa(t.dig)
	// представление типа myType в виде строки
}

func main() {
	a := new(myType)
	a.dig = 5
	fmt.Printf("%T", a.String1())
} */
