## General Overview
- The producer creates 101 tasks (Task 0 to Task 100) and sends them to a shared channel.
- The consumer receives tasks from that channel and processes them.
- Both routines run concurrently.
- Once all tasks are processed, the program exits cleanly.

## How It Works

- The producer loop generates tasks and sends them to the tasks channel.
- The consumer reads from the channel and prints each task as it processes it.
- The sync.WaitGroup ensures the program waits until both producer and consumer complete.
- The channel is closed once the producer finishes sending all tasks.


## Code Overview
  ```GO
  func producer(tasks chan<- Task)
  ```
- Producer function that generates tasks and sends them to the channel.
- It creates 101 tasks (from 0 to 100).
- Uses a loop with fmt.Sprintf() to generate task names like "Task 0", "Task 1", etc.
- Sends tasks to the tasks channel.
- Breaks the loop after task 100.

```Go
func consumer(tasks <-chan Task)
```
- Consumer function that listens on the tasks channel.
- Processes tasks as they come in (prints task name).
- The loop automatically exits when the channel is closed.

