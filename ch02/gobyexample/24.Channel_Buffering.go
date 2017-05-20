/*
Title: 24.Channel_Buffering.go
Author: OpenSource
Date: 2017-05-21
Description: For Study

By default channels are unbuffered,
meaning that they will only accept sends (chan <-)
if there is a corresponding receive (<- chan) ready to receive the sent value. 
Buffered channels accept a limited number of values without a corresponding receiver for those values.

*/

package main 


import "fmt"


func main(){
    fmt.Println("24.Channel_Buffering.go---------Start------------\n\n")
    
    // Here we make a channel of strings buffering up to 2 values.
    fmt.Println("messages := make(chan string, 2)")
    messages := make(chan string, 2)


    // Because this channel is buffered, 
    // we can send these values into the channel without a corresponding concurrent receive.
    fmt.Println("messages <- \"Buffered\"")
    fmt.Println("messages <- \"Channel\"")
    messages <- "Buffered"
    messages <- "Channel"

    // Later we can receive these two values as usual.
    fmt.Println("<- messages =>", <-messages)
    fmt.Println("<- messages =>", <-messages)

    fmt.Println("\n\n24.Channel_Buffering.go----------End-------------")
}



