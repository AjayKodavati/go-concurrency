package main

import "fmt"

func generator(in ...int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for _, num := range in {
			out <- num
		}
	}()
	return out
}

func multiply(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			out <- n * 2
		}
	}()

	return out
}

func add(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for n := range in {
			out <- n + 2
		}
	}()

	return out
}

func main() {
	// Connect the pipeline: generator → multiply → add
	nums := generator(1, 2, 3)
	multiplied := multiply(nums)
	summed := add(multiplied)

	for result := range summed {
		fmt.Println(result) 
	}
}