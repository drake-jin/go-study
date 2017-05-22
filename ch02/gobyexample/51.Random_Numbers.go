/*
Title: 51.Random_Numbers.go
Author: OpenSource
Date: 2017-05-23
Description: For Study


Go’s math/rand package provides pseudorandom number generation.
*/

package main 

import "time"
import "fmt"
import "math/rand"

func main(){
    fmt.Println("51.Random_Numbers.go---------Start------------\n\n")
    
    // For example, rand.Intn returns a random int n, 0 <= n < 100.
    fmt.Println(rand.Intn(100), "|", rand.Intn(100))
    // rand.Float64 returns a float64 f, 0.0 <= f < 1.0.
    rmt.Println(rand.Float64())
    // This can be used to generate random floats in other ranges, for example 5.0 <= f' < 10.0.
    fmt.Println((rand.Float64()*5)+5, "| ", (rand.Float64()*5)+5)

    // The default number generator is deterministic, 
    // so it’ll produce the same sequence of numbers each time by default. 
    // To produce varying sequences, give it a seed that changes.
    // Note that this is not safe to use for random numbers you intend to be secret, 
    // use crypto/rand for those.
    s1 := rand.NewSource(time.Now().UnixNano())
    r1 := rand.New(s1)
    // Call the resulting rand.Rand just like the functions on the rand package.
    fmt.Println(r1.Intn(100), "|", r1.Intn(100))

    // If you seed a source with the same number, it produces the same sequence of random numbers.
    s2 := rand.NewSource(42)
    r2 := rand.New(s2)
    fmt.Println(r2.Intn(100),"|",r2.Intn(100))

    s3 := rand.NewSource(42)
    r3 := rand.New(s3)
    fmt.Println(r3.Intn(100), "|", r3.Intn(100))

    fmt.Println("\n\n51.Random_Numbers.go-----------End------------")
    // See the math/rand package docs for references on other random quantities that Go can provide.
}
