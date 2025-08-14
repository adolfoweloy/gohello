package basics

import (
	"strings"
	"testing"
)

func TestDemoBasicFunctions(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("DemoBasicFunctions panicked: %v", r)
		}
	}()
	
	DemoBasicFunctions()
}

func TestDemoMultipleReturns(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("DemoMultipleReturns panicked: %v", r)
		}
	}()
	
	DemoMultipleReturns()
}

func TestDemoNamedReturns(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("DemoNamedReturns panicked: %v", r)
		}
	}()
	
	DemoNamedReturns()
}

func TestDemoVariadicFunctions(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("DemoVariadicFunctions panicked: %v", r)
		}
	}()
	
	DemoVariadicFunctions()
}

func TestDemoFunctionTypes(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("DemoFunctionTypes panicked: %v", r)
		}
	}()
	
	DemoFunctionTypes()
}

func TestDemoClosures(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("DemoClosures panicked: %v", r)
		}
	}()
	
	DemoClosures()
}

func TestDemoRecursion(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("DemoRecursion panicked: %v", r)
		}
	}()
	
	DemoRecursion()
}

func TestDemoHigherOrderFunctions(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("DemoHigherOrderFunctions panicked: %v", r)
		}
	}()
	
	DemoHigherOrderFunctions()
}

func TestGreet(t *testing.T) {
	result := greet("Alice")
	expected := "Hello, Alice!"
	if result != expected {
		t.Errorf("greet('Alice'): expected '%s', got '%s'", expected, result)
	}
}

func TestAdd(t *testing.T) {
	result := add(10, 20)
	expected := 30
	if result != expected {
		t.Errorf("add(10, 20): expected %d, got %d", expected, result)
	}
}

func TestFormatPersonInfo(t *testing.T) {
	result := formatPersonInfo("Bob", 25, true)
	expected := "Bob is 25 years old and is active"
	if result != expected {
		t.Errorf("formatPersonInfo('Bob', 25, true): expected '%s', got '%s'", expected, result)
	}
	
	result2 := formatPersonInfo("Alice", 30, false)
	expected2 := "Alice is 30 years old and is inactive"
	if result2 != expected2 {
		t.Errorf("formatPersonInfo('Alice', 30, false): expected '%s', got '%s'", expected2, result2)
	}
}

func TestDivide(t *testing.T) {
	quotient, remainder := divide(17, 5)
	if quotient != 3 || remainder != 2 {
		t.Errorf("divide(17, 5): expected (3, 2), got (%d, %d)", quotient, remainder)
	}
	
	quotient2, remainder2 := divide(20, 4)
	if quotient2 != 5 || remainder2 != 0 {
		t.Errorf("divide(20, 4): expected (5, 0), got (%d, %d)", quotient2, remainder2)
	}
}

func TestSafeDivide(t *testing.T) {
	// Test normal division
	result, err := safeDivide(10, 2)
	if err != nil {
		t.Errorf("safeDivide(10, 2): unexpected error %v", err)
	}
	if result != 5.0 {
		t.Errorf("safeDivide(10, 2): expected 5.0, got %f", result)
	}
	
	// Test division by zero
	_, err = safeDivide(10, 0)
	if err == nil {
		t.Errorf("safeDivide(10, 0): expected error for division by zero")
	}
	if !strings.Contains(err.Error(), "division by zero") {
		t.Errorf("safeDivide(10, 0): expected 'division by zero' error, got '%s'", err.Error())
	}
}

func TestGetPersonDetails(t *testing.T) {
	name, age := getPersonDetails()
	if name != "John Doe" {
		t.Errorf("getPersonDetails(): expected name 'John Doe', got '%s'", name)
	}
	if age != 30 {
		t.Errorf("getPersonDetails(): expected age 30, got %d", age)
	}
}

func TestCalculateRectangle(t *testing.T) {
	area, perimeter := calculateRectangle(5, 3)
	expectedArea := 15.0
	expectedPerimeter := 16.0
	
	if area != expectedArea {
		t.Errorf("calculateRectangle(5, 3): expected area %f, got %f", expectedArea, area)
	}
	if perimeter != expectedPerimeter {
		t.Errorf("calculateRectangle(5, 3): expected perimeter %f, got %f", expectedPerimeter, perimeter)
	}
}

