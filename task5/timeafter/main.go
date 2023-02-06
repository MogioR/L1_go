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

	Метод через канал time.After
*/

func main() {
	// Указываем N соответствующего типа
	N := time.Duration(5)
	// Запрашиваем сигнал в канал timeout после прошествия n секунд.
	timeout := time.After(N * time.Second)

	nums := make(chan int32)

	// Пишем числа в канал
	go func() {
		for {
			select {
			// Закрываем канал по прошествию таймера
			case <-timeout:
				close(nums)
				return
			default:
				nums <- rand.Int31()
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()

	// Читаем из канала пока он открыт
	for num := range nums {
		fmt.Printf("%d\n", num)
	}
	fmt.Printf("Time end")
}
