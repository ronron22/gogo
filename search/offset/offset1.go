package main

import (
	"os"
	"fmt"
	"log"
	"bufio"
)

func main() {

fd, err := os.Open("input.txt")
if err != nil { //error handler
    log.Println(err)
    return
}

reader := bufio.NewReader(fd) // creates a new reader

_, err = reader.Discard(64) // discard the following 64 bytes
if err != nil { // error handler
    log.Println(err)
    return
}

// use isPrefix if is needed, this example doesn't use it
// read line until a new line is found
line, _, err := reader.ReadLine() 
if err != nil { // error handler
    log.Println(err)
    return
}
fmt.Println(string(line))

}
