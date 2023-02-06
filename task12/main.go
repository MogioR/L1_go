package main

import "fmt"

/*
	Имеется последовательность строк - (cat, cat, dog,
	cat, tree) создать для нее собственное множество.
*/

func main() {
	words := []string{"cat", "cat", "dog", "cat", "tree"}
	wordsSet := make(map[string]bool)

	// Если слово не в множестве, то добовляем
	for _, word := range words {
		if !wordsSet[word] {
			wordsSet[word] = true
		}
	}

	//Вывод
	for key := range wordsSet {
		fmt.Printf("%s ", key)
	}
}
