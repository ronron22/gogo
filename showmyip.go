package main


import (
	"fmt"
	"log"
	"net/http"
	"github.com/Tomasen/realip"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("X-tagada: ", "pouet")
	w.WriteHeader(http.StatusOK)
	// recup du xforwardedfor 
	clientIP := realip.FromRequest(r)
	w.Header().Set("Content-Length", fmt.Sprint(len(clientIP)))
	fmt.Fprint(w, clientIP)
}

func main() {
	http.HandleFunc("/ip", rootHandler)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
