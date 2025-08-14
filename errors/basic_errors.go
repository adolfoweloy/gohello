// Package errors demonstrates Go's error handling patterns including
// basic error handling, custom errors, and panic/recover mechanisms.
package errors

import (
	stderrors "errors"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// DemoBasicErrors demonstrates basic error handling in Go
func DemoBasicErrors() {
	fmt.Println("\n=== Basic Error Handling Demo ===")

	// Function that returns value and error
	result, err := divide(10, 2)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("10 ÷ 2 = %.2f\n", result)
	}

	// Error case
	_, err = divide(10, 0)
	if err != nil {
		fmt.Printf("Error dividing by zero: %v\n", err)
	}

	// Multiple error checks
	fmt.Println("\nProcessing numbers:")
	numbers := []string{"10", "20", "abc", "30", "0"}
	for i, numStr := range numbers {
		if num, err := parseAndDouble(numStr); err != nil {
			fmt.Printf("  [%d] Error processing '%s': %v\n", i, numStr, err)
		} else {
			fmt.Printf("  [%d] '%s' doubled = %d\n", i, numStr, num)
		}
	}

	// Chaining operations with errors
	fmt.Println("\nChaining operations:")
	if result, err := processChain("100"); err != nil {
		fmt.Printf("Chain failed: %v\n", err)
	} else {
		fmt.Printf("Chain result: %d\n", result)
	}

	if _, err := processChain("invalid"); err != nil {
		fmt.Printf("Chain failed with invalid input: %v\n", err)
	}
}

// divide performs division with error handling
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, stderrors.New("division by zero")
	}
	return a / b, nil
}

// parseAndDouble parses a string and doubles the number
func parseAndDouble(s string) (int, error) {
	num, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("failed to parse '%s' as integer: %w", s, err)
	}
	return num * 2, nil
}

// processChain demonstrates error propagation through multiple function calls
func processChain(input string) (int, error) {
	// Step 1: Parse
	num, err := strconv.Atoi(input)
	if err != nil {
		return 0, fmt.Errorf("parse error: %w", err)
	}

	// Step 2: Validate
	if err := validateRange(num); err != nil {
		return 0, fmt.Errorf("validation error: %w", err)
	}

	// Step 3: Process
	result, err := complexCalculation(num)
	if err != nil {
		return 0, fmt.Errorf("calculation error: %w", err)
	}

	return result, nil
}

// validateRange validates that a number is in acceptable range
func validateRange(num int) error {
	if num < 0 {
		return stderrors.New("number must be non-negative")
	}
	if num > 1000 {
		return stderrors.New("number must be <= 1000")
	}
	return nil
}

// complexCalculation performs a calculation that might fail
func complexCalculation(num int) (int, error) {
	if num%7 == 0 && num != 0 {
		return 0, stderrors.New("number divisible by 7 not supported")
	}
	return num * num, nil
}

// DemoErrorTypes demonstrates different error creation methods
func DemoErrorTypes() {
	fmt.Println("\n=== Error Types Demo ===")

	// stderrors.New
	err1 := stderrors.New("simple error message")
	fmt.Printf("stderrors.New: %v\n", err1)

	// fmt.Errorf
	value := 42
	err2 := fmt.Errorf("error with value %d", value)
	fmt.Printf("fmt.Errorf: %v\n", err2)

	// fmt.Errorf with error wrapping
	originalErr := stderrors.New("original error")
	wrappedErr := fmt.Errorf("wrapped error: %w", originalErr)
	fmt.Printf("Wrapped error: %v\n", wrappedErr)

	// Unwrapping errors
	fmt.Printf("Unwrapped: %v\n", stderrors.Unwrap(wrappedErr))

	// stderrors.Is for error comparison
	if stderrors.Is(wrappedErr, originalErr) {
		fmt.Println("wrappedErr contains originalErr")
	}

	// Multiple operations with error handling
	fmt.Println("\nFile operation simulation:")
	if err := simulateFileOperations("test.txt"); err != nil {
		fmt.Printf("File operation failed: %v\n", err)
		
		// Check for specific error types
		if stderrors.Is(err, io.EOF) {
			fmt.Println("Reached end of file")
		}
	}
}

