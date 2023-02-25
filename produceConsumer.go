package main

import (
	"fmt"
)

type producer struct {
	msg  *chan int
	done *chan bool
}

type consumer struct {
	msg *chan int
}

func NewProducer(msg *chan int, done *chan bool) *producer {
	return &producer{msg: msg, done: done}
}
func NewConsumer(msg *chan int) *consumer {
	return &consumer{msg: msg}
}

func (p *producer) produce(max int) {
	fmt.Println("Producer started")
	for i := 0; i < max; i++ {
		fmt.Println("Producer is producing", i)
		*p.msg <- i
	}
	*p.done <- true
	fmt.Println("Producer Done")
}

func (c *consumer) consume() {
	for {
		fmt.Println("Consumer consuming")
		fmt.Println("consumer consuming", <-*c.msg)
	}
}

func main() {
	msg := make(chan int)
	done := make(chan bool)
	go NewProducer(&msg, &done).produce(5)
	go NewConsumer(&msg).consume()
	<-done
}
