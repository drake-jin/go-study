/*
Title: 49.Epoch.go
Author: OpenSource
Date: 2017-05-23
Description: For Study

A common requirement in programs is getting the number of seconds, milliseconds, or nanoseconds 
since the Unix epoch. Here’s how to do it in Go.


*/

package main

import "fmt"
import "time"

func main() {
    
    fmt.Println("49.Epoch.go---------Start------------\n\n")

    // Use time.Now with Unix or UnixNano to get elapsed time 
    // since the Unix epoch in seconds or nanoseconds, respectively.
    now := time.Now()
    secs := now.Unix()
    nanos := now.UnixNano()
    fmt.Println(now)

    // Note that there is no UnixMillis, so to get the milliseconds since epoch you’ll need to manually divide from nanoseconds.
    millis := nanos / 1000000
    
    fmt.Println(secs)
    fmt.Println(millis)
    fmt.Println(nanos)

    // You can also convert integer seconds or nanoseconds since the epoch into the corresponding time.
    fmt.Println(time.Unix(secs, 0))
    fmt.Println(time.Unix(0, nanos))
    
    fmt.Println("\n\n49.Epoch.go-----------End------------")
    // Next we’ll look at another time-related task: time parsing and formatting.
}













