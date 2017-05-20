/*
Title: 6.If_Else.go
Author: OpenSource
Date: 2017-05-20
Description: For Study

Branching with if and else in Go is straight-forward.
*/
package main

import "fmt"

func main(){
    fmt.Println("6.If_Else.go--------Start-----------")
    
    fmt.Println("\n if 7%2 == 0 {} else {}")
    // Here’s a basic example.
    if 7%2 == 0 {
        fmt.Println("7 is even")
    }else {
        fmt.Println("7 is odd")
    }

    fmt.Println("\n if 8%2==0 {}")
    // You can have an if statement without an else.
    if 8%2 == 0 {
        fmt.Println(" 8 is divisible by 4")
    }

    fmt.Println("\n if num := 9; num < 0 {}else if num < 10 {} else {}")
    // A statement can precede conditionals; 
    // any variables declared in this statement are available in all branches.
    if num := 9; num < 0 {
        fmt.Println(num, " is negative")
    }else if num < 10 {
        fmt.Println(num, " has 1 digit")
    }else {
        fmt.Println(num, " has multiple digits")
    }

    // Note that you don’t need parentheses around conditions in Go, 
    // but that the braces are required.
    // There is no ternary if in Go,  [ternary](http://en.wikipedia.org/wiki/%3F:)
    // so you’ll need to use a full if statement even for basic conditions.
    fmt.Println("6.If_Else.go--------End-----------")
}





