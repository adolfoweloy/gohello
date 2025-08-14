// Package concurrency demonstrates Go's concurrency features
package concurrency

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// DemonstrateGoroutines shows basic goroutine usage
func DemonstrateGoroutines() {
	fmt.Println("=== Goroutines ===")
	
	// Sequential execution
	fmt.Println("Sequential execution:")
	start := time.Now()
	task("A", 1)
	task("B", 1)
	task("C", 1)
	fmt.Printf("Sequential time: %v\n", time.Since(start))
	
	// Concurrent execution
	fmt.Println("\nConcurrent execution:")
	start = time.Now()
	go task("X", 1)
	go task("Y", 1)
	go task("Z", 1)
	
	// Wait a bit for goroutines to complete
	time.Sleep(2 * time.Second)
	fmt.Printf("Concurrent time: %v\n", time.Since(start))
	
	// Demonstrate main goroutine ending
	fmt.Println("\nGoroutines continuing after function returns:")
	go longRunningTask("Background-1", 3)
	go longRunningTask("Background-2", 3)
	
	// These will be cut off when the function returns
	time.Sleep(1 * time.Second)
	fmt.Println("Function ending...")
}

// task simulates work
func task(name string, seconds int) {
	for i := 1; i <= seconds; i++ {
		fmt.Printf("  Task %s: step %d\n", name, i)
		time.Sleep(1 * time.Second)
	}
	fmt.Printf("  Task %s: completed\n", name)
}

// longRunningTask simulates a long-running task
func longRunningTask(name string, seconds int) {
	for i := 1; i <= seconds; i++ {
		fmt.Printf("  Long task %s: step %d/%d\n", name, i, seconds)
		time.Sleep(1 * time.Second)
	}
	fmt.Printf("  Long task %s: completed\n", name)
}

// DemonstrateChannels shows channel communication
func DemonstrateChannels() {
	fmt.Println("\n=== Channels ===")
	
	// Basic channel usage
	fmt.Println("Basic channel communication:")
	ch := make(chan string)
	
	// Send data in a goroutine
	go func() {
		time.Sleep(1 * time.Second)
		ch <- "Hello from goroutine!"
	}()
	
	// Receive data
	message := <-ch
	fmt.Printf("Received: %s\n", message)
	
	// Buffered channel
	fmt.Println("\nBuffered channel:")
	buffered := make(chan int, 3)
	
	// Send values without blocking
	buffered <- 1
	buffered <- 2
	buffered <- 3
	
	// Receive values
	fmt.Printf("Received: %d\n", <-buffered)
	fmt.Printf("Received: %d\n", <-buffered)
	fmt.Printf("Received: %d\n", <-buffered)
	
	// Channel as parameter
	fmt.Println("\nChannel as parameter:")
	numbers := make(chan int, 5)
	go producer(numbers)
	consumer(numbers)
	
	// Closing channels
	fmt.Println("\nClosing channels:")
	demonstrateChannelClosing()
}

// producer sends numbers to a channel
func producer(ch chan<- int) { // Send-only channel
	for i := 1; i <= 5; i++ {
		fmt.Printf("  Producing: %d\n", i)
		ch <- i
		time.Sleep(500 * time.Millisecond)
	}
	close(ch)
}

// consumer receives numbers from a channel
func consumer(ch <-chan int) { // Receive-only channel
	for num := range ch {
		fmt.Printf("  Consuming: %d\n", num)
		time.Sleep(200 * time.Millisecond)
	}
	fmt.Println("  Consumer finished")
}

// demonstrateChannelClosing shows proper channel closing
func demonstrateChannelClosing() {
	ch := make(chan int, 3)
	
	// Send some values
	go func() {
		for i := 1; i <= 3; i++ {
			ch <- i
			time.Sleep(100 * time.Millisecond)
		}
		close(ch) // Close the channel when done
	}()
	
	// Receive until channel is closed
	for {
		value, ok := <-ch
		if !ok {
			fmt.Println("  Channel closed")
			break
		}
		fmt.Printf("  Received: %d\n", value)
	}
	
	// Alternative: using range (automatically stops when channel is closed)
	fmt.Println("  Using range:")
	ch2 := make(chan string, 3)
	go func() {
		messages := []string{"Hello", "World", "Go"}
		for _, msg := range messages {
			ch2 <- msg
		}
		close(ch2)
	}()
	
	for msg := range ch2 {
		fmt.Printf("  Received: %s\n", msg)
	}
}

// DemonstrateSelect shows select statement usage
func DemonstrateSelect() {
	fmt.Println("\n=== Select Statement ===")
	
	// Basic select
	fmt.Println("Basic select:")
	ch1 := make(chan string)
	ch2 := make(chan string)
	
	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "Message from channel 1"
	}()
	
	go func() {
		time.Sleep(500 * time.Millisecond)
		ch2 <- "Message from channel 2"
	}()
	
	// Select will choose the first available channel
	select {
	case msg1 := <-ch1:
		fmt.Printf("  Received from ch1: %s\n", msg1)
	case msg2 := <-ch2:
		fmt.Printf("  Received from ch2: %s\n", msg2)
	}
	
	// Select with timeout
	fmt.Println("\nSelect with timeout:")
	demonstrateSelectTimeout()
	
	// Select with default case
	fmt.Println("\nSelect with default case:")
	demonstrateSelectDefault()
	
	// Fan-in pattern
	fmt.Println("\nFan-in pattern:")
	demonstrateFanIn()
}

// demonstrateSelectTimeout shows timeout handling with select
func demonstrateSelectTimeout() {
	ch := make(chan string)
	
	// This goroutine will take too long
	go func() {
		time.Sleep(2 * time.Second)
		ch <- "Delayed message"
	}()
	
	select {
	case msg := <-ch:
		fmt.Printf("  Received: %s\n", msg)
	case <-time.After(1 * time.Second):
		fmt.Println("  Timeout! No message received within 1 second")
	}
}

