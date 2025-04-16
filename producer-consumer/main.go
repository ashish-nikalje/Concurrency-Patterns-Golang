package main

import (
	"fmt"
	"sync"
)

type Task struct {
	ID       int
	Taskname string
}

func producer(tasks chan<- Task) {
	i := 0 // initialize i to 0
	for {
		if i > 100 {
			break // exit the loop if i exceeds 100
		}

		task := Task{
			ID:       1,
			Taskname: fmt.Sprintf("Task %d", i),
		}
		tasks <- task
		i++ // increment i
	}
}

func consumer(tasks <-chan Task) {
	for task := range tasks {
		println("processing task: ", task.Taskname)
	}
}

func main() {
	tasks := make(chan Task)
	wg := sync.WaitGroup{}

	wg.Add(1) // add to WaitGroup
	go func() {
		defer wg.Done() // decrement the WaitGroup counter when the goroutine exits
		producer(tasks)
		close(tasks) // close the tasks channel after producing all tasks
	}()

	wg.Add(1) // add to WaitGroup
	go func() {
		defer wg.Done() // decrement the WaitGroup counter when the goroutine exits
		consumer(tasks)
	}()

	wg.Wait() // wait for all goroutines to finish
	println("all tasks processed.")
}
