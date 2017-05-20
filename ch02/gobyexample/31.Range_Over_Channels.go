/*
Title: 31.Range_Over_Channels.go
Author: OpenSource
Date: 2017-05-21
Description: For Study

In a previous example[exam-11] we saw how for and range provide iteration over basic data structures.
We can also use this syntax to iterate over values received from a channel.

*/

package main

import "fmt"

func main(){
    fmt.Println("31.Range_Over_Channels.go---------Start------------\n\n")

    // We’ll iterate over 2 values in the queue channel.
    queue := make(chan string, 2)
    queue <- "one"
    queue <- "two"
    close(queue) // buffer close
    
    // This range iterates over each element as it’s received from queue.
    // Because we closed the channel above, the iteration terminates after receiving the 2 elements.
    for elem := range queue{
        fmt.Println(elem)
    }
    
    fmt.Println("\n\n31.Range_Over_Channels.go----------End-------------")
    // This example also showed 
    // that it’s possible to close a non-empty channel 
    // but still have the remaining values be received.
}
