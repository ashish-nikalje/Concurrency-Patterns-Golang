## General Overview

- This example demonstrates the **Select Statement** concurrency pattern in Go.
- It listens to multiple channels and handles inputs dynamically based on which channel becomes ready first.
- The select statement allows for non-blocking operations, making it ideal for situations where you need to listen to multiple inputs concurrently.

---

## How It Works

- Two channels are created, `ch1` and `ch2`, each representing an asynchronous task.
- Two goroutines are started to send data to these channels at different times (with `time.Sleep` to simulate delayed data).
- A `select` statement is used to listen for messages on both channels.
- When a message is ready from either channel, it is processed, and the program continues.
- The `select` statement blocks until one of the channels sends data, allowing the program to handle whichever task is ready first.

---

## Code Overview

```go
func main()
```
- Creates two channels (`ch1` and `ch2`).
- Starts two goroutines that send messages to the channels after different delays.
- The select statement listens for messages on both channels and processes the first one that is ready.

## Sample Output
```
Received: Message from ch2
Received: Message from ch1
```