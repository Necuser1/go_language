package main

import (
	"fmt"
)

func main() {
	var arr [100]int
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	for i := 1; i < n-1; i++ {
		if arr[i] > arr[i-1] && arr[i] > arr[i+1] {
			fmt.Println(arr[i])
		}
	}
}
