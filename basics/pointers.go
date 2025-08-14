package basics

import "fmt"

// DemonstratePointers shows pointer basics in Go
func DemonstratePointers() {
	fmt.Println("\n=== Pointers ===")
	
	// Basic pointer usage
	x := 42
	p := &x // p is a pointer to x
	
	fmt.Printf("x = %d\n", x)
	fmt.Printf("Address of x: %p\n", &x)
	fmt.Printf("Pointer p: %p\n", p)
	fmt.Printf("Value at p: %d\n", *p) // dereferencing
	
	// Modify value through pointer
	*p = 100
	fmt.Printf("After *p = 100, x = %d\n", x)
	
	// Zero value of pointer is nil
	var ptr *int
	fmt.Printf("Zero value pointer: %v\n", ptr)
	
	if ptr == nil {
		fmt.Println("Pointer is nil")
	}
	
	// Pointer to different types
	name := "Go"
	namePtr := &name
	fmt.Printf("Name: %s, Pointer: %p, Value: %s\n", name, namePtr, *namePtr)
	
	// Function with pointer parameters
	fmt.Println("\n--- Pointer Parameters ---")
	demonstratePointerParams()
	
	// Pointer arithmetic is not allowed in Go (unlike C/C++)
	// This is a safety feature
	fmt.Println("\n--- Arrays and Pointers ---")
	demonstrateArrayPointers()
}

// increment modifies the value through a pointer
func increment(ptr *int) {
	*ptr++
}

// incrementCopy takes a copy of the value (no modification to original)
func incrementCopy(val int) {
	val++
}

// demonstratePointerParams shows the difference between pass by value and pass by pointer
func demonstratePointerParams() {
	num := 5
	fmt.Printf("Original num: %d\n", num)
	
	// Pass by value - no change to original
	incrementCopy(num)
	fmt.Printf("After incrementCopy: %d\n", num)
	
	// Pass by pointer - modifies original
	increment(&num)
	fmt.Printf("After increment: %d\n", num)
}

// demonstrateArrayPointers shows arrays and pointers
func demonstrateArrayPointers() {
	// Arrays are passed by value in Go
	arr := [3]int{1, 2, 3}
	fmt.Printf("Original array: %v\n", arr)
	
	// Pointer to array
	arrPtr := &arr
	fmt.Printf("Array pointer: %p\n", arrPtr)
	
	// Modify through pointer
	(*arrPtr)[0] = 10
	fmt.Printf("After modifying through pointer: %v\n", arr)
	
	// Slices contain pointers internally
	slice := []int{1, 2, 3}
	fmt.Printf("Original slice: %v\n", slice)
	
	modifySlice(slice)
	fmt.Printf("After modifySlice: %v\n", slice)
}

// modifySlice demonstrates that slices are reference types
func modifySlice(s []int) {
	if len(s) > 0 {
		s[0] = 99 // This modifies the original slice
	}
}

// Person struct for demonstrating struct pointers
type Person struct {
	Name string
	Age  int
}

// UpdateAge modifies a person's age through a pointer
func (p *Person) UpdateAge(newAge int) {
	p.Age = newAge
}

// GetInfo returns person info (value receiver)
func (p Person) GetInfo() string {
	return fmt.Sprintf("%s is %d years old", p.Name, p.Age)
}

// DemonstrateStructPointers shows pointers with structs
func DemonstrateStructPointers() {
	fmt.Println("\n--- Struct Pointers ---")
	
	// Create a struct
	person := Person{Name: "Alice", Age: 30}
	fmt.Printf("Original: %s\n", person.GetInfo())
	
	// Pointer to struct
	personPtr := &person
	
	// Access fields through pointer (Go automatically dereferences)
	fmt.Printf("Name through pointer: %s\n", personPtr.Name)
	
	// Modify through pointer
	personPtr.Age = 31
	fmt.Printf("After direct modification: %s\n", person.GetInfo())
	
	// Method with pointer receiver
	person.UpdateAge(32)
	fmt.Printf("After UpdateAge method: %s\n", person.GetInfo())
	
	// Create struct using new
	newPerson := new(Person)
	newPerson.Name = "Bob"
	newPerson.Age = 25
	fmt.Printf("New person: %s\n", newPerson.GetInfo())
}