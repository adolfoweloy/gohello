package oop

import (
	"io"
	"sort"
	"strings"
	"testing"
)

func TestDemoBasicInterfaces(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("DemoBasicInterfaces panicked: %v", r)
		}
	}()
	
	DemoBasicInterfaces()
}

func TestDemoInterfaceAssertion(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("DemoInterfaceAssertion panicked: %v", r)
		}
	}()
	
	DemoInterfaceAssertion()
}

func TestDemoEmptyInterface(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("DemoEmptyInterface panicked: %v", r)
		}
	}()
	
	DemoEmptyInterface()
}

func TestDemoInterfaceComposition(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("DemoInterfaceComposition panicked: %v", r)
		}
	}()
	
	DemoInterfaceComposition()
}

func TestDemoInterfaceNil(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("DemoInterfaceNil panicked: %v", r)
		}
	}()
	
	DemoInterfaceNil()
}

func TestDemoPolymorphism(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("DemoPolymorphism panicked: %v", r)
		}
	}()
	
	DemoPolymorphism()
}

func TestDemoInterfaceSlices(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("DemoInterfaceSlices panicked: %v", r)
		}
	}()
	
	DemoInterfaceSlices()
}

func TestDemoInterfaceBestPractices(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("DemoInterfaceBestPractices panicked: %v", r)
		}
	}()
	
	DemoInterfaceBestPractices()
}

func TestShapeInterface(t *testing.T) {
	rect := Rectangle{Width: 4, Height: 3}
	circle := Circle{Radius: 2}
	triangle := Triangle{Base: 6, Height: 4}
	
	// Test that all types implement Shape interface
	var shapes []Shape = []Shape{rect, circle, triangle}
	
	expectedAreas := []float64{12.0, 12.566, 12.0} // approximate for circle
	
	for i, shape := range shapes {
		area := shape.Area()
		if i == 1 { // circle - check approximate value
			if area < 12.5 || area > 12.6 {
				t.Errorf("Circle area should be around 12.57, got %.2f", area)
			}
		} else {
			if area != expectedAreas[i] {
				t.Errorf("Shape %d: expected area %.2f, got %.2f", i, expectedAreas[i], area)
			}
		}
	}
}

func TestDrawableInterface(t *testing.T) {
	rect := Rectangle{Width: 4, Height: 3}
	circle := Circle{Radius: 2}
	triangle := Triangle{Base: 6, Height: 4}
	
	// Test that all types implement Drawable interface
	var drawables []Drawable = []Drawable{rect, circle, triangle}
	
	for _, drawable := range drawables {
		// Should be able to call all Drawable methods
		area := drawable.Area()
		perimeter := drawable.Perimeter()
		
		if area <= 0 {
			t.Errorf("Area should be positive, got %.2f", area)
		}
		
		if perimeter <= 0 {
			t.Errorf("Perimeter should be positive, got %.2f", perimeter)
		}
		
		// Draw method should not panic
		drawable.Draw()
	}
}

func TestTriangleImplementation(t *testing.T) {
	triangle := Triangle{Base: 6, Height: 4}
	
	expectedArea := 12.0 // 0.5 * 6 * 4
	if triangle.Area() != expectedArea {
		t.Errorf("Triangle area: expected %.2f, got %.2f", expectedArea, triangle.Area())
	}
	
	expectedPerimeter := 18.0 // 3 * 6 (equilateral approximation)
	if triangle.Perimeter() != expectedPerimeter {
		t.Errorf("Triangle perimeter: expected %.2f, got %.2f", expectedPerimeter, triangle.Perimeter())
	}
}

func TestTypeAssertion(t *testing.T) {
	var shape Shape = Rectangle{Width: 5, Height: 3}
	
	// Successful type assertion
	if rect, ok := shape.(Rectangle); ok {
		if rect.Width != 5 || rect.Height != 3 {
			t.Errorf("Type assertion failed to preserve values")
		}
	} else {
		t.Errorf("Type assertion should succeed for Rectangle")
	}
	
	// Failed type assertion
	if _, ok := shape.(Circle); ok {
		t.Errorf("Type assertion should fail for Circle when shape is Rectangle")
	}
}

