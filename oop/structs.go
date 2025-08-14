// Package oop demonstrates object-oriented programming concepts in Go
package oop

import "fmt"

// Rectangle represents a geometric rectangle
type Rectangle struct {
	Width  float64
	Height float64
}

// Area calculates the area of the rectangle (value receiver)
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Perimeter calculates the perimeter of the rectangle (value receiver)
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Scale scales the rectangle by a factor (pointer receiver)
func (r *Rectangle) Scale(factor float64) {
	r.Width *= factor
	r.Height *= factor
}

// String returns a string representation of the rectangle
func (r Rectangle) String() string {
	return fmt.Sprintf("Rectangle(%.2f x %.2f)", r.Width, r.Height)
}

// Circle represents a geometric circle
type Circle struct {
	Radius float64
}

// Area calculates the area of the circle
func (c Circle) Area() float64 {
	return 3.14159 * c.Radius * c.Radius
}

// Perimeter calculates the perimeter (circumference) of the circle
func (c Circle) Perimeter() float64 {
	return 2 * 3.14159 * c.Radius
}

// String returns a string representation of the circle
func (c Circle) String() string {
	return fmt.Sprintf("Circle(radius: %.2f)", c.Radius)
}

// DemonstrateStructsAndMethods shows struct and method usage
func DemonstrateStructsAndMethods() {
	fmt.Println("=== Structs and Methods ===")
	
	// Create structs
	rect := Rectangle{Width: 5.0, Height: 3.0}
	circle := Circle{Radius: 2.5}
	
	fmt.Printf("Rectangle: %s\n", rect)
	fmt.Printf("  Area: %.2f\n", rect.Area())
	fmt.Printf("  Perimeter: %.2f\n", rect.Perimeter())
	
	fmt.Printf("\nCircle: %s\n", circle)
	fmt.Printf("  Area: %.2f\n", circle.Area())
	fmt.Printf("  Perimeter: %.2f\n", circle.Perimeter())
	
	// Demonstrate pointer receiver
	fmt.Printf("\nBefore scaling: %s\n", rect)
	rect.Scale(2.0)
	fmt.Printf("After scaling by 2: %s\n", rect)
	fmt.Printf("New area: %.2f\n", rect.Area())
	
	// Method on pointer
	rectPtr := &Rectangle{Width: 4.0, Height: 6.0}
	fmt.Printf("\nPointer rectangle: %s\n", rectPtr)
	rectPtr.Scale(0.5)
	fmt.Printf("After scaling: %s\n", rectPtr)
}

// Person represents a person with basic information
type Person struct {
	FirstName string
	LastName  string
	Age       int
	Email     string
}

// FullName returns the full name of the person
func (p Person) FullName() string {
	return p.FirstName + " " + p.LastName
}

// IsAdult checks if the person is an adult (18 or older)
func (p Person) IsAdult() bool {
	return p.Age >= 18
}

// UpdateEmail updates the person's email (pointer receiver to modify)
func (p *Person) UpdateEmail(newEmail string) {
	p.Email = newEmail
}

// HaveBirthday increments the person's age
func (p *Person) HaveBirthday() {
	p.Age++
}

// String returns a string representation of the person
func (p Person) String() string {
	return fmt.Sprintf("%s (%d years old, %s)", p.FullName(), p.Age, p.Email)
}

// Employee embeds Person and adds work-related fields
type Employee struct {
	Person     // Embedded struct
	EmployeeID string
	Department string
	Salary     float64
}

// GetInfo returns employee information
func (e Employee) GetInfo() string {
	return fmt.Sprintf("Employee %s - %s, Department: %s, Salary: $%.2f", 
		e.EmployeeID, e.FullName(), e.Department, e.Salary)
}

// Promote gives the employee a raise
func (e *Employee) Promote(salaryIncrease float64) {
	e.Salary += salaryIncrease
}

// DemonstrateEmbedding shows struct embedding (composition)
func DemonstrateEmbedding() {
	fmt.Println("\n=== Struct Embedding ===")
	
	// Create a person
	person := Person{
		FirstName: "John",
		LastName:  "Doe",
		Age:       28,
		Email:     "john.doe@example.com",
	}
	
	fmt.Printf("Person: %s\n", person)
	fmt.Printf("Is adult: %t\n", person.IsAdult())
	
	// Create an employee (embeds Person)
	employee := Employee{
		Person: Person{
			FirstName: "Jane",
			LastName:  "Smith",
			Age:       32,
			Email:     "jane.smith@company.com",
		},
		EmployeeID: "EMP001",
		Department: "Engineering",
		Salary:     75000.0,
	}
	
	fmt.Printf("\nEmployee: %s\n", employee.GetInfo())
	
	// Access embedded fields directly
	fmt.Printf("Employee full name: %s\n", employee.FullName()) // Method from embedded Person
	fmt.Printf("Employee age: %d\n", employee.Age)             // Field from embedded Person
	
	// Modify embedded fields
	employee.HaveBirthday()
	employee.UpdateEmail("jane.smith@newcompany.com")
	employee.Promote(5000.0)
	
	fmt.Printf("After updates: %s\n", employee.GetInfo())
	fmt.Printf("New age: %d\n", employee.Age)
	fmt.Printf("New email: %s\n", employee.Email)
}

// Manager embeds Employee and adds management fields
type Manager struct {
	Employee    // Embedded struct
	TeamSize    int
	TeamMembers []string
}

// AddTeamMember adds a team member
func (m *Manager) AddTeamMember(member string) {
	m.TeamMembers = append(m.TeamMembers, member)
	m.TeamSize = len(m.TeamMembers)
}

// GetTeamInfo returns information about the manager's team
func (m Manager) GetTeamInfo() string {
	return fmt.Sprintf("Manager %s leads a team of %d: %v", 
		m.FullName(), m.TeamSize, m.TeamMembers)
}

// DemonstrateNestedEmbedding shows multiple levels of embedding
func DemonstrateNestedEmbedding() {
	fmt.Println("\n=== Nested Embedding ===")
	
	manager := Manager{
		Employee: Employee{
			Person: Person{
				FirstName: "Alice",
				LastName:  "Johnson",
				Age:       35,
				Email:     "alice.johnson@company.com",
			},
			EmployeeID: "MGR001",
			Department: "Engineering",
			Salary:     95000.0,
		},
		TeamSize:    0,
		TeamMembers: []string{},
	}
	
	fmt.Printf("Manager: %s\n", manager.GetInfo())
	
	// Add team members
	manager.AddTeamMember("Bob Wilson")
	manager.AddTeamMember("Carol Brown")
	manager.AddTeamMember("David Miller")
	
	fmt.Printf("Team info: %s\n", manager.GetTeamInfo())
	
	// Access methods from all embedded levels
	fmt.Printf("Full name: %s\n", manager.FullName())     // From Person
	fmt.Printf("Is adult: %t\n", manager.IsAdult())       // From Person
	fmt.Printf("Employee info: %s\n", manager.GetInfo())  // From Employee
}