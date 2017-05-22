/*
Title: 41.Panic.go
Author: OpenSource
Date: 2017-05-22
Description: For Study

A panic typically means something went unexpectedly wrong. 
Mostly we use it to fail fast on errors that shouldn’t occur during normal operation, 
or that we aren’t prepared to handle gracefully.   

*/
package main 

import "os"
import "fmt"

func main(){
    fmt.Println("41.Panic.go---------Start------------\n\n")
    
    // We’ll use panic throughout this site to check for unexpected errors. 
    // This is the only program on the site designed to panic.
    fmt.Println("panic(\"a Problem\")")
    panic("A problem")

    // A common use of panic is to abort if a function returns an error value 
    // that we don’t know how to (or want to) handle. 
    // Here’s an example of panicking if we get an unexpected error when creating a new file. 
    _, err := os.Create("/tmp/file")
    if err != nil {
        panic(err)
    }

    fmt.Println("\n\n41.Panic.go-----------End------------")
    // A common use of panic is to abort if a function returns an error value 
    // that we don’t know how to (or want to) handle. Here’s an example of panicking 
    // if we get an unexpected error when creating a new file.
}

