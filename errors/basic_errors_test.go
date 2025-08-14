package errors

import (
	stderrors "errors"
	"fmt"
	"io"
	"strings"
	"testing"
)

func TestDemoBasicErrors(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("DemoBasicErrors panicked: %v", r)
		}
	}()
	
	DemoBasicErrors()
}

func TestDemoErrorTypes(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("DemoErrorTypes panicked: %v", r)
		}
	}()
	
	DemoErrorTypes()
}

func TestDemoErrorSentinel(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("DemoErrorSentinel panicked: %v", r)
		}
	}()
	
	DemoErrorSentinel()
}

func TestDemoErrorWrapping(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("DemoErrorWrapping panicked: %v", r)
		}
	}()
	
	DemoErrorWrapping()
}

func TestDemoMultipleErrors(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("DemoMultipleErrors panicked: %v", r)
		}
	}()
	
	DemoMultipleErrors()
}

func TestDemoErrorInterfaces(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("DemoErrorInterfaces panicked: %v", r)
		}
	}()
	
	DemoErrorInterfaces()
}

func TestDemoErrorBestPractices(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("DemoErrorBestPractices panicked: %v", r)
		}
	}()
	
	DemoErrorBestPractices()
}

func TestDivide(t *testing.T) {
	// Test successful division
	result, err := divide(10, 2)
	if err != nil {
		t.Errorf("Expected no error for valid division, got: %v", err)
	}
	if result != 5.0 {
		t.Errorf("Expected result 5.0, got %.2f", result)
	}
	
	// Test division by zero
	_, err = divide(10, 0)
	if err == nil {
		t.Errorf("Expected error for division by zero")
	}
	if err.Error() != "division by zero" {
		t.Errorf("Expected 'division by zero' error, got: %v", err)
	}
}

func TestParseAndDouble(t *testing.T) {
	// Test valid input
	result, err := parseAndDouble("5")
	if err != nil {
		t.Errorf("Expected no error for valid input, got: %v", err)
	}
	if result != 10 {
		t.Errorf("Expected result 10, got %d", result)
	}
	
	// Test invalid input
	_, err = parseAndDouble("abc")
	if err == nil {
		t.Errorf("Expected error for invalid input")
	}
	if !strings.Contains(err.Error(), "failed to parse") {
		t.Errorf("Error should contain 'failed to parse', got: %v", err)
	}
}

func TestProcessChain(t *testing.T) {
	// Test valid input
	result, err := processChain("25")
	if err != nil {
		t.Errorf("Expected no error for valid input, got: %v", err)
	}
	if result != 625 { // 25^2
		t.Errorf("Expected result 625, got %d", result)
	}
	
	// Test invalid parse
	_, err = processChain("abc")
	if err == nil {
		t.Errorf("Expected parse error")
	}
	if !strings.Contains(err.Error(), "parse error") {
		t.Errorf("Error should contain 'parse error', got: %v", err)
	}
	
	// Test range validation error
	_, err = processChain("2000")
	if err == nil {
		t.Errorf("Expected validation error for large number")
	}
	if !strings.Contains(err.Error(), "validation error") {
		t.Errorf("Error should contain 'validation error', got: %v", err)
	}
	
	// Test calculation error (divisible by 7)
	_, err = processChain("14")
	if err == nil {
		t.Errorf("Expected calculation error for number divisible by 7")
	}
	if !strings.Contains(err.Error(), "calculation error") {
		t.Errorf("Error should contain 'calculation error', got: %v", err)
	}
}

func TestValidateRange(t *testing.T) {
	// Test valid range
	err := validateRange(50)
	if err != nil {
		t.Errorf("Expected no error for valid range, got: %v", err)
	}
	
	// Test negative number
	err = validateRange(-5)
	if err == nil {
		t.Errorf("Expected error for negative number")
	}
	
	// Test too large number
	err = validateRange(2000)
	if err == nil {
		t.Errorf("Expected error for too large number")
	}
}

