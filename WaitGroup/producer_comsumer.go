package main

import (
	"fmt"
)

func producer(job chan string, slice []string) {
	for _, value := range slice {
		job <- value

	}
	close(job)
}
func consumer(job chan string, done chan<- bool) {
	for value := range job {
		fmt.Printf(value + "\n")
	}
	done <- true
}
func main() {
	job := make(chan string)
	done := make(chan bool)
	var slice = []string{
		"The world itself's",
		"just one big hoax.",
		"Spamming each other with our",
		"running commentary of bullshit,",
	}
	go producer(job, slice)
	go consumer(job, done)
	<-done
}
