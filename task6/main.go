package main

import (
	"context"
	"fmt"
	"os"
	"runtime"
	"time"
)

/*
	Реализовать все возможные способы остановки выполнения горутины.
*/

func main() {

	// 1. Конец выполнения (или вызов return)
	go func() {
	}()

	// 2. Сообщение о закрытии через канал (или закрытие канала с данными)
	exit := make(chan bool)
	go func(exit chan bool) {
		for {
			select {
			case <-exit:
				return
			default:
				time.Sleep(500 * time.Millisecond)
			}
		}
	}(exit)

	exit <- true

	// 3. Отмена контекста
	//Обьявляем контекст с возможностью отмены
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				time.Sleep(500 * time.Millisecond)
			}
		}
	}(ctx)

	time.Sleep(3 * time.Second)
	cancel()

	// 4. runtime.goexit - Останавливает данную рутину не трогая остальны

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("%d ", i)
		}
	}()

	runtime.Goexit()
	// Out:
	// 1 2 3 4 5
	// DEADLOCK

	// 5. Закрытие потока выполнения из корого запущена горутина
	os.Exit(0)

}
