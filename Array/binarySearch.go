package main

import (
	"fmt"
)

func BinaryS(arr []int, l, r, value int) int {
	var mid int
	for r > l {
		mid = r + l/2
		if arr[mid] == value {
			return mid
		}
		if value > arr[mid] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return -1
}
func main() {
	arr := []int{}
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var a int
		fmt.Scan(&a)
		arr = append(arr, a)
	}
	var value int
	fmt.Scan(&value)
	result := BinaryS(arr, 0, n-1, value)
	fmt.Println(result)
}
