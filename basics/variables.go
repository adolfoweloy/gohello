// Package basics demonstrates fundamental Go concepts including variables,
// data types, and basic operations.
package basics

import "fmt"

// DemoVariables demonstrates different ways to declare and use variables in Go
func DemoVariables() {
	fmt.Println("\n=== Variables and Data Types Demo ===")

	// Variable declarations with explicit types
	var name string = "Gopher"
	var age int = 10
	var isActive bool = true
	var height float64 = 5.9

	fmt.Printf("Explicit types: %s (age: %d, active: %t, height: %.1f)\n", 
		name, age, isActive, height)

	// Type inference
	var city = "San Francisco"  // string inferred
	var population = 884363     // int inferred
	var latitude = 37.7749     // float64 inferred

	fmt.Printf("Type inference: %s (population: %d, latitude: %.4f)\n", 
		city, population, latitude)

	// Short variable declarations (only inside functions)
	country := "United States"
	founded := 1776
	isIndependent := true

	fmt.Printf("Short declarations: %s (founded: %d, independent: %t)\n", 
		country, founded, isIndependent)

	// Multiple variable declarations
	var (
		firstName = "John"
		lastName  = "Doe"
		email     = "john.doe@example.com"
	)

	fmt.Printf("Multiple declarations: %s %s <%s>\n", firstName, lastName, email)

	// Multiple assignments
	x, y, z := 1, 2, 3
	fmt.Printf("Multiple assignments: x=%d, y=%d, z=%d\n", x, y, z)

	// Zero values (default values for uninitialized variables)
	var defaultInt int
	var defaultString string
	var defaultBool bool
	var defaultFloat float64

	fmt.Printf("Zero values: int=%d, string='%s', bool=%t, float=%.1f\n", 
		defaultInt, defaultString, defaultBool, defaultFloat)
}

// DemoArrays demonstrates arrays in Go
func DemoArrays() {
	fmt.Println("\n=== Arrays Demo ===")

	// Array declaration with size
	var numbers [5]int
	numbers[0] = 10
	numbers[1] = 20
	numbers[2] = 30
	numbers[3] = 40
	numbers[4] = 50

	fmt.Printf("Array by index assignment: %v\n", numbers)

	// Array literal
	colors := [3]string{"red", "green", "blue"}
	fmt.Printf("Array literal: %v\n", colors)

	// Array with inferred size
	fruits := [...]string{"apple", "banana", "orange", "grape"}
	fmt.Printf("Inferred size array: %v (length: %d)\n", fruits, len(fruits))

	// Accessing array elements
	fmt.Printf("First fruit: %s, Last fruit: %s\n", fruits[0], fruits[len(fruits)-1])

	// Arrays are value types - copying creates a new array
	originalArray := [3]int{1, 2, 3}
	copiedArray := originalArray
	copiedArray[0] = 99

	fmt.Printf("Original: %v, Copied: %v\n", originalArray, copiedArray)
}

// DemoSlices demonstrates slices in Go
func DemoSlices() {
	fmt.Println("\n=== Slices Demo ===")

	// Slice declaration
	var numbers []int
	fmt.Printf("Empty slice: %v (length: %d, capacity: %d)\n", 
		numbers, len(numbers), cap(numbers))

	// Slice literal
	colors := []string{"red", "green", "blue"}
	fmt.Printf("Slice literal: %v (length: %d, capacity: %d)\n", 
		colors, len(colors), cap(colors))

	// Creating slice with make
	dynamicSlice := make([]int, 3, 5) // length=3, capacity=5
	fmt.Printf("Made slice: %v (length: %d, capacity: %d)\n", 
		dynamicSlice, len(dynamicSlice), cap(dynamicSlice))

	// Appending to slices
	numbers = append(numbers, 1, 2, 3, 4, 5)
	fmt.Printf("After append: %v (length: %d, capacity: %d)\n", 
		numbers, len(numbers), cap(numbers))

	// Slice operations
	subset := numbers[1:4] // elements at index 1, 2, 3
	fmt.Printf("Subset [1:4]: %v\n", subset)

	firstThree := numbers[:3] // first 3 elements
	fmt.Printf("First three [:3]: %v\n", firstThree)

	lastTwo := numbers[3:] // from index 3 to end
	fmt.Printf("Last two [3:]: %v\n", lastTwo)

	// Slices share underlying array
	original := []int{1, 2, 3, 4, 5}
	view := original[1:4]
	view[0] = 99

	fmt.Printf("Original after modifying view: %v\n", original)
	fmt.Printf("View: %v\n", view)

	// Copy slices to avoid sharing
	source := []int{1, 2, 3}
	destination := make([]int, len(source))
	copy(destination, source)
	destination[0] = 99

	fmt.Printf("Source: %v, Destination: %v\n", source, destination)
}