func TestComplexCalculation(t *testing.T) {
	// Test valid calculation
	result, err := complexCalculation(5)
	if err != nil {
		t.Errorf("Expected no error for valid input, got: %v", err)
	}
	if result != 25 {
		t.Errorf("Expected result 25, got %d", result)
	}
	
	// Test number divisible by 7
	_, err = complexCalculation(14)
	if err == nil {
		t.Errorf("Expected error for number divisible by 7")
	}
	
	// Test zero (special case - divisible by 7 but allowed)
	result, err = complexCalculation(0)
	if err != nil {
		t.Errorf("Expected no error for zero, got: %v", err)
	}
	if result != 0 {
		t.Errorf("Expected result 0, got %d", result)
	}
}

func TestErrorWrapping(t *testing.T) {
	// Test error wrapping with fmt.Errorf
	originalErr := stderrors.New("original error")
	wrappedErr := fmt.Errorf("wrapped: %w", originalErr)
	
	// Test that wrapping preserves the original error
	if !stderrors.Is(wrappedErr, originalErr) {
		t.Errorf("Wrapped error should contain original error")
	}
	
	// Test unwrapping
	unwrapped := stderrors.Unwrap(wrappedErr)
	if unwrapped != originalErr {
		t.Errorf("Unwrapped error should be the original error")
	}
}

func TestSimulateFileOperations(t *testing.T) {
	// Test successful operation
	err := simulateFileOperations("test.txt")
	if err != nil {
		t.Errorf("Expected no error for normal file, got: %v", err)
	}
	
	// Test file not found
	err = simulateFileOperations("missing_file.txt")
	if err == nil {
		t.Errorf("Expected error for missing file")
	}
	if !strings.Contains(err.Error(), "file not found") {
		t.Errorf("Error should contain 'file not found', got: %v", err)
	}
	
	// Test permission error
	err = simulateFileOperations("restricted_file.txt")
	if err == nil {
		t.Errorf("Expected error for restricted file")
	}
	if !strings.Contains(err.Error(), "permission denied") {
		t.Errorf("Error should contain 'permission denied', got: %v", err)
	}
	
	// Test EOF error
	err = simulateFileOperations("empty_file.txt")
	if err == nil {
		t.Errorf("Expected error for empty file")
	}
	if !stderrors.Is(err, io.EOF) {
		t.Errorf("Error should wrap io.EOF")
	}
}

func TestSentinelErrors(t *testing.T) {
	// Test each sentinel error
	testCases := []struct {
		operation string
		expected  error
	}{
		{"read", nil},
		{"write", ErrPermissionDenied},
		{"delete", ErrResourceNotFound},
		{"timeout", ErrTimeout},
		{"invalid", ErrInvalidOperation},
	}
	
	for _, tc := range testCases {
		err := performOperation(tc.operation)
		if tc.expected == nil {
			if err != nil {
				t.Errorf("Operation %s: expected no error, got %v", tc.operation, err)
			}
		} else {
			if !stderrors.Is(err, tc.expected) {
				t.Errorf("Operation %s: expected %v, got %v", tc.operation, tc.expected, err)
			}
		}
	}
}

func TestTopLevelFunction(t *testing.T) {
	// Test error propagation through function chain
	err := topLevelFunction("test")
	if err == nil {
		t.Errorf("Expected error from function chain")
	}
	
	// Verify error chain contains expected errors
	if !stderrors.Is(err, ErrResourceNotFound) {
		t.Errorf("Error chain should contain ErrResourceNotFound")
	}
	
	// Check error message contains context from each level
	errMsg := err.Error()
	if !strings.Contains(errMsg, "top level failed") {
		t.Errorf("Error should contain top level context")
	}
	if !strings.Contains(errMsg, "middle level failed") {
		t.Errorf("Error should contain middle level context")
	}
	if !strings.Contains(errMsg, "bottom level error") {
		t.Errorf("Error should contain bottom level context")
	}
}

