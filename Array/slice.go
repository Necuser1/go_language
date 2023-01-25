package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	
	arr = append(arr[0:2], arr[2:5]...)
	fmt.Println(arr)
}
