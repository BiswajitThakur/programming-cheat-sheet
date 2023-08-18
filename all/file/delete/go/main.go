package main

// delete file and empty directory

import (
    "fmt"
    "os"
)

func main() {
    file_name := "enemy.txt"
    err := os.Remove(file_name)
    if err != nil {
        panic(err)
    } else {
        fmt.Println("Successfully deleted file: ", file_name)
    }
}

