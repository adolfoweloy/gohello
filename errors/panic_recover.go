// Package errors demonstrates Go's panic and recover mechanisms including
// panic handling, graceful recovery, and best practices.
package errors

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

// DemoPanicBasics demonstrates basic panic and recover usage
func DemoPanicBasics() {
	fmt.Println("\n=== Panic and Recover Basics Demo ===")

	// Example 1: Function that might panic
	fmt.Println("1. Basic panic handling:")
	result := SafeFunction(10, 2)
	fmt.Printf("   Result: %d\n", result)

	result = SafeFunction(10, 0) // This will panic internally
	fmt.Printf("   Result: %d\n", result)

	// Example 2: Recover in defer
	fmt.Println("\n2. Defer with recover:")
	DemoRecover()
	fmt.Println("   Program continues after panic")

	// Example 3: Multiple panic scenarios
	fmt.Println("\n3. Multiple panic scenarios:")
	testCases := [][]int{
		{10, 2},   // Normal case
		{10, 0},   // Division by zero
		{-1, 5},   // Invalid input
	}

	for i, tc := range testCases {
		result := ProcessNumbers(tc[0], tc[1])
		fmt.Printf("   Test %d: %v -> %s\n", i+1, tc, result)
	}
}

// SafeFunction demonstrates panic recovery
func SafeFunction(a, b int) int {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("   Recovered from panic: %v\n", r)
		}
	}()

	if b == 0 {
		panic("division by zero")
	}
	return a / b
}

// DemoRecover demonstrates basic recover usage
func DemoRecover() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("   Recovered: %v\n", r)
		}
	}()

	fmt.Printf("   Before panic\n")
	panic("something went wrong")
	fmt.Printf("   This won't be printed\n") // Unreachable
}

// ProcessNumbers safely processes numbers with detailed error information
func ProcessNumbers(a, b int) string {
	defer func() {
		if r := recover(); r != nil {
			// In real code, you might log this or handle it differently
		}
	}()

	// Validate inputs
	if a < 0 {
		panic(fmt.Sprintf("negative input not allowed: %d", a))
	}

	if b == 0 {
		panic("cannot divide by zero")
	}

	result := a / b
	return fmt.Sprintf("success: %d", result)
}

// DemoGracefulRecovery demonstrates graceful error recovery
func DemoGracefulRecovery() {
	fmt.Println("\n=== Graceful Recovery Demo ===")

	// Process a batch of operations, recovering from individual failures
	operations := []func() string{
		func() string { return "Operation 1: " + safeOperation(10, 2) },
		func() string { return "Operation 2: " + safeOperation(10, 0) }, // Will panic
		func() string { return "Operation 3: " + safeOperation(15, 3) },
		func() string { return "Operation 4: " + safeOperation(-5, 2) }, // Will panic
		func() string { return "Operation 5: " + safeOperation(20, 4) },
	}

	results := ProcessBatchOperations(operations)
	fmt.Println("Batch processing results:")
	for i, result := range results {
		fmt.Printf("  %d. %s\n", i+1, result)
	}

	// Web server-like error handling
	fmt.Println("\nWeb request simulation:")
	requests := []string{"valid", "invalid", "timeout", "success"}
	for _, req := range requests {
		response := HandleWebRequest(req)
		fmt.Printf("  Request '%s': %s\n", req, response)
	}
}

// safeOperation performs an operation that might panic
func safeOperation(a, b int) string {
	if a < 0 {
		panic(fmt.Sprintf("negative input: %d", a))
	}
	if b == 0 {
		panic("division by zero")
	}
	return fmt.Sprintf("result = %d", a/b)
}

// ProcessBatchOperations processes operations with individual recovery
func ProcessBatchOperations(operations []func() string) []string {
	results := make([]string, len(operations))

	for i, op := range operations {
		func(index int) {
			defer func() {
				if r := recover(); r != nil {
					results[index] = fmt.Sprintf("ERROR: %v", r)
				}
			}()
			results[index] = op()
		}(i)
	}

	return results
}

// HandleWebRequest simulates web request handling with panic recovery
func HandleWebRequest(request string) (response string) {
	defer func() {
		if r := recover(); r != nil {
			// Log error in real application
			fmt.Printf("    [ERROR] Request panic: %v\n", r)
			response = "500 Internal Server Error"
		}
	}()

	switch request {
	case "valid":
		return "200 OK"
	case "invalid":
		panic("400 Bad Request: Invalid format")
	case "timeout":
		panic("504 Gateway Timeout")
	case "success":
		return "200 OK: Data processed"
	default:
		return "404 Not Found"
	}
}

// DemoStackTrace demonstrates capturing stack traces during panic
func DemoStackTrace() {
	fmt.Println("\n=== Stack Trace Demo ===")

	// Demonstrate stack trace capture
	fmt.Println("Stack trace example:")
	result := CallChainWithPanic()
	fmt.Printf("Final result: %s\n", result)

	// Demonstrate manual stack trace
	fmt.Println("\nManual stack trace:")
	ShowCurrentStack()
}

