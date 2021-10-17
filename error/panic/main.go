package main	

import (
	"fmt"
	"os"
)

func divideByZero() {
	defer func() {
                if r := recover() ; r != nil {
                        fmt.Printf("a panig is observed with value %v \n",r)
                }
        }()
	var a , b int 
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Println(a/b)

}

func assersionError() {
	defer func() {
                if r := recover() ; r != nil {
                        fmt.Printf("a panig is observed with value %v \n",r)
                }
        }()
	var str interface{}
	fmt.Scan(&str)
	fmt.Println("enter value 1 for executing error and 2 for normal flow")
	var a int 
	fmt.Scan(&a)
	if a ==1 {
		result := str.([]interface{})
		fmt.Println(result)
	} else {
		result := str.(string)
		fmt.Println(result)
	}
	fmt.Println("normal execution done! ")
}
func main() {
	consoleinput := os.Args
	errorType := (consoleinput[1])
	if errorType == "1" {
		divideByZero()
	}
	if errorType == "2" {
		assersionError()
	}
	var b int 
	fmt.Scan(&b)

}
