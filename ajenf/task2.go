package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan struct{})

	go func() {
		for {
			select {
			case <-done:
				fmt.Println("горутина остановлена")
				return
			default:
				fmt.Println("сообщение раз в секунду")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	time.Sleep(3 * time.Second)
	done <- struct{}{}
	time.Sleep(100 * time.Millisecond)
}