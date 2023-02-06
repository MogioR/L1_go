package main

import (
	"fmt"
	"sync"
)

/*
	Реализовать конкурентную запись данных в map.

	Метод sync map
*/

func main() {
	N := 5

	testMap := sync.Map{}

	wg := sync.WaitGroup{}
	wg.Add(N)
	for i := 0; i < N; i++ {
		go func(i int, wg *sync.WaitGroup) {
			// Так как эта имплементация потоко безопасна просто пишем в карту
			testMap.Store(i, "item: "+fmt.Sprint(i))
			wg.Done()
		}(i, &wg)
	}

	wg.Wait()

	testMap.Range(func(key, value interface{}) bool {
		fmt.Printf("%d: %s\n", key, value)
		return true
	})
}
