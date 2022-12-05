package main

import (
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
}

/*
	Adding 1 goroutine and calling Wait function but there is no goroutine to wait So we will get deadlock in Code
*/
