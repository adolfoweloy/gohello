// Package basics demonstrates fundamental Go programming concepts
package basics

import "fmt"

// DemonstrateVariables shows different ways to declare and use variables in Go
func DemonstrateVariables() {
	fmt.Println("=== Variables and Data Types ===")
	
	// Variable declarations
	var name string = "Go Programming"
	var age int = 14 // Go was first released in 2009
	var isOpen bool = true
	
	// Short variable declaration
	language := "Go"
	version := 1.21
	
	// Multiple variable declarations
	var (
		width  int = 100
		height int = 200
	)
	
	// Zero values
	var zeroInt int
	var zeroString string
	var zeroBool bool
	
	fmt.Printf("Name: %s, Age: %d, IsOpen: %t\n", name, age, isOpen)
	fmt.Printf("Language: %s, Version: %.2f\n", language, version)
	fmt.Printf("Dimensions: %dx%d\n", width, height)
	fmt.Printf("Zero values - Int: %d, String: '%s', Bool: %t\n", zeroInt, zeroString, zeroBool)
}

// DemonstrateArraysAndSlices shows arrays and slices usage
func DemonstrateArraysAndSlices() {
	fmt.Println("\n=== Arrays and Slices ===")
	
	// Arrays - fixed size
	var numbers [5]int = [5]int{1, 2, 3, 4, 5}
	fruits := [3]string{"apple", "banana", "orange"}
	
	fmt.Printf("Array numbers: %v\n", numbers)
	fmt.Printf("Array fruits: %v\n", fruits)
	fmt.Printf("Array length: %d\n", len(numbers))
	
	// Slices - dynamic arrays
	var slice []int
	slice = append(slice, 10, 20, 30)
	
	colors := []string{"red", "green", "blue"}
	colors = append(colors, "yellow", "purple")
	
	fmt.Printf("Slice: %v, Length: %d, Capacity: %d\n", slice, len(slice), cap(slice))
	fmt.Printf("Colors: %v\n", colors)
	
	// Slice operations
	fmt.Printf("First 3 colors: %v\n", colors[:3])
	fmt.Printf("Colors from index 2: %v\n", colors[2:])
	fmt.Printf("Colors from 1 to 3: %v\n", colors[1:4])
}

// DemonstrateMaps shows map usage
func DemonstrateMaps() {
	fmt.Println("\n=== Maps ===")
	
	// Map declaration and initialization
	ages := make(map[string]int)
	ages["Alice"] = 30
	ages["Bob"] = 25
	ages["Charlie"] = 35
	
	// Map literal
	capitals := map[string]string{
		"France": "Paris",
		"Italy":  "Rome",
		"Spain":  "Madrid",
	}
	
	fmt.Printf("Ages: %v\n", ages)
	fmt.Printf("Capitals: %v\n", capitals)
	
	// Check if key exists
	if age, exists := ages["Alice"]; exists {
		fmt.Printf("Alice's age: %d\n", age)
	}
	
	// Delete from map
	delete(ages, "Bob")
	fmt.Printf("After deleting Bob: %v\n", ages)
	
	// Iterate over map
	fmt.Println("Capitals:")
	for country, capital := range capitals {
		fmt.Printf("  %s: %s\n", country, capital)
	}
}