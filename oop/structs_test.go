package oop

import (
	"strings"
	"testing"
)

func TestDemoBasicStructs(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("DemoBasicStructs panicked: %v", r)
		}
	}()
	
	DemoBasicStructs()
}

func TestDemoValueReceivers(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("DemoValueReceivers panicked: %v", r)
		}
	}()
	
	DemoValueReceivers()
}

func TestDemoPointerReceivers(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("DemoPointerReceivers panicked: %v", r)
		}
	}()
	
	DemoPointerReceivers()
}

func TestDemoMethodsOnDifferentTypes(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("DemoMethodsOnDifferentTypes panicked: %v", r)
		}
	}()
	
	DemoMethodsOnDifferentTypes()
}

func TestDemoComplexStructs(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("DemoComplexStructs panicked: %v", r)
		}
	}()
	
	DemoComplexStructs()
}

func TestDemoMethodSets(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("DemoMethodSets panicked: %v", r)
		}
	}()
	
	DemoMethodSets()
}

func TestDemoStructComparison(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("DemoStructComparison panicked: %v", r)
		}
	}()
	
	DemoStructComparison()
}

func TestDemoConstructorPatterns(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("DemoConstructorPatterns panicked: %v", r)
		}
	}()
	
	DemoConstructorPatterns()
}

func TestPersonStruct(t *testing.T) {
	person := Person{
		Name:  "Alice",
		Age:   30,
		Email: "alice@example.com",
	}
	
	if person.Name != "Alice" {
		t.Errorf("Expected name 'Alice', got '%s'", person.Name)
	}
	
	if person.Age != 30 {
		t.Errorf("Expected age 30, got %d", person.Age)
	}
	
	if person.Email != "alice@example.com" {
		t.Errorf("Expected email 'alice@example.com', got '%s'", person.Email)
	}
}

func TestPersonDisplay(t *testing.T) {
	person := Person{Name: "Bob", Age: 25, Email: "bob@example.com"}
	display := person.Display()
	
	expected := "Bob (age: 25, email: bob@example.com)"
	if display != expected {
		t.Errorf("Expected display '%s', got '%s'", expected, display)
	}
}

func TestPersonIsAdult(t *testing.T) {
	adult := Person{Name: "Adult", Age: 25}
	child := Person{Name: "Child", Age: 15}
	
	if !adult.IsAdult() {
		t.Errorf("25-year-old should be adult")
	}
	
	if child.IsAdult() {
		t.Errorf("15-year-old should not be adult")
	}
}

func TestPersonGetEmailDomain(t *testing.T) {
	person := Person{Email: "user@example.com"}
	domain := person.GetEmailDomain()
	
	if domain != "example.com" {
		t.Errorf("Expected domain 'example.com', got '%s'", domain)
	}
	
	// Test invalid email
	invalidPerson := Person{Email: "invalid-email"}
	invalidDomain := invalidPerson.GetEmailDomain()
	
	if invalidDomain != "" {
		t.Errorf("Expected empty domain for invalid email, got '%s'", invalidDomain)
	}
}

func TestValueVsPointerReceivers(t *testing.T) {
	person := Person{Name: "Test", Age: 20, Email: "test@example.com"}
	
	// Test value receiver (should not modify original)
	originalAge := person.Age
	person.SetAgeValueReceiver(25)
	
	if person.Age != originalAge {
		t.Errorf("Value receiver should not modify original struct")
	}
	
	// Test pointer receiver (should modify original)
	person.SetAge(25)
	
	if person.Age != 25 {
		t.Errorf("Pointer receiver should modify original struct")
	}
}

func TestPersonPointerMethods(t *testing.T) {
	person := Person{Name: "Test", Age: 20, Email: "old@example.com"}
	
	// Test SetEmail
	person.SetEmail("new@example.com")
	if person.Email != "new@example.com" {
		t.Errorf("SetEmail should update email")
	}
	
	// Test HaveBirthday
	originalAge := person.Age
	person.HaveBirthday()
	if person.Age != originalAge+1 {
		t.Errorf("HaveBirthday should increment age by 1")
	}
	
	// Test UpdateInfo
	person.UpdateInfo("New Name", "updated@example.com")
	if person.Name != "New Name" {
		t.Errorf("UpdateInfo should update name")
	}
	if person.Email != "updated@example.com" {
		t.Errorf("UpdateInfo should update email")
	}
}

func TestMyStringMethods(t *testing.T) {
	str := MyString("Hello World")
	
	if str.Length() != 11 {
		t.Errorf("Expected length 11, got %d", str.Length())
	}
	
	if str.Upper() != "HELLO WORLD" {
		t.Errorf("Expected 'HELLO WORLD', got '%s'", str.Upper())
	}
	
	if str.WordCount() != 2 {
		t.Errorf("Expected word count 2, got %d", str.WordCount())
	}
}

