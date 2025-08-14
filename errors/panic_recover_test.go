package errors

import (
	"strings"
	"testing"
)

func TestDemoPanicBasics(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("DemoPanicBasics panicked: %v", r)
		}
	}()
	
	DemoPanicBasics()
}

func TestDemoGracefulRecovery(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("DemoGracefulRecovery panicked: %v", r)
		}
	}()
	
	DemoGracefulRecovery()
}

func TestDemoStackTrace(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("DemoStackTrace panicked: %v", r)
		}
	}()
	
	DemoStackTrace()
}

func TestDemoPanicInGoroutines(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("DemoPanicInGoroutines panicked: %v", r)
		}
	}()
	
	DemoPanicInGoroutines()
}

func TestDemoResourceCleanup(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("DemoResourceCleanup panicked: %v", r)
		}
	}()
	
	DemoResourceCleanup()
}

func TestDemoPanicBestPractices(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("DemoPanicBestPractices panicked: %v", r)
		}
	}()
	
	DemoPanicBestPractices()
}

func TestDemoMustPattern(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("DemoMustPattern panicked: %v", r)
		}
	}()
	
	DemoMustPattern()
}

func TestSafeFunction(t *testing.T) {
	// Test normal operation
	result := SafeFunction(10, 2)
	if result != 5 {
		t.Errorf("Expected result 5, got %d", result)
	}
	
	// Test panic recovery - function should return 0 due to panic
	result = SafeFunction(10, 0)
	if result != 0 {
		t.Errorf("Expected result 0 after panic recovery, got %d", result)
	}
}

func TestProcessNumbers(t *testing.T) {
	testCases := []struct {
		a, b     int
		expected string
	}{
		{10, 2, "success: 5"},
		{15, 3, "success: 5"},
		{20, 4, "success: 5"},
	}
	
	for _, tc := range testCases {
		result := ProcessNumbers(tc.a, tc.b)
		if result != tc.expected {
			t.Errorf("ProcessNumbers(%d, %d): expected %s, got %s", 
				tc.a, tc.b, tc.expected, result)
		}
	}
	
	// Test panic cases - these should be handled gracefully
	panicCases := []struct {
		a, b int
	}{
		{10, 0},  // Division by zero
		{-5, 2},  // Negative input
	}
	
	for _, tc := range panicCases {
		result := ProcessNumbers(tc.a, tc.b)
		// Should not be a success message
		if strings.Contains(result, "success") {
			t.Errorf("ProcessNumbers(%d, %d): expected error, got success: %s", 
				tc.a, tc.b, result)
		}
	}
}

func TestProcessBatchOperations(t *testing.T) {
	operations := []func() string{
		func() string { return "op1: success" },
		func() string { panic("op2 failed") },
		func() string { return "op3: success" },
		func() string { panic("op4 failed") },
		func() string { return "op5: success" },
	}
	
	results := ProcessBatchOperations(operations)
	
	if len(results) != 5 {
		t.Errorf("Expected 5 results, got %d", len(results))
	}
	
	// Check successful operations
	if results[0] != "op1: success" {
		t.Errorf("Operation 1 should succeed")
	}
	
	if results[2] != "op3: success" {
		t.Errorf("Operation 3 should succeed")
	}
	
	if results[4] != "op5: success" {
		t.Errorf("Operation 5 should succeed")
	}
	
	// Check failed operations
	if !strings.Contains(results[1], "ERROR") {
		t.Errorf("Operation 2 should have error, got: %s", results[1])
	}
	
	if !strings.Contains(results[3], "ERROR") {
		t.Errorf("Operation 4 should have error, got: %s", results[3])
	}
}