// simulateFileOperations simulates file operations with different error types
func simulateFileOperations(filename string) error {
	// Simulate file not found
	if strings.Contains(filename, "missing") {
		return fmt.Errorf("file not found: %s", filename)
	}

	// Simulate permission error
	if strings.Contains(filename, "restricted") {
		return fmt.Errorf("permission denied: %s", filename)
	}

	// Simulate EOF
	if strings.Contains(filename, "empty") {
		return fmt.Errorf("reading file %s: %w", filename, io.EOF)
	}

	// Success case
	fmt.Printf("File operation on %s successful\n", filename)
	return nil
}

// DemoErrorSentinel demonstrates sentinel errors pattern
func DemoErrorSentinel() {
	fmt.Println("\n=== Sentinel Errors Demo ===")

	// Using predefined sentinel errors
	operations := []string{"read", "write", "delete", "invalid"}
	
	for _, op := range operations {
		if err := performOperation(op); err != nil {
			fmt.Printf("Operation '%s' failed: %v\n", op, err)
			
			// Check for specific sentinel errors
			switch {
			case stderrors.Is(err, ErrPermissionDenied):
				fmt.Println("  → Need higher privileges")
			case stderrors.Is(err, ErrResourceNotFound):
				fmt.Println("  → Resource doesn't exist")
			case stderrors.Is(err, ErrInvalidOperation):
				fmt.Println("  → Operation not supported")
			default:
				fmt.Println("  → Unknown error type")
			}
		} else {
			fmt.Printf("Operation '%s' succeeded\n", op)
		}
	}
}

// Sentinel errors - predefined error variables
var (
	ErrPermissionDenied  = stderrors.New("permission denied")
	ErrResourceNotFound  = stderrors.New("resource not found")
	ErrInvalidOperation  = stderrors.New("invalid operation")
	ErrTimeout          = stderrors.New("operation timeout")
)

// performOperation simulates operations that can fail with sentinel errors
func performOperation(operation string) error {
	switch operation {
	case "read":
		return nil // Success
	case "write":
		return ErrPermissionDenied
	case "delete":
		return ErrResourceNotFound
	case "timeout":
		return ErrTimeout
	default:
		return ErrInvalidOperation
	}
}

// DemoErrorWrapping demonstrates error wrapping and unwrapping
func DemoErrorWrapping() {
	fmt.Println("\n=== Error Wrapping Demo ===")

	// Simulate nested function calls with error wrapping
	if err := topLevelFunction("test"); err != nil {
		fmt.Printf("Top level error: %v\n", err)
		
		// Unwrap to see the chain
		fmt.Println("\nError chain:")
		currentErr := err
		level := 0
		for currentErr != nil {
			fmt.Printf("  Level %d: %v\n", level, currentErr)
			currentErr = stderrors.Unwrap(currentErr)
			level++
		}
		
		// Check if it contains specific errors
		if stderrors.Is(err, ErrResourceNotFound) {
			fmt.Println("\nRoot cause: Resource not found")
		}
	}

	// Demonstrate error.As for type assertions
	fmt.Println("\nError type assertion:")
	if err := functionWithCustomError(); err != nil {
		var customErr *CustomError
		if stderrors.As(err, &customErr) {
			fmt.Printf("Custom error details: Code=%d, Message=%s\n", 
				customErr.Code, customErr.Message)
		}
	}
}

// topLevelFunction starts a chain of function calls
func topLevelFunction(input string) error {
	if err := middleFunction(input); err != nil {
		return fmt.Errorf("top level failed: %w", err)
	}
	return nil
}

// middleFunction is in the middle of the call chain
func middleFunction(input string) error {
	if err := bottomFunction(input); err != nil {
		return fmt.Errorf("middle level failed: %w", err)
	}
	return nil
}

// bottomFunction is at the bottom of the call chain
func bottomFunction(input string) error {
	if input == "test" {
		return fmt.Errorf("bottom level error: %w", ErrResourceNotFound)
	}
	return nil
}

// functionWithCustomError returns a custom error type
func functionWithCustomError() error {
	return &CustomError{
		Code:    404,
		Message: "Resource not found",
		Details: "The requested resource could not be located",
	}
}