// CallChainWithPanic demonstrates panic propagation through call stack
func CallChainWithPanic() string {
	defer func() {
		if r := recover(); r != nil {
			// Capture stack trace
			stack := make([]byte, 4096)
			length := runtime.Stack(stack, false)
			fmt.Printf("Panic recovered: %v\n", r)
			fmt.Printf("Stack trace:\n%s\n", stack[:length])
		}
	}()

	return levelOne()
}

func levelOne() string {
	return levelTwo()
}

func levelTwo() string {
	return levelThree()
}

func levelThree() string {
	panic("Something went wrong in level three")
}

// ShowCurrentStack shows the current call stack
func ShowCurrentStack() {
	stack := make([]byte, 2048)
	length := runtime.Stack(stack, false)
	fmt.Printf("Current stack:\n%s\n", stack[:length])
}

// DemoPanicInGoroutines demonstrates panic handling in goroutines
func DemoPanicInGoroutines() {
	fmt.Println("\n=== Panic in Goroutines Demo ===")

	// Note: In real applications, you'd use sync.WaitGroup or channels
	// This is simplified for demonstration

	fmt.Println("Starting goroutines with panic handling:")

	// Goroutine that panics
	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("  Goroutine 1 recovered: %v\n", r)
			}
		}()
		panic("goroutine 1 panic")
	}()

	// Goroutine that succeeds
	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("  Goroutine 2 recovered: %v\n", r)
			}
		}()
		fmt.Println("  Goroutine 2 completed successfully")
	}()

	// Goroutine with different panic
	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("  Goroutine 3 recovered: %v\n", r)
			}
		}()
		var slice []int
		_ = slice[100] // Index out of range panic
	}()

	// Give goroutines time to complete (in real code, use proper synchronization)
	// time.Sleep(100 * time.Millisecond) // Commented out to avoid dependency

	fmt.Println("Main goroutine continues")
}

// DemoResourceCleanup demonstrates resource cleanup with panic recovery
func DemoResourceCleanup() {
	fmt.Println("\n=== Resource Cleanup Demo ===")

	// Simulate file operations with panic recovery
	fmt.Println("File operation simulation:")
	ProcessFile("valid_file.txt")
	ProcessFile("corrupted_file.txt")
	ProcessFile("locked_file.txt")

	// Simulate database operations
	fmt.Println("\nDatabase operation simulation:")
	ProcessDatabase("SELECT * FROM users")
	ProcessDatabase("INVALID SQL QUERY")
	ProcessDatabase("SELECT * FROM missing_table")
}

// ProcessFile simulates file processing with cleanup
func ProcessFile(filename string) {
	file := &MockFile{Name: filename}

	defer func() {
		// Always cleanup, even if panic occurs
		file.Close()
		if r := recover(); r != nil {
			fmt.Printf("  File processing failed for '%s': %v\n", filename, r)
		}
	}()

	file.Open()

	// Simulate different error conditions
	switch {
	case strings.Contains(filename, "corrupted"):
		panic("file corrupted: unable to read")
	case strings.Contains(filename, "locked"):
		panic("file locked by another process")
	default:
		fmt.Printf("  File '%s' processed successfully\n", filename)
	}
}

// MockFile simulates a file resource
type MockFile struct {
	Name   string
	isOpen bool
}

func (f *MockFile) Open() {
	f.isOpen = true
	fmt.Printf("    Opened file: %s\n", f.Name)
}

func (f *MockFile) Close() {
	if f.isOpen {
		f.isOpen = false
		fmt.Printf("    Closed file: %s\n", f.Name)
	}
}

// ProcessDatabase simulates database operations with cleanup
func ProcessDatabase(query string) {
	conn := &MockDBConnection{}

	defer func() {
		conn.Close()
		if r := recover(); r != nil {
			fmt.Printf("  Database query failed: %v\n", r)
		}
	}()

	conn.Connect()

	// Simulate query processing
	if strings.Contains(query, "INVALID") {
		panic("syntax error in SQL query")
	}
	if strings.Contains(query, "missing_table") {
		panic("table does not exist")
	}

	fmt.Printf("  Query executed successfully: %s\n", query)
}

// MockDBConnection simulates a database connection
type MockDBConnection struct {
	isConnected bool
}

func (c *MockDBConnection) Connect() {
	c.isConnected = true
	fmt.Printf("    Database connected\n")
}

func (c *MockDBConnection) Close() {
	if c.isConnected {
		c.isConnected = false
		fmt.Printf("    Database connection closed\n")
	}
}

