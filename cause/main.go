package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"
)

func main() {
	ctx, cancel := context.WithCancelCause(context.Background())
	go func() {
		time.Sleep(300 * time.Millisecond)
		cancel(errors.New("mirror 2 answered first"))
	}()

	req, _ := http.NewRequestWithContext(ctx, "GET", "http://localhost:9091/search?mirror=1&ms=1500", nil)
	_, err := http.DefaultClient.Do(req)

	fmt.Println("request error:", err)
	fmt.Println("ctx.Err():    ", ctx.Err())
	fmt.Println("Cause(ctx):   ", context.Cause(ctx))
}
