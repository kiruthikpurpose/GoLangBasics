package main

import "fmt"

// Define an interface
type Shape interface {
    Area() float64
}

// Define a struct
type Rectangle struct {
    length  float64
    width   float64
}

// Method to calculate area
func (r Rectangle) Area() float64 {
    return r.length * r.width
}

func main() {
    // Using structs, methods, and interfaces
    rect := Rectangle{length: 5.0, width: 3.0}
    fmt.Println("Rectangle:", rect)
    fmt.Println("Area:", rect.Area())

    var shape Shape
    shape = rect
    fmt.Println("Shape Area:", shape.Area())
}
