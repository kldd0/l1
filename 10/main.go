package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	seq := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	sort.Float64s(seq)

	groups := make(map[int][]float64)
	// возьмем минимальный элемент, так как отсортировали посл-ть => он первый
	// дальше получим число, объединяющее группу, чтобы сверять отклонение
	// если отклонение становится >= 10 => пересчитываем заглавное число группы
	// повторяем алгоритм
	mn := float64(10 * int(seq[0]/10))
	for _, v := range seq {
		if math.Abs(v-mn) >= 10 {
			// обновление числа группы
			mn = float64(10 * int(v/10))
		}
		groups[int(mn)] = append(groups[int(mn)], v)
	}
	fmt.Println(groups)
}
