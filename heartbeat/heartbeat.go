package main

import (
	"fmt"
	"time"
)

func worker(ch chan<- struct{}) {
	for {
		time.Sleep(2 * time.Second);
		fmt.Println("alive......")
		ch <- struct{}{}
	}
}

func monitor(ch <-chan struct{}, timeout time.Duration) {
	timer := time.NewTimer(timeout)
	
	for {
		select {
		case <-ch:
			fmt.Println("worker is alive.......!")
			if !timer.Stop() {
				<-timer.C
			}
			timer.Reset(timeout)
		case <-timer.C:
			fmt.Println("Monitor: No heartbeat received!")
			return
		}
	}
}

func main() {
	heartbeat := make(chan struct{})

	go worker(heartbeat)
	monitor(heartbeat, 5*time.Second)
}