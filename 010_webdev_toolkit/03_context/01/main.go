package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", foo)
	http.ListenAndServe(":8081", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx = context.WithValue(ctx, "userID", 777)
	ctx = context.WithValue(ctx, "fname", "Bond")

	results, err := dbAccess(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusRequestTimeout)
		// return
	}

	fmt.Fprintln(w, results)
}

func dbAccess(ctx context.Context) (int, error) {

	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	ch := make(chan int)

	go func() {
		uid := ctx.Value("userID").(int)
		time.Sleep(10 * time.Second)
		if ctx.Err() != nil {
			fmt.Println(ctx.Err())
			return
		}
		ch <- uid
	}()

	select {
	case <-ctx.Done():
		fmt.Println("1")
		return 0, ctx.Err()
	case i := <-ch:
		fmt.Println("2")
		return i, nil
	}
}
