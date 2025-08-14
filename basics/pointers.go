// Package basics demonstrates Go pointers including basic pointer operations,
// pointer dereferencing, pass by reference, and pointer safety.
package basics

import (
	"fmt"
	"unsafe"
)

// DemoBasicPointers demonstrates basic pointer concepts
func DemoBasicPointers() {
	fmt.Println("\n=== Basic Pointers Demo ===")

	// Basic variable
	x := 42
	fmt.Printf("Variable x: value=%d, address=%p\n", x, &x)

	// Pointer declaration and initialization
	var p *int
	fmt.Printf("Uninitialized pointer p: %p (nil pointer)\n", p)

	// Point to x
	p = &x
	fmt.Printf("Pointer p pointing to x: %p\n", p)
	fmt.Printf("Value at pointer p: %d\n", *p)

	// Modify value through pointer
	*p = 100
	fmt.Printf("After modifying through pointer: x=%d, *p=%d\n", x, *p)

	// Create pointer directly
	y := 25
	q := &y
	fmt.Printf("Direct pointer creation: y=%d, *q=%d\n", y, *q)

	// Pointer to pointer
	pp := &p
	fmt.Printf("Pointer to pointer: pp=%p, *pp=%p, **pp=%d\n", pp, *pp, **pp)
}

// DemoPointerTypes demonstrates pointers with different data types
func DemoPointerTypes() {
	fmt.Println("\n=== Pointer Types Demo ===")

	// String pointer
	name := "Alice"
	namePtr := &name
	fmt.Printf("String: %s, pointer: %p, dereferenced: %s\n", name, namePtr, *namePtr)

	// Boolean pointer
	isActive := true
	activePtr := &isActive
	fmt.Printf("Bool: %t, pointer: %p, dereferenced: %t\n", isActive, activePtr, *activePtr)

	// Slice pointer
	numbers := []int{1, 2, 3, 4, 5}
	numbersPtr := &numbers
	fmt.Printf("Slice: %v, pointer: %p\n", numbers, numbersPtr)
	fmt.Printf("Dereferenced slice: %v\n", *numbersPtr)
	
	// Modify slice through pointer
	*numbersPtr = append(*numbersPtr, 6)
	fmt.Printf("After append through pointer: %v\n", numbers)

	// Map pointer
	scores := map[string]int{"Alice": 95, "Bob": 87}
	scoresPtr := &scores
	fmt.Printf("Map: %v, pointer: %p\n", scores, scoresPtr)
	(*scoresPtr)["Charlie"] = 92
	fmt.Printf("After adding through pointer: %v\n", scores)
}

// DemoPassByValue demonstrates pass by value (default in Go)
func DemoPassByValue() {
	fmt.Println("\n=== Pass By Value Demo ===")

	original := 10
	fmt.Printf("Before function call: %d\n", original)

	modifyValue(original)
	fmt.Printf("After modifyValue: %d (unchanged)\n", original)

	// With slice (reference type behavior)
	slice := []int{1, 2, 3}
	fmt.Printf("Before modifySlice: %v\n", slice)

	modifySlice(slice)
	fmt.Printf("After modifySlice: %v (modified - slice header copied, but data shared)\n", slice)

	// Reassigning slice doesn't affect original
	fmt.Printf("Before reassignSlice: %v\n", slice)
	reassignSlice(slice)
	fmt.Printf("After reassignSlice: %v (unchanged - new slice assignment)\n", slice)
}

// modifyValue attempts to modify an integer (won't work - pass by value)
func modifyValue(x int) {
	x = 999
	fmt.Printf("  Inside modifyValue: %d\n", x)
}

// modifySlice modifies slice elements (works - slice data is shared)
func modifySlice(s []int) {
	if len(s) > 0 {
		s[0] = 999
	}
	fmt.Printf("  Inside modifySlice: %v\n", s)
}

// reassignSlice reassigns the slice (doesn't affect original)
func reassignSlice(s []int) {
	s = []int{888, 777, 666}
	fmt.Printf("  Inside reassignSlice: %v\n", s)
}

// DemoPassByReference demonstrates pass by reference using pointers
func DemoPassByReference() {
	fmt.Println("\n=== Pass By Reference Demo ===")

	value := 10
	fmt.Printf("Before function call: %d\n", value)

	modifyByPointer(&value)
	fmt.Printf("After modifyByPointer: %d (modified through pointer)\n", value)

	// With struct
	person := Person{Name: "Alice", Age: 30}
	fmt.Printf("Before modifyPerson: %+v\n", person)

	modifyPersonByValue(person)
	fmt.Printf("After modifyPersonByValue: %+v (unchanged)\n", person)

	modifyPersonByPointer(&person)
	fmt.Printf("After modifyPersonByPointer: %+v (modified)\n", person)
}

