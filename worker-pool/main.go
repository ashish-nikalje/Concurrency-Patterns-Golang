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

type Worker struct {
	ID       int       // Worker ID
	TaskChan chan Task // Channel to receive tasks
}

type WorkerPool struct {
	Workers  []Worker  // Slice of workers
	TaskChan chan Task // Channel to receive tasks
	Wg       sync.WaitGroup
}

func NewWorkerPool(numWorkers int) *WorkerPool { // NewWorkerPool creates a new worker pool with the specified number of workers
	pool := &WorkerPool{
		TaskChan: make(chan Task),
		Wg:       sync.WaitGroup{},
	}

	for i := 0; i < numWorkers; i++ { // Initialize each worker in the pool
		pool.Workers[i] = Worker{
			ID:       i,
			TaskChan: make(chan Task),
		}
	}

	return pool
}

func (wp *WorkerPool) Start(numWorkers int) {
	for i := 0; i < numWorkers; i++ {
		worker := Worker{ID: i} // Create a new worker
		wp.Wg.Add(1)            // Add to WaitGroup

		go func(w Worker) {
			defer wp.Wg.Done()              // Decrement the WaitGroup counter when the worker exits
			for task := range wp.TaskChan { // Listen for tasks on the TaskChan
				fmt.Printf("Worker %d processing task %d: %s\n", w.ID, task.ID, task.Taskname) // Process the tas
				time.Sleep(300 * time.Millisecond)                                             // added to simulate work
			}
			fmt.Printf("Worker %d exiting.\n", w.ID) // Print when worker exits
		}(worker)
	}
}

func (wp *WorkerPool) AddTask(task Task) {
	wp.TaskChan <- task
}

func (wp *WorkerPool) Stop() { // Stop stops the worker pool and waits for all workers to finish
	close(wp.TaskChan)
	for _, worker := range wp.Workers {
		close(worker.TaskChan)
	}

	wp.Wg.Wait()
}

func main() {
	workerCount := 3
	fmt.Printf("Starting worker pool with %d workers...\n", workerCount)
	workerPool := NewWorkerPool(workerCount) // Create a new worker pool with the specified number of workers
	workerPool.Start(workerCount)            // Start the worker pool

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

	for _, task := range tasks {
		workerPool.AddTask(task)
	}

	workerPool.Stop()
}
