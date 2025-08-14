// Package errors demonstrates advanced custom error patterns including
// structured errors, error collections, and domain-specific error types.
package errors

import (
	"fmt"
	"time"
)

// DemoCustomErrorTypes demonstrates various custom error implementations
func DemoCustomErrorTypes() {
	fmt.Println("\n=== Custom Error Types Demo ===")

	// HTTP-like errors
	httpErrors := []error{
		&HTTPError{StatusCode: 400, Message: "Bad Request", Details: "Invalid JSON payload"},
		&HTTPError{StatusCode: 401, Message: "Unauthorized", Details: "Token expired"},
		&HTTPError{StatusCode: 404, Message: "Not Found", Details: "User not found"},
		&HTTPError{StatusCode: 500, Message: "Internal Server Error", Details: "Database connection failed"},
	}

	fmt.Println("HTTP Errors:")
	for _, err := range httpErrors {
		if httpErr, ok := err.(*HTTPError); ok {
			fmt.Printf("  %v\n", httpErr)
			fmt.Printf("    Category: %s\n", httpErr.Category())
			fmt.Printf("    Retryable: %t\n", httpErr.IsRetryable())
		}
	}

	// Database errors
	fmt.Println("\nDatabase Errors:")
	dbErrors := []error{
		&DatabaseError{
			Operation: "INSERT",
			Table:     "users",
			Message:   "Duplicate key violation",
			Code:      "23505",
			Query:     "INSERT INTO users (email) VALUES ('user@example.com')",
		},
		&DatabaseError{
			Operation: "SELECT",
			Table:     "orders",
			Message:   "Connection timeout",
			Code:      "08006",
			Query:     "SELECT * FROM orders WHERE user_id = 123",
		},
	}

	for _, err := range dbErrors {
		if dbErr, ok := err.(*DatabaseError); ok {
			fmt.Printf("  %v\n", dbErr)
			fmt.Printf("    Is Constraint Error: %t\n", dbErr.IsConstraintError())
			fmt.Printf("    Is Connection Error: %t\n", dbErr.IsConnectionError())
		}
	}

	// Business logic errors
	fmt.Println("\nBusiness Logic Errors:")
	businessErrors := []error{
		&BusinessError{
			Domain:    "Payment",
			Code:      "INSUFFICIENT_FUNDS",
			Message:   "Account balance too low",
			UserID:    "user123",
			Details:   map[string]interface{}{"balance": 50.0, "required": 100.0},
			Timestamp: time.Now(),
		},
		&BusinessError{
			Domain:    "Inventory",
			Code:      "OUT_OF_STOCK",
			Message:   "Product not available",
			Details:   map[string]interface{}{"product_id": "prod456", "requested": 5, "available": 0},
			Timestamp: time.Now(),
		},
	}

	for _, err := range businessErrors {
		if bizErr, ok := err.(*BusinessError); ok {
			fmt.Printf("  %v\n", bizErr)
			fmt.Printf("    Domain: %s\n", bizErr.Domain)
			fmt.Printf("    Code: %s\n", bizErr.Code)
			if bizErr.UserID != "" {
				fmt.Printf("    User: %s\n", bizErr.UserID)
			}
		}
	}
}

// HTTPError represents HTTP-related errors
type HTTPError struct {
	StatusCode int
	Message    string
	Details    string
	Headers    map[string]string
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("HTTP %d: %s", e.StatusCode, e.Message)
}

// Category returns the error category based on status code
func (e *HTTPError) Category() string {
	switch {
	case e.StatusCode >= 400 && e.StatusCode < 500:
		return "Client Error"
	case e.StatusCode >= 500:
		return "Server Error"
	default:
		return "Unknown"
	}
}

// IsRetryable indicates if the error is retryable
func (e *HTTPError) IsRetryable() bool {
	// Generally, 5xx errors are retryable, 4xx are not
	return e.StatusCode >= 500
}

// IsClientError checks if it's a client error (4xx)
func (e *HTTPError) IsClientError() bool {
	return e.StatusCode >= 400 && e.StatusCode < 500
}

// IsServerError checks if it's a server error (5xx)
func (e *HTTPError) IsServerError() bool {
	return e.StatusCode >= 500
}

