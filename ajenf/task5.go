package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, stop <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-stop:
			fmt.Printf("Воркер %d завершил текущий запрос\n", id)
			return
		default:
			fmt.Printf("Воркер %d обрабатывает входящий запрос\n", id)
			time.Sleep(300 * time.Millisecond)
		}
	}
}

func main() {
	var wg sync.WaitGroup
	stop := make(chan struct{})

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, stop, &wg)
	}

	time.Sleep(1 * time.Second)
	fmt.Println("Сигнал завершения: ожидание закрытия активных запросов...")
	close(stop)

	wg.Wait()
	fmt.Println("Все горутины корректно остановлены. Сервер завершил работу.")
}