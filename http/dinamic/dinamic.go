package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func getTime(w http.ResponseWriter, r *http.Request) {
	s := time.Now().Format("02/01/2006 03:04:05") // standard format string
	fmt.Fprintf(w, "<h1>Time now: %s<h1>", s)
}

func main() {
	http.HandleFunc("/time", getTime)
	log.Println("Executing...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
