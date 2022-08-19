package main

import (
	"fmt"

	"github.com/dimsonson/YPracticum_Go1/S2T3L4_extern_depend/summodule"
)

func main() {
	if sum := summodule.Add(1, 2); sum != 3 {
		panic(fmt.Sprintf("sum expected to be 3; got %d", sum))
	}

	fmt.Println("Well done!")
}
