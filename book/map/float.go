package main

import "fmt"

func main() {
	//a := 1.0000000000000001
	//b := 1.0000000000000002
	//fmt.Println(a == b) // false
	//f2()
	f3()
}
func f2() {
	a := 0.1
	b := 0.2
	c := 0.3

	if a+b == c {
		fmt.Println("a + b == c")
	}

	if c-a-b == 0 {
		fmt.Println("c - a - b == 0")
	}
}

func f3() {
	fmt.Printf("%b\n", 0b1011&^0b11)
	fmt.Printf("%b\n", 0b10111&^0b111)
	// 0b10111
	// 0b00111
	fmt.Printf("%b\n", 0b10111&^0b101)
	// 0b10111
	// 0b00101
	// 0b10010
}
