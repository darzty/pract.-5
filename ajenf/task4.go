package main

import (
	"fmt"
	"time"
)

func main() {
	highPriority := make(chan string, 10)
	lowPriority := make(chan string, 10)
	done := make(chan struct{})

	go func() {
		for {
			select {
			case task := <-highPriority:
				fmt.Println("Срочная задача:", task)
			default:
				select {
				case task := <-highPriority:
					fmt.Println("Срочная задача:", task)
				case task := <-lowPriority:
					fmt.Println("Обычная задача:", task)
				case <-done:
					return
				}
			}
		}
	}()

	lowPriority <- "Обычный заказ 1"
	lowPriority <- "Обычный заказ 2"
	highPriority <- "СРОЧНЫЙ заказ 1"
	lowPriority <- "Обычный заказ 3"
	highPriority <- "СРОЧНЫЙ заказ 2"

	time.Sleep(500 * time.Millisecond)
	close(done)
}