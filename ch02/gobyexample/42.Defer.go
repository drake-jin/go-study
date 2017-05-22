/*
Title: 42.Defer.go
Author: OpenSource
Date: 2017-05-22
Description: For Study

Defer is used to ensure that a function call is performed later in a program’s execution,
usually for purposes of cleanup. defer is often used where e.g.
ensure and finally would be used in other languages.


*/

package main 

import "fmt"
import "os"


func main() {
    fmt.Println("42.Defer.go---------Start------------\n\n")
    
    // Immediately after getting a file object with createFile, 
    // we defer the closing of that file with closeFile. 
    // This will be executed at the end of the enclosing function (main),
    // after writeFile has finished.
    f := createFile("/tmp/defer.txt")
    defer closeFile(f)
    writeFile(f)

    
    fmt.Println("\n\n42.Defer.go-----------End------------")
    // Suppose we wanted to create a file, write to it,
    // and then close when we’re done. Here’s how we could do that with defer.
}

func createFile(p string) *os.File{
    fmt.Println("------method - call => createFile(p string) *os.File{}") 
    f, err := os.Create(p) 
    if err != nil{
        panic(err)
    }
    return f
}

func writeFile(f *os.File){
    fmt.Println("------method - call => writeFile(f *os.File){} ") 
    fmt.Fprintln(f, "data")

}

func closeFile(f *os.File){
    fmt.Println("------method - call => closeFile(f *os.File{} ") 
    f.Close()
}
