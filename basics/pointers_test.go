package basics

import (
	"testing"
	"unsafe"
)

func TestDemoBasicPointers(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("DemoBasicPointers panicked: %v", r)
		}
	}()
	
	DemoBasicPointers()
}

func TestDemoPointerTypes(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("DemoPointerTypes panicked: %v", r)
		}
	}()
	
	DemoPointerTypes()
}

func TestDemoPassByValue(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("DemoPassByValue panicked: %v", r)
		}
	}()
	
	DemoPassByValue()
}

func TestDemoPassByReference(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("DemoPassByReference panicked: %v", r)
		}
	}()
	
	DemoPassByReference()
}

func TestDemoPointerArithmetic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("DemoPointerArithmetic panicked: %v", r)
		}
	}()
	
	DemoPointerArithmetic()
}

func TestDemoPointerSafety(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("DemoPointerSafety panicked: %v", r)
		}
	}()
	
	DemoPointerSafety()
}

func TestDemoStructPointers(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("DemoStructPointers panicked: %v", r)
		}
	}()
	
	DemoStructPointers()
}

func TestDemoSlicePointers(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("DemoSlicePointers panicked: %v", r)
		}
	}()
	
	DemoSlicePointers()
}

func TestDemoLinkedList(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("DemoLinkedList panicked: %v", r)
		}
	}()
	
	DemoLinkedList()
}

func TestDemoPointerBestPractices(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("DemoPointerBestPractices panicked: %v", r)
		}
	}()
	
	DemoPointerBestPractices()
}

func TestBasicPointerOperations(t *testing.T) {
	// Test pointer creation and dereferencing
	x := 42
	p := &x
	
	if *p != 42 {
		t.Errorf("Expected *p to be 42, got %d", *p)
	}
	
	// Test modification through pointer
	*p = 100
	if x != 100 {
		t.Errorf("Expected x to be 100 after modification through pointer, got %d", x)
	}
	
	// Test pointer comparison
	y := 100
	q := &y
	r := &x
	
	if p != r {
		t.Errorf("Pointers to same variable should be equal")
	}
	
	if p == q {
		t.Errorf("Pointers to different variables should not be equal")
	}
}

func TestNilPointers(t *testing.T) {
	var p *int
	
	if p != nil {
		t.Errorf("Uninitialized pointer should be nil")
	}
	
	// Test nil pointer comparison
	var q *int
	if p != q {
		t.Errorf("Two nil pointers should be equal")
	}
}

func TestPointerToPointer(t *testing.T) {
	x := 42
	p := &x
	pp := &p
	
	if **pp != 42 {
		t.Errorf("Expected **pp to be 42, got %d", **pp)
	}
	
	// Modify through pointer to pointer
	**pp = 99
	if x != 99 {
		t.Errorf("Expected x to be 99 after modification through pointer to pointer, got %d", x)
	}
}

func TestModifyByPointer(t *testing.T) {
	value := 10
	modifyByPointer(&value)
	
	if value != 999 {
		t.Errorf("Expected value to be 999 after modifyByPointer, got %d", value)
	}
}

func TestPersonStructPointers(t *testing.T) {
	person := Person{Name: "Alice", Age: 30}
	
	// Test modification by value (should not change original)
	originalAge := person.Age
	modifyPersonByValue(person)
	if person.Age != originalAge {
		t.Errorf("modifyPersonByValue should not change original struct")
	}
	
	// Test modification by pointer (should change original)
	modifyPersonByPointer(&person)
	if person.Age == originalAge {
		t.Errorf("modifyPersonByPointer should change original struct")
	}
	if person.Name != "Alice Smith" {
		t.Errorf("Expected name to be 'Alice Smith', got '%s'", person.Name)
	}
}

func TestCreatePointer(t *testing.T) {
	ptr := createPointer()
	
	if ptr == nil {
		t.Errorf("createPointer should return non-nil pointer")
	}
	
	if *ptr != 100 {
		t.Errorf("Expected *ptr to be 100, got %d", *ptr)
	}
	
	// Should be safe to modify
	*ptr = 200
	if *ptr != 200 {
		t.Errorf("Expected *ptr to be 200 after modification, got %d", *ptr)
	}
}

func TestStructPointerAccess(t *testing.T) {
	person := Person{Name: "Bob", Age: 25}
	pPtr := &person
	
	// Test automatic dereferencing
	if pPtr.Name != "Bob" {
		t.Errorf("Expected pPtr.Name to be 'Bob', got '%s'", pPtr.Name)
	}
	
	if pPtr.Age != 25 {
		t.Errorf("Expected pPtr.Age to be 25, got %d", pPtr.Age)
	}
	
	// Test explicit dereferencing
	if (*pPtr).Name != "Bob" {
		t.Errorf("Expected (*pPtr).Name to be 'Bob', got '%s'", (*pPtr).Name)
	}
	
	// Test modification through pointer
	pPtr.Age = 26
	if person.Age != 26 {
		t.Errorf("Expected person.Age to be 26 after modification through pointer, got %d", person.Age)
	}
}

