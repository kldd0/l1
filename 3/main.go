package main

import (
	"fmt"
)

func main() {
	arr := []int{2, 4, 6, 8, 10}
	// канал нужен для синхронизации результатов работы горутин
	out := make(chan int, len(arr))
	for _, el := range arr {
		go func(num int, out chan<- int) {
			out <- num * num
		}(el, out)
	}

	res := 0
	for i := 0; i < len(arr); i++ {
		res += <-out
	}
	close(out)
	fmt.Println(res)
}
