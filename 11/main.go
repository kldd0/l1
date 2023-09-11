package main

import "fmt"

func intersection[T comparable](s1, s2 []T) []T {
	// создаем мн-во с уникальными значениями
	set := make(map[T]struct{})

	// добавляем элементы из 1го массива в мн-во
	for _, v := range s1 {
		set[v] = struct{}{}
	}

	res := make([]T, 0)
	// перебираем все элементы 2го массива, сверяем на существование во мн-ве
	for _, v := range s2 {
		if _, ok := set[v]; ok {
			res = append(res, v)
		}
	}

	return res
}

func main() {
	s1 := []int{4, 6, 7, 1, 11, 14}
	s2 := []int{5, 9, 14, 12, 1, 7}

	fmt.Println(intersection(s1, s2))
}
