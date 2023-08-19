package main

import (
    "fmt"
    "regexp"
)

func main() {
    str := "This is string in golang"

    match, err := regexp.MatchString(`\s*stri.g`, str)
    fmt.Println(match) // true
    fmt.Println(err)  // nil

    match, err = regexp.MatchString(`\s*STRI.G`, str)
    fmt.Println(match) // false
    fmt.Println(err) // nil

    match, err = regexp.MatchString(`(?i)\s*STRI.G`, str)
    fmt.Println(match) // true
    fmt.Println(err)  // nil
}
