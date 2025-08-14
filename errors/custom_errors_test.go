package errors

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestDemoCustomErrorTypes(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("DemoCustomErrorTypes panicked: %v", r)
		}
	}()
	
	DemoCustomErrorTypes()
}

func TestDemoErrorCollections(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("DemoErrorCollections panicked: %v", r)
		}
	}()
	
	DemoErrorCollections()
}

func TestDemoStructuredErrors(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("DemoStructuredErrors panicked: %v", r)
		}
	}()
	
	DemoStructuredErrors()
}

func TestDemoErrorContext(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("DemoErrorContext panicked: %v", r)
		}
	}()
	
	DemoErrorContext()
}

func TestHTTPError(t *testing.T) {
	// Test client error
	clientErr := &HTTPError{
		StatusCode: 400,
		Message:    "Bad Request",
		Details:    "Invalid JSON",
	}
	
	if !clientErr.IsClientError() {
		t.Errorf("400 should be a client error")
	}
	
	if clientErr.IsServerError() {
		t.Errorf("400 should not be a server error")
	}
	
	if clientErr.IsRetryable() {
		t.Errorf("400 should not be retryable")
	}
	
	if clientErr.Category() != "Client Error" {
		t.Errorf("Expected 'Client Error', got '%s'", clientErr.Category())
	}
	
	// Test server error
	serverErr := &HTTPError{
		StatusCode: 500,
		Message:    "Internal Server Error",
	}
	
	if !serverErr.IsServerError() {
		t.Errorf("500 should be a server error")
	}
	
	if serverErr.IsClientError() {
		t.Errorf("500 should not be a client error")
	}
	
	if !serverErr.IsRetryable() {
		t.Errorf("500 should be retryable")
	}
	
	if serverErr.Category() != "Server Error" {
		t.Errorf("Expected 'Server Error', got '%s'", serverErr.Category())
	}
	
	// Test error message
	expectedMsg := "HTTP 500: Internal Server Error"
	if serverErr.Error() != expectedMsg {
		t.Errorf("Expected '%s', got '%s'", expectedMsg, serverErr.Error())
	}
}

func TestDatabaseError(t *testing.T) {
	// Test constraint error
	constraintErr := &DatabaseError{
		Operation: "INSERT",
		Table:     "users",
		Message:   "Duplicate key",
		Code:      "23505",
	}
	
	if !constraintErr.IsConstraintError() {
		t.Errorf("Code 23505 should be a constraint error")
	}
	
	if constraintErr.IsConnectionError() {
		t.Errorf("Code 23505 should not be a connection error")
	}
	
	// Test connection error
	connErr := &DatabaseError{
		Operation: "CONNECT",
		Table:     "",
		Message:   "Connection refused",
		Code:      "08006",
	}
	
	if !connErr.IsConnectionError() {
		t.Errorf("Code 08006 should be a connection error")
	}
	
	if connErr.IsConstraintError() {
		t.Errorf("Code 08006 should not be a constraint error")
	}
	
	// Test error message
	expectedMsg := "Database error in INSERT on table 'users': Duplicate key (code: 23505)"
	if constraintErr.Error() != expectedMsg {
		t.Errorf("Expected '%s', got '%s'", expectedMsg, constraintErr.Error())
	}
}

func TestBusinessError(t *testing.T) {
	bizErr := &BusinessError{
		Domain:  "Payment",
		Code:    "INSUFFICIENT_FUNDS",
		Message: "Account balance too low",
		UserID:  "user123",
		Details: map[string]interface{}{
			"balance":  50.0,
			"required": 100.0,
		},
		Timestamp: time.Now(),
	}
	
	expectedMsg := "[Payment] INSUFFICIENT_FUNDS: Account balance too low"
	if bizErr.Error() != expectedMsg {
		t.Errorf("Expected '%s', got '%s'", expectedMsg, bizErr.Error())
	}
	
	if bizErr.Domain != "Payment" {
		t.Errorf("Expected domain 'Payment', got '%s'", bizErr.Domain)
	}
	
	if bizErr.UserID != "user123" {
		t.Errorf("Expected user ID 'user123', got '%s'", bizErr.UserID)
	}
}

