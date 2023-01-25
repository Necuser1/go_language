package main

import (
	"fmt"
	"time"
)

func Portal1(portalData chan string) {
	time.Sleep(3 * time.Second)
	fmt.Println(<-portalData)
}

func Portal2(portal2Data chan string) {
	time.Sleep(9 * time.Second)
	fmt.Println(<-portal2Data)
}

func main() {
	R1 := make(chan string)
	R2 := make(chan string)

	go Portal1(R1)
	go Portal2(R2)

	select {
	case R1 <- "Hello Portal1":
		fmt.Println("Portal 1 has been Updated")
	case R2 <- "Hello portal2":
		fmt.Println("Portal-2 has been updated")
	}
}
