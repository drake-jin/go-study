/*
Title: 15.Closures.go
Author: OpenSource
Date: 2017-05-21
Description: For Study

[Anonymous functions](http://en.wikipedia.org/wiki/Anonymous_function)
[Closures](http://en.wikipedia.org/wiki/Closure_(computer_science))

Go supports anonymous functions, which can form closures. 
Anonymous functions are useful when you want to define 
a function inline without having to name it.

*/

package main 

import "fmt"

// This function intSeq returns another function, 
// which we define anonymously in the body of intSeq. 
// The returned function closes over the variable i to form a closure.
func intSeq() func() int {
    i := 0
    return func() int{
        i += 1
        return i
    }
}



func main() {
    fmt.Println("15.Closures.go--------Start-----------")
    
    // We call intSeq, assigning the result (a function) to nextInt. 
    // This function value captures its own i value, 
    // which will be updated each time we call nextInt.
    fmt.Println("nextInt := intSeq()")
    nextInt := intSeq() 
    
    // See the effect of the closure by calling nextInt a few times.
    fmt.Println("nextInt() =>", nextInt())
    fmt.Println("nextInt() =>", nextInt())
    fmt.Println("nextInt() =>", nextInt())
    fmt.Println("nextInt() =>", nextInt())

    // To confirm that the state is unique to that particular function, create and test a new one.
    fmt.Println("newInts := intSeq()")
    newInts := intSeq()

    fmt.Println("newInts() =>", newInts())
    fmt.Println("newInts() =>", newInts())
    fmt.Println("newInts() =>", newInts())
    fmt.Println("newInts() =>", newInts())

    fmt.Println("They are useing different values ")
    fmt.Println("nextInt() =>", nextInt())
    fmt.Println("nextInt() =>", nextInt())
    fmt.Println("nextInt() =>", nextInt())
    fmt.Println("nextInt() =>", nextInt())

    fmt.Println("newInts() =>", newInts())
    fmt.Println("newInts() =>", newInts())
    fmt.Println("newInts() =>", newInts())
    fmt.Println("newInts() =>", newInts())

    fmt.Println("15.Closures.go---------End------------")
}

