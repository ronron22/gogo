package main

// https://gist.github.com/dustin/1777162
// le modifier pour tourner sous inetd

import (
	"fmt"
	"log"
	"net/http"
	"net"	
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("X-tagada: ", "pouet")
	w.WriteHeader(http.StatusOK)
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Length", fmt.Sprint(len(ip)))
	fmt.Fprint(w, string(ip))
}

func main() {
	http.HandleFunc("/", rootHandler)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
