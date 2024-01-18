package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")

	a := Aiu{Name: "aaa"}
	a1 := &Aiu{Name: "aaa"}
	a2 := new(Aiu)
	a.Name = "bbb"
	a1.Name = "bbb"
	fmt.Println(a.Name)
	fmt.Println(a1.Name)
	fmt.Println(a2.Name)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", a1)
	fmt.Printf("%T\n", a2)
}

type Aiu struct {
	Name string
}
