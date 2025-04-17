package main

import (
	"fmt"
)

func main() {
	tasks := make(chan int)

	// Producer: sending values
	go func() {
		for i := 1; i <= 5; i++ {
			tasks <- i
		}
		close(tasks) // Closing the channel when done
	}()

	// Consumer: receiving values
	for {
		val, ok := <-tasks
		if !ok { // Check if the channel is closed
			fmt.Println("Channel closed, no more tasks.")
			break
		}
		fmt.Println("Processing task:", val)
	}
}
