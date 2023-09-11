package main

import (
	"fmt"
	"math/big"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	a := big.NewInt(0).Exp(big.NewInt(2), big.NewInt(rand.Int63n(100)+20), nil)
	b := big.NewInt(0).Exp(big.NewInt(2), big.NewInt(rand.Int63n(100)+20), nil)
	fmt.Printf("a=%s\nb=%s\n", a.String(), b.String())

	mul := big.NewInt(0).Mul(a, b)
	fmt.Printf("a*b=%s\n", mul.String())

	div := big.NewFloat(0).Quo(big.NewFloat(0).SetInt(a), big.NewFloat(0).SetInt(b))
	fmt.Printf("a/b=%s\n", div.String())

	add := big.NewInt(0).Add(a, b)
	fmt.Printf("a+b=%s\n", add.String())

	sub := big.NewInt(0).Sub(a, b)
	fmt.Printf("a-b=%s\n", sub.String())
}