// demonstrateSelectDefault shows non-blocking operations with default
func demonstrateSelectDefault() {
	ch := make(chan string)
	
	// Try to receive without blocking
	select {
	case msg := <-ch:
		fmt.Printf("  Received: %s\n", msg)
	default:
		fmt.Println("  No message available, continuing...")
	}
	
	// Send a message and try again
	go func() {
		ch <- "Available message"
	}()
	
	time.Sleep(100 * time.Millisecond) // Give goroutine time to send
	
	select {
	case msg := <-ch:
		fmt.Printf("  Received: %s\n", msg)
	default:
		fmt.Println("  Still no message available")
	}
}

// demonstrateFanIn shows fan-in pattern with select
func demonstrateFanIn() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	
	// Two producers
	go func() {
		for i := 1; i <= 3; i++ {
			ch1 <- fmt.Sprintf("Producer 1: message %d", i)
			time.Sleep(500 * time.Millisecond)
		}
		close(ch1)
	}()
	
	go func() {
		for i := 1; i <= 3; i++ {
			ch2 <- fmt.Sprintf("Producer 2: message %d", i)
			time.Sleep(700 * time.Millisecond)
		}
		close(ch2)
	}()
	
	// Fan-in: receive from both channels
	for {
		select {
		case msg1, ok := <-ch1:
			if !ok {
				ch1 = nil // Disable this case
			} else {
				fmt.Printf("  From ch1: %s\n", msg1)
			}
		case msg2, ok := <-ch2:
			if !ok {
				ch2 = nil // Disable this case
			} else {
				fmt.Printf("  From ch2: %s\n", msg2)
			}
		}
		
		// Exit when both channels are closed
		if ch1 == nil && ch2 == nil {
			break
		}
	}
}

// DemonstrateWaitGroups shows WaitGroup usage for synchronization
func DemonstrateWaitGroups() {
	fmt.Println("\n=== WaitGroups ===")
	
	var wg sync.WaitGroup
	
	workers := 3
	jobs := 5
	
	fmt.Printf("Starting %d workers to process %d jobs:\n", workers, jobs)
	
	// Start workers
	for w := 1; w <= workers; w++ {
		wg.Add(1) // Increment the WaitGroup counter
		go worker(w, jobs, &wg)
	}
	
	// Wait for all workers to complete
	wg.Wait()
	fmt.Println("All workers completed!")
	
	// Another example: parallel processing
	fmt.Println("\nParallel processing example:")
	demonstrateParallelProcessing()
}

// worker simulates work being done
func worker(id int, jobs int, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the counter when the goroutine completes
	
	for j := 1; j <= jobs; j++ {
		fmt.Printf("  Worker %d: processing job %d\n", id, j)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	}
	fmt.Printf("  Worker %d: finished\n", id)
}

// demonstrateParallelProcessing shows parallel processing with WaitGroup
func demonstrateParallelProcessing() {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	results := make([]int, len(data))
	var wg sync.WaitGroup
	
	// Process each item in parallel
	for i, value := range data {
		wg.Add(1)
		go func(index, val int) {
			defer wg.Done()
			// Simulate some processing (square the number)
			time.Sleep(100 * time.Millisecond)
			results[index] = val * val
			fmt.Printf("  Processed %d -> %d\n", val, results[index])
		}(i, value)
	}
	
	wg.Wait()
	fmt.Printf("Results: %v\n", results)
}

// WorkerPool demonstrates a worker pool pattern
type WorkerPool struct {
	workers int
	jobs    chan Job
	results chan Result
	wg      sync.WaitGroup
}

// Job represents work to be done
type Job struct {
	ID   int
	Data int
}

// Result represents the result of work
type Result struct {
	Job   Job
	Value int
}

// NewWorkerPool creates a new worker pool
func NewWorkerPool(workers int) *WorkerPool {
	return &WorkerPool{
		workers: workers,
		jobs:    make(chan Job, 100),
		results: make(chan Result, 100),
	}
}

// Start starts the worker pool
func (wp *WorkerPool) Start() {
	for w := 1; w <= wp.workers; w++ {
		wp.wg.Add(1)
		go wp.worker(w)
	}
}

// worker processes jobs from the jobs channel
func (wp *WorkerPool) worker(id int) {
	defer wp.wg.Done()
	for job := range wp.jobs {
		fmt.Printf("  Worker %d: processing job %d (data: %d)\n", id, job.ID, job.Data)
		
		// Simulate work
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		result := job.Data * job.Data
		
		wp.results <- Result{Job: job, Value: result}
	}
	fmt.Printf("  Worker %d: finished\n", id)
}

// AddJob adds a job to the worker pool
func (wp *WorkerPool) AddJob(job Job) {
	wp.jobs <- job
}

// Close closes the jobs channel and waits for workers to finish
func (wp *WorkerPool) Close() {
	close(wp.jobs)
	wp.wg.Wait()
	close(wp.results)
}

// DemonstrateWorkerPool shows worker pool pattern
func DemonstrateWorkerPool() {
	fmt.Println("\n=== Worker Pool ===")
	
	pool := NewWorkerPool(3)
	pool.Start()
	
	// Add jobs
	for i := 1; i <= 10; i++ {
		pool.AddJob(Job{ID: i, Data: i})
	}
	
	// Collect results in a separate goroutine
	go func() {
		for result := range pool.results {
			fmt.Printf("  Result: job %d -> %d\n", result.Job.ID, result.Value)
		}
	}()
	
	// Close the pool
	pool.Close()
	
	// Give time for results to be printed
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Worker pool completed!")
}