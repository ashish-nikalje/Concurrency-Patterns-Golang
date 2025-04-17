## General Overview

- This example demonstrates the **Cancellation with `context.Context`** concurrency pattern in Go.
- A `context.WithTimeout()` is used to automatically cancel task execution after a specified duration.
- Each task is executed in its own goroutine and listens for a cancellation signal using `<-ctx.Done()`.

---

## How It Works

- A single `Task` is defined and launched with a goroutine.
- The main function creates a `context.WithTimeout()` with a 2-second timeout.
- Each task goroutine checks the context’s cancellation signal to determine when to stop processing.
- `time.Sleep()` is used to simulate work and spacing between task executions.
- After 2 seconds, the context is canceled, and the goroutine gracefully exits.

---

## Code Overview

### `main()` Function

```go
ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
defer cancel()
```

- Creates a context that will timeout automatically after 2 seconds.
- The `cancel()` call ensures context resources are released when done.

```GO
for _, task := range tasks {
    time.Sleep(500 * time.Millisecond)
    go processTask(ctx, task)
}
```

- Iterates over tasks and starts each in a separate goroutine.
- Adds a 500ms sleep to simulate task staggering.

```GO
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
```

- Each task prints its processing status every 500ms.
- It listens to the context’s `Done()` channel to know when to cancel.
- Once cancelled, it prints a cancellation message and exits.