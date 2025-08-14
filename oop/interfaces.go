// Package oop demonstrates Go interfaces including interface definition,
// implementation, polymorphism, and interface composition.
package oop

import (
	"fmt"
	"io"
	"sort"
	"strings"
)

// DemoBasicInterfaces demonstrates basic interface concepts
func DemoBasicInterfaces() {
	fmt.Println("\n=== Basic Interfaces Demo ===")

	// Different types implementing the same interface
	var shapes []Shape

	rect := Rectangle{Width: 5, Height: 3}
	circle := Circle{Radius: 4}
	triangle := Triangle{Base: 6, Height: 4}

	shapes = append(shapes, rect, circle, triangle)

	fmt.Println("Calculating areas using interface:")
	totalArea := 0.0
	for i, shape := range shapes {
		area := shape.Area()
		fmt.Printf("  Shape %d: Area = %.2f\n", i+1, area)
		totalArea += area
	}
	fmt.Printf("Total area: %.2f\n", totalArea)

	// Interface with multiple methods
	fmt.Println("\nUsing Drawable interface:")
	drawables := []Drawable{rect, circle, triangle}
	for _, drawable := range drawables {
		drawable.Draw()
		fmt.Printf("  Area: %.2f, Perimeter: %.2f\n", 
			drawable.Area(), drawable.Perimeter())
	}
}

// Shape interface defines what it means to be a shape
type Shape interface {
	Area() float64
}

// Drawable interface extends Shape with additional methods
type Drawable interface {
	Shape // Embedded interface
	Perimeter() float64
	Draw()
}

// Triangle implements both Shape and Drawable interfaces
type Triangle struct {
	Base, Height float64
}

// Area calculates triangle area
func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

// Perimeter calculates triangle perimeter (assuming equilateral for simplicity)
func (t Triangle) Perimeter() float64 {
	return 3 * t.Base // Simplified for equilateral triangle
}

// Draw draws the triangle
func (t Triangle) Draw() {
	fmt.Printf("Drawing triangle with base %.1f and height %.1f", t.Base, t.Height)
}

// Rectangle and Circle already implement Shape from structs.go
// Let's add Draw and Perimeter methods to make them implement Drawable

// Draw draws the rectangle
func (r Rectangle) Draw() {
	fmt.Printf("Drawing rectangle %.1f x %.1f", r.Width, r.Height)
}

// Draw draws the circle  
func (c Circle) Draw() {
	fmt.Printf("Drawing circle with radius %.1f", c.Radius)
}

// DemoInterfaceAssertion demonstrates type assertions and type switches
func DemoInterfaceAssertion() {
	fmt.Println("\n=== Interface Assertion Demo ===")

	var shapes []Shape = []Shape{
		Rectangle{Width: 5, Height: 3},
		Circle{Radius: 4},
		Triangle{Base: 6, Height: 4},
	}

	for i, shape := range shapes {
		fmt.Printf("\nShape %d:\n", i+1)
		
		// Type assertion with ok check
		if rect, ok := shape.(Rectangle); ok {
			fmt.Printf("  Rectangle: %.1f x %.1f\n", rect.Width, rect.Height)
		}

		if circle, ok := shape.(Circle); ok {
			fmt.Printf("  Circle: radius %.1f\n", circle.Radius)
		}

		// Type switch
		switch s := shape.(type) {
		case Rectangle:
			fmt.Printf("  Type switch: Rectangle area %.2f\n", s.Area())
		case Circle:
			fmt.Printf("  Type switch: Circle circumference %.2f\n", s.Circumference())
		case Triangle:
			fmt.Printf("  Type switch: Triangle perimeter %.2f\n", s.Perimeter())
		default:
			fmt.Printf("  Type switch: Unknown shape type\n")
		}

		// Check if shape implements Drawable
		if drawable, ok := shape.(Drawable); ok {
			fmt.Printf("  Shape is drawable: ")
			drawable.Draw()
			fmt.Println()
		} else {
			fmt.Printf("  Shape is not drawable\n")
		}
	}
}

