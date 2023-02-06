package main

import "fmt"

/*
	Дана последовательность чисел: 2,4,6,8,10. Найти сумму их
	квадратов(2^2+3^2+4^2 …) с использованием конкурентных вычислений.
*/

// Горутина для конкурентного выполнения
func sqrGorutine(item int32, channel chan int32) {
	//Пишем результат в канал
	channel <- item * item
}

func getSumSqr(arr []int32) int32 {
	//Создаём канал и переменную для результата
	sum := int32(0)
	sqrs := make(chan int32)

	//Вызываем горутины
	for _, item := range arr {
		go sqrGorutine(item, sqrs)
	}

	//Ждем из канала ответы
	for i := 0; i < len(arr); i++ {
		sum = sum + <-sqrs
	}

	return sum
}

func main() {
	test := []int32{2, 4, 6, 8, 10}
	fmt.Printf("%d ", getSumSqr(test))
}
