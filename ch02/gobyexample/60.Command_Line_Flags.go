/*
Title: 60.Command_Line_Flags.go
Author: OpenSource
Date: 2017-05-23
Description: For Study

Command-line flags are a common way to specify options for command-line programs. 
For example, in wc -l the -l is a command-line flag.

*/

package main 

import "fmt"
import "flag"
// Go provides a flag package supporting basic command-line flag parsing. 
// We’ll use this package to implement our example command-line program.

func main(){
    fmt.Println("60.Command_Line_Flags.go---------Start------------\n\n")

    // Basic flag declarations are available for string, integer, and boolean options.
    // Here we declare a string flag word with a default value "foo" and a short description.
    // This "flag.String" function returns a string pointer (not a string value);
    // we’ll see how to use this pointer below.
    wordPtr := flag.String("word", "foo", "a string")

    // This declares numb and fork flags, using a similar approach to the word flag.
    numbPtr := flag.Int("numb", 42, "an int")
    boolPtr := flag.Bool("fork", false, "a bool")

    // It’s also possible to declare an option 
    // that uses an existing var declared elsewhere in the program. 
    // Note that we need to pass in a pointer to the flag declaration function.
    var svar string
    flag.StringVar(&svar, "svar", "bar", "a string var")

    // Once all flags are declared, call flag.Parse() to execute the command-line parsing.
    flag.Parse()

    // Here we’ll just dump out the parsed options and any trailing positional arguments.
    // Note that we need to dereference the pointers with e.g. 
    // *wordPtr to get the actual option values.
    fmt.Println("word=> ", *wordPtr)
    fmt.Println("numb=> ", *numbPtr)
    fmt.Println("fork=> ", *boolPtr)
    fmt.Println("svar=> ", svar)
    fmt.Println("tail=> ", flag.Args())


    fmt.Println("\n\n60.Command_Line_Flags.go-----------End------------")
    // To experiment with the command-line flags program 
    // it’s best to first compile it and then run the resulting binary directly.
    // Try out the built program by first giving it values for all flags.

    // References, https://gobyexample.com/command-line-flags
    // this source is good at running exam
}
