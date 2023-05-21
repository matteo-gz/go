package main

import "fmt"

type MyError struct{}

func (i MyError) Error() string {
	return "MyError"
}

func main() {
	err := Process()

	// nil
	fmt.Println(err)

	// false
	fmt.Println(err == nil)
}

func Process() error {
	var err *MyError = nil
	// true
	fmt.Println(err == nil)
	return err
}
