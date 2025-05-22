package main

import (
    "context"
    "fmt"
    "time"
)

// Worker function to demonstrate goroutine communication
func worker(id int, jobs <-chan int, results chan<- int) {
    for job := range jobs {
        fmt.Printf("Worker %d started job %d\n", id, job)
        time.Sleep(time.Second) // Simulate work
        fmt.Printf("Worker %d finished job %d\n", id, job)
        results <- job * 2 // Send result back
    }
}

func main() {
    // Buffered channel example
    bufferedChannel := make(chan int, 3)
    bufferedChannel <- 1
    bufferedChannel <- 2
    bufferedChannel <- 3
    close(bufferedChannel)

    fmt.Println("Buffered channel example:")
    for val := range bufferedChannel {
        fmt.Println(val)
    }

    // Unbuffered channel example
    unbufferedChannel := make(chan string)
    go func() {
        unbufferedChannel <- "Hello from goroutine"
    }()
    fmt.Println("\nUnbuffered channel example:")
    fmt.Println(<-unbufferedChannel)

    // Fan-out/Fan-in example
    jobs := make(chan int, 5)
    results := make(chan int, 5)

    for w := 1; w <= 3; w++ {
        go worker(w, jobs, results)
    }

    for j := 1; j <= 5; j++ {
        jobs <- j
    }
    close(jobs)

    fmt.Println("\nFan-out/Fan-in example:")
    for a := 1; a <= 5; a++ {
        fmt.Println("Result:", <-results)
    }

    // Select statement example
    ch1 := make(chan int)
    ch2 := make(chan int)

    go func() {
        time.Sleep(2 * time.Second)
        ch1 <- 1
    }()
    go func() {
        time.Sleep(1 * time.Second)
        ch2 <- 2
    }()

    fmt.Println("\nSelect statement example:")
    select {
    case msg1 := <-ch1:
        fmt.Println("Received from ch1:", msg1)
    case msg2 := <-ch2:
        fmt.Println("Received from ch2:", msg2)
    case <-time.After(3 * time.Second):
        fmt.Println("Timeout")
    }

    // Context with channels example
    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel()

    ch := make(chan int)
    go func() {
        for i := 0; i < 5; i++ {
            select {
            case <-ctx.Done():
                fmt.Println("\nContext canceled, stopping goroutine")
                return
            case ch <- i:
                time.Sleep(500 * time.Millisecond)
            }
        }
        close(ch)
    }()

    fmt.Println("\nContext with channels example:")
    for val := range ch {
        fmt.Println("Received:", val)
    }
}