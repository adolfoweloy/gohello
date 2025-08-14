// Package errors demonstrates error handling in Go
package errors

import (
	"errors"
	"fmt"
	"strconv"
)

// Custom error types

// ValidationError represents a validation error
type ValidationError struct {
	Field   string
	Value   interface{}
	Message string
}

// Error implements the error interface
func (ve ValidationError) Error() string {
	return fmt.Sprintf("validation error for field '%s' with value '%v': %s", 
		ve.Field, ve.Value, ve.Message)
}

// BusinessError represents a business logic error
type BusinessError struct {
	Code    string
	Message string
	Cause   error
}

// Error implements the error interface
func (be BusinessError) Error() string {
	if be.Cause != nil {
		return fmt.Sprintf("[%s] %s: %v", be.Code, be.Message, be.Cause)
	}
	return fmt.Sprintf("[%s] %s", be.Code, be.Message)
}

// Unwrap returns the underlying error (for Go 1.13+ error unwrapping)
func (be BusinessError) Unwrap() error {
	return be.Cause
}

// DemonstrateBasicErrors shows basic error handling patterns
func DemonstrateBasicErrors() {
	fmt.Println("=== Basic Error Handling ===")
	
	// Function that returns an error
	result, err := divide(10, 2)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("10 / 2 = %.2f\n", result)
	}
	
	// Error case
	_, err = divide(10, 0)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	
	// Multiple return values with error
	value, err := parseAndValidate("42")
	if err != nil {
		fmt.Printf("Parse error: %v\n", err)
	} else {
		fmt.Printf("Parsed value: %d\n", value)
	}
	
	// Invalid input
	_, err = parseAndValidate("abc")
	if err != nil {
		fmt.Printf("Parse error: %v\n", err)
	}
}

// divide performs division with error handling
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// parseAndValidate parses a string to int and validates it
func parseAndValidate(s string) (int, error) {
	value, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("failed to parse '%s': %w", s, err)
	}
	
	if value < 0 {
		return 0, fmt.Errorf("value must be non-negative, got %d", value)
	}
	
	return value, nil
}

// DemonstrateCustomErrors shows custom error types
func DemonstrateCustomErrors() {
	fmt.Println("\n=== Custom Errors ===")
	
	// Validation errors
	users := []User{
		{Name: "Alice", Email: "alice@example.com", Age: 25},
		{Name: "", Email: "bob@example.com", Age: 30},
		{Name: "Charlie", Email: "invalid-email", Age: 15},
		{Name: "Diana", Email: "diana@example.com", Age: -5},
	}
	
	for i, user := range users {
		fmt.Printf("Validating user %d:\n", i+1)
		if err := validateUser(user); err != nil {
			fmt.Printf("  Error: %v\n", err)
			
			// Type assertion to handle specific error types
			if ve, ok := err.(ValidationError); ok {
				fmt.Printf("  Field: %s, Value: %v\n", ve.Field, ve.Value)
			}
		} else {
			fmt.Printf("  Valid user: %s\n", user.Name)
		}
	}
	
	// Business errors
	fmt.Println("\nBusiness error examples:")
	account := BankAccount{ID: "ACC001", Balance: 100.0}
	
	// Successful withdrawal
	if err := account.Withdraw(50.0); err != nil {
		handleBusinessError(err)
	} else {
		fmt.Printf("Withdrawal successful. New balance: %.2f\n", account.Balance)
	}
	
	// Failed withdrawal - insufficient funds
	if err := account.Withdraw(100.0); err != nil {
		handleBusinessError(err)
	}
	
	// Failed withdrawal - invalid amount
	if err := account.Withdraw(-10.0); err != nil {
		handleBusinessError(err)
	}
}

// User represents a user
type User struct {
	Name  string
	Email string
	Age   int
}

// validateUser validates a user and returns specific validation errors
func validateUser(user User) error {
	if user.Name == "" {
		return ValidationError{
			Field:   "name",
			Value:   user.Name,
			Message: "name cannot be empty",
		}
	}
	
	if user.Age < 0 {
		return ValidationError{
			Field:   "age",
			Value:   user.Age,
			Message: "age cannot be negative",
		}
	}
	
	if user.Age < 18 {
		return ValidationError{
			Field:   "age",
			Value:   user.Age,
			Message: "user must be at least 18 years old",
		}
	}
	
	// Simple email validation
	if !contains(user.Email, "@") || !contains(user.Email, ".") {
		return ValidationError{
			Field:   "email",
			Value:   user.Email,
			Message: "invalid email format",
		}
	}
	
	return nil
}

