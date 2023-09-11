package main

import (
	"cmp"
	"fmt"
)

func qSort[T cmp.Ordered](arr []T, l, r int) {
	if l < r {
		q := partition(arr, l, r)
		qSort(arr, l, q)
		qSort(arr, q+1, r)
	}
}

func partition[T cmp.Ordered](arr []T, l, r int) int {
	// опорный элемент
	v := arr[(l+r)/2]
	i := l
	j := r
	// проходим индексами по массиву с разных сторон,
	// пока они не перепрыгнут друг друга
	for i <= j {
		// сдвигаем левый индекс, пока не найдем элемент > опорного
		for arr[i] < v {
			i++
		}
		// сдвигаем правый индекс, пока не найдем элемент < опорного
		for arr[j] > v {
			j--
		}
		// обмен элементов должен происходить только если они по разные стороны
		// от опорного элемента => если они перешагивают друг друга,
		// то выходим из цикла
		if i >= j {
			break
		}
		// меняем местами элементы
		arr[i], arr[j] = arr[j], arr[i]
		// двигаем указатели дальше по направлению
		i++
		j--
	}
	return j
}

func main() {
	arr := []int{1, 5, 2, 3, 7, 9, 2, 5}

	qSort[int](arr, 0, len(arr)-1)
	fmt.Println(arr)
}
