/*
Title: 40.Sorting_By_Functions.go
Author: OpenSource
Date: 2017-05-22
Description: For Study

Sometimes we’ll want to sort a collection by something other than its natural order. 
For example, suppose we wanted to sort strings by their length instead of alphabetically. 
Here’s an example of custom sorts in Go.

*/

package main 

import (
    "sort"
    "fmt"
)

// In order to sort by a custom function in Go, we need a corresponding type. 
// Here we’ve created a ByLength type that is just an alias for the builtin []string type.
type ByLength []string

// We implement sort.Interface - Len, Less, and Swap - 
// !!! So Len , Swap , Less function is absolutely needable function.
// !!! cause'  sort.Sort(ByLength(aryStr)) is rquired them
// on our type so we can use the sort package’s generic Sort function. 
// Len and Swap will usually be similar across types 
// and Less will hold the actual custom sorting logic. 

// In our case we want to sort in order of increasing string length,
// so we use len(s[i]) and len(s[j]) here.
func (s ByLength) Len() int{
    return len(s)
}

func (s ByLength) Sawp(i, j int){
    s[i], s[j] = s[j], s[i]
}

func (s ByLength) Less(i, j int) bool {
    return len(s[i]) < len(s[j])
}

// With all of this in place,
// we can now implement our custom sort by casting the original fruits slice to ByLength, 
// and then use sort.Sort on that typed slice.
func main(){
    fmt.Println("40.Sorting_By_Functions.go---------Start------------\n\n")

    fruits := []string{"peach", "banana", "kiwi"}
    fmt.Println(fruits)
    sort.Sort(ByLength(fruits))
    fmt.Println(fruits)


    fmt.Println("40.Sorting_By_Functions.go---------Start------------\n\n")
    // Running our program shows a list sorted by string length, as desired.

    // By following this same pattern of creating a custom type, 
    // implementing the three Interface methods on that type, and then calling sort.Sorting 
    // on a collection of that custom type, we can sort Go slices by arbitrary functions.
}