// DemoMultipleErrors demonstrates handling multiple errors
func DemoMultipleErrors() {
	fmt.Println("\n=== Multiple Errors Demo ===")

	// Validate multiple fields
	user := User{
		Name:  "",
		Email: "invalid-email",
		Age:   -5,
	}

	if errs := validateUser(user); len(errs) > 0 {
		fmt.Println("User validation failed:")
		for i, err := range errs {
			fmt.Printf("  %d. %v\n", i+1, err)
		}
	}

	// Valid user
	validUser := User{
		Name:  "Alice Johnson",
		Email: "alice@example.com",
		Age:   25,
	}

	if errs := validateUser(validUser); len(errs) == 0 {
		fmt.Printf("User %s is valid\n", validUser.Name)
	}

	// Batch operation with multiple errors
	fmt.Println("\nBatch operation results:")
	numbers := []string{"10", "abc", "20", "xyz", "30"}
	results, errors := processBatch(numbers)
	
	fmt.Printf("Successful results: %v\n", results)
	if len(errors) > 0 {
		fmt.Println("Errors encountered:")
		for i, err := range errors {
			fmt.Printf("  %d. %v\n", i+1, err)
		}
	}
}

// User represents a user with validation
type User struct {
	Name  string
	Email string
	Age   int
}

// validateUser validates user fields and returns all errors
func validateUser(user User) []error {
	var errs []error

	if user.Name == "" {
		errs = append(errs, stderrors.New("name is required"))
	}

	if user.Email == "" {
		errs = append(errs, stderrors.New("email is required"))
	} else if !strings.Contains(user.Email, "@") {
		errs = append(errs, stderrors.New("email must contain @"))
	}

	if user.Age < 0 {
		errs = append(errs, stderrors.New("age must be non-negative"))
	} else if user.Age > 150 {
		errs = append(errs, stderrors.New("age must be realistic"))
	}

	return errs
}

// processBatch processes multiple items and collects all errors
func processBatch(items []string) ([]int, []error) {
	var results []int
	var errors []error

	for i, item := range items {
		if num, err := strconv.Atoi(item); err != nil {
			errors = append(errors, fmt.Errorf("item %d ('%s'): %w", i, item, err))
		} else {
			results = append(results, num)
		}
	}

	return results, errors
}

// DemoErrorInterfaces demonstrates error interface and custom implementations
func DemoErrorInterfaces() {
	fmt.Println("\n=== Error Interfaces Demo ===")

	// The error interface is simple: Error() string
	var err error = &CustomError{
		Code:    500,
		Message: "Internal server error",
		Details: "Database connection failed",
	}

	fmt.Printf("Error string: %v\n", err)

	// Type assertion to access custom methods
	if customErr, ok := err.(*CustomError); ok {
		fmt.Printf("Error code: %d\n", customErr.Code)
		fmt.Printf("Error details: %s\n", customErr.Details)
		fmt.Printf("Is client error: %t\n", customErr.IsClientError())
		fmt.Printf("Is server error: %t\n", customErr.IsServerError())
	}

	// Different error types
	errors := []error{
		&ValidationError{Field: "email", Value: "invalid"},
		&NetworkError{Host: "api.example.com", Port: 443, Timeout: true},
		&CustomError{Code: 401, Message: "Unauthorized"},
	}

	fmt.Println("\nDifferent error types:")
	for i, err := range errors {
		fmt.Printf("%d. %v\n", i+1, err)
		
		// Type switching on errors
		switch e := err.(type) {
		case *ValidationError:
			fmt.Printf("   → Validation error on field: %s\n", e.Field)
		case *NetworkError:
			fmt.Printf("   → Network error to %s:%d (timeout: %t)\n", 
				e.Host, e.Port, e.Timeout)
		case *CustomError:
			fmt.Printf("   → Custom error with code: %d\n", e.Code)
		}
	}
}

// CustomError is a custom error type with additional information
type CustomError struct {
	Code    int
	Message string
	Details string
}

// Error implements the error interface
func (e *CustomError) Error() string {
	return fmt.Sprintf("error %d: %s", e.Code, e.Message)
}

// IsClientError checks if error is a client error (4xx)
func (e *CustomError) IsClientError() bool {
	return e.Code >= 400 && e.Code < 500
}

// IsServerError checks if error is a server error (5xx)
func (e *CustomError) IsServerError() bool {
	return e.Code >= 500 && e.Code < 600
}

// ValidationError represents validation errors
type ValidationError struct {
	Field string
	Value string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation failed for field '%s' with value '%s'", e.Field, e.Value)
}

// NetworkError represents network-related errors
type NetworkError struct {
	Host    string
	Port    int
	Timeout bool
}