// DatabaseError represents database-related errors
type DatabaseError struct {
	Operation string // SELECT, INSERT, UPDATE, DELETE
	Table     string
	Message   string
	Code      string // Database error code
	Query     string
	Timestamp time.Time
}

func (e *DatabaseError) Error() string {
	return fmt.Sprintf("Database error in %s on table '%s': %s (code: %s)", 
		e.Operation, e.Table, e.Message, e.Code)
}

// IsConstraintError checks if it's a constraint violation
func (e *DatabaseError) IsConstraintError() bool {
	// PostgreSQL constraint violation codes
	return e.Code == "23505" || e.Code == "23503" || e.Code == "23502"
}

// IsConnectionError checks if it's a connection-related error
func (e *DatabaseError) IsConnectionError() bool {
	// PostgreSQL connection error codes
	return e.Code == "08006" || e.Code == "08003" || e.Code == "08000"
}

// BusinessError represents domain-specific business logic errors
type BusinessError struct {
	Domain    string                 // Business domain (Payment, Inventory, etc.)
	Code      string                 // Business error code
	Message   string                 // Human-readable message
	UserID    string                 // Associated user (if applicable)
	Details   map[string]interface{} // Additional context
	Timestamp time.Time
}

func (e *BusinessError) Error() string {
	return fmt.Sprintf("[%s] %s: %s", e.Domain, e.Code, e.Message)
}

// DemoErrorCollections demonstrates handling multiple errors
func DemoErrorCollections() {
	fmt.Println("\n=== Error Collections Demo ===")

	// Validate a form with multiple fields
	form := FormData{
		Name:     "",
		Email:    "invalid-email",
		Age:      -5,
		Password: "123",
		Website:  "not-a-url",
	}

	if errs := ValidateForm(form); errs.HasErrors() {
		fmt.Printf("Form validation failed with %d errors:\n", errs.Count())
		for field, fieldErrors := range errs.Errors {
			fmt.Printf("  %s:\n", field)
			for _, err := range fieldErrors {
				fmt.Printf("    - %s\n", err)
			}
		}

		// Get all errors as a single error
		fmt.Printf("\nCombined error: %v\n", errs)
	}

	// Batch processing with error collection
	fmt.Println("\nBatch processing with errors:")
	items := []string{"valid1", "", "valid2", "invalid!", "valid3"}
	results, batchErrors := ProcessBatchWithErrors(items)

	fmt.Printf("Successful results: %v\n", results)
	if batchErrors.HasErrors() {
		fmt.Printf("Batch errors:\n%v\n", batchErrors)
	}

	// Service errors collection
	fmt.Println("\nService initialization:")
	service := &ComplexService{}
	if err := service.Initialize(); err != nil {
		fmt.Printf("Service initialization failed:\n%v\n", err)
	}
}

// ErrorCollection manages multiple related errors
type ErrorCollection struct {
	Errors map[string][]string
	Context string
}

// NewErrorCollection creates a new error collection
func NewErrorCollection(context string) *ErrorCollection {
	return &ErrorCollection{
		Errors:  make(map[string][]string),
		Context: context,
	}
}

// Add adds an error for a specific field or category
func (ec *ErrorCollection) Add(field, message string) {
	ec.Errors[field] = append(ec.Errors[field], message)
}

// AddError adds an error object for a specific field
func (ec *ErrorCollection) AddError(field string, err error) {
	ec.Add(field, err.Error())
}

// HasErrors returns true if there are any errors
func (ec *ErrorCollection) HasErrors() bool {
	return len(ec.Errors) > 0
}

// Count returns the total number of errors
func (ec *ErrorCollection) Count() int {
	count := 0
	for _, fieldErrors := range ec.Errors {
		count += len(fieldErrors)
	}
	return count
}

// Error implements the error interface
func (ec *ErrorCollection) Error() string {
	if !ec.HasErrors() {
		return ""
	}

	result := fmt.Sprintf("%s failed with %d error(s):", ec.Context, ec.Count())
	for field, fieldErrors := range ec.Errors {
		for _, err := range fieldErrors {
			result += fmt.Sprintf("\n  %s: %s", field, err)
		}
	}
	return result
}

// FormData represents a form with validation
type FormData struct {
	Name     string
	Email    string
	Age      int
	Password string
	Website  string
}

