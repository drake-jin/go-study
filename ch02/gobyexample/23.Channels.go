/*
Title: 23.Channels.go
Author: OpenSource
Date: 2017-05-21
Description: For Study

Channels are the pipes that connect concurrent goroutines.
You can send values into channels from one goroutine and receive those values into another goroutine.

*/

package main 

import "fmt"

func main(){
    fmt.Println("23.Channels.go---------Start------------\n\n")

    fmt.Println("messages := make(chan string)")
    messages := make(chan string)
    // Create a new channel with make(chan val-type). 
    // Channels are typed by the values they convey.

    fmt.Println("go func(){messages <- \"ping\"}()")
    go func(){ messages <- "ping" }()
    // Send a value into a channel using the channel <- syntax.
    // Here we send "ping" to the messages channel we made above, from a new goroutine.

    fmt.Println("msg := <- messages")
    msg := <- messages
    // The <-channel syntax receives a value from the channel.
    // Here we’ll receive the "ping" message we sent above and print it out.

    fmt.Println("msg =>", msg)
    // When we run the program the "ping" message is successfully passed 
    // from one goroutine to another via our channel.

    fmt.Println("\n\n23.Channels.go---------Start------------\n\n")
    // By default sends and receives block until both the sender and receiver are ready.
    // This property allowed us to wait 
    // at the end of our program for the "ping" message 
    // without having to use any other synchronization.
}
