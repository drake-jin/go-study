/*
Title: 25.Channel_Synchronization.go
Author: OpenSource
Date: 2017-05-21
Description: For Study

We can use channels to synchronize execution across goroutines.
Here’s an example of using a blocking receive to wait for a goroutine to finish.

*/

package main


import "fmt"
import "time"

// This is the function we’ll run in a goroutine.
// The done channel will be used to notify another goroutine that this function’s work is done.
func worker(done chan bool){
    fmt.Print("Working...")
    time.Sleep(time.Second)
    fmt.Println("done")

    done <- true 
    // Send a value to notify that we’re done.
}


func main(){
    fmt.Println("25.Channel_Synchronization.go---------Start------------\n\n")

    // Start a worker goroutine, giving it the channel to notify on.
    done := make(chan bool, 1)
    go worker(done)

    <- done
    // Block until we receive a notification from the worker on the channel.
    // If you removed the <- done line from this program, 
    // the program would exit before the worker even started.

    fmt.Println("\n\n25.Channel_Synchronization.go----------End-------------")
}



