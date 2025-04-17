package main

import (
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
		{ID: 2, Taskname: "Task 2"},
		{ID: 3, Taskname: "Task 3"},
		{ID: 4, Taskname: "Task 4"},
		{ID: 5, Taskname: "Task 5"},
		{ID: 6, Taskname: "Task 6"},
	}

	rateLimit := time.NewTicker(1 * time.Second)
	defer rateLimit.Stop()

	for _, task := range tasks {
		<-rateLimit.C // throttle BEFORE launching
		go processTask(task)
	}

	// Wait long enough for all goroutines to complete
	time.Sleep(7 * time.Second)
}

func processTask(task Task) {
	fmt.Println("Processing:", task.Taskname, "at", time.Now().Format("15:04:05"))
	time.Sleep(500 * time.Millisecond) // simulate work
}
