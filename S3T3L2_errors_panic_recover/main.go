package main

import "fmt"

type input struct {
	a       int
	b       int
	counter int
}

// 1) добавьте функцию f
// 2) добавьте функцию test(a, b, counter int) (err error),
// которая преобразует панику в ошибку
// ...

func main() {
	for _, pars := range []input{
		{10, 5, 3},
		{100, 7, 10},
		{1, 1, 1000},
	} {
		fmt.Printf("(%d, %d, %d) => %v\r\n", pars.a, pars.b, pars.counter, test(pars.a, pars.b, pars.counter))
	}
}

func test(a, b, counter int) (err error) {

	return err
}
func f(a1, b1, counter1 int) (a, b, counter int) {
	return b1, a1 / b1, counter1 - 1
}

/*
Напишите рекурсивную функцию,
которая преобразует параметры следующим образом: f(a,b,i) → f(b,a/b,i-1).
a, b — произвольные целые числа, i — счётчик шагов.

При i, равном 0, нужно вызвать панику.

Преобразуйте аварийную ситуацию в ошибку, которая сообщает,
какая аварийная ситуация произошла: паника или деление на ноль.

Проверьте работу на следующих вариантах (a,b,i):
(10, 5, 3),
(100, 7, 10),
(1, 1, 1000).

Программа должна вывести:
(10, 5, 3) => counter equals 0
(100, 7, 10) => runtime error: integer divide by zero
(1, 1, 1000) => counter equals 0

*/
