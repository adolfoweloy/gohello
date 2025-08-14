package oop

import (
	"fmt"
	"math"
)

// Shape interface defines methods that all shapes must implement
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Drawable interface for shapes that can be drawn
type Drawable interface {
	Draw() string
}

// ShapeDrawer combines Shape and Drawable interfaces
type ShapeDrawer interface {
	Shape
	Drawable
}

// Triangle represents a triangle
type Triangle struct {
	Base   float64
	Height float64
	SideA  float64
	SideB  float64
	SideC  float64
}

// Area calculates the area of the triangle
func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

// Perimeter calculates the perimeter of the triangle
func (t Triangle) Perimeter() float64 {
	return t.SideA + t.SideB + t.SideC
}

// Draw returns a string representation of drawing the triangle
func (t Triangle) Draw() string {
	return fmt.Sprintf("Drawing triangle with base %.2f and height %.2f", t.Base, t.Height)
}

// String returns a string representation of the triangle
func (t Triangle) String() string {
	return fmt.Sprintf("Triangle(base: %.2f, height: %.2f, sides: %.2f, %.2f, %.2f)", 
		t.Base, t.Height, t.SideA, t.SideB, t.SideC)
}

// Square represents a square (Rectangle is already defined in structs.go)
type Square struct {
	Side float64
}

// Area calculates the area of the square
func (s Square) Area() float64 {
	return s.Side * s.Side
}

// Perimeter calculates the perimeter of the square
func (s Square) Perimeter() float64 {
	return 4 * s.Side
}

// Draw returns a string representation of drawing the square
func (s Square) Draw() string {
	return fmt.Sprintf("Drawing square with side %.2f", s.Side)
}

// String returns a string representation of the square
func (s Square) String() string {
	return fmt.Sprintf("Square(side: %.2f)", s.Side)
}

// DemonstrateInterfaces shows interface usage
func DemonstrateInterfaces() {
	fmt.Println("\n=== Interfaces ===")
	
	// Create shapes
	rect := Rectangle{Width: 4.0, Height: 3.0}
	circle := Circle{Radius: 2.0}
	triangle := Triangle{
		Base: 6.0, Height: 4.0,
		SideA: 5.0, SideB: 5.0, SideC: 6.0,
	}
	square := Square{Side: 3.0}
	
	// Store shapes in a slice of Shape interface
	shapes := []Shape{rect, circle, triangle, square}
	
	fmt.Println("Shape calculations:")
	var totalArea, totalPerimeter float64
	
	for i, shape := range shapes {
		area := shape.Area()
		perimeter := shape.Perimeter()
		totalArea += area
		totalPerimeter += perimeter
		
		fmt.Printf("  %d. %T - Area: %.2f, Perimeter: %.2f\n", 
			i+1, shape, area, perimeter)
	}
	
	fmt.Printf("Total area: %.2f\n", totalArea)
	fmt.Printf("Total perimeter: %.2f\n", totalPerimeter)
	
	// Demonstrate drawable shapes
	fmt.Println("\nDrawing shapes:")
	drawableShapes := []Drawable{triangle, square}
	
	for _, drawable := range drawableShapes {
		fmt.Printf("  %s\n", drawable.Draw())
	}
	
	// Demonstrate combined interface
	fmt.Println("\nShape drawers:")
	shapeDrawers := []ShapeDrawer{triangle, square}
	
	for _, sd := range shapeDrawers {
		fmt.Printf("  %s (Area: %.2f)\n", sd.Draw(), sd.Area())
	}
}

// Printer interface for things that can print
type Printer interface {
	Print() string
}

// Scanner interface for things that can scan
type Scanner interface {
	Scan() string
}

// AllInOne interface combines multiple interfaces
type AllInOne interface {
	Printer
	Scanner
	Copy() string
}

// Device represents a physical device
type Device struct {
	Name  string
	Model string
}

// LaserPrinter implements Printer interface
type LaserPrinter struct {
	Device
	PagesPerMinute int
}

// Print implements the Printer interface
func (lp LaserPrinter) Print() string {
	return fmt.Sprintf("%s %s printing at %d pages/min", lp.Name, lp.Model, lp.PagesPerMinute)
}

// FlatbedScanner implements Scanner interface
type FlatbedScanner struct {
	Device
	Resolution int
}

// Scan implements the Scanner interface
func (fs FlatbedScanner) Scan() string {
	return fmt.Sprintf("%s %s scanning at %d DPI", fs.Name, fs.Model, fs.Resolution)
}

// MultiFunction implements all interfaces (AllInOne)
type MultiFunction struct {
	Device
	PrintSpeed int
	ScanDPI    int
}