func TestEmptyInterface(t *testing.T) {
	var empty interface{}
	
	// Can hold any type
	empty = 42
	if val, ok := empty.(int); !ok || val != 42 {
		t.Errorf("Empty interface should hold int")
	}
	
	empty = "hello"
	if val, ok := empty.(string); !ok || val != "hello" {
		t.Errorf("Empty interface should hold string")
	}
	
	empty = Rectangle{Width: 2, Height: 3}
	if val, ok := empty.(Rectangle); !ok || val.Area() != 6 {
		t.Errorf("Empty interface should hold Rectangle")
	}
}

func TestInterfaceCompositionDemo(t *testing.T) {
	processor := &DataProcessor{}
	
	// Test that DataProcessor implements all composed interfaces
	var reader Reader = processor
	var writer Writer = processor
	var closer Closer = processor
	var readWriter ReadWriter = processor
	var readWriteCloser ReadWriteCloser = processor
	
	// Test writing
	data := []byte("test data")
	n, err := writer.Write(data)
	if err != nil {
		t.Errorf("Write failed: %v", err)
	}
	if n != len(data) {
		t.Errorf("Write: expected %d bytes, wrote %d", len(data), n)
	}
	
	// Test reading
	buffer := make([]byte, len(data))
	n, err = reader.Read(buffer)
	if err != nil && err != io.EOF {
		t.Errorf("Read failed: %v", err)
	}
	if n != len(data) {
		t.Errorf("Read: expected %d bytes, read %d", len(data), n)
	}
	
	// Test closing
	err = closer.Close()
	if err != nil {
		t.Errorf("Close failed: %v", err)
	}
	
	// Ensure all interface assignments work
	_ = readWriter
	_ = readWriteCloser
}

func TestAnimalPolymorphism(t *testing.T) {
	dog := Dog{Name: "Buddy"}
	cat := Cat{Name: "Whiskers"}
	bird := Bird{Name: "Tweety"}
	
	// Test that all implement Animal interface
	animals := []Animal{dog, cat, bird}
	
	for _, animal := range animals {
		// Should not panic
		animal.Speak()
		animal.Move()
	}
	
	// Test Pet interface
	if pet, ok := animals[0].(Pet); ok { // Dog
		pet.Play()
	} else {
		t.Errorf("Dog should implement Pet interface")
	}
	
	if pet, ok := animals[1].(Pet); ok { // Cat
		pet.Play()
	} else {
		t.Errorf("Cat should implement Pet interface")
	}
	
	if _, ok := animals[2].(Pet); ok { // Bird
		t.Errorf("Bird should not implement Pet interface")
	}
}

func TestCalculateTotalArea(t *testing.T) {
	shapes := []Shape{
		Rectangle{Width: 2, Height: 3}, // Area: 6
		Rectangle{Width: 4, Height: 1}, // Area: 4
	}
	
	total := CalculateTotalArea(shapes)
	expected := 10.0
	
	if total != expected {
		t.Errorf("Total area: expected %.2f, got %.2f", expected, total)
	}
	
	// Test empty slice
	emptyTotal := CalculateTotalArea([]Shape{})
	if emptyTotal != 0 {
		t.Errorf("Empty slice total should be 0, got %.2f", emptyTotal)
	}
}

func TestFindLargestShape(t *testing.T) {
	shapes := []Shape{
		Rectangle{Width: 2, Height: 3}, // Area: 6
		Rectangle{Width: 4, Height: 2}, // Area: 8 (largest)
		Rectangle{Width: 1, Height: 5}, // Area: 5
	}
	
	largest := FindLargestShape(shapes)
	if largest == nil {
		t.Errorf("FindLargestShape should return a shape")
	}
	
	expectedArea := 8.0
	if largest.Area() != expectedArea {
		t.Errorf("Largest area: expected %.2f, got %.2f", expectedArea, largest.Area())
	}
	
	// Test empty slice
	emptyLargest := FindLargestShape([]Shape{})
	if emptyLargest != nil {
		t.Errorf("FindLargestShape on empty slice should return nil")
	}
}