func TestHandleWebRequest(t *testing.T) {
	testCases := []struct {
		request  string
		expected string
	}{
		{"valid", "200 OK"},
		{"success", "200 OK: Data processed"},
		{"unknown", "404 Not Found"},
	}
	
	for _, tc := range testCases {
		result := HandleWebRequest(tc.request)
		if result != tc.expected {
			t.Errorf("HandleWebRequest(%s): expected %s, got %s", 
				tc.request, tc.expected, result)
		}
	}
	
	// Test panic cases - these should be handled by recover
	panicCases := []string{"invalid", "timeout"}
	for _, req := range panicCases {
		result := HandleWebRequest(req)
		// Function should not panic, should return error response
		if result != "500 Internal Server Error" {
			t.Errorf("HandleWebRequest(%s): expected '500 Internal Server Error', got '%s'", req, result)
		}
	}
}

func TestCallChainWithPanic(t *testing.T) {
	// This should not panic because it's recovered
	result := CallChainWithPanic()
	
	// Since panic is recovered, result should be empty string (zero value)
	if result != "" {
		t.Errorf("Expected empty result after panic recovery, got: %s", result)
	}
}

func TestMockFile(t *testing.T) {
	file := &MockFile{Name: "test.txt"}
	
	// Test initial state
	if file.isOpen {
		t.Errorf("File should be closed initially")
	}
	
	// Test open
	file.Open()
	if !file.isOpen {
		t.Errorf("File should be open after Open()")
	}
	
	// Test close
	file.Close()
	if file.isOpen {
		t.Errorf("File should be closed after Close()")
	}
	
	// Test multiple closes
	file.Close() // Should not cause issues
}

func TestMockDBConnection(t *testing.T) {
	conn := &MockDBConnection{}
	
	// Test initial state
	if conn.isConnected {
		t.Errorf("Connection should be disconnected initially")
	}
	
	// Test connect
	conn.Connect()
	if !conn.isConnected {
		t.Errorf("Connection should be connected after Connect()")
	}
	
	// Test close
	conn.Close()
	if conn.isConnected {
		t.Errorf("Connection should be disconnected after Close()")
	}
	
	// Test multiple closes
	conn.Close() // Should not cause issues
}

func TestProcessFile(t *testing.T) {
	// Test successful file processing
	ProcessFile("valid_file.txt") // Should not panic
	
	// Test error cases - these should be handled gracefully
	ProcessFile("corrupted_file.txt") // Should recover from panic
	ProcessFile("locked_file.txt")    // Should recover from panic
}

func TestProcessDatabase(t *testing.T) {
	// Test successful query
	ProcessDatabase("SELECT * FROM users") // Should not panic
	
	// Test error cases - these should be handled gracefully
	ProcessDatabase("INVALID SQL QUERY")      // Should recover from panic
	ProcessDatabase("SELECT * FROM missing_table") // Should recover from panic
}

func TestPublicAPIFunction(t *testing.T) {
	// Test normal case
	result := PublicAPIFunction("valid_input")
	expected := "processed: valid_input"
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
	
	// Test panic case - should be recovered
	result = PublicAPIFunction("panic_input")
	// Should not panic, should return empty string due to recovery
	if result != "" {
		t.Errorf("Expected empty result after panic recovery, got: %s", result)
	}
}

func TestSafeSliceAccess(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	
	// Test valid access
	value, err := SafeSliceAccess(slice, 2)
	if err != nil {
		t.Errorf("Expected no error for valid index, got: %v", err)
	}
	if value != 3 {
		t.Errorf("Expected value 3, got %d", value)
	}
	
	// Test invalid access
	_, err = SafeSliceAccess(slice, 10)
	if err == nil {
		t.Errorf("Expected error for invalid index")
	}
	
	_, err = SafeSliceAccess(slice, -1)
	if err == nil {
		t.Errorf("Expected error for negative index")
	}
}

func TestSafeStringToInt(t *testing.T) {
	// Test valid conversion
	value, err := SafeStringToInt("123")
	if err != nil {
		t.Errorf("Expected no error for valid string, got: %v", err)
	}
	if value != 123 {
		t.Errorf("Expected value 123, got %d", value)
	}
	
	// Test invalid conversion
	_, err = SafeStringToInt("abc")
	if err == nil {
		t.Errorf("Expected error for invalid string")
	}
}

