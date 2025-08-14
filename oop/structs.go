// Package oop demonstrates Go's object-oriented programming concepts including
// structs, methods, and method receivers (value vs pointer receivers).
package oop

import (
	"fmt"
	"math"
	"strings"
)

// DemoBasicStructs demonstrates basic struct creation and usage
func DemoBasicStructs() {
	fmt.Println("\n=== Basic Structs Demo ===")

	// Creating structs with different initialization methods
	p1 := Person{
		Name: "Alice",
		Age:  30,
		Email: "alice@example.com",
	}
	fmt.Printf("Person 1 (struct literal): %+v\n", p1)

	// Positional initialization (not recommended for readability)
	p2 := Person{"Bob", 25, "bob@example.com"}
	fmt.Printf("Person 2 (positional): %+v\n", p2)

	// Partial initialization
	p3 := Person{Name: "Charlie"}
	fmt.Printf("Person 3 (partial): %+v\n", p3)

	// Using new keyword
	p4 := new(Person)
	p4.Name = "Diana"
	p4.Age = 28
	p4.Email = "diana@example.com"
	fmt.Printf("Person 4 (with new): %+v\n", p4)

	// Anonymous struct
	config := struct {
		Host string
		Port int
		SSL  bool
	}{
		Host: "localhost",
		Port: 8080,
		SSL:  true,
	}
	fmt.Printf("Config (anonymous struct): %+v\n", config)
}

// Person represents a person with basic information
type Person struct {
	Name  string
	Age   int
	Email string
}

// DemoValueReceivers demonstrates methods with value receivers
func DemoValueReceivers() {
	fmt.Println("\n=== Value Receivers Demo ===")

	person := Person{
		Name:  "Alice",
		Age:   30,
		Email: "alice@example.com",
	}

	// Methods with value receivers
	fmt.Printf("Original: %+v\n", person)
	fmt.Printf("Display: %s\n", person.Display())
	fmt.Printf("Is Adult: %t\n", person.IsAdult())
	fmt.Printf("Email Domain: %s\n", person.GetEmailDomain())

	// Try to modify through value receiver (won't work)
	person.SetAgeValueReceiver(35)
	fmt.Printf("After SetAgeValueReceiver(35): %+v (unchanged)\n", person)
}

// Display returns a formatted string representation of the person
func (p Person) Display() string {
	return fmt.Sprintf("%s (age: %d, email: %s)", p.Name, p.Age, p.Email)
}

// IsAdult checks if the person is an adult (18 or older)
func (p Person) IsAdult() bool {
	return p.Age >= 18
}

// GetEmailDomain extracts the domain from the email address
func (p Person) GetEmailDomain() string {
	parts := strings.Split(p.Email, "@")
	if len(parts) == 2 {
		return parts[1]
	}
	return ""
}

// SetAgeValueReceiver attempts to set age using value receiver (won't modify original)
func (p Person) SetAgeValueReceiver(newAge int) {
	p.Age = newAge
	fmt.Printf("  Inside SetAgeValueReceiver: %+v\n", p)
}

// DemoPointerReceivers demonstrates methods with pointer receivers
func DemoPointerReceivers() {
	fmt.Println("\n=== Pointer Receivers Demo ===")

	person := Person{
		Name:  "Bob",
		Age:   25,
		Email: "bob@example.com",
	}

	fmt.Printf("Original: %+v\n", person)

	// Methods with pointer receivers (can modify the struct)
	person.SetAge(30)
	fmt.Printf("After SetAge(30): %+v\n", person)

	person.SetEmail("bob.smith@example.com")
	fmt.Printf("After SetEmail: %+v\n", person)

	person.HaveBirthday()
	fmt.Printf("After HaveBirthday: %+v\n", person)

	// Update multiple fields
	person.UpdateInfo("Robert Smith", "robert@company.com")
	fmt.Printf("After UpdateInfo: %+v\n", person)
}

// SetAge sets the person's age using pointer receiver
func (p *Person) SetAge(newAge int) {
	p.Age = newAge
}

// SetEmail sets the person's email using pointer receiver
func (p *Person) SetEmail(newEmail string) {
	p.Email = newEmail
}

