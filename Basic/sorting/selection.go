package main 

import (
	"fmt"
//	"math"
)


func main() {
	var n int 
	var arr [100]int
	fmt.Scan(&n)
        //maxValue := math.MaxInt64
	for i := 0; i<n; i++ {
		fmt.Scan(&arr[i])
	}
	for i:=0; i<n-1; i++ {
		for j := i; j<n; j++ {
			if arr[j] < arr[i] {
				arr[i],arr[j] = arr[j],arr[i]
			}
		}
	}
	for i := 0; i<n; i++ 	{
		fmt.Println(arr[i])
	}
}
