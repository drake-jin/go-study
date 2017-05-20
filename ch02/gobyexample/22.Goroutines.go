/*
Title: 22.Goroutines.go
Author: OpenSource
Date: 2017-05-21
Description: For Study


A goroutine is a lightweight thread of execution.


*/

package main 

import "fmt"

func f(from string){
    for i := 0; i < 5; i++ {
        fmt.Println("argument (from =>",from ,": i=>", i,")")
    }
}

func main(){
    fmt.Println("22.Goroutines.go---------Start------------\n\n")
    
    fmt.Println("f(\"direct\")")
    f("direct")
    // Suppose we have a function call f(s).
    // Here’s how we’d call that in the usual way, running it synchronously.

    fmt.Println("go f(\"goroutine\")")
    go f("goroutine")
    // To invoke this function in a goroutine, use go f(s). 
    // This new goroutine will execute concurrently with the calling one.

    go func(msg string){
        fmt.Println(msg)
    }("going")
    // You can also start a goroutine for an anonymous function call.

    var input string 
    fmt.Scanln(&input)
    fmt.Println("Done")
    // Our two function calls are running asynchronously in separate goroutines now, 
    // so execution falls through to here.
    // This Scanln code requires we press a key before the program exits.
    // Go routine is not synchronously working Don't forget it 

    fmt.Println("\n\n22.Goroutines.go----------End-------------")
    // When we run this program, we see the output of the blocking call first, 
    // then the interleaved output of the two gouroutines. 
    // This interleaving reflects the goroutines being run concurrently by the Go runtime.
}
