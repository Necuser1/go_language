package main

import (
	"fmt"
)

func main() {
	arr := [5]int{0, 1, 0, 1, 0}
	var arrar [10]int
	fmt.Println(arrar)
	j := -1
	for i := 0; i < 5; i = i + 1 {
		if arr[i] == 0 {
			j = j + 1
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	fmt.Println(arr)
}
