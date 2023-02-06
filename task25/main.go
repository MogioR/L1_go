package main

import (
	"fmt"
	"time"
)

/*
	Реализовать собственную функцию sleep.
*/

// Занимает поток, но работает
func dudleSleep(duration time.Duration) {
	start := time.Now()
	for time.Since(start) < duration {

	}
}

// Работает не занимая поток
func smartSleep(duration time.Duration) {
	<-time.After(duration)
}

func main() {
	now := time.Now()
	dudleSleep(2 * time.Second)
	fmt.Println(time.Since(now))

	now = time.Now()
	smartSleep(2 * time.Second)
	fmt.Println(time.Since(now))
}