// DemoEmptyInterface demonstrates the empty interface (interface{})
func DemoEmptyInterface() {
	fmt.Println("\n=== Empty Interface Demo ===")

	// Empty interface can hold any type
	var values []interface{}
	values = append(values, 42)
	values = append(values, "hello")
	values = append(values, true)
	values = append(values, 3.14)
	values = append(values, Rectangle{Width: 2, Height: 3})

	fmt.Println("Values in empty interface slice:")
	for i, value := range values {
		fmt.Printf("  [%d]: %v (type: %T)\n", i, value, value)
		
		// Type assertions with empty interface
		switch v := value.(type) {
		case int:
			fmt.Printf("       Integer: %d\n", v)
		case string:
			fmt.Printf("       String: %s (length: %d)\n", v, len(v))
		case bool:
			fmt.Printf("       Boolean: %t\n", v)
		case float64:
			fmt.Printf("       Float: %.2f\n", v)
		case Rectangle:
			fmt.Printf("       Rectangle area: %.2f\n", v.Area())
		default:
			fmt.Printf("       Unknown type\n")
		}
	}

	// Function accepting empty interface
	PrintAnything("Hello")
	PrintAnything(123)
	PrintAnything([]int{1, 2, 3})
}

// PrintAnything accepts any type using empty interface
func PrintAnything(value interface{}) {
	fmt.Printf("PrintAnything received: %v (type: %T)\n", value, value)
}

// DemoInterfaceComposition demonstrates composing interfaces
func DemoInterfaceComposition() {
	fmt.Println("\n=== Interface Composition Demo ===")

	// ReadWriter combines Reader and Writer
	var buffer strings.Builder
	
	// strings.Builder implements io.Writer
	message := "Hello, Interface Composition!"
	WriteToWriter(&buffer, message)
	
	// Read from string reader
	reader := strings.NewReader(message)
	ReadFromReader(reader)

	// Using interface composition
	fmt.Println("\nDemonstrating interface composition:")
	processor := &DataProcessor{}
	
	// Implement different interfaces
	data := []byte("test data")
	ProcessData(processor, data)
}

// WriteToWriter writes data to any io.Writer
func WriteToWriter(w io.Writer, data string) {
	n, err := w.Write([]byte(data))
	if err != nil {
		fmt.Printf("Error writing: %v\n", err)
	} else {
		fmt.Printf("Wrote %d bytes to writer\n", n)
	}
}

// ReadFromReader reads data from any io.Reader
func ReadFromReader(r io.Reader) {
	buffer := make([]byte, 100)
	n, err := r.Read(buffer)
	if err != nil && err != io.EOF {
		fmt.Printf("Error reading: %v\n", err)
	} else {
		fmt.Printf("Read %d bytes: %s\n", n, string(buffer[:n]))
	}
}

// Custom interfaces for composition
type Reader interface {
	Read(data []byte) (int, error)
}

type Writer interface {
	Write(data []byte) (int, error)
}

type Closer interface {
	Close() error
}

// ReadWriter composes Reader and Writer
type ReadWriter interface {
	Reader
	Writer
}

// ReadWriteCloser composes ReadWriter and Closer
type ReadWriteCloser interface {
	ReadWriter
	Closer
}

// DataProcessor implements multiple interfaces
type DataProcessor struct {
	data []byte
}

// Read implements Reader interface
func (dp *DataProcessor) Read(data []byte) (int, error) {
	if len(dp.data) == 0 {
		return 0, io.EOF
	}
	n := copy(data, dp.data)
	dp.data = dp.data[n:]
	fmt.Printf("DataProcessor.Read: read %d bytes\n", n)
	return n, nil
}

// Write implements Writer interface
func (dp *DataProcessor) Write(data []byte) (int, error) {
	dp.data = append(dp.data, data...)
	fmt.Printf("DataProcessor.Write: wrote %d bytes\n", len(data))
	return len(data), nil
}

