/*
Title: 9.Slices.go
Author: OpenSource
Date: 2017-05-20
Description: For Study

Slices are a key data type in Go, giving a more powerful interface to sequences than arrays.

*/

package main

import "fmt"

func main() {  
    fmt.Println("9.Slices.go--------Start-----------")
    
    s := make([]string, 3)
    // Unlike arrays, slices are typed only by the elements they contain 
    // (not the number of elements). 
    // To create an empty slice with non-zero length, use the builtin make.
    // Here we make a slice of strings of length 3 (initially zero-valued).
    fmt.Println("s := make([]string,3) =>", s)
   
    fmt.Println(s)
    // We can set and get just like with arrays.
    s[0] = "a"
    s[1] = "b"
    s[2] = "c"

    fmt.Println("s[0] =\"a\", s[1]=\"b\",s[2]=\"c\" => ",s)
    fmt.Println("len(s) => ", len(s))
    // len returns the length of the slice as expected.

    s = append(s, "d")
    s = append(s, "e","f")
    fmt.Println("s = append(s, \"d\",\"e\",\"f\" =>  ", s)
    // In addition to these basic operations, 
    // slices support several more that make them richer than arrays.
    // One is the builtin append, which returns a slice containing one or more new values. 
    // Note that we need to accept a return value from append as we may get a new slice value.
    
    c := make([]string, len(s))
    fmt.Println("c := make([]string, len(s)) =>", c)
    // Slices can also be copy’d. 
    // Here we create an empty slice c of the same length as s and copy into c from s.
    copy(c, s)
    fmt.Println("s =>", s)
    fmt.Println("after copy(c,s) => c ", c)

    // Slices support a “slice” operator with the syntax slice[low:high].
    // For example, this gets a slice of the elements s[2], s[3], and s[4].
    l := s[2:5]
    fmt.Println("l := s[2:5] =>", l)

    // This slices up to (but excluding) s[5].
    l = s[:5]
    fmt.Println("l = s[:5] =>", l)
    
    // And this slices up from (and including) s[2].
    l = s[2:] 
    fmt.Println("l = s[2:] =>", l)

    // We can declare and initialize a variable for slice in a single line as well.
    t := []string{"g","h","i"}
    fmt.Println("t := []string{\"g\",\"h\",\"i\"} =>", t)

    // Slices can be composed into multi-dimensional data structures. 
    // The length of the inner slices can vary, unlike with multi-dimensional arrays.
    twoD := make([][]int, 3)
    for i := 0; i < 3; i ++ {
        innerLen := i + 1 
        twoD[i] = make([]int, innerLen)
        for j:= 0; j < innerLen; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("twoD := make([][]int,3) =>", twoD)
    // Note that while slices are different types than arrays, 
    // they are rendered similarly by fmt.Println.
    // Check this blog that https://blog.golang.org/go-slices-usage-and-internals
    // It's really awesome 
    fmt.Println("9.Slices.go---------End------------")
}

