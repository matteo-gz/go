package main

import (
	"fmt"
	"time"
)

func main() {
	//c := make(chan bool)
	var c chan bool
	fmt.Println(c)
	go func() {
		fmt.Println("1")
		fmt.Println(2)
		for {
			time.Sleep(1 * time.Second)
			fmt.Print("1")
		}
	}()
	<-c
}