func TestSortableShapes(t *testing.T) {
	shapes := SortableShapes{
		Rectangle{Width: 3, Height: 2}, // Area: 6
		Rectangle{Width: 2, Height: 1}, // Area: 2
		Rectangle{Width: 2, Height: 2}, // Area: 4
	}
	
	// Test sort.Interface methods
	if shapes.Len() != 3 {
		t.Errorf("Len: expected 3, got %d", shapes.Len())
	}
	
	if !shapes.Less(1, 0) { // Area 2 < Area 6
		t.Errorf("Less: 2 should be less than 6")
	}
	
	if shapes.Less(0, 1) { // Area 6 > Area 2
		t.Errorf("Less: 6 should not be less than 2")
	}
	
	// Test swap
	originalFirst := shapes[0]
	originalSecond := shapes[1]
	shapes.Swap(0, 1)
	
	if shapes[0] != originalSecond || shapes[1] != originalFirst {
		t.Errorf("Swap failed")
	}
	
	// Test sorting
	shapes = SortableShapes{
		Rectangle{Width: 3, Height: 2}, // Area: 6
		Rectangle{Width: 2, Height: 1}, // Area: 2
		Rectangle{Width: 2, Height: 2}, // Area: 4
	}
	
	sort.Sort(shapes)
	
	// Should be sorted by area: 2, 4, 6
	expectedAreas := []float64{2, 4, 6}
	for i, shape := range shapes {
		if shape.Area() != expectedAreas[i] {
			t.Errorf("After sort[%d]: expected area %.2f, got %.2f", 
				i, expectedAreas[i], shape.Area())
		}
	}
}

func TestFileLogger(t *testing.T) {
	logger := &FileLogger{filename: "test.log"}
	
	// Test that it implements Logger interface
	var l Logger = logger
	
	// Should not panic
	l.Log("test message")
}

func TestUserService(t *testing.T) {
	logger := &FileLogger{filename: "test.log"}
	service := UserService{logger: logger}
	
	// Should not panic
	service.CreateUser("testuser", "test@example.com")
}

func TestInterfaceNilBehavior(t *testing.T) {
	var shape Shape
	
	// Nil interface
	if shape != nil {
		t.Errorf("Nil interface should be equal to nil")
	}
	
	// Interface with nil pointer
	var rect *Rectangle
	shape = rect
	
	if shape == nil {
		t.Errorf("Interface with nil pointer should not be equal to nil")
	}
	
	// Type assertion on interface with nil pointer
	if r, ok := shape.(*Rectangle); ok {
		if r != nil {
			t.Errorf("Type assertion should return nil pointer")
		}
	} else {
		t.Errorf("Type assertion should succeed even with nil pointer")
	}
}

func TestDataProcessorInterfaces(t *testing.T) {
	processor := &DataProcessor{}
	
	// Test Write
	data := []byte("hello world")
	n, err := processor.Write(data)
	if err != nil {
		t.Errorf("Write failed: %v", err)
	}
	if n != len(data) {
		t.Errorf("Write: expected %d bytes, got %d", len(data), n)
	}
	
	// Test Read
	buffer := make([]byte, len(data))
	n, err = processor.Read(buffer)
	if err != nil && err != io.EOF {
		t.Errorf("Read failed: %v", err)
	}
	if n != len(data) {
		t.Errorf("Read: expected %d bytes, got %d", len(data), n)
	}
	if string(buffer) != string(data) {
		t.Errorf("Read data mismatch: expected %s, got %s", string(data), string(buffer))
	}
	
	// Test Close
	err = processor.Close()
	if err != nil {
		t.Errorf("Close failed: %v", err)
	}
	
	// After close, data should be nil
	if processor.data != nil {
		t.Errorf("Data should be nil after close")
	}
}

func TestWriteToWriter(t *testing.T) {
	var buffer strings.Builder
	
	WriteToWriter(&buffer, "test message")
	
	result := buffer.String()
	if result != "test message" {
		t.Errorf("Expected 'test message', got '%s'", result)
	}
}

func TestReadFromReader(t *testing.T) {
	reader := strings.NewReader("test content")
	
	// This will print output, but shouldn't panic
	ReadFromReader(reader)
}

func TestProcessData(t *testing.T) {
	processor := &DataProcessor{}
	data := []byte("test data")
	
	// Should not panic
	ProcessData(processor, data)
}