package main

import (
	"fmt"
	"strings"
)

/*
	Разработать программу, которая переворачивает слова в строке.
	Пример: «snow dog sun — sun dog snow».
*/

func wordsReverse(s string) string {
	// Разделяем по пробелам
	words := strings.Fields(s) // Ну или strings.Split(refString, " ")
	// Переворачиваем
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	// Собираем в строку
	return strings.Join(words, " ")
}

func main() {
	fmt.Printf("%s", wordsReverse("snow dog sun")) // Out: sun dog snow
}
