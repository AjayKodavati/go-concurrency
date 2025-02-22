package main

import (
	"fmt"
	"sync"
	"time"
)

func generator(name string, n int) <-chan string {
	ch := make(chan string)

	go func() {
		defer close(ch)
		for i := 1; i <= n; i++ {
			time.Sleep(1 * time.Second)
			ch <- fmt.Sprintf("%s - %d", name, i)
		}
	}()

	return ch
}

func fanIn(ch... <- chan string) <-chan string {
	merged := make(chan string)
	var wg sync.WaitGroup

	for _, ch := range ch {
		wg.Add(1)
		go func (c <-chan string)  {
			defer wg.Done()
			for n := range c {
				merged <- n
			}
		}(ch)
	}
	
	go func() { 
		wg.Wait() 
		close(merged)
	}()
	return merged
}

func main() {
	ch1 := generator("Source1", 5)
	ch2 := generator("Source2", 5)

	merged := fanIn(ch1, ch2)

	for msg := range merged {
		fmt.Println(msg)
	}
}