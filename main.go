package main

import (
    "bufio"
    "fmt"
    "os"
)

func checkErr(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    mdFilePath := os.Args[1]
    file, err := os.Open(mdFilePath)
    checkErr(err)
    defer file.Close()
    scanner := bufio.NewScanner(file)
    tmpString := []string{""}
    for scanner.Scan() {
        tmpString = append(tmpString, scanner.Text())
    }
    fmt.Println(tmpString)
}
