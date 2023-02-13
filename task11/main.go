package main

import "fmt"

/*
	Реализовать пересечение двух неупорядоченных множеств.
*/

func intersection(a, b map[int32]struct{}) map[int32]struct{} {
	result := make(map[int32]struct{})
	// Если элемент в обоих множествах то добовляем его в результат
	for key := range a {
		if _, ok := b[key]; ok {
			result[key] = struct{}{}
		}
	}
	return result
}

func main() {
	// Множество цифр
	a := make(map[int32]struct{})
	for i := int32(0); i < 10; i++ {
		a[i] = struct{}{}
	}

	// Множество четных чисел от 0 до 20
	b := make(map[int32]struct{})
	for i := int32(0); i < 20; i += 2 {
		b[i] = struct{}{}
	}

	c := intersection(a, b)

	//Вывод
	for key := range c {
		fmt.Printf("%d ", key)
	}
}
