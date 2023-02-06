package main

import "fmt"

/*
	Разработать конвейер чисел. Даны два канала: в первый пишутся
	числа (x) из массива, во второй — результат операции x*2,
	после чего данные из второго канала должны выводиться в stdout.
*/

func main() {
	nums := make(chan int32)
	sqrs := make(chan int32)

	arr := [...]int32{1, 2, 3, 4, 5}

	// Горутина для чтения из массива
	go func(a []int32, nums chan int32) {
		for _, n := range a {
			nums <- n
		}
	}(arr[:], nums)

	// Горутина возведения в квадрат
	go func(nums chan int32, sqrs chan int32) {
		for n := range nums {
			sqrs <- n * n
		}
	}(nums, sqrs)

	for sqr := range sqrs {
		fmt.Printf("%d\n", sqr)
	}
}