// HaveBirthday increments the person's age by 1
func (p *Person) HaveBirthday() {
	p.Age++
	fmt.Printf("  🎂 Happy Birthday %s! Now %d years old.\n", p.Name, p.Age)
}

// UpdateInfo updates multiple fields at once
func (p *Person) UpdateInfo(newName, newEmail string) {
	p.Name = newName
	p.Email = newEmail
}

// DemoMethodsOnDifferentTypes demonstrates methods on various types
func DemoMethodsOnDifferentTypes() {
	fmt.Println("\n=== Methods on Different Types Demo ===")

	// Methods on built-in types (through custom types)
	var name MyString = "John Doe"
	fmt.Printf("MyString: %s\n", name)
	fmt.Printf("Length: %d\n", name.Length())
	fmt.Printf("Uppercase: %s\n", name.Upper())
	fmt.Printf("Word count: %d\n", name.WordCount())

	var number MyInt = 42
	fmt.Printf("\nMyInt: %d\n", number)
	fmt.Printf("Is even: %t\n", number.IsEven())
	fmt.Printf("Square: %d\n", number.Square())
	fmt.Printf("Factorial: %d\n", number.Factorial())

	// Methods on slice types
	numbers := NumberSlice{1, 2, 3, 4, 5}
	fmt.Printf("\nNumberSlice: %v\n", numbers)
	fmt.Printf("Sum: %d\n", numbers.Sum())
	fmt.Printf("Average: %.2f\n", numbers.Average())
	fmt.Printf("Max: %d\n", numbers.Max())
	
	numbers.Scale(2)
	fmt.Printf("After Scale(2): %v\n", numbers)
}

// MyString is a custom string type to demonstrate methods on built-in types
type MyString string

// Length returns the length of the string
func (s MyString) Length() int {
	return len(s)
}

// Upper returns the uppercase version of the string
func (s MyString) Upper() string {
	return strings.ToUpper(string(s))
}

// WordCount returns the number of words in the string
func (s MyString) WordCount() int {
	return len(strings.Fields(string(s)))
}

// MyInt is a custom int type to demonstrate methods on numeric types
type MyInt int

// IsEven checks if the number is even
func (n MyInt) IsEven() bool {
	return n%2 == 0
}

// Square returns the square of the number
func (n MyInt) Square() MyInt {
	return n * n
}

// Factorial returns the factorial of the number
func (n MyInt) Factorial() MyInt {
	if n <= 1 {
		return 1
	}
	result := MyInt(1)
	for i := MyInt(2); i <= n; i++ {
		result *= i
	}
	return result
}

// NumberSlice is a custom slice type to demonstrate methods on slices
type NumberSlice []int

// Sum returns the sum of all numbers in the slice
func (ns NumberSlice) Sum() int {
	total := 0
	for _, num := range ns {
		total += num
	}
	return total
}

// Average returns the average of all numbers in the slice
func (ns NumberSlice) Average() float64 {
	if len(ns) == 0 {
		return 0
	}
	return float64(ns.Sum()) / float64(len(ns))
}

// Max returns the maximum number in the slice
func (ns NumberSlice) Max() int {
	if len(ns) == 0 {
		return 0
	}
	max := ns[0]
	for _, num := range ns {
		if num > max {
			max = num
		}
	}
	return max
}

// Scale multiplies all numbers in the slice by the given factor
func (ns NumberSlice) Scale(factor int) {
	for i := range ns {
		ns[i] *= factor
	}
}