// ValidateForm validates form data and returns all errors
func ValidateForm(form FormData) *ErrorCollection {
	errs := NewErrorCollection("Form validation")

	// Validate name
	if form.Name == "" {
		errs.Add("name", "Name is required")
	} else if len(form.Name) < 2 {
		errs.Add("name", "Name must be at least 2 characters")
	}

	// Validate email
	if form.Email == "" {
		errs.Add("email", "Email is required")
	} else if !isValidEmail(form.Email) {
		errs.Add("email", "Email format is invalid")
	}

	// Validate age
	if form.Age < 0 {
		errs.Add("age", "Age must be non-negative")
	} else if form.Age > 150 {
		errs.Add("age", "Age must be realistic")
	}

	// Validate password
	if form.Password == "" {
		errs.Add("password", "Password is required")
	} else {
		if len(form.Password) < 8 {
			errs.Add("password", "Password must be at least 8 characters")
		}
		if !hasDigit(form.Password) {
			errs.Add("password", "Password must contain at least one digit")
		}
		if !hasUpperCase(form.Password) {
			errs.Add("password", "Password must contain at least one uppercase letter")
		}
	}

	// Validate website
	if form.Website != "" && !isValidURL(form.Website) {
		errs.Add("website", "Website URL is invalid")
	}

	return errs
}

// Helper validation functions (simplified implementations)
func isValidEmail(email string) bool {
	return len(email) > 0 && len(email) < 100 && 
		   email[0] != '@' && email[len(email)-1] != '@' &&
		   countRune(email, '@') == 1
}

func hasDigit(s string) bool {
	for _, r := range s {
		if r >= '0' && r <= '9' {
			return true
		}
	}
	return false
}

func hasUpperCase(s string) bool {
	for _, r := range s {
		if r >= 'A' && r <= 'Z' {
			return true
		}
	}
	return false
}

func isValidURL(url string) bool {
	return len(url) > 7 && (url[:7] == "http://" || url[:8] == "https://")
}

func countRune(s string, r rune) int {
	count := 0
	for _, char := range s {
		if char == r {
			count++
		}
	}
	return count
}

// ProcessBatchWithErrors processes items and collects errors
func ProcessBatchWithErrors(items []string) ([]string, *ErrorCollection) {
	var results []string
	errs := NewErrorCollection("Batch processing")

	for i, item := range items {
		if err := validateItem(item); err != nil {
			errs.AddError(fmt.Sprintf("item_%d", i), err)
		} else {
			results = append(results, fmt.Sprintf("processed_%s", item))
		}
	}

	return results, errs
}

// validateItem validates a single item
func validateItem(item string) error {
	if item == "" {
		return fmt.Errorf("item cannot be empty")
	}
	if len(item) > 10 {
		return fmt.Errorf("item too long (max 10 characters)")
	}
	for _, r := range item {
		if r == '!' {
			return fmt.Errorf("item cannot contain '!' character")
		}
	}
	return nil
}

// ComplexService demonstrates service-level error collection
type ComplexService struct {
	Database *DatabaseService
	Cache    *CacheService
	Logger   *LoggerService
}

// Initialize initializes all service components
func (s *ComplexService) Initialize() error {
	errs := NewErrorCollection("Service initialization")

	// Initialize database
	s.Database = &DatabaseService{}
	if err := s.Database.Connect(); err != nil {
		errs.AddError("database", err)
	}

	// Initialize cache
	s.Cache = &CacheService{}
	if err := s.Cache.Connect(); err != nil {
		errs.AddError("cache", err)
	}

	// Initialize logger
	s.Logger = &LoggerService{}
	if err := s.Logger.Setup(); err != nil {
		errs.AddError("logger", err)
	}

	if errs.HasErrors() {
		return errs
	}
	return nil
}

// Mock services for demonstration
type DatabaseService struct{}
func (ds *DatabaseService) Connect() error {
	return &DatabaseError{Operation: "CONNECT", Message: "Connection refused", Code: "08001"}
}

type CacheService struct{}
func (cs *CacheService) Connect() error {
	return nil // Success
}

type LoggerService struct{}
func (ls *LoggerService) Setup() error {
	return fmt.Errorf("failed to create log directory")
}