func (e *NetworkError) Error() string {
	if e.Timeout {
		return fmt.Sprintf("network timeout connecting to %s:%d", e.Host, e.Port)
	}
	return fmt.Sprintf("network error connecting to %s:%d", e.Host, e.Port)
}

// DemoErrorBestPractices demonstrates error handling best practices
func DemoErrorBestPractices() {
	fmt.Println("\n=== Error Handling Best Practices Demo ===")

	// 1. Always check errors
	fmt.Println("1. Always check errors:")
	if result, err := reliableOperation(); err != nil {
		fmt.Printf("   Operation failed: %v\n", err)
	} else {
		fmt.Printf("   Operation succeeded: %s\n", result)
	}

	// 2. Provide context in error messages
	fmt.Println("\n2. Provide context:")
	if err := saveUserToDatabase(User{Name: "Test"}); err != nil {
		fmt.Printf("   %v\n", err)
	}

	// 3. Use error wrapping to preserve context
	fmt.Println("\n3. Error wrapping:")
	if err := processUserData("invalid"); err != nil {
		fmt.Printf("   %v\n", err)
	}

	// 4. Handle errors at the appropriate level
	fmt.Println("\n4. Appropriate error handling:")
	userService := &UserService{}
	if err := userService.CreateUser("Alice", "alice@example.com"); err != nil {
		fmt.Printf("   Service error: %v\n", err)
	}

	// 5. Use sentinel errors for expected error conditions
	fmt.Println("\n5. Sentinel errors:")
	if err := authenticateUser("guest", "wrongpassword"); err != nil {
		if stderrors.Is(err, ErrInvalidCredentials) {
			fmt.Println("   → Login failed - please check credentials")
		} else {
			fmt.Printf("   → Authentication system error: %v\n", err)
		}
	}

	fmt.Println("\nBest practices summary:")
	fmt.Println("- Always check errors explicitly")
	fmt.Println("- Provide meaningful error messages with context")
	fmt.Println("- Use error wrapping to preserve error chains")
	fmt.Println("- Handle errors at the appropriate abstraction level")
	fmt.Println("- Use sentinel errors for expected conditions")
	fmt.Println("- Don't ignore errors (use _ = err if intentional)")
}

// Sentinel error for authentication
var ErrInvalidCredentials = stderrors.New("invalid credentials")

// reliableOperation simulates an operation that might fail
func reliableOperation() (string, error) {
	// Simulate success
	return "operation completed successfully", nil
}

// saveUserToDatabase simulates saving user with context
func saveUserToDatabase(user User) error {
	if user.Name == "" {
		return fmt.Errorf("failed to save user to database: %w", 
			stderrors.New("name is required"))
	}
	// Simulate database error
	return fmt.Errorf("failed to save user '%s' to database: %w", 
		user.Name, stderrors.New("connection timeout"))
}

// processUserData demonstrates error wrapping
func processUserData(data string) error {
	if err := validateUserData(data); err != nil {
		return fmt.Errorf("processing user data: %w", err)
	}
	return nil
}

// validateUserData validates user data
func validateUserData(data string) error {
	if data == "invalid" {
		return fmt.Errorf("validating user data: %w", 
			stderrors.New("data format is incorrect"))
	}
	return nil
}

// UserService demonstrates service-level error handling
type UserService struct{}

func (us *UserService) CreateUser(name, email string) error {
	if err := us.validateUserInput(name, email); err != nil {
		return fmt.Errorf("user creation failed: %w", err)
	}
	
	if err := us.saveUser(name, email); err != nil {
		return fmt.Errorf("user creation failed: %w", err)
	}
	
	return nil
}

func (us *UserService) validateUserInput(name, email string) error {
	if name == "" {
		return stderrors.New("name is required")
	}
	if !strings.Contains(email, "@") {
		return stderrors.New("invalid email format")
	}
	return nil
}

func (us *UserService) saveUser(name, email string) error {
	// Simulate database operation
	if name == "TestError" {
		return stderrors.New("database connection failed")
	}
	return nil
}

// authenticateUser simulates user authentication
func authenticateUser(username, password string) error {
	validUsers := map[string]string{
		"admin": "admin123",
		"user":  "user123",
	}
	
	if validPassword, exists := validUsers[username]; !exists || validPassword != password {
		return ErrInvalidCredentials
	}
	
	return nil
}