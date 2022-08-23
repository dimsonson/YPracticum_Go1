package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

/*

Напишите функцию MyCheck(input string) error,
которая должна проверять строку на соответствие следующим параметрам:

Строка не должна содержать цифр (found numbers).

Длина должна быть меньше 20 символов (line is too long).

Строка должна иметь два пробела (no two spaces).

В скобках приведены тексты возвращаемых ошибок.
Функция должна находить все возможные ошибки за один вызов.

Результат должен содержать слайс ошибок, по которым строка не прошла проверку, или быть nil.

Подсказка: определите свой тип для слайса ошибок.

*/

// 1) вставьте определение типа для []error
// 2) определите метод Error для вашего типа, который будет выводить
//    все ошибки слайса
// 3) реализуйте функцию MyCheck
//
// ...

func main() {
	for {
		fmt.Printf("Укажите строку (q для выхода): ")
		reader := bufio.NewReader(os.Stdin)
		ret, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			continue
		}
		ret = strings.TrimRight(ret, "\n")
		if ret == `q` {
			break
		}
		if err = MyCheck(ret); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(`Строка прошла проверку`)
		}
	}
}

func MyCheck(input string) error {
	if !strings.Contains(input, " ") {
		return fmt.Errorf("дава пробела")
	}

	if len(input) >= 20 {
		return fmt.Errorf("меньше 20 символов")
	}

	for _, s := range input {
		if unicode.IsDigit(s) {
			return fmt.Errorf("цифра")
		}
	}
	return nil
}

type MyError struct {
	SlErr []error
}

func (sl MyError) Error() string {
	return fmt.Sprintf("%v: %v", te.Time.Format(`2006/01/02 15:04:05`), te.Text)
}
