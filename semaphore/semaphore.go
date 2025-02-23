package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(wg *sync.WaitGroup, id int, sem chan struct{}) {
	defer wg.Done()

	// Aquire semaphore blocking
	sem <- struct{}{}

	fmt.Printf("Worker %d is processing...\n", id)
	time.Sleep(2 * time.Second) // Simulate work

	<-sem 
	fmt.Printf("Worker %d finished\n", id)
}

func main() {
	var wg sync.WaitGroup
	sem := make(chan struct{}, 3)

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go worker(&wg, i, sem)
	}

	wg.Wait()
	fmt.Println("All workers done")
}