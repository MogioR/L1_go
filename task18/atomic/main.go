package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
	Реализовать структуру-счетчик, которая будет инкрементироватьс
	я в конкурентной среде. По завершению программа должна
	выводить итоговое значение счетчика.

	Метод Sync Atomic
*/

type Counter struct {
	data int32
}

func (c *Counter) Get() int32 {
	return atomic.LoadInt32(&c.data)
}

func (c *Counter) Add() {
	atomic.AddInt32(&c.data, 1)
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