func TestErrorCollection(t *testing.T) {
	errs := NewErrorCollection("Test validation")
	
	// Test empty collection
	if errs.HasErrors() {
		t.Errorf("New collection should not have errors")
	}
	
	if errs.Count() != 0 {
		t.Errorf("Empty collection should have count 0")
	}
	
	// Add errors
	errs.Add("field1", "error 1")
	errs.Add("field1", "error 2")
	errs.Add("field2", "error 3")
	
	if !errs.HasErrors() {
		t.Errorf("Collection should have errors after adding")
	}
	
	if errs.Count() != 3 {
		t.Errorf("Expected count 3, got %d", errs.Count())
	}
	
	// Test field-specific errors
	if len(errs.Errors["field1"]) != 2 {
		t.Errorf("field1 should have 2 errors")
	}
	
	if len(errs.Errors["field2"]) != 1 {
		t.Errorf("field2 should have 1 error")
	}
	
	// Test error message
	errMsg := errs.Error()
	if !strings.Contains(errMsg, "Test validation failed") {
		t.Errorf("Error message should contain context")
	}
	
	if !strings.Contains(errMsg, "3 error(s)") {
		t.Errorf("Error message should contain error count")
	}
}

func TestValidateForm(t *testing.T) {
	// Test valid form
	validForm := FormData{
		Name:     "Alice Johnson",
		Email:    "alice@example.com",
		Age:      25,
		Password: "SecurePass123",
		Website:  "https://example.com",
	}
	
	errs := ValidateForm(validForm)
	if errs.HasErrors() {
		t.Errorf("Valid form should not have errors, got: %v", errs)
	}
	
	// Test invalid form
	invalidForm := FormData{
		Name:     "",
		Email:    "invalid-email",
		Age:      -5,
		Password: "123",
		Website:  "not-a-url",
	}
	
	errs = ValidateForm(invalidForm)
	if !errs.HasErrors() {
		t.Errorf("Invalid form should have errors")
	}
	
	// Check specific field errors
	if len(errs.Errors["name"]) == 0 {
		t.Errorf("Should have name validation error")
	}
	
	if len(errs.Errors["email"]) == 0 {
		t.Errorf("Should have email validation error")
	}
	
	if len(errs.Errors["age"]) == 0 {
		t.Errorf("Should have age validation error")
	}
	
	if len(errs.Errors["password"]) == 0 {
		t.Errorf("Should have password validation errors")
	}
	
	if len(errs.Errors["website"]) == 0 {
		t.Errorf("Should have website validation error")
	}
}

func TestValidationHelpers(t *testing.T) {
	// Test isValidEmail
	validEmails := []string{"user@example.com", "test.email@domain.org"}
	for _, email := range validEmails {
		if !isValidEmail(email) {
			t.Errorf("Email '%s' should be valid", email)
		}
	}
	
	invalidEmails := []string{"", "@example.com", "user@", "user@@example.com", "user"}
	for _, email := range invalidEmails {
		if isValidEmail(email) {
			t.Errorf("Email '%s' should be invalid", email)
		}
	}
	
	// Test hasDigit
	if !hasDigit("abc123") {
		t.Errorf("'abc123' should have digit")
	}
	
	if hasDigit("abcdef") {
		t.Errorf("'abcdef' should not have digit")
	}
	
	// Test hasUpperCase
	if !hasUpperCase("abcDef") {
		t.Errorf("'abcDef' should have uppercase")
	}
	
	if hasUpperCase("abcdef") {
		t.Errorf("'abcdef' should not have uppercase")
	}
	
	// Test isValidURL
	validURLs := []string{"http://example.com", "https://example.com"}
	for _, url := range validURLs {
		if !isValidURL(url) {
			t.Errorf("URL '%s' should be valid", url)
		}
	}
	
	invalidURLs := []string{"", "example.com", "ftp://example.com"}
	for _, url := range invalidURLs {
		if isValidURL(url) {
			t.Errorf("URL '%s' should be invalid", url)
		}
	}
}