func TestMyIntMethods(t *testing.T) {
	even := MyInt(4)
	odd := MyInt(5)
	
	if !even.IsEven() {
		t.Errorf("4 should be even")
	}
	
	if odd.IsEven() {
		t.Errorf("5 should not be even")
	}
	
	if even.Square() != 16 {
		t.Errorf("4 squared should be 16, got %d", even.Square())
	}
	
	factorial5 := MyInt(5).Factorial()
	if factorial5 != 120 {
		t.Errorf("5! should be 120, got %d", factorial5)
	}
	
	factorial0 := MyInt(0).Factorial()
	if factorial0 != 1 {
		t.Errorf("0! should be 1, got %d", factorial0)
	}
}

func TestNumberSliceMethods(t *testing.T) {
	numbers := NumberSlice{1, 2, 3, 4, 5}
	
	if numbers.Sum() != 15 {
		t.Errorf("Sum should be 15, got %d", numbers.Sum())
	}
	
	if numbers.Average() != 3.0 {
		t.Errorf("Average should be 3.0, got %.2f", numbers.Average())
	}
	
	if numbers.Max() != 5 {
		t.Errorf("Max should be 5, got %d", numbers.Max())
	}
	
	// Test Scale
	numbers.Scale(2)
	expected := NumberSlice{2, 4, 6, 8, 10}
	for i, v := range expected {
		if numbers[i] != v {
			t.Errorf("After Scale(2), expected %v, got %v", expected, numbers)
			break
		}
	}
	
	// Test empty slice
	empty := NumberSlice{}
	if empty.Sum() != 0 {
		t.Errorf("Empty slice sum should be 0")
	}
	if empty.Average() != 0 {
		t.Errorf("Empty slice average should be 0")
	}
	if empty.Max() != 0 {
		t.Errorf("Empty slice max should be 0")
	}
}

func TestEmployeeStruct(t *testing.T) {
	emp := Employee{
		Person: Person{
			Name:  "John",
			Age:   30,
			Email: "john@company.com",
		},
		ID:         "EMP001",
		Department: "Engineering",
		Salary:     80000,
	}
	
	if emp.Name != "John" {
		t.Errorf("Expected name 'John', got '%s'", emp.Name)
	}
	
	if emp.ID != "EMP001" {
		t.Errorf("Expected ID 'EMP001', got '%s'", emp.ID)
	}
	
	bonus := emp.CalculateBonus()
	expectedBonus := 8000.0
	if bonus != expectedBonus {
		t.Errorf("Expected bonus %.2f, got %.2f", expectedBonus, bonus)
	}
}

func TestEmployeeDisplay(t *testing.T) {
	emp := Employee{
		Person: Person{Name: "Jane", Age: 28, Email: "jane@company.com"},
		ID:     "EMP002",
		Department: "Marketing",
		Salary: 70000,
	}
	
	display := emp.Display()
	if !strings.Contains(display, "Jane") {
		t.Errorf("Display should contain name")
	}
	if !strings.Contains(display, "EMP002") {
		t.Errorf("Display should contain ID")
	}
	if !strings.Contains(display, "Marketing") {
		t.Errorf("Display should contain department")
	}
}

func TestEmployeeGiveRaise(t *testing.T) {
	emp := Employee{
		Person: Person{Name: "Test", Age: 30, Email: "test@company.com"},
		Salary: 50000,
	}
	
	emp.GiveRaise(10) // 10% raise
	expectedSalary := 55000.0
	
	if emp.Salary < expectedSalary-0.01 || emp.Salary > expectedSalary+0.01 {
		t.Errorf("Expected salary %.2f after 10%% raise, got %.2f", expectedSalary, emp.Salary)
	}
}

func TestRectangleMethods(t *testing.T) {
	rect := Rectangle{Width: 4, Height: 3}
	
	expectedArea := 12.0
	if rect.Area() != expectedArea {
		t.Errorf("Expected area %.2f, got %.2f", expectedArea, rect.Area())
	}
	
	expectedPerimeter := 14.0
	if rect.Perimeter() != expectedPerimeter {
		t.Errorf("Expected perimeter %.2f, got %.2f", expectedPerimeter, rect.Perimeter())
	}
}

func TestCircleMethods(t *testing.T) {
	circle := Circle{Radius: 2}
	
	// Area = π * r² = π * 4 ≈ 12.566
	area := circle.Area()
	if area < 12.5 || area > 12.6 {
		t.Errorf("Expected area around 12.57, got %.2f", area)
	}
	
	// Circumference = 2 * π * r = 2 * π * 2 ≈ 12.566
	circumference := circle.Circumference()
	if circumference < 12.5 || circumference > 12.6 {
		t.Errorf("Expected circumference around 12.57, got %.2f", circumference)
	}
}

