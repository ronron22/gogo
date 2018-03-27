package main

import (
	"github.com/hpcloud/tail"
	"fmt"
	"regexp"
	"os"
)

var motif string = os.Args[1]

func main() {
t, err := tail.TailFile("log.txt", tail.Config{Follow: true})
if err != nil {
	fmt.Println("err")
}
for line := range t.Lines {
    //fmt.Println(line.Text)
	matched, err := regexp.MatchString(motif, line.Text)  
	if err != nil {
		fmt.Println(err)
	} else {
		if matched == true {
			fmt.Println(line.Text, "ok")
		}
	}
}		
}
