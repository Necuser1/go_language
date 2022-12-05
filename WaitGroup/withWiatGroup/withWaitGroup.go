package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i = i + 1 {
			print(i)
		}
	}()
	fmt.Println("main Done")
	wg.Wait()
}

func print(i int) {
	fmt.Println(i)
}
