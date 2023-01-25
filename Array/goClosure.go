package main

import "fmt"

func oddNumber() func() int {
	n := 1

	oddvalue := func() int {
		n = n + 2
		return n
	}
	return oddvalue
}

func main() {

	result := oddNumber()
	fmt.Println(result())
	fmt.Println(result())

}
