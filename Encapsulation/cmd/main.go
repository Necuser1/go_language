package main

import (
	"encap/pkg/comman"
	"fmt"
)

func main() {
	fmt.Println("Hello Encapsulation")
	encapsulateDataObj := comman.EncapsulateData{"neeeaj", "123445"}
	fmt.Println(encapsulateDataObj)
}
