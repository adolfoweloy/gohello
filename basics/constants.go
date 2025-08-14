package basics

import "fmt"

// DemonstrateConstants shows how to use constants in Go
func DemonstrateConstants() {
	fmt.Println("\n=== Constants and Iota ===")
	
	// Basic constants
	const pi = 3.14159
	const greeting = "Hello, Go!"
	const maxUsers = 100
	
	fmt.Printf("Pi: %.5f\n", pi)
	fmt.Printf("Greeting: %s\n", greeting)
	fmt.Printf("Max users: %d\n", maxUsers)
	
	// Constants block
	const (
		statusOK    = 200
		statusError = 500
		statusNotFound = 404
	)
	
	fmt.Printf("HTTP Status codes - OK: %d, Error: %d, Not Found: %d\n", 
		statusOK, statusError, statusNotFound)
}

// Weekday represents days of the week using iota
type Weekday int

const (
	Sunday Weekday = iota // 0
	Monday                // 1
	Tuesday               // 2
	Wednesday             // 3
	Thursday              // 4
	Friday                // 5
	Saturday              // 6
)

// String returns the string representation of Weekday
func (w Weekday) String() string {
	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	if w < 0 || int(w) >= len(days) {
		return "Unknown"
	}
	return days[w]
}

// FilePermission demonstrates iota with bit operations
type FilePermission int

const (
	Read FilePermission = 1 << iota // 1 = 001
	Write                           // 2 = 010
	Execute                         // 4 = 100
)

// DemonstrateIota shows iota usage with enums
func DemonstrateIota() {
	fmt.Println("\n=== Iota Examples ===")
	
	// Weekday examples
	today := Tuesday
	weekend := Saturday
	
	fmt.Printf("Today is %s (%d)\n", today, today)
	fmt.Printf("Weekend day: %s (%d)\n", weekend, weekend)
	
	// File permission examples
	fmt.Printf("File permissions:\n")
	fmt.Printf("Read: %d (binary: %08b)\n", Read, Read)
	fmt.Printf("Write: %d (binary: %08b)\n", Write, Write)
	fmt.Printf("Execute: %d (binary: %08b)\n", Execute, Execute)
	
	// Combined permissions
	readWrite := Read | Write
	fullAccess := Read | Write | Execute
	
	fmt.Printf("Read+Write: %d (binary: %08b)\n", readWrite, readWrite)
	fmt.Printf("Full access: %d (binary: %08b)\n", fullAccess, fullAccess)
}