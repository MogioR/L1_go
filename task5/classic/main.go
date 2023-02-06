package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
	Разработать программу, которая будет последовательно
	отправлять значения в канал, а с другой стороны канала
	— читать. По истечению N секунд программа должна
	завершаться.

	Метод через рассчет времени
*/

func main() {
	// Указываем N соответствующего типа
	N := time.Duration(5)
	timeDuration := time.Duration(N * time.Second)
	nums := make(chan int32)

	// Пишем числа в канал
	go func(duration time.Duration) {
		timeStart := time.Now()

		for {
			// Считаем время прошедшее с момента начала горутины
			if time.Since(timeStart) < duration {
				nums <- rand.Int31()
				time.Sleep(100 * time.Millisecond)
			} else {
				close(nums)
				return
			}
		}
	}(timeDuration)

	// Читаем из канала пока он открыт
	for num := range nums {
		fmt.Printf("%d\n", num)
	}
	fmt.Printf("Time end")
}
