package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	srv := &http.Server{Addr: "localhost:9092"}
	http.HandleFunc("/report", func(w http.ResponseWriter, r *http.Request) {
		log.Println("report: started")
		time.Sleep(3 * time.Second)
		fmt.Fprintln(w, "report ready")
		log.Println("report: sent")
	})
	go srv.ListenAndServe()
	log.Println("listening on :9092")

	<-ctx.Done() // Ctrl+C closes this
	log.Println("shutting down, waiting for open requests")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Println("shutdown:", err)
	}
	log.Println("bye")
}
