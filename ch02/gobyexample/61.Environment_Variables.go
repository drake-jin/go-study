/*
Title: 61.Environment_Variables.go
Author: OpenSource
Date: 2017-05-23
Description: For Study

Environment variables are a universal mechanism for conveying configuration information 
to Unix programs. Let’s look at how to set, get, and list environment variables.

*/

package main 

import "os"
import "strings"
import "fmt"

func main(){
    fmt.Println("61.Environment_Variables.go---------Start------------\n\n")

    // To set a key/value pair, use os.Setenv. To get a value for a key, use os.Getenv.
    // This will return an empty string if the key isn’t present in the environment.
    os.Setenv("FOO", "1")
    fmt.Println("FOO =>", os.Getenv("FOO"))
    fmt.Println("BAR =>", os.Getenv("BAR"))
    
    fmt.Println()

    // Use os.Environ to list all key/value pairs in the environment.
    // This returns a slice of strings in the form KEY=value.
    // You can strings.Split them to get the key and value. Here we print all the keys.
    for _, e := range os.Environ(){
        pair := strings.Split(e, "=") 
        fmt.Println(pair[0])
    }


    // Running the program shows that we pick up the value for FOO that we set in the program, but that BAR is empty.
    // The list of keys in the environment will depend on your particular machine.

    // If we set BAR in the environment first, the running program picks that value up.

    fmt.Println("\n\n61.Environment_Variables.go-----------End------------")
}