func TestCustomError(t *testing.T) {
	customErr := &CustomError{
		Code:    404,
		Message: "Not Found",
		Details: "Resource does not exist",
	}
	
	// Test Error() method
	expectedMsg := "error 404: Not Found"
	if customErr.Error() != expectedMsg {
		t.Errorf("Expected error message '%s', got '%s'", expectedMsg, customErr.Error())
	}
	
	// Test IsClientError
	if !customErr.IsClientError() {
		t.Errorf("404 should be a client error")
	}
	
	// Test IsServerError
	if customErr.IsServerError() {
		t.Errorf("404 should not be a server error")
	}
	
	// Test server error
	serverErr := &CustomError{Code: 500, Message: "Internal Error"}
	if !serverErr.IsServerError() {
		t.Errorf("500 should be a server error")
	}
	if serverErr.IsClientError() {
		t.Errorf("500 should not be a client error")
	}
}

func TestFunctionWithCustomError(t *testing.T) {
	err := functionWithCustomError()
	if err == nil {
		t.Errorf("Expected custom error")
	}
	
	// Test type assertion with stderrors.As
	var customErr *CustomError
	if !stderrors.As(err, &customErr) {
		t.Errorf("Error should be a CustomError")
	}
	
	if customErr.Code != 404 {
		t.Errorf("Expected code 404, got %d", customErr.Code)
	}
}

func TestValidateUser(t *testing.T) {
	// Test valid user
	validUser := User{
		Name:  "Alice",
		Email: "alice@example.com",
		Age:   25,
	}
	errs := validateUser(validUser)
	if len(errs) != 0 {
		t.Errorf("Expected no errors for valid user, got %d errors", len(errs))
	}
	
	// Test invalid user
	invalidUser := User{
		Name:  "",
		Email: "invalid-email",
		Age:   -5,
	}
	errs = validateUser(invalidUser)
	if len(errs) == 0 {
		t.Errorf("Expected validation errors for invalid user")
	}
	
	// Check that we get all expected errors
	expectedErrors := 3 // name, email, age
	if len(errs) != expectedErrors {
		t.Errorf("Expected %d errors, got %d", expectedErrors, len(errs))
	}
	
	// Test edge cases
	oldUser := User{
		Name:  "Very Old Person",
		Email: "old@example.com",
		Age:   200,
	}
	errs = validateUser(oldUser)
	if len(errs) == 0 {
		t.Errorf("Expected age validation error for unrealistic age")
	}
}

func TestProcessBatch(t *testing.T) {
	// Test mixed valid and invalid inputs
	inputs := []string{"10", "abc", "20", "xyz", "30"}
	results, errs := processBatch(inputs)
	
	expectedResults := []int{10, 20, 30}
	if len(results) != len(expectedResults) {
		t.Errorf("Expected %d results, got %d", len(expectedResults), len(results))
	}
	
	for i, expected := range expectedResults {
		if results[i] != expected {
			t.Errorf("Result %d: expected %d, got %d", i, expected, results[i])
		}
	}
	
	expectedErrors := 2 // "abc" and "xyz"
	if len(errs) != expectedErrors {
		t.Errorf("Expected %d errors, got %d", expectedErrors, len(errs))
	}
}

func TestValidationError(t *testing.T) {
	validationErr := &ValidationError{
		Field: "email",
		Value: "invalid",
	}
	
	expectedMsg := "validation failed for field 'email' with value 'invalid'"
	if validationErr.Error() != expectedMsg {
		t.Errorf("Expected error message '%s', got '%s'", expectedMsg, validationErr.Error())
	}
}