// DemoStructuredErrors demonstrates structured error handling
func DemoStructuredErrors() {
	fmt.Println("\n=== Structured Errors Demo ===")

	// API errors with structured information
	apiErrors := []error{
		NewAPIError(400, "INVALID_REQUEST", "Missing required field", map[string]interface{}{
			"field": "email",
			"received": nil,
		}),
		NewAPIError(409, "CONFLICT", "Resource already exists", map[string]interface{}{
			"resource_type": "user",
			"resource_id": "user123",
		}),
		NewAPIError(429, "RATE_LIMITED", "Too many requests", map[string]interface{}{
			"limit": 100,
			"window": "1h",
			"retry_after": 3600,
		}),
	}

	for _, err := range apiErrors {
		if apiErr, ok := err.(*APIError); ok {
			fmt.Printf("API Error: %v\n", apiErr)
			fmt.Printf("  Status: %d\n", apiErr.StatusCode)
			fmt.Printf("  Code: %s\n", apiErr.Code)
			fmt.Printf("  Details: %v\n", apiErr.Details)
			fmt.Printf("  Timestamp: %s\n", apiErr.Timestamp.Format(time.RFC3339))
		}
	}

	// Operational errors
	fmt.Println("\nOperational Errors:")
	operationalErrors := []error{
		NewOperationalError("network", "CONNECTION_TIMEOUT", 
			"Failed to connect to external service", map[string]interface{}{
			"service": "payment-gateway",
			"timeout_ms": 5000,
			"retry_count": 3,
		}),
		NewOperationalError("disk", "STORAGE_FULL", 
			"Insufficient disk space", map[string]interface{}{
			"available_mb": 50,
			"required_mb": 100,
			"partition": "/var/log",
		}),
	}

	for _, err := range operationalErrors {
		if opErr, ok := err.(*OperationalError); ok {
			fmt.Printf("Operational Error: %v\n", opErr)
			fmt.Printf("  Category: %s\n", opErr.Category)
			fmt.Printf("  Severity: %s\n", opErr.Severity)
			fmt.Printf("  Details: %v\n", opErr.Details)
		}
	}
}

// APIError represents a structured API error
type APIError struct {
	StatusCode int
	Code       string
	Message    string
	Details    map[string]interface{}
	Timestamp  time.Time
}

// NewAPIError creates a new API error
func NewAPIError(statusCode int, code, message string, details map[string]interface{}) *APIError {
	return &APIError{
		StatusCode: statusCode,
		Code:       code,
		Message:    message,
		Details:    details,
		Timestamp:  time.Now(),
	}
}

func (e *APIError) Error() string {
	return fmt.Sprintf("API Error %d (%s): %s", e.StatusCode, e.Code, e.Message)
}

// OperationalError represents operational/infrastructure errors
type OperationalError struct {
	Category  string // network, disk, memory, cpu, etc.
	Code      string
	Message   string
	Severity  string // low, medium, high, critical
	Details   map[string]interface{}
	Timestamp time.Time
}

// NewOperationalError creates a new operational error
func NewOperationalError(category, code, message string, details map[string]interface{}) *OperationalError {
	severity := "medium" // default
	if category == "disk" || category == "memory" {
		severity = "high"
	}
	
	return &OperationalError{
		Category:  category,
		Code:      code,
		Message:   message,
		Severity:  severity,
		Details:   details,
		Timestamp: time.Now(),
	}
}

func (e *OperationalError) Error() string {
	return fmt.Sprintf("[%s] %s: %s", e.Category, e.Code, e.Message)
}

// DemoErrorContext demonstrates adding context to errors
func DemoErrorContext() {
	fmt.Println("\n=== Error Context Demo ===")

	// Simulate a complex operation with context
	userID := "user123"
	orderID := "order456"
	
	if err := processOrder(userID, orderID); err != nil {
		fmt.Printf("Order processing failed: %v\n", err)
		
		// Extract context from error
		if ctxErr, ok := err.(*ContextualError); ok {
			fmt.Printf("Error context:\n")
			for key, value := range ctxErr.Context {
				fmt.Printf("  %s: %v\n", key, value)
			}
		}
	}

	// Error with operation trace
	fmt.Println("\nOperation trace:")
	if err := complexOperation(); err != nil {
		fmt.Printf("Complex operation failed:\n%v\n", err)
	}
}