// Person struct for demonstration
type Person struct {
	Name string
	Age  int
}

// modifyByPointer modifies value through pointer
func modifyByPointer(x *int) {
	*x = 999
	fmt.Printf("  Inside modifyByPointer: %d\n", *x)
}

// modifyPersonByValue attempts to modify person (won't affect original)
func modifyPersonByValue(p Person) {
	p.Age = 999
	fmt.Printf("  Inside modifyPersonByValue: %+v\n", p)
}

// modifyPersonByPointer modifies person through pointer
func modifyPersonByPointer(p *Person) {
	p.Age = 31
	p.Name = "Alice Smith"
	fmt.Printf("  Inside modifyPersonByPointer: %+v\n", p)
}

// DemoPointerArithmetic demonstrates pointer operations (limited in Go)
func DemoPointerArithmetic() {
	fmt.Println("\n=== Pointer Operations Demo ===")

	// Go doesn't support pointer arithmetic like C/C++
	// But we can demonstrate some safe operations

	var p *int
	fmt.Printf("Nil pointer: %p\n", p)

	x := 42
	p = &x
	fmt.Printf("Pointer to x: %p\n", p)

	// Pointer comparison
	y := 42
	q := &y
	fmt.Printf("Pointer to y: %p\n", q)
	fmt.Printf("p == q: %t (different memory addresses)\n", p == q)
	fmt.Printf("*p == *q: %t (same values)\n", *p == *q)

	// Pointer to same variable
	r := &x
	fmt.Printf("Another pointer to x: %p\n", r)
	fmt.Printf("p == r: %t (same memory address)\n", p == r)

	// Nil pointer comparison
	var nilPtr *int
	fmt.Printf("p == nil: %t\n", p == nil)
	fmt.Printf("nilPtr == nil: %t\n", nilPtr == nil)
}

// DemoPointerSafety demonstrates Go's pointer safety features
func DemoPointerSafety() {
	fmt.Println("\n=== Pointer Safety Demo ===")

	// Creating pointer to local variable (safe in Go due to escape analysis)
	ptr := createPointer()
	fmt.Printf("Pointer from function: %p, value: %d\n", ptr, *ptr)

	// The variable "escapes" to heap, so it's safe to return its pointer
	*ptr = 200
	fmt.Printf("Modified value: %d\n", *ptr)

	// Nil pointer safety (will panic if dereferenced)
	var nilPtr *int
	fmt.Printf("Nil pointer: %p\n", nilPtr)

	// Uncomment to see panic:
	// fmt.Printf("Dereferencing nil pointer: %d\n", *nilPtr) // PANIC!

	// Safe nil pointer check
	if nilPtr != nil {
		fmt.Printf("Value: %d\n", *nilPtr)
	} else {
		fmt.Println("Pointer is nil, cannot dereference")
	}
}

// createPointer creates a pointer to local variable (safe due to escape analysis)
func createPointer() *int {
	localVar := 100
	return &localVar // Variable escapes to heap
}

// DemoStructPointers demonstrates pointers with structs
func DemoStructPointers() {
	fmt.Println("\n=== Struct Pointers Demo ===")

	// Create struct
	p1 := Person{Name: "Bob", Age: 25}
	fmt.Printf("Struct: %+v\n", p1)

	// Pointer to struct
	pPtr := &p1
	fmt.Printf("Pointer: %p\n", pPtr)

	// Access fields through pointer (automatic dereferencing)
	fmt.Printf("Name through pointer: %s\n", pPtr.Name)  // Same as (*pPtr).Name
	fmt.Printf("Age through pointer: %d\n", pPtr.Age)    // Same as (*pPtr).Age

	// Modify through pointer
	pPtr.Age = 26
	pPtr.Name = "Bob Smith"
	fmt.Printf("Modified through pointer: %+v\n", p1)

	// Explicit dereferencing
	(*pPtr).Age = 27
	fmt.Printf("Explicit dereference: %+v\n", p1)

	// Create struct with new
	p2 := new(Person)
	fmt.Printf("Created with new: %+v\n", p2)
	p2.Name = "Charlie"
	p2.Age = 35
	fmt.Printf("After initialization: %+v\n", p2)
}