func TestNetworkError(t *testing.T) {
	// Test timeout error
	timeoutErr := &NetworkError{
		Host:    "api.example.com",
		Port:    443,
		Timeout: true,
	}
	
	if !strings.Contains(timeoutErr.Error(), "timeout") {
		t.Errorf("Timeout error should contain 'timeout'")
	}
	
	// Test general network error
	networkErr := &NetworkError{
		Host:    "api.example.com",
		Port:    443,
		Timeout: false,
	}
	
	if strings.Contains(networkErr.Error(), "timeout") {
		t.Errorf("Non-timeout error should not contain 'timeout'")
	}
}

func TestUserService(t *testing.T) {
	service := &UserService{}
	
	// Test successful user creation
	err := service.CreateUser("Alice", "alice@example.com")
	if err != nil {
		t.Errorf("Expected no error for valid user creation, got: %v", err)
	}
	
	// Test validation error
	err = service.CreateUser("", "alice@example.com")
	if err == nil {
		t.Errorf("Expected error for empty name")
	}
	
	err = service.CreateUser("Alice", "invalid-email")
	if err == nil {
		t.Errorf("Expected error for invalid email")
	}
	
	// Test database error
	err = service.CreateUser("TestError", "test@example.com")
	if err == nil {
		t.Errorf("Expected error for TestError user")
	}
}

func TestAuthenticateUser(t *testing.T) {
	// Test valid credentials
	err := authenticateUser("admin", "admin123")
	if err != nil {
		t.Errorf("Expected no error for valid credentials, got: %v", err)
	}
	
	// Test invalid username
	err = authenticateUser("invalid", "password")
	if !stderrors.Is(err, ErrInvalidCredentials) {
		t.Errorf("Expected ErrInvalidCredentials for invalid username")
	}
	
	// Test invalid password
	err = authenticateUser("admin", "wrongpassword")
	if !stderrors.Is(err, ErrInvalidCredentials) {
		t.Errorf("Expected ErrInvalidCredentials for invalid password")
	}
}

func TestErrorChainUnwrapping(t *testing.T) {
	// Create a chain of wrapped errors
	baseErr := stderrors.New("base error")
	level1 := fmt.Errorf("level 1: %w", baseErr)
	level2 := fmt.Errorf("level 2: %w", level1)
	level3 := fmt.Errorf("level 3: %w", level2)
	
	// Test that stderrors.Is works through the chain
	if !stderrors.Is(level3, baseErr) {
		t.Errorf("stderrors.Is should find base error in wrapped chain")
	}
	
	// Test manual unwrapping
	current := level3
	depth := 0
	for current != nil {
		depth++
		current = stderrors.Unwrap(current)
		if depth > 10 { // Prevent infinite loop
			t.Errorf("Unwrapping chain too deep")
			break
		}
	}
	
	if depth != 4 { // level3, level2, level1, baseErr
		t.Errorf("Expected unwrapping depth 4, got %d", depth)
	}
}

func TestErrorTypeAssertion(t *testing.T) {
	// Test stderrors.As with different error types
	errors := []error{
		&CustomError{Code: 404, Message: "Not Found"},
		&ValidationError{Field: "name", Value: ""},
		&NetworkError{Host: "localhost", Port: 8080, Timeout: true},
	}
	
	for i, err := range errors {
		switch i {
		case 0:
			var customErr *CustomError
			if !stderrors.As(err, &customErr) {
				t.Errorf("Should be able to assert CustomError")
			}
			if customErr.Code != 404 {
				t.Errorf("CustomError code should be 404")
			}
			
		case 1:
			var validationErr *ValidationError
			if !stderrors.As(err, &validationErr) {
				t.Errorf("Should be able to assert ValidationError")
			}
			if validationErr.Field != "name" {
				t.Errorf("ValidationError field should be 'name'")
			}
			
		case 2:
			var networkErr *NetworkError
			if !stderrors.As(err, &networkErr) {
				t.Errorf("Should be able to assert NetworkError")
			}
			if !networkErr.Timeout {
				t.Errorf("NetworkError should have timeout=true")
			}
		}
	}
}