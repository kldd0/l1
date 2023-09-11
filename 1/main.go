package main

import "fmt"

type Human struct {
	Name string
	Age  int
}

func (h Human) Say(s string) {
	fmt.Printf("%s сказал: %s\n", h.Name, s)
}

type Action struct {
	Human
}

func (a Action) Run() {
	fmt.Printf("%s побежал...\n", a.Human.Name)
}

func main() {
	a := Action{
		Human: Human{
			Name: "Dave",
			Age:  23,
		},
	}

	a.Say("Hello")
	a.Run()
}
