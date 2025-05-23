# Concurrency Patterns

This is a personal collection of Go concurrency patterns I’ve implemented to better understand how goroutines, channels, and context work in real-world systems. Each pattern is broken down into simple, focused examples that are easy to run and experiment with.

---

## Patterns Included


1. **[Worker Pool](worker-pool/README.md)** – Multiple workers process tasks from a shared channel.  
2. **[Fan-Out, Fan-In](fan-in-fan-out/README.md)** – Distributes tasks across multiple goroutines and aggregates results.  
3. **[Producer-Consumer](producer-consumer/README.md)** – One set of goroutines produce data, another consumes it.  
4. **[Timeout](timeout/README.md)** – Prevents goroutines from waiting indefinitely using `time.After` or `context.Context`.  
5. **[Rate Limiting](rate-limiting/README.md)** – Controls the frequency of execution using `time.Ticker` or token bucket algorithms.  
6. **[Closing Channels](closing-channel/README.md)** – Signals completion by closing channels to prevent goroutine leaks.  
7. **[Select Statement](select-statement/README.md)** – Listens to multiple channels and handles inputs dynamically.  
8. **[Pipeline Pattern](pipeline/README.md)** – Chains multiple stages of processing using channels for efficiency.  
9. **[Ticker Pattern](ticker/README.md)** – Uses `time.Ticker` to execute periodic tasks.  
10. **[Cancellation with Context](cancellation-signal/README.md)** – Uses `context.Context` to propagate cancellation signals across goroutines.  
11. **[Bounded Work Queue](bounding-queue/readme.md)** – Limits the number of jobs in a queue to prevent system overload.               

---

## Structure

```bash
code@code Concurrency-Patterns-Golang % tree
.
├── cancellation-signal
│   ├── main.go
│   └── README.md
├── closing-channel
│   ├── main.go
│   └── README.md
├── fan-in-fan-out
│   ├── main.go
│   └── README.md
├── pipeline
│   ├── main.go
│   └── README.md
├── producer-consumer
│   ├── main.go
│   └── README.md
├── rate-limiting
│   ├── main.go
│   └── README.md
├── README.md
├── select-statement
│   ├── main.go
│   └── README.md
├── ticker
│   ├── main.go
│   └── README.md
├── timeout
│   ├── main.go
│   └── README.md
└── worker-pool
    ├── main.go
    └── README.md
```
