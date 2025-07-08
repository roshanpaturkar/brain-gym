package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()

	contextTimeout(ctx)
}

func contextTimeout(ctx context.Context) {
	contextWithTimeout, cancel := context.WithTimeout(ctx, 2 * time.Second)

	defer cancel()

	done := make(chan struct{})


	go func ()  {
		time.Sleep(3 * time.Second)
		close(done)
	}()

	select {
		case <- done:
			fmt.Println("API called!")
		case <- contextWithTimeout.Done():
			fmt.Println("API timeout:", contextWithTimeout.Err())
			// add handeling logic
	}
}
