package main

// https://gobyexample.com/interfaces

import (
    "fmt"
    "os"
    "github.com/sparrc/go-ping"
)

type  NetworkTests interface {
    TestPing()  *ping.Statistics
    //TestPing()  string
}

type Target struct {
    MyTarget  string
}

func (h Target) TestPing() *ping.Statistics {

    if  len(h.MyTarget) == 0 {
        fmt.Println("You must specify host target")
        os.Exit(3)
    }

	pinger, err := ping.NewPinger(h.MyTarget)
	if err != nil {
    	    panic(err)
	}
	pinger.Count = 1
    pinger.Run()
    stats := pinger.Statistics()
    //fmt.Println(stats)
    return stats
}

func Testing(u  NetworkTests) {
    fmt.Println(u)
    fmt.Println(u.TestPing())
}

func main() {
    
    h := Target{MyTarget: "google.fr"}
    
    Testing(h)
}
    
