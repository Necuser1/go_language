package main

import (
	"fmt"
	"sync"
)

func hello(c chan int, wg *sync.WaitGroup) {
	fmt.Println(<-c)
	wg.Done()
	fmt.Println(<-c)
	wg.Done()
	wg.Done()
}
func main() {
	wg := &sync.WaitGroup{}
	c := make(chan int)
	wg.Add(3)
	go hello(c, wg)
	c <- 10
	c <- 20
	wg.Wait()
}
