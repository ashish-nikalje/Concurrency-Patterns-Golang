## General Overview

- This example demonstrates the **Closing Channels** pattern and the **comma-ok idiom** for handling channel closure in Go.
- A **producer** sends tasks (integers) to a shared channel and closes it once all tasks are sent.
- A **consumer** receives tasks from the channel and checks if the channel is closed using the **comma-ok idiom**.
- The consumer stops processing once it detects the channel is closed.

---

## How It Works

- The producer goroutine sends values (tasks) into the `tasks` channel.
- After sending all tasks, the producer **closes** the channel to signal no more data will be sent.
- The consumer goroutine continuously reads from the channel until it detects the channel is closed.
- The **comma-ok idiom** is used to check if the channel is closed:
  - If the channel is open, it processes the task.
  - If the channel is closed, it exits the loop and gracefully terminates the consumer.

---

## Code Overview

### Producer Function

```go
go func() {
	for i := 1; i <= 5; i++ {
		tasks <- i
	}
	close(tasks) // Close the channel to signal completion
}()
```

- The producer goroutine sends 5 tasks (integers 1-5) into the `tasks` channel.
- After all tasks are sent, the channel is closed using `close(tasks)`.

### Consumer Function
```Go
for {
	val, ok := <-tasks
	if !ok { // Check if the channel is closed
		fmt.Println("Channel closed, no more tasks.")
		break
	}
	fmt.Println("Processing task:", val)
}
```

- The consumer uses the for loop with the comma-ok idiom to receive tasks from the channel.
- It checks the ok value:
	- If `ok` is `true`, it processes the received task.
	- If `ok` is `false`, it indicates the channel is closed, and the consumer exits the loop.