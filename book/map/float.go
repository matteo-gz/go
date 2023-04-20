package main

import "fmt"

func main() {
	a := 1.0000000000000001
	b := 1.0000000000000002
	fmt.Println(a == b) // false
	f2()
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