func TestProcessBatchWithErrors(t *testing.T) {
	items := []string{"valid1", "", "valid2", "toolongitem!", "valid3"}
	results, errs := ProcessBatchWithErrors(items)
	
	// Should have 3 valid results
	expectedResults := 3
	if len(results) != expectedResults {
		t.Errorf("Expected %d results, got %d", expectedResults, len(results))
	}
	
	// Should have 2 errors (empty item and item with !)
	expectedErrors := 2
	if errs.Count() != expectedErrors {
		t.Errorf("Expected %d errors, got %d", expectedErrors, errs.Count())
	}
}

func TestValidateItem(t *testing.T) {
	// Test valid items
	validItems := []string{"valid", "item123", "short"}
	for _, item := range validItems {
		if err := validateItem(item); err != nil {
			t.Errorf("Item '%s' should be valid, got error: %v", item, err)
		}
	}
	
	// Test invalid items
	invalidCases := []struct {
		item        string
		expectError string
	}{
		{"", "empty"},
		{"thisitemistoolong", "too long"},
		{"item!", "!"},
	}
	
	for _, tc := range invalidCases {
		err := validateItem(tc.item)
		if err == nil {
			t.Errorf("Item '%s' should be invalid", tc.item)
		} else if !strings.Contains(err.Error(), tc.expectError) {
			t.Errorf("Error for '%s' should contain '%s', got: %v", tc.item, tc.expectError, err)
		}
	}
}

func TestComplexService(t *testing.T) {
	service := &ComplexService{}
	err := service.Initialize()
	
	// Should fail due to database connection error
	if err == nil {
		t.Errorf("Service initialization should fail")
	}
	
	// Should be an ErrorCollection
	if errColl, ok := err.(*ErrorCollection); ok {
		if !errColl.HasErrors() {
			t.Errorf("Should have initialization errors")
		}
		
		// Should have database and logger errors (cache succeeds)
		expectedErrors := 2
		if errColl.Count() != expectedErrors {
			t.Errorf("Expected %d errors, got %d", expectedErrors, errColl.Count())
		}
	} else {
		t.Errorf("Error should be ErrorCollection, got %T", err)
	}
}

func TestAPIError(t *testing.T) {
	details := map[string]interface{}{
		"field": "email",
		"value": nil,
	}
	
	apiErr := NewAPIError(400, "INVALID_REQUEST", "Missing field", details)
	
	if apiErr.StatusCode != 400 {
		t.Errorf("Expected status code 400, got %d", apiErr.StatusCode)
	}
	
	if apiErr.Code != "INVALID_REQUEST" {
		t.Errorf("Expected code 'INVALID_REQUEST', got '%s'", apiErr.Code)
	}
	
	if apiErr.Details["field"] != "email" {
		t.Errorf("Expected field 'email' in details")
	}
	
	expectedMsg := "API Error 400 (INVALID_REQUEST): Missing field"
	if apiErr.Error() != expectedMsg {
		t.Errorf("Expected '%s', got '%s'", expectedMsg, apiErr.Error())
	}
	
	// Test timestamp is set
	if apiErr.Timestamp.IsZero() {
		t.Errorf("Timestamp should be set")
	}
}

func TestOperationalError(t *testing.T) {
	details := map[string]interface{}{
		"service": "payment-gateway",
		"timeout": 5000,
	}
	
	opErr := NewOperationalError("network", "CONNECTION_TIMEOUT", "Service timeout", details)
	
	if opErr.Category != "network" {
		t.Errorf("Expected category 'network', got '%s'", opErr.Category)
	}
	
	if opErr.Code != "CONNECTION_TIMEOUT" {
		t.Errorf("Expected code 'CONNECTION_TIMEOUT', got '%s'", opErr.Code)
	}
	
	if opErr.Severity != "medium" {
		t.Errorf("Expected severity 'medium', got '%s'", opErr.Severity)
	}
	
	// Test high severity for disk errors
	diskErr := NewOperationalError("disk", "STORAGE_FULL", "Disk full", nil)
	if diskErr.Severity != "high" {
		t.Errorf("Disk errors should have high severity, got '%s'", diskErr.Severity)
	}
	
	expectedMsg := "[network] CONNECTION_TIMEOUT: Service timeout"
	if opErr.Error() != expectedMsg {
		t.Errorf("Expected '%s', got '%s'", expectedMsg, opErr.Error())
	}
}

