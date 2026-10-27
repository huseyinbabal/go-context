package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

func search(ctx context.Context, mirror, ms int) (string, error) {
	url := fmt.Sprintf("http://localhost:9091/search?mirror=%d&ms=%d", mirror, ms)
	req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	return string(body), err
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	results := make(chan string, 3)
	for mirror, ms := range map[int]int{1: 1500, 2: 300, 3: 2500} {
		go func() {
			if r, err := search(ctx, mirror, ms); err == nil {
				results <- r
			}
		}()
	}

	fmt.Println(<-results)
	cancel() // the other two requests stop here
}
