package main

import (
	"fmt"
)


type  Shape interface {
	drow()
}

type Rectangle struct {
}
type Squre struct {
}
type circle struct {
}

type ShapeFactory struct {
}
func (rect Rectangle)drow() {
	fmt.Println("implementation of rectangle")
}

func (rect Squre)drow() {
        fmt.Println("implementation of Squre")
}

func (rect circle)drow() {
        fmt.Println("implementation of circle")
}

func (shape ShapeFactory)ImplementShapeFactory(shape1 string) Shape{
	var s Shape
	if shape1 == "" {
		return nil
	}
	if shape1 == "Rectangle" {
		s = Rectangle{}
	}
	if shape1 == "Squre" {
                s = Squre{}
        }
	if shape1 == "Circle" {
               s = circle{}
        }
	return s

}

func main() {
	 var str string
 	fmt.Scanf("%s", &str)
	s := ShapeFactory{}
	obj := s.ImplementShapeFactory(str)
	obj.drow()

}
