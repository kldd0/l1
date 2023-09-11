package main

import "fmt"

func main() {
	seq := []string{"cat", "cat", "dog", "cat", "tree"}

	set := make(map[string]struct{})

	for _, v := range seq {
		set[v] = struct{}{}
	}

	fmt.Println(set)
}
