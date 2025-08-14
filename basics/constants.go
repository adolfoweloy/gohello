// Package basics demonstrates fundamental Go concepts including constants,
// iota, and enumeration patterns.
package basics

import "fmt"

// Constants demonstration
const (
	// Basic constants
	AppName    = "GoHello"
	AppVersion = "1.0.0"
	MaxRetries = 3
	
	// Typed constants
	Pi       float64 = 3.14159265359
	E        float64 = 2.71828182846
	IsActive bool    = true
)

// Iota demonstration - automatic enumeration
const (
	// Basic iota usage
	Sunday = iota    // 0
	Monday           // 1
	Tuesday          // 2
	Wednesday        // 3
	Thursday         // 4
	Friday           // 5
	Saturday         // 6
)

// File permissions using iota with bit shifting
const (
	ReadPermission   = 1 << iota // 1 << 0 = 1
	WritePermission              // 1 << 1 = 2
	ExecutePermission            // 1 << 2 = 4
)

// HTTP status codes using iota with custom starting values
const (
	StatusOK                  = 200 + iota // 200
	StatusCreated                          // 201
	StatusAccepted                         // 202
	StatusNonAuthoritativeInfo             // 203
	StatusNoContent                        // 204
)

// File size constants using iota for powers of 1024
const (
	_  = iota // ignore first value by assigning to blank identifier
	KB = 1 << (10 * iota) // 1 << (10*1) = 1024
	MB                    // 1 << (10*2) = 1048576
	GB                    // 1 << (10*3) = 1073741824
	TB                    // 1 << (10*4) = 1099511627776
)

// Color constants using string iota
const (
	Red Color = iota
	Green
	Blue
	Yellow
	Purple
)

// Color type for demonstration
type Color int

// String method for Color type to make it printable
func (c Color) String() string {
	colors := []string{"Red", "Green", "Blue", "Yellow", "Purple"}
	if c < 0 || int(c) >= len(colors) {
		return "Unknown"
	}
	return colors[c]
}

// Priority levels using iota
const (
	Low Priority = iota
	Medium
	High
	Critical
)

// Priority type
type Priority int

func (p Priority) String() string {
	priorities := []string{"Low", "Medium", "High", "Critical"}
	if p < 0 || int(p) >= len(priorities) {
		return "Unknown"
	}
	return priorities[p]
}

// DemoBasicConstants demonstrates basic constant usage
func DemoBasicConstants() {
	fmt.Println("\n=== Basic Constants Demo ===")
	
	fmt.Printf("Application: %s v%s\n", AppName, AppVersion)
	fmt.Printf("Max retries: %d\n", MaxRetries)
	fmt.Printf("Mathematical constants: π=%.5f, e=%.5f\n", Pi, E)
	fmt.Printf("Status: Active = %t\n", IsActive)
	
	// Constants can be used in expressions
	circumference := 2 * Pi * 5
	fmt.Printf("Circumference of circle with radius 5: %.2f\n", circumference)
}

// DemoIotaBasics demonstrates basic iota usage for weekdays
func DemoIotaBasics() {
	fmt.Println("\n=== Iota Basics Demo ===")
	
	fmt.Printf("Weekdays using iota:\n")
	fmt.Printf("Sunday: %d\n", Sunday)
	fmt.Printf("Monday: %d\n", Monday)
	fmt.Printf("Tuesday: %d\n", Tuesday)
	fmt.Printf("Wednesday: %d\n", Wednesday)
	fmt.Printf("Thursday: %d\n", Thursday)
	fmt.Printf("Friday: %d\n", Friday)
	fmt.Printf("Saturday: %d\n", Saturday)
	
	// Using constants in conditionals
	today := Friday
	if today == Friday {
		fmt.Println("TGIF! 🎉")
	}
}

// DemoFilePermissions demonstrates iota with bit operations
func DemoFilePermissions() {
	fmt.Println("\n=== File Permissions Demo ===")
	
	fmt.Printf("Read Permission: %d (binary: %b)\n", ReadPermission, ReadPermission)
	fmt.Printf("Write Permission: %d (binary: %b)\n", WritePermission, WritePermission)
	fmt.Printf("Execute Permission: %d (binary: %b)\n", ExecutePermission, ExecutePermission)
	
	// Combining permissions using bitwise OR
	readWrite := ReadPermission | WritePermission
	fmt.Printf("Read + Write: %d (binary: %b)\n", readWrite, readWrite)
	
	fullPermissions := ReadPermission | WritePermission | ExecutePermission
	fmt.Printf("Full Permissions: %d (binary: %b)\n", fullPermissions, fullPermissions)
	
	// Checking permissions using bitwise AND
	hasRead := fullPermissions&ReadPermission != 0
	hasWrite := fullPermissions&WritePermission != 0
	hasExecute := fullPermissions&ExecutePermission != 0
	
	fmt.Printf("Has read: %t, write: %t, execute: %t\n", hasRead, hasWrite, hasExecute)
}

