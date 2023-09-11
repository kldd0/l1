package main

import (
	"fmt"
	"unicode/utf8"
)

func reverse(s string) string {
	s_b := []rune(s)
	for i, j := 0, len(s_b)-1; i < j; i, j = i+1, j-1 {
		s_b[i], s_b[j] = s_b[j], s_b[i]
	}
	return string(s_b)
}

func main() {
	s := "главрыба"
	s1 := "qwerty"

	sc := utf8.RuneCountInString(s)
	s1c := utf8.RuneCountInString(s1)

	s_r := make([]rune, 0, sc)
	for _, r := range s {
		s_r = append(s_r, r)
	}

	s_r1 := make([]rune, 0, s1c)
	for _, r := range s1 {
		s_r1 = append(s_r1, r)
	}

	rs := reverse(string(s_r))
	rs1 := reverse(string(s_r1))

	fmt.Printf("%s -- %s\n", s, rs)
	fmt.Printf("%s -- %s\n", s1, rs1)
}
