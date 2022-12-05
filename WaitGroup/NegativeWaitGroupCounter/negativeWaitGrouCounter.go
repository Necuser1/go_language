package main

import (
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Done()
}

/*
	here wg.Done() methode call wg.Add(-1) So the program will give a panic of negative WaitGroup Counter
*/
