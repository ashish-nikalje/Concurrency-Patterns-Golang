**In this example:**

- **Fan-Out**: Multiple worker goroutines (workers) are started, each of which processes a subset of tasks concurrently.
- **Fan-In**: After the workers process the tasks, their results are collected and processed in a single channel.

## Features
- The worker pool dynamically scales based on the number of workers.
- The tasks are evenly distributed among the workers.
- The program collects and prints results once all workers complete their tasks.


## Build and run the program:
   ```bash
   go run main.go
   ```

## Code Explanation

- **Task**: Represents the unit of work that needs to be processed. Each task has an ID and a task name.
- **Worker**: A goroutine that processes tasks concurrently from a shared task channel.
- **WaitGroup**: Used to synchronize the completion of all worker goroutines.
- **Channels**: 
  - `tasksChan` is used to send tasks to the workers.
  - `resultsChan` is used to collect the results from all workers after they finish processing.

### Core Functions

#### `worker(id int, tasks <-chan Task, results chan<- string, wg *sync.WaitGroup)`
Each worker processes tasks from the `tasks` channel and sends results to the `results` channel. The worker is responsible for notifying the `WaitGroup` when it's done processing.

#### `fanOutFanIn(tasks []Task, numWorkers int)`
- Distributes the tasks to the workers via the `tasksChan`.
- Collects results from workers via the `resultsChan`.
- Waits for all workers to finish processing using `sync.WaitGroup`.

### Example Output:

```bash
Worker 2 processed task 2: Task 2
Worker 3 processed task 3: Task 3
Worker 1 processed task 1: Task 1
Worker 3 processed task 4: Task 4
Worker 2 processed task 5: Task 5
Worker 1 processed task 6: Task 6
Worker 1 processed task 9: Task 9
Worker 3 processed task 7: Task 7
Worker 3 exiting.
Worker 2 exiting.
Worker 2 processed task 8: Task 8
Worker 1 exiting.
Worker 1 processed task 10: Task 10
```