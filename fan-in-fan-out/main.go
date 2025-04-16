package main

import (
	"fmt"
	"sync"
	"time"
)

type Task struct {
	ID       int
	Taskname string
}

func worker(id int, tasks <-chan Task, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done() // descrement the WaitGroup counter when the worker exits
	for task := range tasks {
		time.Sleep(50 * time.Millisecond) // simulate processing
		results <- fmt.Sprintf("Worker %d processed task %d: %s", id, task.ID, task.Taskname)
	}

	fmt.Printf("Worker %d exiting.\n", id) // print when worker exits
}

func fanOutFanIn(tasks []Task, numWorkers int) {
	tasksChan := make(chan Task, len(tasks))     // channel to send tasks to workers
	resultsChan := make(chan string, len(tasks)) // channel to receive results from workers
	var wg sync.WaitGroup

	// Fan-out: start workers
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)                                 // add to WaitGroup
		go worker(i, tasksChan, resultsChan, &wg) // start worker goroutine
	}

	// Send tasks to workers
	for _, task := range tasks {
		tasksChan <- task
	}
	close(tasksChan) // close the tasks channel after sending all tasks

	// Fan-in: close results channel when all workers are done
	go func() {
		wg.Wait()
		close(resultsChan)
	}()

	// Receive results
	for result := range resultsChan {
		fmt.Println(result)
	}
}

func main() {
	numWorkers := 3 // Number of workers you want to use

	tasks := []Task{
		{ID: 1, Taskname: "Task 1"},
		{ID: 2, Taskname: "Task 2"},
		{ID: 3, Taskname: "Task 3"},
		{ID: 4, Taskname: "Task 4"},
		{ID: 5, Taskname: "Task 5"},
		{ID: 6, Taskname: "Task 6"},
		{ID: 7, Taskname: "Task 7"},
		{ID: 8, Taskname: "Task 8"},
		{ID: 9, Taskname: "Task 9"},
		{ID: 10, Taskname: "Task 10"},
	}

	fanOutFanIn(tasks, numWorkers)
}