func TestValidateAge(t *testing.T) {
	// Test negative age
	isValid, message := validateAge(-5)
	if isValid {
		t.Errorf("validateAge(-5): expected invalid")
	}
	if !strings.Contains(message, "negative") {
		t.Errorf("validateAge(-5): expected 'negative' in message, got '%s'", message)
	}
	
	// Test underage
	isValid, message = validateAge(15)
	if isValid {
		t.Errorf("validateAge(15): expected invalid")
	}
	if !strings.Contains(message, "18") {
		t.Errorf("validateAge(15): expected '18' in message, got '%s'", message)
	}
	
	// Test valid age
	isValid, message = validateAge(25)
	if !isValid {
		t.Errorf("validateAge(25): expected valid")
	}
	if !strings.Contains(message, "valid") {
		t.Errorf("validateAge(25): expected 'valid' in message, got '%s'", message)
	}
}

func TestGetStringStats(t *testing.T) {
	stats := getStringStats("Hello World!")
	
	if stats.Length != 12 {
		t.Errorf("getStringStats('Hello World!'): expected length 12, got %d", stats.Length)
	}
	if stats.Words != 2 {
		t.Errorf("getStringStats('Hello World!'): expected 2 words, got %d", stats.Words)
	}
	if stats.Vowels != 3 {
		t.Errorf("getStringStats('Hello World!'): expected 3 vowels, got %d", stats.Vowels)
	}
	if !stats.HasUpper {
		t.Errorf("getStringStats('Hello World!'): expected HasUpper to be true")
	}
	if !stats.HasLower {
		t.Errorf("getStringStats('Hello World!'): expected HasLower to be true")
	}
	if stats.HasDigits {
		t.Errorf("getStringStats('Hello World!'): expected HasDigits to be false")
	}
	
	// Test string with digits
	stats2 := getStringStats("Test123")
	if !stats2.HasDigits {
		t.Errorf("getStringStats('Test123'): expected HasDigits to be true")
	}
}

func TestSum(t *testing.T) {
	// Test with multiple arguments
	result1 := sum(1, 2, 3, 4, 5)
	expected1 := 15
	if result1 != expected1 {
		t.Errorf("sum(1, 2, 3, 4, 5): expected %d, got %d", expected1, result1)
	}
	
	// Test with no arguments
	result2 := sum()
	expected2 := 0
	if result2 != expected2 {
		t.Errorf("sum(): expected %d, got %d", expected2, result2)
	}
	
	// Test with slice
	numbers := []int{10, 20, 30}
	result3 := sum(numbers...)
	expected3 := 60
	if result3 != expected3 {
		t.Errorf("sum(10, 20, 30): expected %d, got %d", expected3, result3)
	}
}

func TestFormatMessage(t *testing.T) {
	result1 := formatMessage("Error", "File not found")
	expected1 := "[Error] File not found"
	if result1 != expected1 {
		t.Errorf("formatMessage('Error', 'File not found'): expected '%s', got '%s'", expected1, result1)
	}
	
	result2 := formatMessage("Info", "Operation", "completed", "successfully")
	expected2 := "[Info] Operation completed successfully"
	if result2 != expected2 {
		t.Errorf("formatMessage with multiple parts: expected '%s', got '%s'", expected2, result2)
	}
}

func TestMultiply(t *testing.T) {
	result := multiply(6, 7)
	expected := 42
	if result != expected {
		t.Errorf("multiply(6, 7): expected %d, got %d", expected, result)
	}
}

func TestSubtract(t *testing.T) {
	result := subtract(10, 3)
	expected := 7
	if result != expected {
		t.Errorf("subtract(10, 3): expected %d, got %d", expected, result)
	}
}

func TestMakeMultiplier(t *testing.T) {
	doubler := makeMultiplier(2)
	result1 := doubler(5)
	expected1 := 10
	if result1 != expected1 {
		t.Errorf("doubler(5): expected %d, got %d", expected1, result1)
	}
	
	tripler := makeMultiplier(3)
	result2 := tripler(4)
	expected2 := 12
	if result2 != expected2 {
		t.Errorf("tripler(4): expected %d, got %d", expected2, result2)
	}
}

