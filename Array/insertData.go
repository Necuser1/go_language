package main

import (
	"fmt"
)

func appendSlice(arr []int, element, index int) []int {

	arr = append(arr[:index], append([]int{element}, arr[index:]...)...)

	return arr

}
func main() {
	arr1 := []int{32, 57, 35, 22}
	element := 10
	index := 2
	result := appendSlice(arr1, element, index)
	fmt.Println(result)
}
