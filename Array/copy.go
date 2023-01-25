package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6}

	slice = append(slice, 0)
	fmt.Println(slice[3:])
	fmt.Println(slice[2:])
	copies := copy(slice[3:], slice[2:])
	fmt.Println(copies)
	fmt.Println(slice)

}
