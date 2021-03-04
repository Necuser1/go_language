package main

import(
	"fmt"
)

func main() {
	var num int 
	fmt.Scan(&num)
	var arr [10000]int
	for i := 0; i<num ; i=i+1 {
		fmt.Scan(&arr[i])
	}
	
	max_result := 0
	max_end_here := 0 
	for i := 0; i<num ; i=i+1 {
		max_end_here = max_end_here+arr[i] 
		
		if max_result < max_end_here {
			max_result = max_end_here
		}
		
		if max_end_here < 0 {
			max_end_here = 0
		}
	}
	fmt.Println(max_result)
}
