package main

import "fmt"

func main() {
	a, b := 1, 2

	fmt.Printf("a=%d [%p], b=%d [%p]\n", a, &a, b, &b)

	b, a = a, b

	fmt.Printf("a=%d [%p], b=%d [%p]\n", a, &a, b, &b)
}
