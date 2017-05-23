/*
Title: 65.Exit.go
Author: OpenSource
Date: 2017-05-23
Description: For Study

Use os.Exit to immediately exit with a given status.

*/
package main

import "fmt"
import "os"

func main(){
    fmt.Println("65.Exit.go---------Start------------\n\n")

    defer fmt.Println("defer !! print !!!!!!!!!")
    // defers will not be run when using os.Exit, so this fmt.Println will never be called.

    os.Exit(3)
    // Exit with status 3.

    fmt.Println("\n\n65.Exit.go-----------End------------")
    // Note that unlike e.g. C, Go does not use an integer return value 
    // from main to indicate exit status. 
    // If you’d like to exit with a non-zero status you should use os.Exit.

    // If you run exit.go using go run, the exit will be picked up by go and printed.
    // By building and executing a binary you can see the status in the terminal.
}

