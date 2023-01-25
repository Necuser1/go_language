package main

import (
	"fmt"
)

func greet(c <-chan int) {

	fmt.Println(<-c)

}
func main() {
	c := make(chan int)

	//fmt.Println("main started")

	go greet(c)

	fmt.Println(c)

	c <- 10

	//fmt.Println("main stopped! ")

}
