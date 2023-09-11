package main

import (
	"fmt"
	"time"
)

func sleep(d time.Duration) {
	for {
		select {
		case <-time.After(d * time.Second):
			return
		}
	}
}

func main() {
	fmt.Println("program started", time.Now())
	sleep(3)
	fmt.Println("program finished", time.Now())
}
