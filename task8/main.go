package main

import "fmt"

/*
	Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.
*/

func setBit(a int64, pos uint, value bool) int64 {
	// Основы логики...
	if value {
		return a | (1 << pos)
	} else {
		return a &^ (1 << pos)
	}
}

func main() {
	fmt.Printf("%b\n", setBit(int64(16), 2, true))  // Out: 10100
	fmt.Printf("%b\n", setBit(int64(16), 4, false)) // Out: 0
}
