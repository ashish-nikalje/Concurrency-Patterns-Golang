## How It Works

### Producer Goroutine:

 - Generates 5 tasks and sends them to the task queue.
 - The producer waits for 500ms between producing each task.
 - After all tasks are produced, the queue is closed.

### Consumer Goroutine:

 - Consumes tasks from the task queue and starts a new goroutine to process each task concurrently.
 - Each task is processed in the processTask() function, where it sleeps for 1 second to simulate processing time.
 
### Synchronization:
 - A sync.WaitGroup is used to ensure that the main program waits until all tasks are processed before exiting.


```Go
taskQueue := make(chan Task, 3) // Bounded queue with max capacity of 3
var wg sync.WaitGroup

// Producer: sending tasks
go func() {
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

wg.Wait() // Wait for all tasks to complete

```

 - Task Queue: A buffered channel (taskQueue) is used to hold tasks. The buffer size is set to 3, meaning the producer will be blocked if there are more than 3 tasks in the queue.
 - Producer Goroutine: A goroutine is launched to produce 5 tasks and send them to the taskQueue. Each task is produced with a delay of 500ms between each.
 - Consumer Goroutine: A separate goroutine consumes tasks from the queue. For each task, a new goroutine is spawned to process the task.

### Sample output

```
Starting bounded queue example...
Produced: Task 1
Processing: Task 1
Produced: Task 2
Processing: Task 2
Completed: Task 1
Produced: Task 3
Processing: Task 3
Completed: Task 2
Produced: Task 4
Processing: Task 4
Produced: Task 5
Processing: Task 5
Completed: Task 3
Completed: Task 4
Completed: Task 5
```