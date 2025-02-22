package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- "result"
	}()

	//each case in select must be a channel send or recieve, it block until one of the cases proceed, if multiple cases are ready
	// it chooses one at random, using the default statement it makes the select statement non-blocking
	select {
	case res := <-ch:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout")
	}
}