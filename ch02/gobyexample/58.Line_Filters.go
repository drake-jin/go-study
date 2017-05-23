/*
Title: 58.Line_Filters.go
Author: OpenSource
Date: 2017-05-23
Description: For Study

A line filter is a common type of program that reads input on stdin, processes it, 
and then prints some derived result to stdout. grep and sed are common line filters.

*/

package main 

import (
    "bufio"
    "fmt"
    "os"
    "strings")

func main(){
    fmt.Println("58.Line_Filters.go---------Start------------\n\n")
    // Here’s an example line filter in Go that writes a capitalized version of all input text.
    // You can use this pattern to write your own Go line filters.
    

    // Wrapping the unbuffered os.Stdin with a buffered scanner 
    // gives us a convenient Scan method that advances the scanner to the next token; which 
    // is the next line in the default scanner.
    scanner := bufio.NewScanner(os.Stdin)

    // Text returns the current token, here the next line, from the input.
    for scanner.Scan(){
        // Write out the uppercased line.
        ucl := strings.ToUpper(scanner.Text())
        fmt.Println(ucl)
    }
    
    // Check for errors during Scan. End of file is expected and not reported by Scan as an error.
    if err := scanner.Err(); err != nil{
        fmt.Fprintln(os.Stderr, "error: ",err)
        os.Exit(1)
    }
    // To try out our line filter, first make a file with a few lowercase lines.
    // Then use the line filter to get uppercase lines.
    fmt.Println("\n\n58.Line_Filters.go-----------End------------")

}


