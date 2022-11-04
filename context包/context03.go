package main

import (
	"context"
	"fmt"
	"time"
)

// WithDeadline

func main() {
	d := time.Now().Add(5 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), d) // ctx d秒后会过期

	defer cancel()
	select {
	case <-time.After(10 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}

}
