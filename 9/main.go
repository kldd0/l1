package main

import "fmt"

func readFromArr(arr []int, c chan<- int) {
	for _, n := range arr {
		c <- n
	}
}

func mulByTwo(in <-chan int, out chan<- int) {
	for val := range in {
		out <- val * 2
	}
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}

	inp := make(chan int)
	out := make(chan int)

	// первая функция для переноса чисел из массива в канал
	// запуск в таком виде для соблюдения "The Channel Closing Principle"
	go func() {
		readFromArr(arr, inp)
		close(inp)
	}()
	// аналогично, функция для переноса данных из 1го канала во 2й,
	// которая попутно осуществляет удваивание
	go func() {
		mulByTwo(inp, out)
		close(out)
	}()

	for val := range out {
		fmt.Println(val)
	}
}
