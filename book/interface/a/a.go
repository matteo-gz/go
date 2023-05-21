package main

import "fmt"

type Coder interface {
	code()
}

type Gopher struct {
	name string
}

func (g Gopher) code() {
	fmt.Printf("%s is coding\n", g.name)
}

func main() {
	var c Coder           // interface
	fmt.Println(c == nil) //true

	//c: <nil>, <nil>
	fmt.Printf("c: %T, %v\n", c, c)

	var g *Gopher
	
	// true
	fmt.Println(g == nil)

	c = g

	//false
	fmt.Println(c == nil)

	// c: *main.Gopher, <nil>
	fmt.Printf("c: %T, %v\n", c, c)
}