// Close implements Closer interface
func (dp *DataProcessor) Close() error {
	dp.data = nil
	fmt.Println("DataProcessor.Close: closed")
	return nil
}

// ProcessData works with any ReadWriteCloser
func ProcessData(rwc ReadWriteCloser, data []byte) {
	// Write data
	rwc.Write(data)
	
	// Read data back
	buffer := make([]byte, len(data))
	rwc.Read(buffer)
	
	// Close
	rwc.Close()
}

// DemoInterfaceNil demonstrates nil interfaces
func DemoInterfaceNil() {
	fmt.Println("\n=== Interface Nil Demo ===")

	var shape Shape
	fmt.Printf("Nil interface: %v (== nil: %t)\n", shape, shape == nil)

	// Nil interface type assertion panics
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from panic: %v\n", r)
		}
	}()

	// This would panic: shape.Area()
	fmt.Println("Calling method on nil interface would panic (commented out)")

	// Interface with nil concrete value
	var rect *Rectangle
	shape = rect // shape is not nil, but contains nil pointer
	fmt.Printf("Interface with nil pointer: %v (== nil: %t)\n", shape, shape == nil)

	// This would also panic because rect is nil
	// shape.Area()

	// Safe way to check
	if shape != nil {
		if r, ok := shape.(*Rectangle); ok && r != nil {
			fmt.Printf("Safe area calculation: %.2f\n", r.Area())
		} else {
			fmt.Println("Shape interface contains nil pointer")
		}
	}
}

// DemoPolymorphism demonstrates polymorphism with interfaces
func DemoPolymorphism() {
	fmt.Println("\n=== Polymorphism Demo ===")

	// Create different animals
	animals := []Animal{
		Dog{Name: "Buddy"},
		Cat{Name: "Whiskers"},
		Bird{Name: "Tweety"},
	}

	fmt.Println("Animal sounds:")
	for _, animal := range animals {
		animal.Speak()
		animal.Move()
		
		// Use interface methods
		if pet, ok := animal.(Pet); ok {
			pet.Play()
		}
		
		fmt.Println()
	}

	// Feeding animals
	fmt.Println("Feeding animals:")
	for _, animal := range animals {
		FeedAnimal(animal)
	}
}

// Animal interface defines animal behavior
type Animal interface {
	Speak()
	Move()
}

// Pet interface for animals that can be pets
type Pet interface {
	Animal // Embedded interface
	Play()
}

// Dog implements Animal and Pet interfaces
type Dog struct {
	Name string
}

func (d Dog) Speak() {
	fmt.Printf("%s says: Woof!\n", d.Name)
}

func (d Dog) Move() {
	fmt.Printf("%s runs on four legs\n", d.Name)
}

func (d Dog) Play() {
	fmt.Printf("%s plays fetch\n", d.Name)
}

// Cat implements Animal and Pet interfaces
type Cat struct {
	Name string
}

func (c Cat) Speak() {
	fmt.Printf("%s says: Meow!\n", c.Name)
}

func (c Cat) Move() {
	fmt.Printf("%s walks gracefully\n", c.Name)
}

func (c Cat) Play() {
	fmt.Printf("%s plays with yarn\n", c.Name)
}

// Bird implements only Animal interface
type Bird struct {
	Name string
}

func (b Bird) Speak() {
	fmt.Printf("%s says: Tweet!\n", b.Name)
}

func (b Bird) Move() {
	fmt.Printf("%s flies in the sky\n", b.Name)
}

// FeedAnimal feeds any animal
func FeedAnimal(animal Animal) {
	fmt.Printf("Feeding the animal...\n")
	animal.Speak() // Happy sound after eating
}

