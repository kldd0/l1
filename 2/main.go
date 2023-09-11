package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := []int{2, 4, 6, 8, 10}
	// WaitGroup нужна для ожидания завершения работы всех горутин,
	// иначе программа может завершиться раньше
	wg := sync.WaitGroup{}
	wg.Add(len(arr))
	for _, el := range arr {
		go func(num int) {
			fmt.Println(num * num)
			wg.Done()
		}(el)
	}
	wg.Wait()
}
