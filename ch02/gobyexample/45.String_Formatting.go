/*
Title: 45.String_Formatting.go
Author: OpenSource
Date: 2017-05-23
Description: For Study

Go offers excellent support for string formatting in the printf tradition. 
Here are some examples of common string formatting tasks.

*/

package main

import "fmt"
import "os"

type point struct{
    x, y int
}


func main() {
    fmt.Println("45.String_Formatting.go---------Start------------\n\n")

    p := point{1, 2}
    
    // Go offers several printing “verbs” designed to format general Go values.
    // For example, this prints an instance of our point struct.
    fmt.Printf("%v \n", p) //{1 2} 

    // If the value is a struct, the %+v variant will include the struct’s field names.
    fmt.Printf("%+v \n", p) // {x:1 y:2}

    // The %#v variant prints a Go syntax representation of the value, i.e.
    // the source code snippet that would produce that value.
    fmt.Printf("%#v \n", p) // main.point{x:1, y:2}

    // To print the type of a value, use %T
    fmt.Printf("%T \n", p) // main.point

    // Formatting booleans is straight-forward
    fmt.Printf("%t \n", true) // true

    // There are many options for formatting integers. Use %d for standard, base-10 formatting.
    fmt.Printf("%d \n", 123) // 123

    // This prints a binary representation.
    fmt.Printf("%b \n", 14) // 1110

    // This prints the character corresponding to the given integer.
    fmt.Printf("%c \n", 33) // !

    // %x provides hex encoding.
    fmt.Printf("%x \n", 456) // 1cb

    // There are also several formatting options for floats. For basic decimal formatting use %f.
    fmt.Printf("%f \n", 78.9) // 78.900000

    // %e and %E format the float in (slightly different versions of) scientific notation.
    fmt.Printf("%e \n", 12340000.0) // 1.234.000e+07
    fmt.Printf("%E \n", 12340000.0) // 1.234.000E+07

    // For basic string printing use %s.
    fmt.Printf("%s \n", "\"strings\"") // "strings"

    // To double-quote strings as in Go source, use %q.
    fmt.Printf("%q \n", "\"strings\"") // "\"strings\""

    // As with integers seen earlier, %x renders the string in base-16,
    // with two output characters per byte of input.
    fmt.Printf("%x \n", "hex this") // 6865782074686973

    // To print a representation of a pointer, use %p.
    fmt.Printf("%p \n", &p) // 0xc42000e2c0

    // When formatting numbers you will often want to control the width and precision 
    // of the resulting figure. To specify the width of an integer, 
    // use a number after the % in the verb. 
    // By default the result will be right-justified and padded with spaces.
    fmt.Printf("|%6d|%6d| \n", 12, 345) // |    12|   345|

    // You can also specify the width of printed floats, 
    // though usually you’ll also want to restrict the decimal precision at the same time 
    // with the width.precision syntax.
    fmt.Printf("|%6.2f|%6.2f| \n", 1.2, 3.45) // |  1.20|  3.45|

    // To left-justify, use the - flag.
    fmt.Printf("|%-6.2f|%-6.2f| \n", 1.2, 3.45) // |1.20  |3.45  |

    // You may also want to control width when formatting strings, 
    // especially to ensure that they align in table-like output. For basic right-justified width.
    fmt.Printf("|%6s|%6s| \n", "foo", "b") // |   foo|     b|

    // To left-justify use the - flag as with numbers.
    fmt.Printf("|%-6s|%-6s| \n", "foo", "b") // |foo   |b     |

    // So far we’ve seen Printf, which prints the formatted string to os.Stdout.
    // Sprintf formats and returns a string without printing it anywhere.
    s := fmt.Sprintf("a %s", "string") // a string
    fmt.Println(s)

    // You can format+print to io.Writers other than os.Stdout using Fprintf.
    fmt.Fprintf(os.Stderr, "an %s \n", "error") // an error

    fmt.Println("\n\n45.String_Formatting.go-----------End------------")
}















