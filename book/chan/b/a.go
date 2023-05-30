package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan bool)
	fmt.Println(c, c == nil)
	go func() {
		fmt.Println("1-")
		fmt.Println("1-", c, c == nil)
		b := <-c
		fmt.Println(b)
	}()
	go func() {
		fmt.Println("2-")
		fmt.Println("2-", c, c == nil)
		c = make(chan bool)
		fmt.Println("2-", c, c == nil)
		b := <-c
		fmt.Println(b)
	}()
	//
	c <- true
	time.Sleep(5 * time.Second)
	fmt.Println("?")
}