// Print implements the Printer interface
func (mf MultiFunction) Print() string {
	return fmt.Sprintf("%s %s printing at %d pages/min", mf.Name, mf.Model, mf.PrintSpeed)
}

// Scan implements the Scanner interface
func (mf MultiFunction) Scan() string {
	return fmt.Sprintf("%s %s scanning at %d DPI", mf.Name, mf.Model, mf.ScanDPI)
}

// Copy implements copy functionality
func (mf MultiFunction) Copy() string {
	return fmt.Sprintf("%s %s copying document", mf.Name, mf.Model)
}

// DemonstrateInterfaceComposition shows interface composition
func DemonstrateInterfaceComposition() {
	fmt.Println("\n=== Interface Composition ===")
	
	// Create devices
	printer := LaserPrinter{
		Device:         Device{Name: "HP", Model: "LaserJet Pro"},
		PagesPerMinute: 25,
	}
	
	scanner := FlatbedScanner{
		Device:     Device{Name: "Epson", Model: "Perfection V600"},
		Resolution: 6400,
	}
	
	multiFunc := MultiFunction{
		Device:     Device{Name: "Canon", Model: "PIXMA TR8520"},
		PrintSpeed: 15,
		ScanDPI:    4800,
	}
	
	// Use as specific interfaces
	printers := []Printer{printer, multiFunc}
	scanners := []Scanner{scanner, multiFunc}
	
	fmt.Println("Printers:")
	for _, p := range printers {
		fmt.Printf("  %s\n", p.Print())
	}
	
	fmt.Println("\nScanners:")
	for _, s := range scanners {
		fmt.Printf("  %s\n", s.Scan())
	}
	
	// Use as combined interface
	fmt.Println("\nAll-in-one devices:")
	allInOnes := []AllInOne{multiFunc}
	
	for _, aio := range allInOnes {
		fmt.Printf("  Print: %s\n", aio.Print())
		fmt.Printf("  Scan: %s\n", aio.Scan())
		fmt.Printf("  Copy: %s\n", aio.Copy())
	}
}

// demonstrateTypeAssertion shows type assertion and type switches
func DemonstrateTypeAssertion() {
	fmt.Println("\n=== Type Assertion ===")
	
	// Create a slice of interfaces
	var devices []interface{} = []interface{}{
		LaserPrinter{Device: Device{Name: "HP", Model: "LaserJet"}, PagesPerMinute: 20},
		FlatbedScanner{Device: Device{Name: "Epson", Model: "Scanner"}, Resolution: 1200},
		MultiFunction{Device: Device{Name: "Canon", Model: "All-in-One"}, PrintSpeed: 10, ScanDPI: 2400},
		"Not a device",
		42,
	}
	
	for i, device := range devices {
		fmt.Printf("\nDevice %d:\n", i+1)
		
		// Type switch
		switch d := device.(type) {
		case LaserPrinter:
			fmt.Printf("  Laser Printer: %s\n", d.Print())
		case FlatbedScanner:
			fmt.Printf("  Scanner: %s\n", d.Scan())
		case MultiFunction:
			fmt.Printf("  Multi-function: %s\n", d.Print())
			fmt.Printf("  Multi-function: %s\n", d.Scan())
		default:
			fmt.Printf("  Unknown device type: %T (%v)\n", d, d)
		}
		
		// Type assertion with check
		if printer, ok := device.(Printer); ok {
			fmt.Printf("  As Printer: %s\n", printer.Print())
		}
		
		if scanner, ok := device.(Scanner); ok {
			fmt.Printf("  As Scanner: %s\n", scanner.Scan())
		}
	}
}

// EmptyInterface demonstrates the empty interface
func DemonstrateEmptyInterface() {
	fmt.Println("\n=== Empty Interface ===")
	
	// Empty interface can hold any type
	var anything interface{}
	
	values := []interface{}{
		42,
		"Hello, Go!",
		3.14159,
		true,
		[]int{1, 2, 3},
		map[string]int{"a": 1, "b": 2},
		Rectangle{Width: 5, Height: 3},
	}
	
	for i, value := range values {
		anything = value
		fmt.Printf("Value %d: %v (type: %T)\n", i+1, anything, anything)
		
		// Type assertion examples
		switch v := anything.(type) {
		case int:
			fmt.Printf("  Integer: %d squared is %d\n", v, v*v)
		case string:
			fmt.Printf("  String: length is %d\n", len(v))
		case float64:
			fmt.Printf("  Float: rounded is %.0f\n", math.Round(v))
		case Rectangle:
			fmt.Printf("  Rectangle: area is %.2f\n", v.Area())
		}
	}
}