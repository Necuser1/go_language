package main

import (
	"fmt"
)

func main() {
	arr := make([]int, 0)
	var n int
	fmt.Println("Enter a number:")
	fmt.Scan(&n)
	fmt.Println(n)
	for i := 0; i < n; i++ {
		var a int
		fmt.Scan(&a)
		arr = append(arr, a)
	}
	fmt.Println(arr)
}
