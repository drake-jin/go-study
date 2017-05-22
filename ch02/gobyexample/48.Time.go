/*
Title: 48.Time.go
Author: OpenSource
Date: 2017-05-23
Description: For Study

Go offers extensive support for times and durations; here are some examples.

*/

package main

import "fmt"
import "time"

func main(){
    
    fmt.Println("48.Time.go---------Start------------\n\n")

    p := fmt.Println

    // We’ll start by getting the current time.
    now := time.Now()
    p(now)
    
    // You can build a time struct by providing the year, month, day, etc. 
    // Times are always associated with a Location, i.e. time zone.
    then := time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)

    p(then)

    // You can extract the various components of the time value as expected.
    p("then.Year() =>", then.Year())
    p("then.Month() => ", then.Month())
    p("then.Day() => ", then.Day())
    p("then.Hour() => ", then.Hour())
    p("then.Minute() => ", then.Minute())
    p("then.Second() =>", then.Second())
    p("then.Nanosecond() => ", then.Nanosecond())
    p("then.Location() =>", then.Location())

    // The Monday-Sunday Weekday is also available.
    p("then.Weekday() => ", then.Weekday())

    // These methods compare two times, 
    // testing if the first occurs before, after, or at the same time as the second, respectively.
    p("then.Before(now) =>", then.Before(now))
    p("then.After(now) =>", then.After(now))
    p("then.Equal(now) =>", then.Equal(now))

    // The Sub methods returns a Duration representing the interval between two times.    
    diff := now.Sub(then)   
    p("diff := now.Sub(then) =>", diff) 

    // We can compute the length of the duration in various units.   
    p("diff.Hours() =>", diff.Hours())   
    p("diff.Minutes() =>", diff.Minutes())   
    p("diff.Seconds() =>", diff.Seconds())   
    p("diff.Nanoseconds() =>", diff.Nanoseconds())
    // You can use Add to advance a time by a given duration,
    // or with a - to move backwards by a duration.   
    p("then.Add(diff) =>", then.Add(diff))   
    p("then.Add(-diff) =>", then.Add(-diff))

    fmt.Println("\n\n48.Time.go------------End------------\n\n")
}
