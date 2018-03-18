package main

import (
    "fmt"
    "flag"
    "github.com/sparrc/go-ping"
    "os"
)

type Target struct {
    Target  string
}

var target *Target

func init() {
    target = &Target{}
    flag.StringVar(&target.Target, "cible", "", "cible du ping")
}

func main() {

    flag.Parse()
    
    //fmt.Print("target: ", target.Target)

    if  len(target.Target) == 0 {
        fmt.Println("You must specify host target")
        os.Exit(3)
    }


	pinger, err := ping.NewPinger(target.Target)
	if err != nil {
    	    panic(err)
	}
	pinger.Count = 1
    pinger.Run()
    stats := pinger.Statistics()
    fmt.Println(stats)
}

