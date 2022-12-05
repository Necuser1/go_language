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
		var wg_new sync.WaitGroup
		wg_new.Add(3)
		go func() {
			wg_new.Done()
			fmt.Println("The frist thread...")
		}()
		go func() {
			wg_new.Done()
			fmt.Println("Second Thread....")
		}()
		go func() {
			wg_new.Done()
			fmt.Println("Third Thread....")
		}()
		wg_new.Wait()
	}()
	wg.Wait()
}
