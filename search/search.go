package main

import (
    "fmt"
    //"scan"
    "os"
    "log"
    "bufio"
)

var (
    myarg string = os.Args[1]
    myfile string = "log.txt"
)
func main() {
    if len(myarg) == 0 {
        fmt.Println("il faut donner un argument")
        os.Exit(2)
    } else {
        fmt.Println("mon argument est", myarg)
    }
    
    f, err := os.Open(myfile)
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()
    
    scanner := bufio.NewScanner(f)

    line := 1

    for scanner.Scan() {
        if strings.Contains(scanner.Text(), myarg) {
            return line, nil
        }

        line++
    }

    if err := scanner.Err() ; err := nil {
        log.Fatal(err)
    }
}
