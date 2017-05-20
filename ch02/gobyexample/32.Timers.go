/*
Title: 32.Timers.go
Author: OpenSource
Date: 2017-05-21
Description: For Study

We often want to execute Go code at some point in the future, or repeatedly at some interval. 
Go’s built-in timer and ticker features make both of these tasks easy. 
We’ll look first at timers and then at tickers.[exam-33]

*/

package main 

import "time"
import "fmt"

func main(){
    fmt.Println("32.Timers.go---------Start------------\n\n")

    timer1 := time.NewTimer(time.Second * 2)
    // Timers represent a single event in the future.
    // You tell the timer how long you want to wait,
    // and it provides a channel that will be notified at that time. This timer will wait 2 seconds.
    
    fmt.Println("<- timer1.C",<- timer1.C)
    fmt.Println("Timer 1 expired")
    // The <-timer1.C blocks on the timer’s channel C until it sends a value indicating 
    // that the timer expired.

    timer2 := time.NewTimer(time.Second)
    // If you just wanted to wait, you could have used time.Sleep.
    // One reason a timer may be useful is that you can cancel the timer before it expires.
    // Here’s an example of that.

    go func() {
        fmt.Println("<- timer2.C", <- timer2.C)
        fmt.Println("Timer 2 expired")
    }()

    stop2 := timer2.Stop()
    if stop2 {
        fmt.Println("Timer 2 stopped")
    }

    fmt.Println("\n\n32.Timers.go----------End-------------n")
    // The first timer will expire ~2s after we start the program, 
    // but the second should be stopped before it has a chance to expire.
}

