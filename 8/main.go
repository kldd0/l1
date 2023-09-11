package main

import "fmt"

func SetBit(num int64, i int, bVal int) int64 {
	var bitMask int64 = 1 << i

	if bVal == 0 {
		// &^ - AND NOT op => bitMask инвертируется
		// 0 -> 1, 1 -> 0, дальше операция AND
		num &^= bitMask
	} else {
		// операция OR, вне зависимости от значения бита, он будет 1
		num |= bitMask
	}

	return num
}

func main() {
	var num int64
	fmt.Printf("num: ")
	fmt.Scan(&num)
	fmt.Printf("bits: %b\n", num)

	var i int
	fmt.Printf("i: ")
	fmt.Scan(&i)

	var bVal int
	fmt.Printf("bit value: ")
	fmt.Scan(&bVal)

	num = SetBit(num, i, bVal)

	fmt.Printf("val: %d, bits: %b\n", num, num)
}
