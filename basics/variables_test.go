package basics

import (
	"reflect"
	"testing"
)

func TestDemoVariables(t *testing.T) {
	// Test that DemoVariables runs without panic
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("DemoVariables panicked: %v", r)
		}
	}()
	
	DemoVariables()
}

func TestDemoArrays(t *testing.T) {
	// Test that DemoArrays runs without panic
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("DemoArrays panicked: %v", r)
		}
	}()
	
	DemoArrays()
}

func TestDemoSlices(t *testing.T) {
	// Test that DemoSlices runs without panic
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("DemoSlices panicked: %v", r)
		}
	}()
	
	DemoSlices()
}

func TestDemoMaps(t *testing.T) {
	// Test that DemoMaps runs without panic
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("DemoMaps panicked: %v", r)
		}
	}()
	
	DemoMaps()
}

func TestDemoStructBasics(t *testing.T) {
	// Test that DemoStructBasics runs without panic
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("DemoStructBasics panicked: %v", r)
		}
	}()
	
	DemoStructBasics()
}

func TestGetTypeInfo(t *testing.T) {
	typeInfo := GetTypeInfo()
	
	// Test that we have information for basic types
	expectedTypes := []string{"bool", "string", "int", "float64", "byte", "rune"}
	
	for _, expectedType := range expectedTypes {
		if _, exists := typeInfo[expectedType]; !exists {
			t.Errorf("Expected type info for %s, but not found", expectedType)
		}
	}
	
	// Test that descriptions are not empty
	for typeName, description := range typeInfo {
		if description == "" {
			t.Errorf("Description for type %s is empty", typeName)
		}
	}
}

func TestPersonInfo(t *testing.T) {
	person := PersonInfo{
		Name: "Test Person",
		Age:  25,
	}
	
	if person.Name != "Test Person" {
		t.Errorf("Expected name 'Test Person', got '%s'", person.Name)
	}
	
	if person.Age != 25 {
		t.Errorf("Expected age 25, got %d", person.Age)
	}
}

func TestVariableTypes(t *testing.T) {
	// Test different variable declaration methods
	
	// Explicit type declaration
	var name string = "Gopher"
	if reflect.TypeOf(name).Kind() != reflect.String {
		t.Errorf("Expected string type, got %v", reflect.TypeOf(name))
	}
	
	// Type inference
	age := 10
	if reflect.TypeOf(age).Kind() != reflect.Int {
		t.Errorf("Expected int type, got %v", reflect.TypeOf(age))
	}
	
	// Multiple assignment
	x, y := 1, 2
	if x != 1 || y != 2 {
		t.Errorf("Expected x=1, y=2, got x=%d, y=%d", x, y)
	}
}

func TestArrayOperations(t *testing.T) {
	// Test array creation and access
	numbers := [3]int{1, 2, 3}
	
	if len(numbers) != 3 {
		t.Errorf("Expected array length 3, got %d", len(numbers))
	}
	
	if numbers[0] != 1 || numbers[2] != 3 {
		t.Errorf("Array values not as expected: %v", numbers)
	}
	
	// Test array copying (arrays are value types)
	original := [2]int{1, 2}
	copied := original
	copied[0] = 99
	
	if original[0] != 1 {
		t.Errorf("Original array should not be modified when copy is changed")
	}
}

func TestSliceOperations(t *testing.T) {
	// Test slice creation
	numbers := []int{1, 2, 3, 4, 5}
	
	if len(numbers) != 5 {
		t.Errorf("Expected slice length 5, got %d", len(numbers))
	}
	
	// Test append
	numbers = append(numbers, 6)
	if len(numbers) != 6 || numbers[5] != 6 {
		t.Errorf("Append operation failed")
	}
	
	// Test slicing
	subset := numbers[1:4]
	if len(subset) != 3 || subset[0] != 2 || subset[2] != 4 {
		t.Errorf("Slicing operation failed: %v", subset)
	}
	
	// Test that slices share underlying array
	original := []int{1, 2, 3}
	view := original[0:2]
	view[0] = 99
	
	if original[0] != 99 {
		t.Errorf("Slice should share underlying array")
	}
}

func TestMapOperations(t *testing.T) {
	// Test map creation and operations
	scores := make(map[string]int)
	scores["Alice"] = 95
	scores["Bob"] = 87
	
	if len(scores) != 2 {
		t.Errorf("Expected map length 2, got %d", len(scores))
	}
	
	// Test value access
	if scores["Alice"] != 95 {
		t.Errorf("Expected Alice's score to be 95, got %d", scores["Alice"])
	}
	
	// Test existence check
	_, exists := scores["Charlie"]
	if exists {
		t.Errorf("Charlie should not exist in map")
	}
	
	// Test deletion
	delete(scores, "Bob")
	if len(scores) != 1 {
		t.Errorf("Expected map length 1 after deletion, got %d", len(scores))
	}
}