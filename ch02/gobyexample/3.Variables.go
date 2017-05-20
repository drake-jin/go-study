/*
Title: 3.Variables.go
Author: OpenSource
Date: 2017-05-20 
Description: For Study

In Go, variables are explicitly declared and used by the compiler to e.g. 
check type-correctness of function calls.


*/

package main

import "fmt"

func main(){
    fmt.Println("3.Variables.go--------Start-----------")
    
    var a string = "initial"
    //var declares 1 or more variables.
    fmt.Println("var a sring = \"initial\"= ",a)

    var b, c int = 1, 2
    //You can declare multiple variables at once.
    fmt.Println("var b, c int = 1, 2 =",b,c)

    var d = true
    //Go will infer the type of initialized variables.
    fmt.Println("var d = true = ",d)

    var e int
    //Variables declared without a corresponding initialization are zero-valued. 
    // For example, the zero value for an int is 0.
    fmt.Println("var e int = ",e)

    f := "shortu"
    //The := syntax is shorthand for declaring and initializing a variable, 
    //e.g. for var f string = "short" in this case.
    fmt.Println("f := \"short\"", f)

    fmt.Println("3.Variables.go--------End-----------")
}
