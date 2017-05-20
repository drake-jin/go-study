/*
Title: 7.Switch.go
Author: OpenSource
Date: 2017-05-20
Description: For Study

Switch statements express conditionals across many branches.

*/

package main

import "fmt"
import "time"

func main() {
    fmt.Println("7.Switch.go--------Start-----------")
   
    fmt.Println("\n i := 2 switch i {} ")
    // Here’s a basic switch.
    i := 2
    fmt.Print("Write ", i, " as ")
    switch i {
        case 1:
            fmt.Println("One")
        case 2:
            fmt.Println("Two")
        case 3:
            fmt.Println("Three")
    }

    fmt.Println("\n switch time.Now().Weekday(){} ")
    // You can use commas to separate multiple expressions in the same case statement.
    // We use the optional default case in this example as well.
    switch time.Now().Weekday() {
        case time.Saturday, time.Sunday:
            fmt.Println("It's the weekend")
        default:
            fmt.Println("It's a weekday")
    }

    fmt.Println("\n switch time.Now().Weekday(){}")
    // switch without an expression is an alternate way to express if/else logic.
    // Here we also show how the case expressions can be non-constants.
    t := time.Now()
    switch{
        case t.Hour() < 12:
            fmt.Println("It's before noon")
        default:
            fmt.Println("It's after noon")
    }
    
    fmt.Println("\n whatAmI := func(i interface{}){ switch t := i.(type){} } ")
    // A type switch compares types instead of values. 
    // You can use this to discover the the type of an interface value. 
    // In this example, the variable t will have the type corresponding to its clause.
    whatAmI := func(i interface{}){
        switch t := i.(type){
            case bool:
                fmt.Println("I'm a bool")
            case int:
                fmt.Println("I'm an int")
            default:
                fmt.Println("Don't know type %T \n",t)
        }
    }
    
    fmt.Println("whatAmI(true)")
    whatAmI(true)
    fmt.Println("whatAmI(1)")
    whatAmI(1)
    fmt.Println("whatAmI(\"hey\")")
    whatAmI("hey")

    fmt.Println("7.Switch.go--------End-----------")
}
