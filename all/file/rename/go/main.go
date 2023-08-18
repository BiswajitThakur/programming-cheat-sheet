package main

import (
    "fmt"
    "os"
)

func rename(old_name string, new_name string) error {
    return os.Rename(old_name, new_name)
}

func main() {
    old_name := "tmp.old"
    new_name := "tmp.new"
    
    err := rename(old_name, new_name)

    if err != nil {
        fmt.Println("Failed to rename")
        panic(err)
    } else {
        fmt.Println("Successfully renamed")
    }
}
