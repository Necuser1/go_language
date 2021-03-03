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
