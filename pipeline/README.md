## General Overview

- This example demonstrates the **Pipeline Pattern** concurrency pattern in Go.
- It processes a list of tasks through multiple stages: converting strings to uppercase, adding a suffix, and printing the result.
- Data flows through each stage, and each stage runs concurrently in a separate goroutine.

---

## How It Works

- **Stage 1**: The input data (a list of strings) is read and passed into a channel (`input`).
- **Stage 2**: The first stage transforms the strings to uppercase and sends the result to the `upperCaseChan`.
- **Stage 3**: The second stage appends a suffix (`_SUFFIX`) to each string and sends the result to `suffixChan`.
- **Final Stage**: The final stage prints the processed data to the console.

---

## Code Overview

```go
func main()
```

- Initializes the input data and creates channels for each stage.
- Starts a goroutine for each stage in the pipeline, processing data concurrently.
- Uses time.Sleep() to ensure all goroutines complete before the program exits.



```Go
func toUpperCase(input <-chan string, output chan<- string)
```

- Reads strings from the input channel, converts them to uppercase, and sends the result to the output channel.


```GO
func addSuffix(input <-chan string, output chan<- string)

```

- Reads strings from the input channel, appends a suffix (_SUFFIX), and sends the result to the output channel.
