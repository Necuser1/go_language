package main

import (
	"fmt"
)

func FindSqure(c chan int) {
	for i := 1; i < 10; i = i + 1 {
		c <- i * i
	}
	close(c)
}
func main() {
	c := make(chan int)
	go FindSqure(c)

	for {
		val, ok := <-c
		if ok == false {
			break
		} else {
			fmt.Println(val)
		}
	}
	fmt.Println("main stopped!")
}