func TestNewPerson(t *testing.T) {
	// Valid person
	person, err := NewPerson("Alice", "alice@example.com", 25)
	if err != nil {
		t.Errorf("Should create valid person without error: %v", err)
	}
	if person.Name != "Alice" {
		t.Errorf("Expected name 'Alice', got '%s'", person.Name)
	}
	
	// Invalid name
	_, err = NewPerson("", "alice@example.com", 25)
	if err == nil {
		t.Errorf("Should return error for empty name")
	}
	
	// Invalid age
	_, err = NewPerson("Alice", "alice@example.com", -5)
	if err == nil {
		t.Errorf("Should return error for negative age")
	}
	
	// Invalid email
	_, err = NewPerson("Alice", "invalid-email", 25)
	if err == nil {
		t.Errorf("Should return error for invalid email")
	}
}

func TestNewEmployee(t *testing.T) {
	// Valid employee
	emp, err := NewEmployee("Bob", "bob@company.com", 30, "EMP001", "Engineering", 75000)
	if err != nil {
		t.Errorf("Should create valid employee without error: %v", err)
	}
	if emp.ID != "EMP001" {
		t.Errorf("Expected ID 'EMP001', got '%s'", emp.ID)
	}
	
	// Invalid person data
	_, err = NewEmployee("", "bob@company.com", 30, "EMP001", "Engineering", 75000)
	if err == nil {
		t.Errorf("Should return error for invalid person data")
	}
	
	// Invalid employee ID
	_, err = NewEmployee("Bob", "bob@company.com", 30, "", "Engineering", 75000)
	if err == nil {
		t.Errorf("Should return error for empty employee ID")
	}
	
	// Invalid salary
	_, err = NewEmployee("Bob", "bob@company.com", 30, "EMP001", "Engineering", -1000)
	if err == nil {
		t.Errorf("Should return error for negative salary")
	}
}

func TestNewEmployeeWithOptions(t *testing.T) {
	opts := EmployeeOptions{
		Name:       "Charlie",
		Email:      "charlie@company.com",
		Age:        35,
		ID:         "EMP003",
		Department: "Sales",
		Salary:     65000,
	}
	
	emp := NewEmployeeWithOptions(opts)
	
	if emp.Name != "Charlie" {
		t.Errorf("Expected name 'Charlie', got '%s'", emp.Name)
	}
	
	if emp.ID != "EMP003" {
		t.Errorf("Expected ID 'EMP003', got '%s'", emp.ID)
	}
	
	if emp.Department != "Sales" {
		t.Errorf("Expected department 'Sales', got '%s'", emp.Department)
	}
}

func TestStructComparison(t *testing.T) {
	p1 := Person{Name: "Alice", Age: 30, Email: "alice@example.com"}
	p2 := Person{Name: "Alice", Age: 30, Email: "alice@example.com"}
	p3 := Person{Name: "Bob", Age: 25, Email: "bob@example.com"}
	
	if p1 != p2 {
		t.Errorf("Identical structs should be equal")
	}
	
	if p1 == p3 {
		t.Errorf("Different structs should not be equal")
	}
}

func TestStructCopying(t *testing.T) {
	original := Person{Name: "Alice", Age: 30, Email: "alice@example.com"}
	copy := original
	
	copy.Age = 31
	
	if original.Age != 30 {
		t.Errorf("Original struct should not be modified when copy is changed")
	}
	
	if copy.Age != 31 {
		t.Errorf("Copy should be modified independently")
	}
}

func TestMethodSetAccess(t *testing.T) {
	// Value type should access both value and pointer receiver methods
	p1 := Person{Name: "Alice", Age: 30, Email: "alice@example.com"}
	
	// Value receiver method
	display := p1.Display()
	if display == "" {
		t.Errorf("Should be able to call value receiver method on value")
	}
	
	// Pointer receiver method (Go automatically takes address)
	p1.SetAge(31)
	if p1.Age != 31 {
		t.Errorf("Should be able to call pointer receiver method on value")
	}
	
	// Pointer type should access both value and pointer receiver methods
	p2 := &Person{Name: "Bob", Age: 25, Email: "bob@example.com"}
	
	// Value receiver method
	display2 := p2.Display()
	if display2 == "" {
		t.Errorf("Should be able to call value receiver method on pointer")
	}
	
	// Pointer receiver method
	p2.SetAge(26)
	if p2.Age != 26 {
		t.Errorf("Should be able to call pointer receiver method on pointer")
	}
}