// contains checks if a string contains a substring
func contains(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

// BankAccount represents a bank account
type BankAccount struct {
	ID      string
	Balance float64
}

// Withdraw withdraws money from the account
func (ba *BankAccount) Withdraw(amount float64) error {
	if amount <= 0 {
		return BusinessError{
			Code:    "INVALID_AMOUNT",
			Message: "withdrawal amount must be positive",
		}
	}
	
	if amount > ba.Balance {
		return BusinessError{
			Code:    "INSUFFICIENT_FUNDS",
			Message: fmt.Sprintf("insufficient funds: requested %.2f, available %.2f", amount, ba.Balance),
		}
	}
	
	ba.Balance -= amount
	return nil
}

// handleBusinessError handles business errors specifically
func handleBusinessError(err error) {
	if be, ok := err.(BusinessError); ok {
		switch be.Code {
		case "INSUFFICIENT_FUNDS":
			fmt.Printf("  Transaction declined: %s\n", be.Message)
		case "INVALID_AMOUNT":
			fmt.Printf("  Invalid request: %s\n", be.Message)
		default:
			fmt.Printf("  Business error [%s]: %s\n", be.Code, be.Message)
		}
	} else {
		fmt.Printf("  Unexpected error: %v\n", err)
	}
}

// DemonstratePanicRecover shows panic and recover mechanisms
func DemonstratePanicRecover() {
	fmt.Println("\n=== Panic and Recover ===")
	
	// Safe function that recovers from panic
	fmt.Println("Calling risky function with recovery:")
	callRiskyFunctionSafely(5)
	callRiskyFunctionSafely(0)
	callRiskyFunctionSafely(-1)
	
	fmt.Println("\nDemonstrating defer with panic recovery:")
	demonstrateDefer()
	
	fmt.Println("\nProgram continues after panic recovery!")
}

// riskyFunction simulates a function that might panic
func riskyFunction(value int) int {
	if value == 0 {
		panic("zero value not allowed")
	}
	if value < 0 {
		panic(fmt.Sprintf("negative value not allowed: %d", value))
	}
	return 100 / value
}

// callRiskyFunctionSafely calls risky function with panic recovery
func callRiskyFunctionSafely(value int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("  Recovered from panic with value %d: %v\n", value, r)
		}
	}()
	
	result := riskyFunction(value)
	fmt.Printf("  Result for %d: %d\n", value, result)
}

// demonstrateDefer shows defer execution order and panic recovery
func demonstrateDefer() {
	defer fmt.Println("  Defer 1: This runs last")
	defer fmt.Println("  Defer 2: This runs second")
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("  Defer 3: Recovered from panic: %v\n", r)
		}
	}()
	defer fmt.Println("  Defer 4: This runs first")
	
	fmt.Println("  About to panic...")
	panic("demonstration panic")
	fmt.Println("  This line will never execute")
}

// DemonstrateErrorWrapping shows Go 1.13+ error wrapping
func DemonstrateErrorWrapping() {
	fmt.Println("\n=== Error Wrapping (Go 1.13+) ===")
	
	// Create a chain of wrapped errors
	err := processData("invalid-data")
	if err != nil {
		fmt.Printf("Final error: %v\n", err)
		
		// Unwrap errors to find the root cause
		unwrapped := errors.Unwrap(err)
		if unwrapped != nil {
			fmt.Printf("Unwrapped error: %v\n", unwrapped)
		}
		
		// Check for specific error types using errors.As
		var ve ValidationError
		if errors.As(err, &ve) {
			fmt.Printf("Found validation error: %v\n", ve)
		}
		
		// Check for specific error values using errors.Is
		if errors.Is(err, ErrInvalidFormat) {
			fmt.Println("Error is related to invalid format")
		}
	}
}

// Predefined errors
var (
	ErrInvalidFormat = errors.New("invalid format")
	ErrNotFound      = errors.New("not found")
)

// processData simulates data processing with error wrapping
func processData(data string) error {
	if err := validateFormat(data); err != nil {
		return fmt.Errorf("failed to process data: %w", err)
	}
	return nil
}

// validateFormat validates data format
func validateFormat(data string) error {
	if data == "invalid-data" {
		ve := ValidationError{
			Field:   "data",
			Value:   data,
			Message: "data format is invalid",
		}
		return fmt.Errorf("validation failed: %w", ve)
	}
	return nil
}