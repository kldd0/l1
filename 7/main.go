package main

import (
	"fmt"
	"sync"
)

func main() {
	// concurrent write in map
	m := make(map[int]struct{})
	mapMutex := sync.Mutex{}
	wg := sync.WaitGroup{}

	wg.Add(10)
	for i := 0; i < 10; i++ {
		i := i
		go func() {
			mapMutex.Lock()
			defer wg.Done()
			defer mapMutex.Unlock()
			m[i] = struct{}{}
		}()
	}
	wg.Wait()

	fmt.Printf("map: %+v", m)
}
