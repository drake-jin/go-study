/*
Title: 8.Arrays.go
Author: OpenSource
Date: 2017-05-20
Description: For Study

In Go, an array is a numbered sequence of elements of a specific length.

*/

package main 

import "fmt"

func main() {
    fmt.Println("8.Arrays.go--------Start-----------")

    // Here we create an array a that will hold exactly 5 ints. 
    // The type of elements and length are both part of the array’s type. 
    // By default an array is zero-valued, which for ints means 0s.
    var a [5]int
    fmt.Println("var a [5]int => ",a)
  
    a[4] = 100 
    fmt.Println("a[4] = 100 =>", a)
    fmt.Println("a[4] =>",a[4])
    // We can set a value at an index using the array[index] = value syntax, 
    // and get a value with array[index].
    
    fmt.Println("len: ", len(a))
    // The builtin len returns the length of an array.

    b := [5]int{1, 2, 3, 4, 5}
    // Use this syntax to declare and initialize an array in one line.
    fmt.Println("b := [5]int{1,2,3,4,5}", b)

    var twoD [2][3]int 
    // Array types are one-dimensional, 
    // but you can compose types to build multi-dimensional data structures.
    for i := 0; i < 2; i++ {
        for j := 0; j < 3; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("var two [2][3]int", twoD)
    // Note that arrays appear in the form [v1 v2 v3 ...] when printed with fmt.Println.

    fmt.Println("8.Arrays.go--------End----------")
}


