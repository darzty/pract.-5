package main

import (
	"fmt"
)

func main() {
	ch := make(chan string, 1)

	select {
	case msg := <-ch:
		fmt.Println("Получено:", msg)
	default:
		fmt.Println("Данных нет")
	}
}