package main

import "fmt"

func merge(frist, second []int) []int {
	final := []int{}
	i := 0
	j := 0
	for i < len(frist) && j < len(second) {
		if frist[i] >= second[j] {
			final = append(final, second[j])
			j++
		} else {
			final = append(final, frist[i])
			i++
		}
	}
	for ; i < len(frist); i++ {
		final = append(final, frist[i])
	}
	for ; j < len(second); j++ {
		final = append(final, second[j])
	}
	return final
}

func main() {
	a := []int{23, 45, 67, 60, 61, 123, 234, 678}
	b := []int{1, 2, 3, 4, 5, 567, 7654, 7896}
	result := merge(a, b)
	fmt.Printf(result)
}
