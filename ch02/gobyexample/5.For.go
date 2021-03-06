/*
Title: 5.For.go
Author: OpenSource
Date: 2017-05-20
Description: For Study

for is Go’s only looping construct. Here are three basic types of for loops.
*/

package main

import "fmt"

func main(){ 
    fmt.Println("5.For.go--------Start-----------")
    
    i := 1
    
    fmt.Println("for i <= 3 {} ")
    // The most basic type, with a single condition.
    for i <= 3 {
        fmt.Println(i)
        i = i + 1
    }

    fmt.Println("for j:= 7; j <= 9; j++ {} ")
    // A classic initial/condition/after for loop
    for j := 7; j <= 9; j++{
        fmt.Println(j)
    }

    fmt.Println("for {} ")
    // for without a condition will loop repeatedly until you break out of the loop 
    // or return from the enclosing function.
    for {
        fmt.Println("loop")
        break
    }

    fmt.Println("for n:= 0; n <= 5; n++ {} ")
    // You can also continue to the next iteration of the loop.
    for n := 0 ; n <= 5; n++ {
        if n%2 ==0 {
            continue
        }
        fmt.Println(n)
    }
    // We’ll see some other for forms later when we look at 
    // range statements, channels, and other data structures.
    fmt.Println("5.For.go--------End-----------")
}
