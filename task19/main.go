package main

import "fmt"

/*
	Разработать программу, которая переворачивает по
	даваемую на ход строку (например: «главрыба — абырвалг»).
	Символы могут быть unicode.
*/

// Плохо работает с языками где один символ записывается нескольким рунами, например с руной '
func stringReverse(s string) string {
	runes := []rune(s)

	//Меняем руны местами
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	fmt.Printf("%s", stringReverse("Abbat 😀"))
}
