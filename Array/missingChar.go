package main

import (
	"fmt"
)

func main() {
	var str1, str2 string
	fmt.Scan(&str1)
	fmt.Scan(&str2)
	stringMap := make(map[string]int)
	for _, value := range str2 {
		stringMap[string(value)] = 1
	}
	for _, value := range str1 {
		if _, ok := stringMap[string(value)]; ok == false {
			fmt.Println(string(value))
		}
	}
}
