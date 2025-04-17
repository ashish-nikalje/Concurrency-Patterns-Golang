package main

import (
	"fmt"
	"sync"
	"time"
)

type Task struct {
	ID   int
	Name string
}

func main() {
	fmt.Println("Starting bounded queue example...")
	taskQueue := make(chan Task, 3) // Bounded queue with max capacity of 3
	var wg sync.WaitGroup

	// Producer: sending tasks
	wg.Add(1)
	go func() {
		defer wg.Done()

		for i := 1; i <= 5; i++ {
			taskQueue <- Task{ID: i, Name: fmt.Sprintf("Task %d", i)} // Adds tasks to the queue
			fmt.Printf("Produced: Task %d\n", i)
			time.Sleep(500 * time.Millisecond)
		}
		close(taskQueue) // Close the channel after all tasks are sent
	}()

	// Consumer: processing tasks
	go func() {
		for task := range taskQueue {
			wg.Add(1)
			go processTask(task, &wg)
		}
	}()

	wg.Wait()
}

func processTask(task Task, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Processing: %s\n", task.Name)
	time.Sleep(1 * time.Second) // Simulate task processing
	fmt.Printf("Completed: %s\n", task.Name)
}
