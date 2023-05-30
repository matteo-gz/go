package main

import (
	"fmt"
	"time"
)

func main() {
	var c chan bool
	fmt.Println(c, c == nil)
	go func() {
		b := <-c
		fmt.Println(b)
	}()
	//c = make(chan bool)
	c <- true
	time.Sleep(5 * time.Second)
	fmt.Println("?")
}
