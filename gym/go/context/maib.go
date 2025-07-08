package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {
	// ctx := context.Background()

	// contextTimeout(ctx)
	// contextWithValue(ctx)

	http.HandleFunc("/hello", helloHandler)
	http.ListenAndServe(":8080", nil)
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

func contextWithValue(ctx context.Context) {
	type key int
	const UserKey key = 0

	contextWithValue := context.WithValue(ctx, UserKey, "333221")
	// contextWithValue := context.WithValue(ctx, UserKey, nil)

	if userId, ok := contextWithValue.Value(UserKey).(string); ok {
		fmt.Println("User authenticated with userId:", userId)
	} else {
		fmt.Println("Protected rout, userId is missing!")
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	select {
		case <- time.After(3*time.Second):
			fmt.Println("API response!")
		case <- ctx.Done():
			fmt.Println("Context expired!")
			http.Error(w, "Request contect timeout", http.StatusRequestTimeout)
			return
	}
}
