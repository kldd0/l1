package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := make(chan int)

	// запускаем в отдельной горутине запись в канал
	go func() {
		for {
			ch <- rand.Intn(100)
			time.Sleep(time.Millisecond * 300)
		}
	}()

	var d int
	fmt.Scan(&d)

	// устанавливаем таймер на заданное время
	timer := time.NewTicker(time.Second * time.Duration(d))

	for {
		select {
		case data := <-ch:
			fmt.Printf("val: %d\n", data)
		// после срабатывания таймера завершаем программу
		case <-timer.C:
			fmt.Println("Time is over")
			return
		}
	}
}
