package main


import (
	"fmt"
)

func main() {

	var n int
	fmt.Scan(&n)
	var arr [10]int
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	for i := 0 ; i < n; i++ {
		fmt.Println(arr[i])
	}
}