func TestMakeCounter(t *testing.T) {
	counter := makeCounter()
	
	// Test first call
	result1 := counter()
	if result1 != 1 {
		t.Errorf("counter() first call: expected 1, got %d", result1)
	}
	
	// Test second call
	result2 := counter()
	if result2 != 2 {
		t.Errorf("counter() second call: expected 2, got %d", result2)
	}
	
	// Test third call
	result3 := counter()
	if result3 != 3 {
		t.Errorf("counter() third call: expected 3, got %d", result3)
	}
	
	// Test that different counters are independent
	counter2 := makeCounter()
	result4 := counter2()
	if result4 != 1 {
		t.Errorf("counter2() first call: expected 1, got %d", result4)
	}
}

func TestMakeFunctions(t *testing.T) {
	funcs := makeFunctions()
	
	if len(funcs) != 3 {
		t.Errorf("makeFunctions(): expected 3 functions, got %d", len(funcs))
	}
	
	// Test that each function returns the correct value
	expected := []int{0, 1, 4} // 0*0, 1*1, 2*2
	for i, f := range funcs {
		result := f()
		if result != expected[i] {
			t.Errorf("function %d: expected %d, got %d", i, expected[i], result)
		}
	}
}

func TestFactorial(t *testing.T) {
	testCases := []struct {
		input    int
		expected int
	}{
		{0, 1},
		{1, 1},
		{2, 2},
		{3, 6},
		{4, 24},
		{5, 120},
	}
	
	for _, tc := range testCases {
		result := factorial(tc.input)
		if result != tc.expected {
			t.Errorf("factorial(%d): expected %d, got %d", tc.input, tc.expected, result)
		}
	}
}

func TestFibonacci(t *testing.T) {
	testCases := []struct {
		input    int
		expected int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
		{8, 21},
	}
	
	for _, tc := range testCases {
		result := fibonacci(tc.input)
		if result != tc.expected {
			t.Errorf("fibonacci(%d): expected %d, got %d", tc.input, tc.expected, result)
		}
	}
}

func TestSumTree(t *testing.T) {
	// Test empty tree
	result1 := sumTree(nil)
	if result1 != 0 {
		t.Errorf("sumTree(nil): expected 0, got %d", result1)
	}
	
	// Test single node
	single := &TreeNode{Value: 5}
	result2 := sumTree(single)
	if result2 != 5 {
		t.Errorf("sumTree(single node): expected 5, got %d", result2)
	}
	
	// Test complex tree
	tree := &TreeNode{
		Value: 1,
		Left: &TreeNode{
			Value: 2,
			Left:  &TreeNode{Value: 4},
			Right: &TreeNode{Value: 5},
		},
		Right: &TreeNode{
			Value: 3,
			Left:  &TreeNode{Value: 6},
			Right: &TreeNode{Value: 7},
		},
	}
	
	result3 := sumTree(tree)
	expected3 := 1 + 2 + 3 + 4 + 5 + 6 + 7 // 28
	if result3 != expected3 {
		t.Errorf("sumTree(complex tree): expected %d, got %d", expected3, result3)
	}
}

func TestMap(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	result := Map(input, func(x int) int { return x * x })
	expected := []int{1, 4, 9, 16, 25}
	
	if len(result) != len(expected) {
		t.Errorf("Map: expected length %d, got %d", len(expected), len(result))
	}
	
	for i, v := range expected {
		if result[i] != v {
			t.Errorf("Map[%d]: expected %d, got %d", i, v, result[i])
		}
	}
}

func TestFilter(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6}
	result := Filter(input, func(x int) bool { return x%2 == 0 })
	expected := []int{2, 4, 6}
	
	if len(result) != len(expected) {
		t.Errorf("Filter: expected length %d, got %d", len(expected), len(result))
	}
	
	for i, v := range expected {
		if result[i] != v {
			t.Errorf("Filter[%d]: expected %d, got %d", i, v, result[i])
		}
	}
}

func TestReduce(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	
	// Test sum
	result1 := Reduce(input, 0, func(acc, x int) int { return acc + x })
	expected1 := 15
	if result1 != expected1 {
		t.Errorf("Reduce (sum): expected %d, got %d", expected1, result1)
	}
	
	// Test product
	result2 := Reduce(input, 1, func(acc, x int) int { return acc * x })
	expected2 := 120
	if result2 != expected2 {
		t.Errorf("Reduce (product): expected %d, got %d", expected2, result2)
	}
}