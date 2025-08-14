package basics

import (
	"testing"
)

func TestDemoBasicConstants(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("DemoBasicConstants panicked: %v", r)
		}
	}()
	
	DemoBasicConstants()
}

func TestDemoIotaBasics(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("DemoIotaBasics panicked: %v", r)
		}
	}()
	
	DemoIotaBasics()
}

func TestDemoFilePermissions(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("DemoFilePermissions panicked: %v", r)
		}
	}()
	
	DemoFilePermissions()
}

func TestDemoHTTPStatus(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("DemoHTTPStatus panicked: %v", r)
		}
	}()
	
	DemoHTTPStatus()
}

func TestDemoFileSizes(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("DemoFileSizes panicked: %v", r)
		}
	}()
	
	DemoFileSizes()
}

func TestDemoCustomTypes(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("DemoCustomTypes panicked: %v", r)
		}
	}()
	
	DemoCustomTypes()
}

func TestBasicConstants(t *testing.T) {
	if AppName != "GoHello" {
		t.Errorf("Expected AppName to be 'GoHello', got '%s'", AppName)
	}
	
	if AppVersion != "1.0.0" {
		t.Errorf("Expected AppVersion to be '1.0.0', got '%s'", AppVersion)
	}
	
	if MaxRetries != 3 {
		t.Errorf("Expected MaxRetries to be 3, got %d", MaxRetries)
	}
	
	if !IsActive {
		t.Errorf("Expected IsActive to be true")
	}
}

func TestWeekdayConstants(t *testing.T) {
	// Test that weekday constants are properly defined
	expectedValues := map[int]string{
		Sunday:    "Sunday",
		Monday:    "Monday",
		Tuesday:   "Tuesday",
		Wednesday: "Wednesday",
		Thursday:  "Thursday",
		Friday:    "Friday",
		Saturday:  "Saturday",
	}
	
	for day, expectedName := range expectedValues {
		if day < 0 || day > 6 {
			t.Errorf("Day constant %s has invalid value %d", expectedName, day)
		}
	}
	
	// Test sequential values
	if Monday != Sunday+1 {
		t.Errorf("Monday should be Sunday+1")
	}
	
	if Saturday != Friday+1 {
		t.Errorf("Saturday should be Friday+1")
	}
}

func TestFilePermissionConstants(t *testing.T) {
	// Test bit values
	if ReadPermission != 1 {
		t.Errorf("Expected ReadPermission to be 1, got %d", ReadPermission)
	}
	
	if WritePermission != 2 {
		t.Errorf("Expected WritePermission to be 2, got %d", WritePermission)
	}
	
	if ExecutePermission != 4 {
		t.Errorf("Expected ExecutePermission to be 4, got %d", ExecutePermission)
	}
	
	// Test bit operations
	readWrite := ReadPermission | WritePermission
	if readWrite != 3 {
		t.Errorf("Expected ReadPermission|WritePermission to be 3, got %d", readWrite)
	}
	
	fullPerms := ReadPermission | WritePermission | ExecutePermission
	if fullPerms != 7 {
		t.Errorf("Expected full permissions to be 7, got %d", fullPerms)
	}
}

func TestHTTPStatusConstants(t *testing.T) {
	if StatusOK != 200 {
		t.Errorf("Expected StatusOK to be 200, got %d", StatusOK)
	}
	
	if StatusCreated != 201 {
		t.Errorf("Expected StatusCreated to be 201, got %d", StatusCreated)
	}
	
	if StatusAccepted != 202 {
		t.Errorf("Expected StatusAccepted to be 202, got %d", StatusAccepted)
	}
	
	if StatusNoContent != 204 {
		t.Errorf("Expected StatusNoContent to be 204, got %d", StatusNoContent)
	}
}

func TestFileSizeConstants(t *testing.T) {
	if KB != 1024 {
		t.Errorf("Expected KB to be 1024, got %d", KB)
	}
	
	if MB != 1024*1024 {
		t.Errorf("Expected MB to be %d, got %d", 1024*1024, MB)
	}
	
	if GB != 1024*1024*1024 {
		t.Errorf("Expected GB to be %d, got %d", 1024*1024*1024, GB)
	}
	
	expectedTB := uint64(1024 * 1024 * 1024 * 1024)
	if TB != expectedTB {
		t.Errorf("Expected TB to be %d, got %d", expectedTB, TB)
	}
}