// DemoComplexStructs demonstrates more complex struct patterns
func DemoComplexStructs() {
	fmt.Println("\n=== Complex Structs Demo ===")

	// Nested structs
	emp := Employee{
		Person: Person{
			Name:  "Jane Smith",
			Age:   28,
			Email: "jane@company.com",
		},
		ID:         "EMP001",
		Department: "Engineering",
		Salary:     75000,
		Manager:    nil,
	}

	fmt.Printf("Employee: %+v\n", emp)
	fmt.Printf("Display: %s\n", emp.Display())
	fmt.Printf("Annual bonus: $%.2f\n", emp.CalculateBonus())

	// Create manager
	manager := &Employee{
		Person: Person{
			Name:  "John Manager",
			Age:   40,
			Email: "john@company.com",
		},
		ID:         "MGR001",
		Department: "Engineering",
		Salary:     95000,
		Manager:    nil,
	}

	emp.Manager = manager
	fmt.Printf("Employee with manager: %s reports to %s\n", emp.Name, emp.Manager.Name)

	// Geometric shapes with different implementations
	rect := Rectangle{Width: 5, Height: 3}
	circle := Circle{Radius: 4}

	fmt.Printf("\nRectangle: %+v\n", rect)
	fmt.Printf("Area: %.2f, Perimeter: %.2f\n", rect.Area(), rect.Perimeter())

	fmt.Printf("\nCircle: %+v\n", circle)
	fmt.Printf("Area: %.2f, Circumference: %.2f\n", circle.Area(), circle.Circumference())
}

// Employee represents an employee with embedded Person struct
type Employee struct {
	Person     // Embedded struct
	ID         string
	Department string
	Salary     float64
	Manager    *Employee
}

// Display returns employee information (overrides Person.Display)
func (e Employee) Display() string {
	return fmt.Sprintf("%s (ID: %s, Dept: %s, Salary: $%.0f)", 
		e.Person.Display(), e.ID, e.Department, e.Salary)
}

// CalculateBonus calculates annual bonus (10% of salary)
func (e Employee) CalculateBonus() float64 {
	return e.Salary * 0.10
}

// GiveRaise gives a percentage raise to the employee
func (e *Employee) GiveRaise(percentage float64) {
	e.Salary *= (1 + percentage/100)
	fmt.Printf("  💰 %s received a %.1f%% raise! New salary: $%.0f\n", 
		e.Name, percentage, e.Salary)
}

// Rectangle represents a rectangle shape
type Rectangle struct {
	Width, Height float64
}

// Area calculates the area of the rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Perimeter calculates the perimeter of the rectangle
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Circle represents a circle shape
type Circle struct {
	Radius float64
}

// Area calculates the area of the circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Perimeter calculates the perimeter (circumference) of the circle
func (c Circle) Perimeter() float64 {
	return c.Circumference()
}

// Circumference calculates the circumference of the circle
func (c Circle) Circumference() float64 {
	return 2 * math.Pi * c.Radius
}

// DemoMethodSets demonstrates method sets and receiver types
func DemoMethodSets() {
	fmt.Println("\n=== Method Sets Demo ===")

	// Value and pointer method accessibility
	p1 := Person{Name: "Alice", Age: 30, Email: "alice@example.com"}
	p2 := &Person{Name: "Bob", Age: 25, Email: "bob@example.com"}

	fmt.Println("Value receiver methods work on both value and pointer:")
	fmt.Printf("p1.Display(): %s\n", p1.Display())
	fmt.Printf("p2.Display(): %s\n", p2.Display())

	fmt.Println("\nPointer receiver methods work on both value and pointer:")
	p1.SetAge(31) // Go automatically takes address: (&p1).SetAge(31)
	fmt.Printf("p1 after SetAge: %+v\n", p1)

	p2.SetAge(26) // Direct call on pointer
	fmt.Printf("p2 after SetAge: %+v\n", p2)

	// Demonstrating the difference
	fmt.Println("\nDifference between value and pointer receivers:")
	
	// Value receiver - copy semantics
	p3 := Person{Name: "Charlie", Age: 20, Email: "charlie@example.com"}
	p3.SetAgeValueReceiver(99) // Doesn't modify original
	fmt.Printf("After value receiver: %+v\n", p3)

	// Pointer receiver - reference semantics
	p3.SetAge(99) // Modifies original
	fmt.Printf("After pointer receiver: %+v\n", p3)
}

