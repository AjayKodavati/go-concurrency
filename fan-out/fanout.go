package main

import (
	"fmt"
)

// worker pool: A fixed pool of goroutines processes tasks from a shared channel,
func worker(jobId int, jobs <-chan int, result chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", jobId, job)
		result <- jobId * 2;
	}
}

func main() {
	// shared channel among multiple go-routines.
	jobs := make(chan int, 100)
	result := make(chan int, 100)

	//create worker go-routines
	for i := 0; i < 3; i++ {
		go worker(i, jobs, result)
	}

	// creating jobs for processing
	for i := 0; i < 5; i++ {
		jobs <- i + 1
	}
	close(jobs)

	for i := 0; i < 5; i++ {
		fmt.Printf("result: %v\n", <-result)
	}
}
