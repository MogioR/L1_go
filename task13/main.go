package main

import "fmt"

/*
	Поменять местами два числа без создания временной переменной.
*/

func main() {
	// Классика
	a := 4
	b := 8

	a = a + b
	b = a - b
	a = a - b

	fmt.Printf("%d %d\n", a, b) // Out: 8 4

	// Python style
	a = 4
	b = 8

	a, b = b, a

	fmt.Printf("%d %d\n", a, b) // Out: 8 4

	// Классика через умножение (с 0 не работает)
	a = 4
	b = 8

	a = a * b
	b = a / b
	a = a / b

	fmt.Printf("%d %d\n", a, b) // Out: 8 4

	// Xor
	a = 4
	b = 8

	a = a ^ b
	b = a ^ b
	a = a ^ b

	fmt.Printf("%d %d\n", a, b) // Out: 8 4
}