// DemoPanicBestPractices demonstrates best practices for panic/recover
func DemoPanicBestPractices() {
	fmt.Println("\n=== Panic Best Practices Demo ===")

	fmt.Println("1. Use panic for truly exceptional conditions:")
	ValidateConfigFile("valid_config.json")
	ValidateConfigFile("invalid_config.json")

	fmt.Println("\n2. Recover at package boundaries:")
	result := PublicAPIFunction("valid_input")
	fmt.Printf("   API result: %s\n", result)

	result = PublicAPIFunction("panic_input")
	fmt.Printf("   API result: %s\n", result)

	fmt.Println("\n3. Don't recover from all panics:")
	DemonstrateSelectiveRecovery()

	fmt.Println("\n4. Provide meaningful panic messages:")
	DemonstrateMeaningfulPanics()

	fmt.Println("\nBest practices summary:")
	fmt.Println("- Use panic for programming errors, not expected errors")
	fmt.Println("- Recover at package boundaries to protect callers")
	fmt.Println("- Don't recover from runtime panics (let them crash)")
	fmt.Println("- Provide clear, actionable panic messages")
	fmt.Println("- Always clean up resources in defer functions")
}

// ValidateConfigFile demonstrates panic for configuration errors
func ValidateConfigFile(filename string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("   Config validation failed: %v\n", r)
		}
	}()

	if strings.Contains(filename, "invalid") {
		panic(fmt.Sprintf("critical configuration error in %s: missing required field 'database_url'", filename))
	}

	fmt.Printf("   Config file %s is valid\n", filename)
}

// PublicAPIFunction demonstrates recovery at package boundaries
func PublicAPIFunction(input string) string {
	defer func() {
		if r := recover(); r != nil {
			// Log the error in real applications
			fmt.Printf("   [INTERNAL ERROR] %v\n", r)
		}
	}()

	return internalFunction(input)
}

func internalFunction(input string) string {
	if input == "panic_input" {
		panic("internal processing error")
	}
	return fmt.Sprintf("processed: %s", input)
}

// DemonstrateSelectiveRecovery shows when NOT to recover
func DemonstrateSelectiveRecovery() {
	defer func() {
		if r := recover(); r != nil {
			panicStr := fmt.Sprintf("%v", r)
			
			// Only recover from application panics, not runtime panics
			if strings.Contains(panicStr, "application error") {
				fmt.Printf("   Recovered from application error: %v\n", r)
			} else {
				fmt.Printf("   Runtime panic detected - re-panicking: %v\n", r)
				panic(r) // Re-panic runtime errors
			}
		}
	}()

	// This would be recovered
	// panic("application error: business logic failed")

	// This should crash the program (commented out for safety)
	// var slice []int
	// _ = slice[100] // Runtime panic - should not be recovered

	fmt.Println("   Selective recovery demonstration completed")
}

// DemonstrateMeaningfulPanics shows how to create useful panic messages
func DemonstrateMeaningfulPanics() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("   Meaningful panic caught: %v\n", r)
		}
	}()

	// Example of a meaningful panic message
	userID := "user123"
	orderID := "order456"

	// Bad: panic("error")
	// Good: Detailed panic with context
	panic(fmt.Sprintf("order processing failed: user_id=%s, order_id=%s, reason=insufficient_balance, available=50.00, required=100.00", 
		userID, orderID))
}

// Utility functions for safe operations

// SafeSliceAccess safely accesses slice elements
func SafeSliceAccess(slice []int, index int) (int, error) {
	defer func() {
		if r := recover(); r != nil {
			// Convert panic to error
		}
	}()

	if index < 0 || index >= len(slice) {
		return 0, fmt.Errorf("index %d out of range for slice of length %d", index, len(slice))
	}

	return slice[index], nil
}

// SafeStringToInt safely converts string to int
func SafeStringToInt(s string) (int, error) {
	defer func() {
		if r := recover(); r != nil {
			// Convert panic to error if needed
		}
	}()

	return strconv.Atoi(s)
}

// SafeMapAccess safely accesses map values
func SafeMapAccess(m map[string]interface{}, key string) (interface{}, bool) {
	defer func() {
		if r := recover(); r != nil {
			// Handle potential map access panics
		}
	}()

	value, exists := m[key]
	return value, exists
}

// MustOperation demonstrates "must" pattern for operations that should never fail
func MustOperation(operation string) string {
	switch operation {
	case "parse_config":
		// In real code, this might parse embedded config
		return "config_parsed"
	case "invalid_operation":
		panic(fmt.Sprintf("critical error: operation '%s' is not supported in this build", operation))
	default:
		return fmt.Sprintf("operation_%s_completed", operation)
	}
}

// SafeChannelSend safely sends to a channel without panicking
func SafeChannelSend(ch chan<- string, value string) (sent bool) {
	defer func() {
		if r := recover(); r != nil {
			sent = false
		}
	}()

	select {
	case ch <- value:
		return true
	default:
		return false // Channel is full
	}
}

// DemoMustPattern demonstrates the "must" pattern
func DemoMustPattern() {
	fmt.Println("\n=== Must Pattern Demo ===")

	// Operations that should always succeed
	result1 := MustOperation("parse_config")
	fmt.Printf("Must operation result: %s\n", result1)

	// Handle must operations that might fail
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Must operation failed: %v\n", r)
		}
	}()

	result2 := MustOperation("invalid_operation")
	fmt.Printf("This won't be printed: %s\n", result2)
}