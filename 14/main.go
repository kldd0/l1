package main

import "fmt"

func main() {
	n := 1234
	s := "qwerty"
	b := true
	c := make(chan struct{})
	arr := []interface{}{n, s, b, c}

	for _, v := range arr {
		switch t := v.(type) {
		case int:
			fmt.Printf("%v is type of %T\n", t, t)
		case string:
			fmt.Printf("%v is type of %T\n", t, t)
		case bool:
			fmt.Printf("%v is type of %T\n", t, t)
		case chan struct{}:
			fmt.Printf("%v is type of %T\n", t, t)
		default:
			fmt.Println("Uknown type")
		}
	}
}
