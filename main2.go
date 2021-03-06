package main

// https://gobyexample.com/interfaces

import (
    "fmt"
    "os"
    "github.com/sparrc/go-ping"
    "flag"
    "net"
)

type  NetworkTests interface {
    TestPing() *ping.Statistics
    TestDns() []net.IP 
}

type Target struct {
    MyTarget  string
}

var target *Target

func init() {
    target = &Target{}
    flag.StringVar(&target.MyTarget, "cible", "", "cible du ping")
}

func (h Target) TestPing() *ping.Statistics {

	pinger, err := ping.NewPinger(h.MyTarget)
	if err != nil {
    	    panic(err)
	}
	pinger.Count = 1
    pinger.Run()
    stats := pinger.Statistics()
    return stats
}

func (h Target) TestDns() []net.IP {
    ipaddress, err := net.LookupIP(h.MyTarget)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Could not get IPs: %v\n", err)
        os.Exit(2)
    }
    return ipaddress
}

func Testing(u  NetworkTests) {
    fmt.Println("ping de ", u)
    fmt.Println(u.TestPing())
    fmt.Println("lookup de ", u)
    fmt.Println(u.TestDns())
}

func main() {
    
    flag.Parse()

    if  len(target.MyTarget) == 0 {
        fmt.Println("You must specify host target")
        os.Exit(2)
    }

    h := Target{MyTarget: target.MyTarget}
    
    Testing(h)
}
    
