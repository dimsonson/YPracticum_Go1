package main

import (
	"bufio"
	"fmt"
	"os"
)

//cat test.txt | go run test.go

func main() {
	//var a int
	s := make([]int, 10)
	now := bufio.NewReader(os.Stdin)
	//now.Scan()
	//r := now.Text()
	n, err := fmt.Fscan(now, s[0:3])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fscanf: %v\n", err)
	}
	fmt.Println(n)
	fmt.Println(s)

	/* for {
		fmt.Scan(&a)
		s = append(s, a)
		fmt.Println(s)
	} */
}

//

// переписать на for ******************************

/* package main

import "fmt"

func input(x []int, err error) []int {
	if err != nil {
		return x
	}
	var d int
	n, err := fmt.Scanf("%d", &d)
	if n == 1 {
		x = append(x, d)
	}
	return input(x, err)
}

func main() {
	fmt.Println("Enter input:")
	x := input([]int{}, nil)
	fmt.Println("Input:", x)
} */