// DemoStructComparison demonstrates struct comparison and copying
func DemoStructComparison() {
	fmt.Println("\n=== Struct Comparison and Copying Demo ===")

	p1 := Person{Name: "Alice", Age: 30, Email: "alice@example.com"}
	p2 := Person{Name: "Alice", Age: 30, Email: "alice@example.com"}
	p3 := Person{Name: "Bob", Age: 25, Email: "bob@example.com"}

	// Struct comparison (all fields must be comparable)
	fmt.Printf("p1 == p2: %t\n", p1 == p2)
	fmt.Printf("p1 == p3: %t\n", p1 == p3)

	// Struct copying
	p4 := p1 // Copy
	p4.Age = 31
	fmt.Printf("Original p1: %+v\n", p1)
	fmt.Printf("Copied p4: %+v\n", p4)

	// Pointer to struct
	p5 := &p1
	p5.Age = 32
	fmt.Printf("p1 after modifying through pointer: %+v\n", p1)

	// Struct with non-comparable fields (slice, map, function)
	type Config struct {
		Name     string
		Values   []string    // slice - not comparable
		Settings map[string]int // map - not comparable
	}

	config1 := Config{
		Name:     "app",
		Values:   []string{"a", "b"},
		Settings: map[string]int{"x": 1},
	}

	config2 := Config{
		Name:     "app",
		Values:   []string{"a", "b"},
		Settings: map[string]int{"x": 1},
	}

	// This would cause compile error: config1 == config2
	fmt.Printf("Config 1: %+v\n", config1)
	fmt.Printf("Config 2: %+v\n", config2)
	fmt.Println("Note: Structs with slices/maps are not comparable with ==")
}

// Helper functions for creating structs

// NewPerson creates a new Person with validation
func NewPerson(name, email string, age int) (*Person, error) {
	if name == "" {
		return nil, fmt.Errorf("name cannot be empty")
	}
	if age < 0 {
		return nil, fmt.Errorf("age cannot be negative")
	}
	if !strings.Contains(email, "@") {
		return nil, fmt.Errorf("invalid email format")
	}

	return &Person{
		Name:  name,
		Age:   age,
		Email: email,
	}, nil
}

// NewEmployee creates a new Employee with validation
func NewEmployee(name, email string, age int, id, department string, salary float64) (*Employee, error) {
	person, err := NewPerson(name, email, age)
	if err != nil {
		return nil, err
	}

	if id == "" {
		return nil, fmt.Errorf("employee ID cannot be empty")
	}
	if salary < 0 {
		return nil, fmt.Errorf("salary cannot be negative")
	}

	return &Employee{
		Person:     *person,
		ID:         id,
		Department: department,
		Salary:     salary,
	}, nil
}

// DemoConstructorPatterns demonstrates constructor patterns in Go
func DemoConstructorPatterns() {
	fmt.Println("\n=== Constructor Patterns Demo ===")

	// Using constructor functions
	person, err := NewPerson("Alice Smith", "alice@example.com", 30)
	if err != nil {
		fmt.Printf("Error creating person: %v\n", err)
	} else {
		fmt.Printf("Created person: %+v\n", person)
	}

	// Error case
	_, err = NewPerson("", "invalid-email", -5)
	if err != nil {
		fmt.Printf("Expected error: %v\n", err)
	}

	// Creating employee
	emp, err := NewEmployee("Bob Johnson", "bob@company.com", 28, "EMP002", "Sales", 65000)
	if err != nil {
		fmt.Printf("Error creating employee: %v\n", err)
	} else {
		fmt.Printf("Created employee: %s\n", emp.Display())
	}

	// Functional options pattern (alternative constructor style)
	opts := EmployeeOptions{
		Name:       "Charlie Brown",
		Email:      "charlie@company.com",
		Age:        35,
		ID:         "EMP003",
		Department: "Marketing",
		Salary:     70000,
	}

	emp2 := NewEmployeeWithOptions(opts)
	fmt.Printf("Employee with options: %s\n", emp2.Display())
}

// EmployeeOptions represents options for creating an employee
type EmployeeOptions struct {
	Name       string
	Email      string
	Age        int
	ID         string
	Department string
	Salary     float64
	Manager    *Employee
}

// NewEmployeeWithOptions creates an employee using options pattern
func NewEmployeeWithOptions(opts EmployeeOptions) *Employee {
	return &Employee{
		Person: Person{
			Name:  opts.Name,
			Age:   opts.Age,
			Email: opts.Email,
		},
		ID:         opts.ID,
		Department: opts.Department,
		Salary:     opts.Salary,
		Manager:    opts.Manager,
	}
}