package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	data, err := ReadTextFile("myconfig.yaml")
	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("Файл не найден")
		return
	}
	fmt.Println(data)
}

func ReadTextFile(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		// возвратим ошибку на русском языке
		return ``, fmt.Errorf(`не удалось прочитать файл (%s): %w`, filename, err)
	}
	return string(data), nil
}
