/*
Title: 18.Stucts.go
Author: OpenSource
Date: 2017-05-21
Description: For Study

Go’s structs are typed collections of fields. 
They’re useful for grouping data together to form records.


*/


package main 

import "fmt"

// This person struct type has name and age fields.
type person struct {
    name string 
    age int 

}

func main() {
    fmt.Println("18.Structs.go--------Start-----------")

    // This syntax creates a new struct.
    fmt.Println(person{"bob", 20})
    
    // You can name the fields when initializing a struct.
    fmt.Println(person{name:"Alice", age: 30})

    // Omitted fields will be zero-valued.
    fmt.Println(person{name: "Fred"})

    // An & prefix yields a pointer to the struct.
    fmt.Println(&person{name: "Ann", age: 40})

    // Access struct fields with a dot.
    s := person{name: "Sean", age: 50}
    fmt.Println(s.name)
    fmt.Println(s.age)

    // You can also use dots with struct pointers - the pointers are automatically dereferenced.
    sp := &s
    fmt.Println("------------------- sp := &s")
    fmt.Println(sp.name)
    fmt.Println(sp.age)

    // Structs are mutable.
    sp.age = 51
    fmt.Println("------------------- sp.age = 51")
    fmt.Println(sp.name)
    fmt.Println(sp.age)

    fmt.Println("18.Structs.go---------End------------")
}

