/*
Title: 36.Atomic_Counters.go
Author: OpenSource
Date: 2017-05-21
Description: For Study

The primary mechanism for managing state in Go is communication over channels.
We saw this for example with worker pools. [exam - 34]
There are a few other options for managing state though. 
Here we’ll look at using the sync/atomic package for atomic counters accessed 
by multiple goroutines.

*/


package main 

import "fmt"
import "time"
import "sync/atomic"

func main(){ 
    fmt.Println("36.Atomic_Counters.go---------Start------------\n\n")
    
    var ops uint64 = 0 
    // We’ll use an unsigned integer to represent our (always-positive) counter.
    
    for i := 0 ; i < 50 ; i++{
        go func(){
            // To simulate concurrent updates,
            // we’ll start 50 goroutines that each increment the counter about once a millisecond.
            for {
                atomic.AddUint64(&ops, 1)
                // To atomically increment the counter we use AddUint64,
                // giving it the memory address of our ops counter with the & syntax.
                time.Sleep(time.Millisecond)  
                // Wait a bit between increments.
            }
        }()
    }
    time.Sleep(time.Second)
    // Wait a second to allow some ops to accumulate.

    // In order to safely use the counter while it’s still being updated by other goroutines,
    // we extract a copy of the current value into opsFinal via LoadUint64. 
    // As above we need to give this function the memory address &ops 
    // from which to fetch the value.

    opsFinal := atomic.LoadUint64(&ops)
    fmt.Println("opsFinal =>", opsFinal)

    fmt.Println("\n\n36.Atomic_Counters.go----------End-------------")
    // Next we’ll look at mutexes, another tool for managing state.
    // Running the program shows that we executed about 40,000 operations.
}

