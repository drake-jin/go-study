/*
Title: 26.Channel_Directions.go
Author: OpenSource
Date: 2017-05-21
Description: For Study

When using channels as function parameters, 
you can specify if a channel is meant to only send or receive values. 
This specificity increases the type-safety of the program.

*/


package main


import "fmt"


// This ping function only accepts a channel for sending values.
// It would be a compile-time error to try to receive on this channel.
func ping(pings chan <- string, msg string){
    pings <- msg
}


// The pong function accepts one channel for receives (pings) and a second for sends (pongs).
func pong(pings <- chan string, pongs chan <- string){
    msg := <- pings
    pongs <- msg
}

func main(){
    fmt.Println("26.Channel_Directions.go---------Start------------\n\n")

    fmt.Println("pings := make(chan string,1)") 
    fmt.Println("pongs := make(chan string,1)") 
    pings := make(chan string, 1)
    pongs := make(chan string, 1)
   
    fmt.Println("ping(pings, \"passed messages\")")
    ping(pings, "passed messages")
    
    fmt.Println("pong(pings, pongs)")
    pong(pings, pongs)
    
    fmt.Println("<-pongs =>", <-pongs)

    fmt.Println("\n\n26.Channel_Directions.go----------End-------------")
}














