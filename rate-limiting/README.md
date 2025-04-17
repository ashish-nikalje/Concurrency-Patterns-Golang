## General Overview

- This example demonstrates the **Rate Limiting** concurrency pattern in Go.
- It processes a list of tasks at a controlled rate — one task per second.
- The rate limit is enforced using `time.Ticker`.
- Each task is handled concurrently in a separate goroutine.

---

## How It Works

- A `Ticker` is created with a 1-second interval.
- Before launching each goroutine, the program waits for the next tick from the ticker.
- This ensures that tasks are **throttled** and start at a controlled rate.
- `processTask()` simulates task work and prints output with timestamps.
- A `time.Sleep()` at the end of `main()` ensures all goroutines finish before the program exits.

---

## Code Overview

```go
func main()
```

- Initializes a list of tasks.
- Creates a time.Ticker to limit execution to 1 task per second.

For each task:

    - Waits for the ticker.
    - Launches the task as a goroutine.
    - Sleeps at the end to allow all tasks to complete.

```Go
func processTask(task Task)
```

- Simulates processing with a time.Sleep().
- Prints the task name along with the timestamp to show when it was processed.

## Sample output
```
Processing: Task 1 at 19:43:48
Processing: Task 2 at 19:43:49
Processing: Task 3 at 19:43:50
Processing: Task 4 at 19:43:51
Processing: Task 5 at 19:43:52
Processing: Task 6 at 19:43:53
```