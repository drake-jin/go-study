/*
Title: 14.Variadic_Functions.go
Author: OpenSource
Date: 2017-05-21
Description: For Study

[Variadic functions](http://en.wikipedia.org/wiki/Variadic_function)

Variadic functions can be called with any number of trailing arguments. 
For example, fmt.Println is a common variadic function.


*/

package main

import "fmt"

// Here’s a function that will take an arbitrary number of ints as arguments.
func sum(nums ...int){
    fmt.Print(nums, "  ")
    total := 0 

    for _, num := range nums {
        total += num
    }
    fmt.Println(total)
}

func main() {
    fmt.Println("14.Variadic_Functions.go--------Start-----------")
    
    // Variadic functions can be called in the usual way with individual arguments.
    fmt.Println("sum(1,2)--------")
    sum(1, 2)
    fmt.Println("sum(1,2,3)--------")
    sum(1, 2, 3)

    // If you already have multiple args in a slice,
    // apply them to a variadic function using func(slice...) like this.
    nums := []int{1,2,3,4,5} 
    fmt.Println("nums := []int{1,2,3,4,5}")
    fmt.Println("sum(nums...)--------")
    sum(nums...)
    fmt.Println("14.Variadic_Functions.go---------End------------")
}


