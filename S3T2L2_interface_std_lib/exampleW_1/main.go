package main

import (
	"fmt"
	"time"

	"github.com/dimsonson/YPracticum_Go1/S3T2L2_interface_std_lib/example/randbyte"
)

func main() {

	// создаём генератор случайных чисел
	generator := randbyte.New(time.Now().UnixNano()) // в качестве затравки передаём ему
	//текущее время, и при каждом запуске оно будет разным.

	buf := make([]byte, 16)

	for i := 0; i < 5; i++ {
		n, _ := generator.Read(buf) // единственный доступный метод, но он нам и нужен.
		fmt.Printf("Generate bytes: %v size(%d)\n", buf, n)
	}

}
