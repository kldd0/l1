package main

import (
	"errors"
	"fmt"
)

func SetBit(num int64, i int, bVal int) (int64, error) {
    if i >= 64 {
        return 0, errors.New("Invalid bit index, out of range [0..64]")
    }

    if i == 63 && bVal == 1 {
        return 0, errors.New("Integer overflow")
    }

	var bitMask int64 = 1 << i

	if bVal == 0 {
		// &^ - AND NOT op => bitMask инвертируется
		// 0 -> 1, 1 -> 0, дальше операция AND
		num &^= bitMask
	} else {
		// операция OR, вне зависимости от значения бита, он будет 1
		num |= bitMask
	}

	return num, nil
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

    if bVal < 0 || bVal > 1 {
        fmt.Printf("Error: bVal required be in [0, 1]\n")
        return
    }

    num, err := SetBit(num, i, bVal)
    if err != nil {
        fmt.Printf("Error: %s\n", err)
        return
    }

	fmt.Printf("val: %d, bits: %b\n", num, num)
}
