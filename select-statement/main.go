package main

import (
	"fmt"
	"time"
)

func main() {
	// Creating channels
	ch1 := make(chan string)
	ch2 := make(chan string)

	// Goroutine to send data to ch1 after 2 seconds
	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "Message from ch1"
	}()

	// Goroutine to send data to ch2 after 1 second
	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "Message from ch2"
	}()

	// Use select to listen to multiple channels
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("Received:", msg1)
		case msg2 := <-ch2:
			fmt.Println("Received:", msg2)
		}
	}
}
