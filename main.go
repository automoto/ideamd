package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
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
    str := strings.Join(tmpString, " ")
    //for index, line := range tmpString {
    //    h, isHeadline := pkg.CheckHeadline(line)
    //    if isHeadline {
    //        // add check if this an epic
    //    }
    //}

    fmt.Println(tmpString)
}