// ContextualError provides rich context about where and why an error occurred
type ContextualError struct {
	Operation string
	Message   string
	Context   map[string]interface{}
	Cause     error
	Timestamp time.Time
}

// NewContextualError creates a new contextual error
func NewContextualError(operation, message string, context map[string]interface{}) *ContextualError {
	return &ContextualError{
		Operation: operation,
		Message:   message,
		Context:   context,
		Timestamp: time.Now(),
	}
}

// WithCause adds a cause to the contextual error
func (e *ContextualError) WithCause(cause error) *ContextualError {
	e.Cause = cause
	return e
}

func (e *ContextualError) Error() string {
	result := fmt.Sprintf("%s failed: %s", e.Operation, e.Message)
	if e.Cause != nil {
		result += fmt.Sprintf(" (caused by: %v)", e.Cause)
	}
	return result
}

// Unwrap returns the underlying cause
func (e *ContextualError) Unwrap() error {
	return e.Cause
}

// processOrder simulates order processing with contextual errors
func processOrder(userID, orderID string) error {
	ctx := map[string]interface{}{
		"user_id":    userID,
		"order_id":   orderID,
		"service":    "order-processor",
		"version":    "1.2.3",
	}

	// Simulate validation failure
	if userID == "user123" {
		return NewContextualError("order_validation", "Invalid user account", ctx).
			WithCause(fmt.Errorf("user account suspended"))
	}

	return nil
}

// complexOperation demonstrates operation tracing
func complexOperation() error {
	trace := NewOperationTrace("complex_operation")
	
	// Step 1
	trace.AddStep("step1", "Validate input")
	if err := validateInput(); err != nil {
		return trace.WithError("step1", err)
	}
	
	// Step 2
	trace.AddStep("step2", "Process data")
	if err := processData(); err != nil {
		return trace.WithError("step2", err)
	}
	
	// Step 3
	trace.AddStep("step3", "Save results")
	if err := saveResults(); err != nil {
		return trace.WithError("step3", err)
	}
	
	return nil
}

// OperationTrace tracks the steps of a complex operation
type OperationTrace struct {
	Operation string
	Steps     []TraceStep
	StartTime time.Time
}

// TraceStep represents a single step in an operation
type TraceStep struct {
	Name        string
	Description string
	StartTime   time.Time
	EndTime     time.Time
	Error       error
}

// NewOperationTrace creates a new operation trace
func NewOperationTrace(operation string) *OperationTrace {
	return &OperationTrace{
		Operation: operation,
		StartTime: time.Now(),
	}
}

// AddStep adds a new step to the trace
func (ot *OperationTrace) AddStep(name, description string) {
	step := TraceStep{
		Name:        name,
		Description: description,
		StartTime:   time.Now(),
	}
	ot.Steps = append(ot.Steps, step)
}

// WithError marks the current step as failed and returns a traced error
func (ot *OperationTrace) WithError(stepName string, err error) error {
	// Find the step and mark it as failed
	for i := range ot.Steps {
		if ot.Steps[i].Name == stepName {
			ot.Steps[i].EndTime = time.Now()
			ot.Steps[i].Error = err
			break
		}
	}
	
	return &TracedError{
		Operation: ot.Operation,
		Steps:     ot.Steps,
		FailedAt:  stepName,
		Cause:     err,
	}
}

// TracedError contains the full trace of a failed operation
type TracedError struct {
	Operation string
	Steps     []TraceStep
	FailedAt  string
	Cause     error
}

func (e *TracedError) Error() string {
	result := fmt.Sprintf("Operation '%s' failed at step '%s': %v\n", 
		e.Operation, e.FailedAt, e.Cause)
	
	result += "Execution trace:\n"
	for _, step := range e.Steps {
		status := "✓"
		if step.Error != nil {
			status = "✗"
		}
		result += fmt.Sprintf("  %s %s: %s\n", status, step.Name, step.Description)
	}
	
	return result
}

// Unwrap returns the underlying cause
func (e *TracedError) Unwrap() error {
	return e.Cause
}

// Mock functions for demonstration
func validateInput() error {
	return nil
}

func processData() error {
	return fmt.Errorf("data processing failed: invalid format")
}

func saveResults() error {
	return nil
}