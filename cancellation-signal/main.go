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

func main() {
	tasks := []Task{
		{ID: 1, Taskname: "Task 1"},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	for _, task := range tasks {
		time.Sleep(500 * time.Millisecond)
		go processTask(ctx, task)
	}

	// Wait long enough for all goroutines to complete
	time.Sleep(5 * time.Second)
}

func processTask(ctx context.Context, task Task) {
	for {
		time.Sleep(500 * time.Millisecond)
		select {
		case <-ctx.Done():
			fmt.Println("Cancelled:", task.Taskname)
			return
		default:
			fmt.Println("Processing:", task.Taskname, "at", time.Now().Format("15:04:05"))
		}
	}
}
