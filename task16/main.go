package main

import "fmt"

/*
	Реализовать быструю сортировку массива (quicksort) встроенными
	методами языка.
*/

func partition(arr []int, left, right int) int {
	pivot := arr[right] // Выбираем последний элемент опорным
	smallerPtr := left  // Указатель на конец подмассива меньших

	// Проходим по диапазону сортировки
	for i := left; i < right; i++ {
		//Если элемент меньше или равен опорному
		if arr[i] <= pivot {
			arr[smallerPtr], arr[i] = arr[i], arr[smallerPtr] // Переносим его в подмассив меньших (левую часть)
			smallerPtr += 1                                   // Сдвигаем указатель
		}
	}

	// Переносим опорный элемент в конец подмассива меньших
	// Так как он больше или равен всем в левой части, но меньше всех в правой
	arr[smallerPtr], arr[right] = arr[right], arr[smallerPtr]
	return smallerPtr
}

func quickSortRange(arr []int, left, right int) {
	if left < right {
		// Делим массив на 2 отсортированные части
		pivot := partition(arr, left, right)

		// Сортируем получившиеся части
		quickSortRange(arr, left, pivot-1)
		quickSortRange(arr, pivot+1, right)
	}
}

func quickSort(arr []int) {
	quickSortRange(arr, 0, len(arr)-1)
}

func main() {
	arr := []int{1, 6, 3, 7, 3, 7, 234, 76, 321, 63, 63, -3, 3, 0, 3, -34, 32, 34, 43}
	quickSort(arr)

	for _, val := range arr {
		fmt.Printf("%d ", val)
	}
}
