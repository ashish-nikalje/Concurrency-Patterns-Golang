package main

import (
	"fmt"
	"time"
)

func main() {
	// Create a ticker that ticks every 2 seconds
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop() // Stop the ticker when done

	// Run the task periodically until the program is manually stopped
	for i := 0; i < 5; i++ {
		select {
		case <-ticker.C:
			// Perform a periodic task
			fmt.Println("Performing periodic task", i+1, time.Now().Format("15:04:05"))
		}
	}

	// Finish program execution
	fmt.Println("Program complete.")
}
