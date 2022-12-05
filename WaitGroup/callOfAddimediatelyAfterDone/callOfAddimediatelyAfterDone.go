package main

import (
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		time.Sleep(10 * time.Second)
		wg.Done()
		wg.Add(1)
	}()
	wg.Wait()
}

/*
	After call of wg.Done() imediately wg.Wait() do something and we can not use wg.Add() imediately by doing this we
	get error of WaitGroup is reused before previous Wait has returned
*/
