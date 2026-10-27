package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

// resize stands in for work that never touches the network: 200 ms per image.
func resize(ctx context.Context, images int) (int, error) {
	for i := range images {
		select {
		case <-ctx.Done():
			return i, ctx.Err()
		default:
		}
		time.Sleep(200 * time.Millisecond)
	}
	return images, nil
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	done, err := resize(ctx, 20)
	fmt.Printf("resized %d of 20 images\n", done)
	fmt.Println("timed out:", errors.Is(err, context.DeadlineExceeded))
	fmt.Println("canceled: ", errors.Is(err, context.Canceled))
}
