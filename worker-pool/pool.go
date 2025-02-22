package main

import (
	"fmt"
	"time"
)

// Worker function that processes jobs
func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job)
		time.Sleep(time.Second) 
		results <- job * 2      
	}
}

func main() {
	numWorkers := 3 
	numJobs := 5    

	jobs := make(chan int, numJobs)   // unbuffered channels, shared job queue 
	results := make(chan int, numJobs) 

	for i := 1; i <= numWorkers; i++ {
		go worker(i, jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs) 

	// Collect results
	for r := 1; r <= numJobs; r++ {
		fmt.Println("Result:", <-results)
	}
}
