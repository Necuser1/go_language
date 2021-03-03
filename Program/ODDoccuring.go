/*
	C/C++ Program to Find the Number Occurring Odd Number of Times

	Given an array arr[] consisting of positive integers that occur even number of times, except one number, which occurs odd number of times. The task is to find this odd number of times occurring number.
*/

package main 

import (
	"fmt"
)

func main() {
	var number int
	fmt.Scan(&number)
	var arr [100]int
	for i := 0 ;i<number; i = i+1 {
		fmt.Scan(&arr[i])
	}
	var sum int = 0
	for i := 0;i<number ; i=i+1 {
		sum = sum^arr[i]
	} 
	fmt.Println(sum)
}
