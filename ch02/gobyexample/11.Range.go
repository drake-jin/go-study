/*
Title: 11.Range.go
Author: OpenSource
Date: 2017-05-21
Description: For Study

range iterates over elements in a variety of data structures. 
Let’s see how to use range with some of the data structures we’ve already learned.

*/

package main 

import "fmt"

func main() {
    fmt.Println("11.Range.go--------Start-----------")
    
    nums := []int{1, 2, 3}
    sum := 0 
    // Here we use range to sum the numbers in a slice. Arrays work like this too.

    for _, num := range nums {
        sum += num
    }
    fmt.Println("sum: ", sum)

    // range on arrays and slices provides both the index and value for each entry.
    // Above we didn’t need the index, so we ignored it with the blank identifier _. 
    // Sometimes we actually want the indexes though.
    for i, num := range nums {
        if num == 3 {
            fmt.Println("index: ",i)
            fmt.Println("num: ",num)
        }
    }
    
    // range on map iterates over key/value pairs.
    kvs := map[string]string{"a": "apple", "b": "banana"}
    for k, v := range kvs {
        fmt.Println("%s -> %s\n => (k, v) :", k, v)
    }

    // range can also iterate over just the keys of a map.
    for k := range kvs {
        fmt.Println("key :", k)
    }

    // range on strings iterates over Unicode code points. 
    // The first value is the starting byte index of the rune and the second the rune itself.
    for i, c := range "go"{
        fmt.Println(i, c)
    }

    fmt.Println("11.Range.go---------End------------")
}
