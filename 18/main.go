package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	i  int
	mu sync.Mutex
}

func (c *Counter) Inc() {
	// захватываем мьютекс
	c.mu.Lock()
	// инкрементируем счетчик
	c.i++
	// отпускаем мьютекс
	c.mu.Unlock()
}

func main() {
	counter := Counter{}
	wg := sync.WaitGroup{}
	for i := 0; i < 124; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Inc()
		}()
	}
	wg.Wait()
	fmt.Printf("counter value: %d\n", counter.i)
}
