## General Overview

- A single task is processed concurrently in a separate goroutine.
- The main routine waits for either the task to complete or the timeout to trigger.
- The `context.WithTimeout` ensures the task doesn't run longer than a specified duration.
- This pattern is useful to avoid resource leaks and keep systems responsive.

---

## How It Works

- A `Task` struct holds the task metadata.
- The `PrintTaskDetails` function simulates processing (with `time.Sleep`) and returns results via a channel.
- The function runs in a separate goroutine to simulate asynchronous behavior.
- A `select` block in `main` waits for either:
  - The task result from the channel, or
  - A timeout/cancellation signal from the context.
- If the task completes before the timeout, results are printed.
- If the timeout occurs first, the program gracefully handles it and exits.

---

## Code Overview

```go
func PrintTaskDetails(ctx context.Context, task Task, done chan<- string)
```

## Sample Output
If the task takes longer than the timeout:

 `Timeout or cancelled: context deadline exceeded`

If the task finishes before the timeout:
```Go
Task completed:
Task ID: 1
Task Name: Example Task
```