// DemoMaps demonstrates maps in Go
func DemoMaps() {
	fmt.Println("\n=== Maps Demo ===")

	// Map declaration
	var scores map[string]int
	fmt.Printf("Uninitialized map: %v\n", scores)

	// Initialize with make
	scores = make(map[string]int)
	scores["Alice"] = 95
	scores["Bob"] = 87
	scores["Charlie"] = 92

	fmt.Printf("Scores map: %v\n", scores)

	// Map literal
	capitals := map[string]string{
		"USA":    "Washington, D.C.",
		"France": "Paris",
		"Japan":  "Tokyo",
		"Brazil": "Brasília",
	}

	fmt.Printf("Capitals: %v\n", capitals)

	// Accessing map values
	fmt.Printf("Capital of Japan: %s\n", capitals["Japan"])

	// Check if key exists
	capital, exists := capitals["Germany"]
	if exists {
		fmt.Printf("Capital of Germany: %s\n", capital)
	} else {
		fmt.Println("Germany not found in map")
	}

	// Delete from map
	delete(capitals, "France")
	fmt.Printf("After deleting France: %v\n", capitals)

	// Iterate over map
	fmt.Println("All capitals:")
	for country, capital := range capitals {
		fmt.Printf("  %s: %s\n", country, capital)
	}

	// Map length
	fmt.Printf("Number of countries: %d\n", len(capitals))
}

// PersonInfo demonstrates struct as a custom data type
type PersonInfo struct {
	Name string
	Age  int
}

// DemoStructBasics demonstrates basic struct usage
func DemoStructBasics() {
	fmt.Println("\n=== Basic Structs Demo ===")

	// Struct literal
	person1 := PersonInfo{
		Name: "Alice",
		Age:  30,
	}

	fmt.Printf("Person 1: %+v\n", person1)

	// Positional initialization
	person2 := PersonInfo{"Bob", 25}
	fmt.Printf("Person 2: %+v\n", person2)

	// Field access
	fmt.Printf("%s is %d years old\n", person1.Name, person1.Age)

	// Modify fields
	person1.Age = 31
	fmt.Printf("After birthday: %+v\n", person1)

	// Anonymous struct
	config := struct {
		Host string
		Port int
	}{
		Host: "localhost",
		Port: 8080,
	}

	fmt.Printf("Config: %+v\n", config)
}

// GetTypeInfo returns information about Go's basic types
func GetTypeInfo() map[string]string {
	return map[string]string{
		"bool":       "Boolean type (true/false)",
		"string":     "String type (UTF-8 encoded)",
		"int":        "Platform-dependent integer (32 or 64 bit)",
		"int8":       "8-bit signed integer (-128 to 127)",
		"int16":      "16-bit signed integer (-32,768 to 32,767)",
		"int32":      "32-bit signed integer (-2^31 to 2^31-1)",
		"int64":      "64-bit signed integer (-2^63 to 2^63-1)",
		"uint":       "Platform-dependent unsigned integer",
		"uint8":      "8-bit unsigned integer (0 to 255)",
		"uint16":     "16-bit unsigned integer (0 to 65,535)",
		"uint32":     "32-bit unsigned integer (0 to 2^32-1)",
		"uint64":     "64-bit unsigned integer (0 to 2^64-1)",
		"byte":       "Alias for uint8",
		"rune":       "Alias for int32 (Unicode code point)",
		"float32":    "32-bit floating point",
		"float64":    "64-bit floating point",
		"complex64":  "Complex number with float32 real and imaginary parts",
		"complex128": "Complex number with float64 real and imaginary parts",
	}
}