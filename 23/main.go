package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(arr, cap(arr))

	var i int
	fmt.Scan(&i)
	arr = append(arr[:i], arr[i+1:]...)
	fmt.Println(arr, cap(arr))
	arr = append(arr, 13)
	fmt.Println(arr, cap(arr))
}
