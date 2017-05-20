/*
Title: 13.Multiple_Return_Values.go
Author: OpenSource
Date: 2017-05-21
Description: For Study


Functions are central in Go. 
We’ll learn about functions with a few different examples.
*/


package main 

import "fmt"

// The (int, int) in this function signature shows that the function returns 2 ints.
func vals()(int,int){
    return 3, 7
}

func main () {
    fmt.Println("13.Multiple_Return_Values.go--------Start-----------")
    
    // Here we use the 2 different return values from the call with multiple assignment.
    a, b := vals()
    fmt.Println("a, b := vals() => ", a, b)
    fmt.Println("a => ", a)
    fmt.Println("b => ", b)

    // If you only want a subset of the returned values, use the blank identifier _.
    _, c := vals()
    fmt.Println("_, c :=vals() => c ", c)

    fmt.Println("13.Multiple_Return_Values.go---------End------------")
}

