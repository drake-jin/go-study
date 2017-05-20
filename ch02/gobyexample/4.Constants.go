/*
Title: 4.Constants.go
Author: OpenSource
Date: 2017-05-20
Description: For Study

Go supports constants of character, string, boolean, and numeric values.

*/

package main

import "fmt"
import "math"

const s string = "constant"
//const declares a constant value.

func main(){
    
    fmt.Println("4.Constants.go--------Start-----------")
    
    const n = 500000000
    // A const statement can appear anywhere a var statement can. 
    const d = 3e20 / n 
    // Constant expressions perform arithmetic with arbitrary precision.
    
    
    fmt.Println("const n = 500000000 \n const d = 3e20 / n =",d)
    fmt.Println("const n = 500000000 \n const d = 3e20 / n \n int64(d) =",int64(d))
    //A numeric constant has no type until it’s given one, such as by an explicit cast.

    fmt.Println("math.Sin(n) = ",math.Sin(n))
    // A number can be given a type by using it in a context that requires one,
    // such as a variable assignment or function call. 
    // For example, here math.Sin expects a float64.
    fmt.Println("4.Constants.go--------End-----------")
}
