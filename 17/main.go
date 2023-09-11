package main

import (
	"fmt"
)

func binSearch(arr []int, x int) int {
	// задаем границы поиска
	l := 0
	r := len(arr) - 1
	//
	for l <= r {
		// вычисляем индекс серединного элемента между границами
		m := (l + r) / 2
		// проверяем, если нашли нужный элемент
		if arr[m] == x {
			return m
		}
		// если число, которое ищем больше середины => ищем в правой
		if arr[m] < x {
			l = m + 1
		}
		// если число, которое ищем меньше середины => ищем в левой
		if arr[m] > x {
			r = m - 1
		}
	}
	// если не нашли, возвращаем -1
	return -1
}

func main() {
	arr := []int{1, 2, 3, 3, 6, 7, 9, 21, 62}

	target := 9
	res := binSearch(arr, target)
	if res != -1 {
		fmt.Printf("target: %d res: %d res_index: %d\narr: %+v\n", target, arr[res], res, arr)
	} else {
		fmt.Println("элемент не найден")
	}
}
