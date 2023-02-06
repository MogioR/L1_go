package main

import (
	"fmt"
	"strings"
)

/*
	К каким негативным последствиям может привести данный
	фрагмент кода, и как это исправить? Приведите корректный
	пример реализации.

	var justString string
	func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100]
	}

	func main() {
	someFunc()
	}

	Проблемы:
	В данном отрезке есть 2 проблемы:
	* утечка памяти (используется не вся выделенная в функции память)
	* используется глобальная переменная
*/

func createHugeString(size int) string {
	return strings.Repeat("F", size)
}

// Вместо глобальной переменной строку можно присваивать или возвращать по ссылке
func someFunc() string {
	v := createHugeString(1 << 10)
	// Так как мы больше не ссылаемся на v, её можно удалить
	return string([]byte(v)[:100])
}

func main() {
	val := someFunc()
	fmt.Printf("%s", val)
}
