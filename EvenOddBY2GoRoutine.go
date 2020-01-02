package main

import (
	"sync"
	"fmt"
	"runtime"
	"time"
)

type waitg struct {
	wg sync.WaitGroup
	evenc chan int
	oddc  chan int
}

func (w1 *waitg)start() {
	w1.evenc=make(chan int)
	w1.oddc=make(chan int)
}

func (w1 *waitg)EvenWriter() {
	var count1 int =1
	for i:=0;i<10;i++ {
		fmt.Println("This is even")
		time.Sleep(2 * time.Second)
		fmt.Println(<-w1.evenc)
		w1.oddc <- count1
		count1=count1+2
	}
	w1.wg.Done()	
}

func (w1 *waitg)OddWriter() {
	var count2 int =0
	for j:=0;j<10;j++ {
		w1.evenc <- count2
		count2=count2+2
		fmt.Println("This is odd number")
		time.Sleep(2 * time.Second)
		fmt.Println(<-w1.oddc)
	}
	w1.wg.Done()
}

func main() {
	var obj waitg
	obj.start()
	go obj.EvenWriter()
	go obj.OddWriter()
	obj.wg.Add(2)
	obj.wg.Wait()
	fmt.Println("The even and odd printing has been finished")
	fmt.Println(runtime.NumCPU())
}