func TestSafeMapAccess(t *testing.T) {
	m := map[string]interface{}{
		"key1": "value1",
		"key2": 42,
	}
	
	// Test existing key
	value, exists := SafeMapAccess(m, "key1")
	if !exists {
		t.Errorf("Key 'key1' should exist")
	}
	if value != "value1" {
		t.Errorf("Expected 'value1', got %v", value)
	}
	
	// Test non-existing key
	_, exists = SafeMapAccess(m, "nonexistent")
	if exists {
		t.Errorf("Key 'nonexistent' should not exist")
	}
}

func TestMustOperation(t *testing.T) {
	// Test successful operation
	result := MustOperation("parse_config")
	if result != "config_parsed" {
		t.Errorf("Expected 'config_parsed', got %s", result)
	}
	
	// Test default operation
	result = MustOperation("other_operation")
	expected := "operation_other_operation_completed"
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
	
	// Test panic operation
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic for invalid_operation")
		}
	}()
	
	MustOperation("invalid_operation") // Should panic
}

func TestSafeChannelSend(t *testing.T) {
	// Test with buffered channel
	ch := make(chan string, 2)
	
	// Should succeed
	sent := SafeChannelSend(ch, "value1")
	if !sent {
		t.Errorf("Should be able to send to buffered channel")
	}
	
	sent = SafeChannelSend(ch, "value2")
	if !sent {
		t.Errorf("Should be able to send second value to buffered channel")
	}
	
	// Should fail (channel full)
	sent = SafeChannelSend(ch, "value3")
	if sent {
		t.Errorf("Should not be able to send to full channel")
	}
	
	// Verify values were sent
	select {
	case value := <-ch:
		if value != "value1" {
			t.Errorf("Expected 'value1', got %s", value)
		}
	default:
		t.Errorf("Channel should contain value1")
	}
}

func TestPanicRecoveryBehavior(t *testing.T) {
	// Test that defer with recover works as expected
	var recovered bool
	var panicValue interface{}
	
	func() {
		defer func() {
			if r := recover(); r != nil {
				recovered = true
				panicValue = r
			}
		}()
		
		panic("test panic")
	}()
	
	if !recovered {
		t.Errorf("Should have recovered from panic")
	}
	
	if panicValue != "test panic" {
		t.Errorf("Expected panic value 'test panic', got %v", panicValue)
	}
}

func TestResourceCleanupWithPanic(t *testing.T) {
	var cleanupCalled bool
	
	func() {
		defer func() {
			cleanupCalled = true
			recover() // Ignore panic for test
		}()
		
		panic("resource error")
	}()
	
	if !cleanupCalled {
		t.Errorf("Cleanup should be called even when panic occurs")
	}
}

func TestNestedPanicRecovery(t *testing.T) {
	var outerRecovered, innerRecovered bool
	
	defer func() {
		if r := recover(); r != nil {
			outerRecovered = true
		}
	}()
	
	func() {
		defer func() {
			if r := recover(); r != nil {
				innerRecovered = true
				// Don't re-panic, handle it here
			}
		}()
		
		panic("inner panic")
	}()
	
	if !innerRecovered {
		t.Errorf("Inner function should have recovered")
	}
	
	if outerRecovered {
		t.Errorf("Outer function should not have seen panic after inner recovery")
	}
}

func TestValidateConfigFile(t *testing.T) {
	// Should not panic for valid config
	ValidateConfigFile("valid_config.json")
	
	// Should handle panic for invalid config
	ValidateConfigFile("invalid_config.json")
}

func TestDemonstrateSelectiveRecovery(t *testing.T) {
	// This should not panic since it handles the demonstration
	DemonstrateSelectiveRecovery()
}

func TestDemonstrateMeaningfulPanics(t *testing.T) {
	// This should handle its own panic
	DemonstrateMeaningfulPanics()
}