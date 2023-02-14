package main

import (
	"fmt"
	"sync"
)

/*
	Реализовать структуру-счетчик, которая будет инкрементироватьс
	я в конкурентной среде. По завершению программа должна
	выводить итоговое значение счетчика.

	Метод Mutex
*/

type Counter struct {
	data  int32
	mutex sync.RWMutex
}

func (c *Counter) Get() int32 {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.data
}

func (c *Counter) Add() {
	c.mutex.Lock()
	c.data += 1
	c.mutex.Unlock()
}

func main() {
	counter := Counter{}
	wg := sync.WaitGroup{}

	N := 1000

	wg.Add(N)
	for i := 0; i < N; i++ {
		go func(c *Counter, wg *sync.WaitGroup) {
			c.Add()
			wg.Done()
		}(&counter, &wg)
	}

	wg.Wait()
	fmt.Println(counter.Get())
}