func TestNewStruct(t *testing.T) {
	person := new(Person)
	
	if person == nil {
		t.Errorf("new(Person) should return non-nil pointer")
	}
	
	// Test zero values
	if person.Name != "" {
		t.Errorf("Expected person.Name to be empty string, got '%s'", person.Name)
	}
	
	if person.Age != 0 {
		t.Errorf("Expected person.Age to be 0, got %d", person.Age)
	}
	
	// Test assignment
	person.Name = "Charlie"
	person.Age = 35
	
	if person.Name != "Charlie" {
		t.Errorf("Expected person.Name to be 'Charlie', got '%s'", person.Name)
	}
	
	if person.Age != 35 {
		t.Errorf("Expected person.Age to be 35, got %d", person.Age)
	}
}

func TestSliceElementPointers(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	
	// Get pointers to elements
	firstPtr := &numbers[0]
	lastPtr := &numbers[len(numbers)-1]
	
	if *firstPtr != 1 {
		t.Errorf("Expected *firstPtr to be 1, got %d", *firstPtr)
	}
	
	if *lastPtr != 5 {
		t.Errorf("Expected *lastPtr to be 5, got %d", *lastPtr)
	}
	
	// Modify through pointers
	*firstPtr = 99
	*lastPtr = 88
	
	if numbers[0] != 99 {
		t.Errorf("Expected numbers[0] to be 99, got %d", numbers[0])
	}
	
	if numbers[4] != 88 {
		t.Errorf("Expected numbers[4] to be 88, got %d", numbers[4])
	}
}

func TestLinkedListOperations(t *testing.T) {
	// Create simple linked list: 1 -> 2 -> nil
	head := &Node{Data: 1}
	head.Next = &Node{Data: 2}
	
	// Test structure
	if head.Data != 1 {
		t.Errorf("Expected head.Data to be 1, got %d", head.Data)
	}
	
	if head.Next == nil {
		t.Errorf("Expected head.Next to be non-nil")
	}
	
	if head.Next.Data != 2 {
		t.Errorf("Expected head.Next.Data to be 2, got %d", head.Next.Data)
	}
	
	if head.Next.Next != nil {
		t.Errorf("Expected head.Next.Next to be nil")
	}
	
	// Test adding to end
	addToEnd(head, 3)
	
	if head.Next.Next == nil {
		t.Errorf("Expected third node to exist after addToEnd")
	}
	
	if head.Next.Next.Data != 3 {
		t.Errorf("Expected third node data to be 3, got %d", head.Next.Next.Data)
	}
}

func TestSwapValues(t *testing.T) {
	a := 10
	b := 20
	
	SwapValues(&a, &b)
	
	if a != 20 {
		t.Errorf("Expected a to be 20 after swap, got %d", a)
	}
	
	if b != 10 {
		t.Errorf("Expected b to be 10 after swap, got %d", b)
	}
}

func TestCreateIntPointer(t *testing.T) {
	ptr := CreateIntPointer(42)
	
	if ptr == nil {
		t.Errorf("CreateIntPointer should return non-nil pointer")
	}
	
	if *ptr != 42 {
		t.Errorf("Expected *ptr to be 42, got %d", *ptr)
	}
}

func TestIsNilPointer(t *testing.T) {
	var nilPtr *int
	validPtr := CreateIntPointer(10)
	
	if !IsNilPointer(nilPtr) {
		t.Errorf("IsNilPointer should return true for nil pointer")
	}
	
	if IsNilPointer(validPtr) {
		t.Errorf("IsNilPointer should return false for valid pointer")
	}
}

func TestSafeDereference(t *testing.T) {
	// Test with nil pointer
	var nilPtr *int
	result := SafeDereference(nilPtr, 999)
	if result != 999 {
		t.Errorf("SafeDereference with nil pointer: expected 999, got %d", result)
	}
	
	// Test with valid pointer
	value := 42
	validPtr := &value
	result = SafeDereference(validPtr, 999)
	if result != 42 {
		t.Errorf("SafeDereference with valid pointer: expected 42, got %d", result)
	}
}

func TestPointerAddresses(t *testing.T) {
	x := 10
	y := 20
	
	ptrX := &x
	ptrY := &y
	
	// Pointers to different variables should have different addresses
	if ptrX == ptrY {
		t.Errorf("Pointers to different variables should not be equal")
	}
	
	// Multiple pointers to same variable should be equal
	anotherPtrX := &x
	if ptrX != anotherPtrX {
		t.Errorf("Multiple pointers to same variable should be equal")
	}
}

func TestPointerSizes(t *testing.T) {
	var intPtr *int
	var stringPtr *string
	var structPtr *Person
	
	// All pointers should have the same size (8 bytes on 64-bit systems)
	intPtrSize := unsafe.Sizeof(intPtr)
	stringPtrSize := unsafe.Sizeof(stringPtr)
	structPtrSize := unsafe.Sizeof(structPtr)
	
	if intPtrSize != stringPtrSize {
		t.Errorf("All pointer types should have same size")
	}
	
	if stringPtrSize != structPtrSize {
		t.Errorf("All pointer types should have same size")
	}
}