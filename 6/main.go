package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// #1 "return"
	go func() {
		fmt.Println("g1 s")
		defer fmt.Println("g1 e")
		// code
		return
	}()

	// #2 "exit channel"
	quit := make(chan struct{})
	go func() {
		fmt.Println("g2 s")
		for {
			select {
			case <-quit:
				fmt.Println("g2 e")
			default:
				// code
			}
		}
	}()
	quit <- struct{}{}

	// #3 "closing channel for reading"
	ch := make(chan int)
	go func() {
		fmt.Println("g3 s")
		for data := range ch {
			fmt.Printf("val: %d", data)
			time.Sleep(time.Millisecond * 200)
		}
		fmt.Println("g3 e")
	}()
	go func() {
		time.Sleep(time.Second * 2)
		close(ch)
	}()

	// #4 "context/timer"
	ctx, _ := context.WithTimeout(context.Background(), time.Second*2)

	go func() {
		fmt.Println("g4 s")
		for {
			select {
			case <-ctx.Done():
				fmt.Println("g4 e")
				return
			}
		}
	}()

	time.Sleep(time.Second * 4)
}
