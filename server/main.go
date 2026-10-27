package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

// /search?mirror=N&ms=D answers after D milliseconds, unless the client gives up first.
func main() {
	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		mirror := r.URL.Query().Get("mirror")
		ms, _ := strconv.Atoi(r.URL.Query().Get("ms"))
		start := time.Now()
		select {
		case <-time.After(time.Duration(ms) * time.Millisecond):
			log.Printf("mirror %s: answered after %v", mirror, time.Since(start).Round(time.Millisecond))
			fmt.Fprintf(w, "result from mirror %s", mirror)
		case <-r.Context().Done():
			log.Printf("mirror %s: client gone after %v, stopped", mirror, time.Since(start).Round(time.Millisecond))
		}
	})
	log.Fatal(http.ListenAndServe("localhost:9091", nil))
}
