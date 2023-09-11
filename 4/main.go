package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func worker(id int, in <-chan int, wg *sync.WaitGroup) {
	for d := range in {
		fmt.Printf("[worker %d] message: %d\n", id, d)
	}
	wg.Done()
}

func main() {
	var workerNum int
	fmt.Print("number of workers: ")
	fmt.Scan(&workerNum)
	// WaitGroup нужна для ожидания завершения всех горутин после завершения
	var wg sync.WaitGroup
	wg.Add(workerNum)
	// канал, в который поступают данные
	in := make(chan int)
	// запуск воркеров
	for i := 1; i <= workerNum; i++ {
		go worker(i, in, &wg)
	}
	// отдельный канал для перехвата SIGINT
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	// генерация рандомных данных в канал
	go func() {
		// timer := time.NewTicker(time.Millisecond*200)

		for {
			in <- rand.Intn(100)
			time.Sleep(time.Millisecond * 200)
		}
	}()
	// ожидание завершения программы
	<-sigint
	// закрытие канала
	close(in)
	// ожидание завершения работы всех горутин
	wg.Wait()
	fmt.Println("stopped")
}
