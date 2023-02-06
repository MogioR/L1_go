package main

import "fmt"

/*
	Реализовать бинарный поиск встроенными методами языка.
*/

func binarySearch(arr []int, el int) int {
	left := 0
	right := len(arr) - 1

	//Пока есть элементы
	for left <= right {
		// Ищем середину
		pos := int((right + left) / 2)
		if arr[pos] == el {
			return pos // Если нашли элемент возвращаем позицию
		} else if arr[pos] > el {
			right = pos - 1 // Ищем среди меньшей половины (смещаем границу влево)
		} else {
			left = pos + 1 // Ищем среди большей половины (смещаем границу вправо)
		}
	}
	return -1 // Элемент не найден
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 5, 6, 7, 7, 7, 10, 11, 11, 23, 43, 23, 122}
	pos := binarySearch(arr, 7)

	fmt.Printf("%d", pos)
}
