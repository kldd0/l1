package main

import (
	"fmt"
	"strings"
)

func reverse(s string) string {
	s_arr := strings.Split(s, " ")
	for i, j := 0, len(s_arr)-1; i < j; i, j = i+1, j-1 {
		s_arr[i], s_arr[j] = s_arr[j], s_arr[i]
	}
	return strings.Join(s_arr, " ")
}

func main() {
	s := "snow dog sun"
	rs := reverse(s)
	fmt.Printf("%s -- %s\n", s, rs)
}
