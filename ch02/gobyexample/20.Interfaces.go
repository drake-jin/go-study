/*
Title: 20.Interfaces.go
Author: OpenSource
Date: 2017-05-21
Description: For Study

Interfaces are named collections of method signatures.

*/


package main

import "fmt"
import "math"


// define struct
// For our example we’ll implement this interface on rect and circle types.
type rect struct {
    width, height float64
}

type circle struct {
    radius float64
}


// type geometry
type geometry interface{
    // Here’s a basic interface for geometric shapes.
    area() float64
    perim() float64
}


// impl methods
// To implement an interface in Go, we just need to implement all the methods in the interface. Here we implement geometry on rects.
func (r rect) area() float64 {  
    return r.width * r.height
}

func (r rect) perim() float64 {
    return 2*r.width +  2*r.height
}

// The implementation for circles.
func (c circle) area() float64 {
    return math.Pi * c.radius * c.radius
}

// The implementation for circles.
func (c circle) perim() float64 {
    return 2 * math.Pi * c.radius
}

// If a variable has an interface type, 
// then we can call methods that are in the named interface. 
// Here’s a generic measure function taking advantage of this to work on any geometry.
func measure (g geometry ) {
    fmt.Println("measure")
    fmt.Println(g)
    fmt.Println(g.area())
    fmt.Println(g.perim())
}


func main() {
    fmt.Println("20.Interfaces.go--------Start-----------\n\n")
    r := rect{width: 3, height: 4}
    c := circle{radius: 5}

    measure(r)
    measure(c)
    fmt.Println("20.Interfaces.go---------End------------\n\n")
}

