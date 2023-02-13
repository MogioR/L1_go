package main

import "fmt"

/*
	Имеется последовательность строк - (cat, cat, dog,
	cat, tree) создать для нее собственное множество.
*/

func main() {
	words := []string{"cat", "cat", "dog", "cat", "tree"}
	wordsSet := make(map[string]struct{})

	// Если слово не в множестве, то добовляем
	for _, word := range words {
		if _, ok := wordsSet[word]; !ok {
			wordsSet[word] = struct{}{}
		}
	}

	//Вывод
	for key := range wordsSet {
		fmt.Printf("%s ", key)
	}
}
