package main

import (
	"fmt"
	"sync"
)

/*
	Написать программу, которая конкурентно рассчитает значение
	квадратов чисел взятых из массива (2,4,6,8,10) и выведет
	их квадраты в stdout.

	Метод с использованием груп ожидания.
*/

// Горутина для конкурентного выполнения
func sqrGorutine(item int32, result *int32, wg *sync.WaitGroup) {
	// Вычисляем значение
	*result = item * item

	// Сообщаем группе о готовности
	defer wg.Done()
}

func mapSqr(arr []int32) []int32 {
	// Обьявляем группу ожидания и задаем ей размер равный массиву
	var wg sync.WaitGroup
	wg.Add(len(arr))
	result := make([]int32, len(arr))

	// Создаём горутины для вычисления каждого элемента, сообщая им куда возвращать результат во избежание конфликтов (и сохранения порядка)
	for index, item := range arr {
		go sqrGorutine(item, &result[index], &wg)
	}

	// Ждем выполнения
	wg.Wait()
	return result
}

func main() {
	test := []int32{2, 4, 6, 8, 10}
	out := mapSqr(test)

	for _, item := range out {
		fmt.Printf("%d ", item)
	}
}
