/*
Title: 16.Recursion.go
Author: OpenSource
Date: 2017-05-21
Description: For Study

Go supports recursive functions. Here’s a classic factorial example.
*/

package main 

import "fmt"

// This fact function calls itself until it reaches the base case of fact(0).
func fact(n int) int {
    if n == 0 {
        return 1
    }
    return n * fact(n-1)
}

func main(){
    fmt.Println("16.Recursion.go--------Start-----------")
    
    fmt.Println("7! (seven factorial) ", fact(7))

    fmt.Println("16.Recursion.go---------End------------")
}

