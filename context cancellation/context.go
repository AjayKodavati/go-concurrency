package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context, id int) {
	for {
		select { // select case statements should be send/ receive
		case <-ctx.Done():
			fmt.Printf("Worker %d canceled\n", id)
            return
		default:
            fmt.Printf("Worker %d working...\n", id)
            time.Sleep(1 * time.Second)
		}
	}
}

//The WithCancel, WithDeadline, and WithTimeout functions take a Context (the parent) and return a derived Context
//  (the child) and a CancelFunc. Calling the CancelFunc directly cancels the child and its children, 
// removes the parent's reference to the child, and stops any associated timers. Failing to call the CancelFunc leaks the child and its children until the parent is canceled.

//A Context may be canceled to indicate that work done on its behalf should stop.
//  A Context with a deadline is canceled after the deadline passes. 
// When a Context is canceled, all Contexts derived from it are also canceled.

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go worker(ctx, 1)
	go worker(ctx, 2)

	time.Sleep(3 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)
}