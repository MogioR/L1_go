package main

import (
	"fmt"
	"sync"
)

/*
	Реализовать конкурентную запись данных в map.

	Метод mutex
*/

func main() {
	N := 5

	testMap := make(map[int]string)
	mapMutex := sync.Mutex{}

	wg := sync.WaitGroup{}
	wg.Add(N)
	for i := 0; i < N; i++ {
		go func(i int, wg *sync.WaitGroup, mutex *sync.Mutex) {
			//Создаём критическую секцию
			//Блокируем выполнение
			mutex.Lock()
			// Пишем в карту
			testMap[i] = "item: " + fmt.Sprint(i)
			//Разрешаем выполнение
			mutex.Unlock()
			wg.Done()
		}(i, &wg, &mapMutex)
	}
	// Ожидаем выполнения
	wg.Wait()

	for key, value := range testMap {
		fmt.Printf("%d: %s\n", key, value)
	}
}