func TestColorType(t *testing.T) {
	// Test Color String method
	if Red.String() != "Red" {
		t.Errorf("Expected Red.String() to be 'Red', got '%s'", Red.String())
	}
	
	if Blue.String() != "Blue" {
		t.Errorf("Expected Blue.String() to be 'Blue', got '%s'", Blue.String())
	}
	
	// Test invalid color
	invalidColor := Color(999)
	if invalidColor.String() != "Unknown" {
		t.Errorf("Expected invalid color to return 'Unknown', got '%s'", invalidColor.String())
	}
}

func TestPriorityType(t *testing.T) {
	// Test Priority String method
	if Low.String() != "Low" {
		t.Errorf("Expected Low.String() to be 'Low', got '%s'", Low.String())
	}
	
	if Critical.String() != "Critical" {
		t.Errorf("Expected Critical.String() to be 'Critical', got '%s'", Critical.String())
	}
	
	// Test priority ordering
	if Low >= Medium {
		t.Errorf("Low priority should be less than Medium")
	}
	
	if Critical <= High {
		t.Errorf("Critical priority should be greater than High")
	}
}

func TestGetConstantTypes(t *testing.T) {
	constants := GetConstantTypes()
	
	expectedKeys := []string{
		"string_constant", "int_constant", "float_constant",
		"bool_constant", "iota_constant", "custom_type",
	}
	
	for _, key := range expectedKeys {
		if _, exists := constants[key]; !exists {
			t.Errorf("Expected constant type '%s' not found", key)
		}
	}
	
	// Test specific values
	if constants["string_constant"] != AppName {
		t.Errorf("Expected string_constant to be AppName")
	}
	
	if constants["int_constant"] != MaxRetries {
		t.Errorf("Expected int_constant to be MaxRetries")
	}
}

func TestIsWeekend(t *testing.T) {
	// Test weekend days
	if !IsWeekend(Sunday) {
		t.Errorf("Sunday should be weekend")
	}
	
	if !IsWeekend(Saturday) {
		t.Errorf("Saturday should be weekend")
	}
	
	// Test weekdays
	weekdays := []int{Monday, Tuesday, Wednesday, Thursday, Friday}
	for _, day := range weekdays {
		if IsWeekend(day) {
			t.Errorf("Day %d should not be weekend", day)
		}
	}
}

func TestGetDayName(t *testing.T) {
	expectedDays := map[int]string{
		0: "Sunday",
		1: "Monday",
		2: "Tuesday",
		3: "Wednesday",
		4: "Thursday",
		5: "Friday",
		6: "Saturday",
	}
	
	for day, expectedName := range expectedDays {
		if GetDayName(day) != expectedName {
			t.Errorf("Expected day %d to be '%s', got '%s'", 
				day, expectedName, GetDayName(day))
		}
	}
	
	// Test invalid days
	if GetDayName(-1) != "Invalid day" {
		t.Errorf("Expected 'Invalid day' for -1")
	}
	
	if GetDayName(7) != "Invalid day" {
		t.Errorf("Expected 'Invalid day' for 7")
	}
}

func TestFormatFileSize(t *testing.T) {
	testCases := []struct {
		bytes    uint64
		expected string
	}{
		{512, "512 bytes"},
		{1024, "1.00 KB"},
		{1536, "1.50 KB"},
		{1048576, "1.00 MB"},
		{1073741824, "1.00 GB"},
		{1099511627776, "1.00 TB"},
		{2147483648, "2.00 GB"},
	}
	
	for _, tc := range testCases {
		result := FormatFileSize(tc.bytes)
		if result != tc.expected {
			t.Errorf("FormatFileSize(%d): expected '%s', got '%s'", 
				tc.bytes, tc.expected, result)
		}
	}
}

func TestConstantImmutability(t *testing.T) {
	// Constants should be immutable - this is enforced at compile time
	// We can test that they maintain their values
	originalAppName := AppName
	originalMaxRetries := MaxRetries
	
	// Call functions that use constants
	DemoBasicConstants()
	
	// Values should remain unchanged
	if AppName != originalAppName {
		t.Errorf("AppName changed after function call")
	}
	
	if MaxRetries != originalMaxRetries {
		t.Errorf("MaxRetries changed after function call")
	}
}

func TestMathematicalConstants(t *testing.T) {
	// Test that Pi is approximately correct
	if Pi < 3.14 || Pi > 3.15 {
		t.Errorf("Pi constant seems incorrect: %f", Pi)
	}
	
	// Test that E is approximately correct
	if E < 2.71 || E > 2.72 {
		t.Errorf("E constant seems incorrect: %f", E)
	}
}