// DemoInterfaceSlices demonstrates working with slices of interfaces
func DemoInterfaceSlices() {
	fmt.Println("\n=== Interface Slices Demo ===")

	// Create slice of different shapes
	shapes := []Shape{
		Rectangle{Width: 4, Height: 3},
		Circle{Radius: 2},
		Triangle{Base: 5, Height: 3},
	}

	// Calculate total area
	totalArea := CalculateTotalArea(shapes)
	fmt.Printf("Total area of all shapes: %.2f\n", totalArea)

	// Find largest shape
	largest := FindLargestShape(shapes)
	fmt.Printf("Largest shape area: %.2f\n", largest.Area())

	// Sort shapes by area
	sortableShapes := make(SortableShapes, len(shapes))
	copy(sortableShapes, shapes)
	
	fmt.Println("\nShapes before sorting:")
	for i, shape := range sortableShapes {
		fmt.Printf("  %d: Area %.2f\n", i+1, shape.Area())
	}

	sort.Sort(sortableShapes)
	
	fmt.Println("\nShapes after sorting by area:")
	for i, shape := range sortableShapes {
		fmt.Printf("  %d: Area %.2f\n", i+1, shape.Area())
	}
}

// CalculateTotalArea sums areas of all shapes
func CalculateTotalArea(shapes []Shape) float64 {
	total := 0.0
	for _, shape := range shapes {
		total += shape.Area()
	}
	return total
}

// FindLargestShape finds the shape with largest area
func FindLargestShape(shapes []Shape) Shape {
	if len(shapes) == 0 {
		return nil
	}
	
	largest := shapes[0]
	for _, shape := range shapes[1:] {
		if shape.Area() > largest.Area() {
			largest = shape
		}
	}
	return largest
}

// SortableShapes implements sort.Interface for []Shape
type SortableShapes []Shape

func (s SortableShapes) Len() int {
	return len(s)
}

func (s SortableShapes) Less(i, j int) bool {
	return s[i].Area() < s[j].Area()
}

func (s SortableShapes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// DemoInterfaceBestPractices demonstrates interface best practices
func DemoInterfaceBestPractices() {
	fmt.Println("\n=== Interface Best Practices Demo ===")

	// 1. Accept interfaces, return concrete types
	logger := &FileLogger{filename: "app.log"}
	
	// Function accepts interface
	LogMessage(logger, "Application started")
	LogMessage(logger, "Processing data")
	
	// 2. Define interfaces at the point of use
	// (UserService defines what it needs, not what Logger provides)
	service := UserService{logger: logger}
	service.CreateUser("alice", "alice@example.com")
	
	// 3. Keep interfaces small and focused
	// Good: io.Reader, io.Writer (single method)
	// Bad: Large interfaces with many methods
	
	// 4. Interface segregation - prefer many small interfaces
	var processor DataProcessor
	
	// Can be used as different interfaces
	var writer Writer = &processor
	var reader Reader = &processor
	var closer Closer = &processor
	
	WriteToInterface(writer, []byte("test"))
	ReadFromInterface(reader)
	CloseInterface(closer)

	fmt.Println("\nInterface best practices demonstrated:")
	fmt.Println("1. Accept interfaces, return concrete types")
	fmt.Println("2. Define interfaces where they're used")
	fmt.Println("3. Keep interfaces small (1-3 methods)")
	fmt.Println("4. Use interface segregation")
}

// Logger interface - small and focused
type Logger interface {
	Log(message string)
}

// FileLogger implements Logger
type FileLogger struct {
	filename string
}

func (fl *FileLogger) Log(message string) {
	fmt.Printf("FileLogger: Writing '%s' to %s\n", message, fl.filename)
}

// LogMessage accepts any Logger interface
func LogMessage(logger Logger, message string) {
	logger.Log(message)
}

// UserService defines the interface it needs (not what FileLogger provides)
type UserService struct {
	logger interface {
		Log(message string)
	}
}

func (us *UserService) CreateUser(name, email string) {
	us.logger.Log(fmt.Sprintf("Creating user: %s (%s)", name, email))
	// ... user creation logic
	us.logger.Log("User created successfully")
}

// Helper functions for interface segregation
func WriteToInterface(w Writer, data []byte) {
	w.Write(data)
}

func ReadFromInterface(r Reader) {
	buffer := make([]byte, 10)
	r.Read(buffer)
}

func CloseInterface(c Closer) {
	c.Close()
}