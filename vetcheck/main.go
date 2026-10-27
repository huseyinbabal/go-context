package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	fmt.Println(ctx.Err())
}
