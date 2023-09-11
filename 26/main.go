package main

import "fmt"

func containsOnlyUniqueSymbols(s string) bool {
	// используем map для проверки на уникальность
	m := make(map[rune]struct{}, len(s))

	for _, r := range s {
		if _, ok := m[r]; ok {
			return false
		}
		m[r] = struct{}{}
	}
	return true
}

func main() {
	s1 := "abcd"
	s2 := "abCdefAaf"
	s3 := "aabcd"

	fmt.Printf("%s -- %t\n", s1, containsOnlyUniqueSymbols(s1))
	fmt.Printf("%s -- %t\n", s2, containsOnlyUniqueSymbols(s2))
	fmt.Printf("%s -- %t\n", s3, containsOnlyUniqueSymbols(s3))
}
