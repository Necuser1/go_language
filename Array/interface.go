package main

import (
	"fmt"
)

type inter interface {
	getcolor() string
}

type Dog struct {
	clolor string
}

type Cat struct {
	color string
}

func (d Dog) getcolor() string {
	return d.clolor
}
func (c Cat) getcolor() string {

	return c.color

}

func main() {
	var obj inter
	obj = &Dog{"Black"}
	fmt.Println(obj.getcolor())
	obj = &Cat{"White"}
	fmt.Println(obj.getcolor())
}
