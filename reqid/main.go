package main

import (
	"context"
	"crypto/rand"
	"fmt"
	"log"
	"net/http"
)

type ctxKey struct{} // an unexported key type: no other package can collide with it

func withRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := rand.Text()[:8]
		ctx := context.WithValue(r.Context(), ctxKey{}, id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func requestID(ctx context.Context) string {
	id, _ := ctx.Value(ctxKey{}).(string)
	return id
}

func search(ctx context.Context, q string) string {
	log.Printf("[%s] searching for %q", requestID(ctx), q)
	return "3 results"
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, search(r.Context(), r.URL.Query().Get("q")))
	})
	log.Println("listening on :9190")
	log.Fatal(http.ListenAndServe("localhost:9190", withRequestID(mux)))
}
