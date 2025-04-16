## In this implementation:

- A fixed number of workers are created.
- A shared channel (TaskChan) is used to distribute tasks to workers.
- Each worker listens to the shared channel for tasks and processes them.
- The program waits for all workers to finish their tasks using a sync.WaitGroup.
- The program waits for all tasks to complete before exiting.

## How It Works
### Worker Pool Creation:
 - A WorkerPool is initialized with a specific number of workers.
 - Workers are assigned a channel to receive tasks and start processing.

### Task Dispatching:
 - Tasks are added to a shared task channel (TaskChan).
 - Workers consume tasks from the channel concurrently.

### Graceful Shutdown:
 - The Stop() method closes the shared task channel and waits for all workers to finish processing.



## Functions:

```Go
NewWorkerPool(numWorkers int): Initializes the pool with the specified number of workers.

Start(numWorkers int): Starts the worker pool and assigns tasks to workers.

AddTask(task Task): Adds a new task to the task queue for workers.

Stop(): Stops the worker pool, closes the channels, and waits for all workers to finish.
```