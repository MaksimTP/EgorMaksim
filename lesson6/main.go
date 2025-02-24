package main

import "fmt"

type Human struct {
	Age int
}

func (h *Human) AddAge(value int) {
	h.Age += value
}

// func AddAge(h Human, value int) {}

func main() {
	h := Human{
		Age: 1,
	}

	h.AddAge(10)

	fmt.Printf("%#v", h)
}