func TestContextualError(t *testing.T) {
	context := map[string]interface{}{
		"user_id":  "user123",
		"order_id": "order456",
	}
	
	ctxErr := NewContextualError("order_processing", "Validation failed", context)
	
	if ctxErr.Operation != "order_processing" {
		t.Errorf("Expected operation 'order_processing', got '%s'", ctxErr.Operation)
	}
	
	if ctxErr.Context["user_id"] != "user123" {
		t.Errorf("Expected user_id 'user123' in context")
	}
	
	// Test with cause
	cause := fmt.Errorf("account suspended")
	ctxErrWithCause := ctxErr.WithCause(cause)
	
	if ctxErrWithCause.Unwrap() != cause {
		t.Errorf("Unwrap should return the cause")
	}
	
	errMsg := ctxErrWithCause.Error()
	if !strings.Contains(errMsg, "order_processing failed") {
		t.Errorf("Error message should contain operation")
	}
	
	if !strings.Contains(errMsg, "caused by") {
		t.Errorf("Error message should contain cause")
	}
}

func TestOperationTrace(t *testing.T) {
	trace := NewOperationTrace("test_operation")
	
	if trace.Operation != "test_operation" {
		t.Errorf("Expected operation 'test_operation', got '%s'", trace.Operation)
	}
	
	// Add steps
	trace.AddStep("step1", "First step")
	trace.AddStep("step2", "Second step")
	
	if len(trace.Steps) != 2 {
		t.Errorf("Expected 2 steps, got %d", len(trace.Steps))
	}
	
	// Test error creation
	stepErr := fmt.Errorf("step failed")
	tracedErr := trace.WithError("step2", stepErr)
	
	if tracedErr == nil {
		t.Errorf("Should return traced error")
	}
	
	if tErr, ok := tracedErr.(*TracedError); ok {
		if tErr.Operation != "test_operation" {
			t.Errorf("Traced error should preserve operation name")
		}
		
		if tErr.FailedAt != "step2" {
			t.Errorf("Should mark correct failed step")
		}
		
		if tErr.Unwrap() != stepErr {
			t.Errorf("Should unwrap to original error")
		}
		
		// Test error message contains trace
		errMsg := tErr.Error()
		if !strings.Contains(errMsg, "step1") {
			t.Errorf("Error message should contain all steps")
		}
		
		if !strings.Contains(errMsg, "step2") {
			t.Errorf("Error message should contain failed step")
		}
	} else {
		t.Errorf("Should return TracedError, got %T", tracedErr)
	}
}

func TestProcessOrder(t *testing.T) {
	// Test error case
	err := processOrder("user123", "order456")
	if err == nil {
		t.Errorf("Should return error for user123")
	}
	
	if ctxErr, ok := err.(*ContextualError); ok {
		if ctxErr.Operation != "order_validation" {
			t.Errorf("Expected operation 'order_validation'")
		}
		
		if ctxErr.Context["user_id"] != "user123" {
			t.Errorf("Should preserve user_id in context")
		}
		
		if ctxErr.Context["order_id"] != "order456" {
			t.Errorf("Should preserve order_id in context")
		}
	} else {
		t.Errorf("Should return ContextualError, got %T", err)
	}
}

func TestComplexOperation(t *testing.T) {
	// This should fail at processData step
	err := complexOperation()
	if err == nil {
		t.Errorf("Complex operation should fail")
	}
	
	if tracedErr, ok := err.(*TracedError); ok {
		if tracedErr.Operation != "complex_operation" {
			t.Errorf("Should preserve operation name")
		}
		
		if tracedErr.FailedAt != "step2" {
			t.Errorf("Should fail at step2 (processData)")
		}
		
		// Should have 3 steps even though it failed at step 2
		if len(tracedErr.Steps) != 3 {
			t.Errorf("Should have all steps in trace, got %d", len(tracedErr.Steps))
		}
	} else {
		t.Errorf("Should return TracedError, got %T", err)
	}
}