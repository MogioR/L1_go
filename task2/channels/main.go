package main

import "fmt"

/*
	Написать программу, которая конкурентно рассчитает значение
	квадратов чисел взятых из массива (2,4,6,8,10) и выведет
	их квадраты в stdout.

	Метод с использованием каналов.
*/

// Горутина для конкурентного выполнения
func sqrGorutine(item int32, channel chan int32) {
	//Пишем результат в канал
	channel <- item * item
}

func mapSqr(arr []int32) []int32 {
	//Создаём канал и слайс для результата
	result := make([]int32, 0)
	sqrs := make(chan int32)

	//Вызываем горутины
	for _, item := range arr {
		go sqrGorutine(item, sqrs)
	}

	//Ждем из канала ответы
	for i := 0; i < len(arr); i++ {
		result = append(result, <-sqrs)
	}

	return result
}

func main() {
	test := []int32{2, 4, 6, 8, 10}
	out := mapSqr(test)

	for _, item := range out {
		fmt.Printf("%d ", item)
	}
}
