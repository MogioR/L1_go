package main

import "fmt"

/*
	Реализовать пересечение двух неупорядоченных множеств.
*/

func intersection(a, b map[int32]bool) map[int32]bool {
	result := make(map[int32]bool)
	// Если элемент в обоих множествах то добовляем его в результат
	for key := range a {
		if b[key] {
			result[key] = true
		}
	}
	return result
}

func main() {
	// Множество цифр
	a := make(map[int32]bool)
	for i := int32(0); i < 10; i++ {
		a[i] = true
	}

	// Множество четных чисел от 0 до 20
	b := make(map[int32]bool)
	for i := int32(0); i < 20; i += 2 {
		b[i] = true
	}

	c := intersection(a, b)

	//Вывод
	for key := range c {
		fmt.Printf("%d ", key)
	}
}