// DemoSlicePointers demonstrates pointers with slices
func DemoSlicePointers() {
	fmt.Println("\n=== Slice Pointers Demo ===")

	// Slice elements are automatically dereferenced
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Printf("Original slice: %v\n", numbers)

	// Getting pointer to slice elements
	firstPtr := &numbers[0]
	lastPtr := &numbers[len(numbers)-1]

	fmt.Printf("First element: %d, pointer: %p\n", numbers[0], firstPtr)
	fmt.Printf("Last element: %d, pointer: %p\n", numbers[len(numbers)-1], lastPtr)

	// Modify through element pointers
	*firstPtr = 99
	*lastPtr = 88
	fmt.Printf("After modification: %v\n", numbers)

	// Slice of pointers
	var ptrs []*int
	for i := range numbers {
		ptrs = append(ptrs, &numbers[i])
	}

	fmt.Printf("Slice of pointers: %v\n", ptrs)
	fmt.Println("Values through pointers:")
	for i, ptr := range ptrs {
		fmt.Printf("  [%d]: %d\n", i, *ptr)
	}

	// Modify through slice of pointers
	*ptrs[1] = 777
	fmt.Printf("After modifying through pointer slice: %v\n", numbers)
}

// Node represents a linked list node
type Node struct {
	Data int
	Next *Node
}

// DemoLinkedList demonstrates pointers in data structures
func DemoLinkedList() {
	fmt.Println("\n=== Linked List Demo ===")

	// Create linked list: 1 -> 2 -> 3 -> nil
	head := &Node{Data: 1}
	head.Next = &Node{Data: 2}
	head.Next.Next = &Node{Data: 3}

	fmt.Println("Linked list:")
	printList(head)

	// Add element at beginning
	newHead := &Node{Data: 0, Next: head}
	fmt.Println("After adding 0 at beginning:")
	printList(newHead)

	// Add element at end
	addToEnd(newHead, 4)
	fmt.Println("After adding 4 at end:")
	printList(newHead)
}

// printList prints all elements in linked list
func printList(head *Node) {
	current := head
	for current != nil {
		fmt.Printf("  %d", current.Data)
		if current.Next != nil {
			fmt.Print(" -> ")
		}
		current = current.Next
	}
	fmt.Println()
}

// addToEnd adds element at end of linked list
func addToEnd(head *Node, data int) {
	// Find last node
	current := head
	for current.Next != nil {
		current = current.Next
	}
	
	// Add new node
	current.Next = &Node{Data: data}
}

// DemoPointerBestPractices demonstrates pointer best practices
func DemoPointerBestPractices() {
	fmt.Println("\n=== Pointer Best Practices Demo ===")

	// 1. Always check for nil before dereferencing
	var ptr *int
	if ptr != nil {
		fmt.Printf("Value: %d\n", *ptr)
	} else {
		fmt.Println("Pointer is nil - safe check prevented panic")
	}

	// 2. Use pointers for large structs to avoid copying
	largeStruct := LargeStruct{Data: [1000]int{}}
	fmt.Printf("Large struct size: %d bytes\n", len(largeStruct.Data)*8) // rough estimate

	processLargeStructByValue(largeStruct)    // Copies entire struct
	processLargeStructByPointer(&largeStruct) // Passes only pointer (8 bytes on 64-bit)

	// 3. Use pointers when you need to modify the original
	counter := 0
	fmt.Printf("Counter before: %d\n", counter)
	incrementByPointer(&counter)
	fmt.Printf("Counter after: %d\n", counter)

	// 4. Return pointers carefully (Go handles escape analysis)
	valuePtr := getValuePointer()
	fmt.Printf("Value from pointer: %d\n", *valuePtr)
}

// LargeStruct demonstrates when to use pointers for performance
type LargeStruct struct {
	Data [1000]int
}

// processLargeStructByValue copies the entire struct (expensive)
func processLargeStructByValue(ls LargeStruct) {
	// Process struct - entire struct is copied
	fmt.Println("  Processing large struct by value (expensive copy)")
}

// processLargeStructByPointer uses pointer (efficient)
func processLargeStructByPointer(ls *LargeStruct) {
	// Process struct - only pointer is passed
	fmt.Println("  Processing large struct by pointer (efficient)")
}

// incrementByPointer increments value through pointer
func incrementByPointer(x *int) {
	*x++
}

// getValuePointer returns pointer to local variable (safe due to escape analysis)
func getValuePointer() *int {
	value := 42
	return &value // Variable escapes to heap automatically
}

// Helper functions for testing

// SwapValues swaps two integer values using pointers
func SwapValues(a, b *int) {
	*a, *b = *b, *a
}

// GetAddressDifference returns the difference between two pointer addresses (for testing)
func GetAddressDifference(a, b *int) uintptr {
	return uintptr(unsafe.Pointer(a)) - uintptr(unsafe.Pointer(b))
}

// CreateIntPointer creates a pointer to an integer with given value
func CreateIntPointer(value int) *int {
	return &value
}

// IsNilPointer checks if a pointer is nil
func IsNilPointer(ptr *int) bool {
	return ptr == nil
}

// SafeDereference safely dereferences a pointer, returning default value if nil
func SafeDereference(ptr *int, defaultValue int) int {
	if ptr == nil {
		return defaultValue
	}
	return *ptr
}