
## **General Overview**

- This example demonstrates the **Ticker Pattern** in Go.
- It uses a `time.Ticker` to execute a periodic task every 2 seconds.
- The program runs the task 5 times, prints the task number and the current time, and then exits.
- The ticker ensures that the task runs at regular intervals.

---

## **How It Works**

- A `time.Ticker` is created with a 2-second interval.
- The program waits for a tick from the ticker and executes a task each time it ticks.
- The task prints the task number and the current time when it is executed.
- The program stops after 5 executions of the periodic task.
- `defer ticker.Stop()` is used to stop the ticker when the program exits, ensuring resources are cleaned up.

---

## **Code Overview**

```go
func main()
```

- Initializes a `time.Ticker` with a 2-second interval (`time.NewTicker(2 * time.Second)`).
- Uses a `select` statement to listen for ticks from the ticker and executes the periodic task each time a tick occurs.
- The task prints a message indicating the current task number and the timestamp.
- After 5 iterations, the program stops the ticker and exits.

---

## **Sample Output**

```
Performing periodic task 1 14:30:30
Performing periodic task 2 14:30:32
Performing periodic task 3 14:30:34
Performing periodic task 4 14:30:36
Performing periodic task 5 14:30:38
Program complete.
```

---

## **Key Concepts**

- **Ticker**: A `time.Ticker` generates ticks at fixed intervals, which can be used to trigger periodic tasks.
- **Periodicity**: The periodic task is executed at regular intervals (every 2 seconds in this case).
- **Graceful Cleanup**: The `defer ticker.Stop()` ensures that the ticker is stopped when no longer needed to prevent resource leaks.

---