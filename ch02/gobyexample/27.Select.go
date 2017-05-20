/*
Title: 27.Select.go
Author: OpenSource
Date: 2017-05-21
Description: For Study

Go’s select lets you wait on multiple channel operations. 
Combining goroutines and channels with select is a powerful feature of Go.

*/

package main 

import "time"
import "fmt"

func main(){
    fmt.Println("27.Select.go---------Start------------\n\n")
   
    fmt.Println("c1 := make(chan string)")
    fmt.Println("c2 := make(chan string)")
    c1 := make(chan string)
    c2 := make(chan string)
    // For our example we’ll select across two channels.


    // Each channel will receive a value after some amount of time, to simulate e.g.
    // blocking RPC operations executing in concurrent goroutines.
    go func(){
        time.Sleep(time.Second * 1)
        c1 <- "one"
    }()

    go func(){
        time.Sleep(time.Second * 2)
        c2 <- "two"
    }()
    
    // We’ll use select to await both of these values simultaneously, 
    // printing each one as it arrives
    for i := 0; i < 2; i++ {
        select{
            case msg1 := <- c1:
                fmt.Println("received", msg1)

            case msg2 := <- c2:
                fmt.Println("received", msg2)
        }
    }
    // We receive the values "one" and then "two" as expected.
    // Note that the total execution time is only ~2 seconds 
    // since both the 1 and 2 second Sleeps execute concurrently.

    fmt.Println("27.Select.go----------End-------------\n\n")
}


