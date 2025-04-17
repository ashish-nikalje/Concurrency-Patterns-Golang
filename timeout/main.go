package main

import (
	"context"
	"fmt"
	"time"
)

type Task struct {
	ID       uint
	Taskname string
}

func PrintTaskDetails(ctx context.Context, task Task, done chan<- string) {
	time.Sleep(3 * time.Second)
	result := fmt.Sprintf("Task ID: %d\nTask Name: %s", task.ID, task.Taskname)

	// Try to send result if context not canceled
	select {
	case <-ctx.Done():
		// Context timed out or canceled
		return
	case done <- result:
	}
}

func main() {
	task := Task{
		ID:       1,
		Taskname: "Example Task",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// Create a channel to receive the result
	resultChan := make(chan string)

	go PrintTaskDetails(ctx, task, resultChan)

	// Wait for either result or timeout
	select {
	case res := <-resultChan:
		fmt.Println("Task completed:\n" + res)
	case <-ctx.Done():
		fmt.Println("Timeout or cancelled:", ctx.Err())
	}
}