// DemoHTTPStatus demonstrates iota with custom starting values
func DemoHTTPStatus() {
	fmt.Println("\n=== HTTP Status Codes Demo ===")
	
	fmt.Printf("HTTP Status Codes:\n")
	fmt.Printf("OK: %d\n", StatusOK)
	fmt.Printf("Created: %d\n", StatusCreated)
	fmt.Printf("Accepted: %d\n", StatusAccepted)
	fmt.Printf("Non-Authoritative Info: %d\n", StatusNonAuthoritativeInfo)
	fmt.Printf("No Content: %d\n", StatusNoContent)
}

// DemoFileSizes demonstrates iota with exponential values
func DemoFileSizes() {
	fmt.Println("\n=== File Sizes Demo ===")
	
	fmt.Printf("File size constants:\n")
	fmt.Printf("1 KB = %d bytes\n", KB)
	fmt.Printf("1 MB = %d bytes (%.2f KB)\n", MB, float64(MB)/float64(KB))
	fmt.Printf("1 GB = %d bytes (%.2f MB)\n", GB, float64(GB)/float64(MB))
	fmt.Printf("1 TB = %d bytes (%.2f GB)\n", TB, float64(TB)/float64(GB))
	
	// Example usage
	fileSize := 2.5 * GB
	fmt.Printf("Example file size: %.1f GB = %.0f bytes\n", 
		fileSize/float64(GB), fileSize)
}

// DemoCustomTypes demonstrates iota with custom types
func DemoCustomTypes() {
	fmt.Println("\n=== Custom Types with Iota Demo ===")
	
	// Color enumeration
	fmt.Printf("Colors:\n")
	fmt.Printf("Red: %s (%d)\n", Red, Red)
	fmt.Printf("Green: %s (%d)\n", Green, Green)
	fmt.Printf("Blue: %s (%d)\n", Blue, Blue)
	fmt.Printf("Yellow: %s (%d)\n", Yellow, Yellow)
	fmt.Printf("Purple: %s (%d)\n", Purple, Purple)
	
	// Priority enumeration
	fmt.Printf("\nPriorities:\n")
	fmt.Printf("Low: %s (%d)\n", Low, Low)
	fmt.Printf("Medium: %s (%d)\n", Medium, Medium)
	fmt.Printf("High: %s (%d)\n", High, High)
	fmt.Printf("Critical: %s (%d)\n", Critical, Critical)
	
	// Using in logic
	taskPriority := High
	if taskPriority >= High {
		fmt.Printf("Task has %s priority - needs immediate attention!\n", taskPriority)
	}
}

// GetConstantTypes returns examples of different constant types
func GetConstantTypes() map[string]interface{} {
	return map[string]interface{}{
		"string_constant":   AppName,
		"int_constant":      MaxRetries,
		"float_constant":    Pi,
		"bool_constant":     IsActive,
		"iota_constant":     Friday,
		"custom_type":       High,
	}
}

// IsWeekend checks if a given day is weekend
func IsWeekend(day int) bool {
	return day == Sunday || day == Saturday
}

// GetDayName returns the name of the day for a given number
func GetDayName(day int) string {
	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	if day < 0 || day >= len(days) {
		return "Invalid day"
	}
	return days[day]
}

// FormatFileSize formats a size in bytes to human readable format
func FormatFileSize(bytes uint64) string {
	switch {
	case bytes >= TB:
		return fmt.Sprintf("%.2f TB", float64(bytes)/float64(TB))
	case bytes >= GB:
		return fmt.Sprintf("%.2f GB", float64(bytes)/float64(GB))
	case bytes >= MB:
		return fmt.Sprintf("%.2f MB", float64(bytes)/float64(MB))
	case bytes >= KB:
		return fmt.Sprintf("%.2f KB", float64(bytes)/float64(KB))
	default:
		return fmt.Sprintf("%d bytes", bytes)
	}
}