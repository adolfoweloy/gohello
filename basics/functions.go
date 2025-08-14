// Package basics demonstrates Go functions including multiple return values,
// named returns, variadic functions, and function types.
package basics

import (
	"errors"
	"fmt"
	"strings"
)

// DemoBasicFunctions demonstrates basic function syntax and usage
func DemoBasicFunctions() {
	fmt.Println("\n=== Basic Functions Demo ===")

	// Simple function call
	greeting := greet("Alice")
	fmt.Println(greeting)

	// Function with multiple parameters
	sum := add(10, 20)
	fmt.Printf("10 + 20 = %d\n", sum)

	// Function with different parameter types
	info := formatPersonInfo("Bob", 25, true)
	fmt.Println(info)
}

// greet is a simple function that takes a string and returns a string
func greet(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

// add takes two integers and returns their sum
func add(a, b int) int {
	return a + b
}

// formatPersonInfo demonstrates a function with multiple different parameter types
func formatPersonInfo(name string, age int, isActive bool) string {
	status := "inactive"
	if isActive {
		status = "active"
	}
	return fmt.Sprintf("%s is %d years old and is %s", name, age, status)
}

// DemoMultipleReturns demonstrates functions with multiple return values
func DemoMultipleReturns() {
	fmt.Println("\n=== Multiple Return Values Demo ===")

	// Function returning multiple values
	quotient, remainder := divide(17, 5)
	fmt.Printf("17 ÷ 5 = %d remainder %d\n", quotient, remainder)

	// Function returning value and error
	result, err := safeDivide(10, 2)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("10 ÷ 2 = %.2f\n", result)
	}

	// Error case
	_, err = safeDivide(10, 0)
	if err != nil {
		fmt.Printf("Error dividing by zero: %v\n", err)
	}

	// Ignoring return values with blank identifier
	name, _ := getPersonDetails()
	fmt.Printf("Name: %s (age ignored)\n", name)

	// Using all return values
	fullName, personAge := getPersonDetails()
	fmt.Printf("Full details: %s, age %d\n", fullName, personAge)
}

// divide returns both quotient and remainder
func divide(dividend, divisor int) (int, int) {
	quotient := dividend / divisor
	remainder := dividend % divisor
	return quotient, remainder
}

// safeDivide performs division with error handling
func safeDivide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// getPersonDetails returns multiple values for demonstration
func getPersonDetails() (string, int) {
	return "John Doe", 30
}

// DemoNamedReturns demonstrates named return values
func DemoNamedReturns() {
	fmt.Println("\n=== Named Return Values Demo ===")

	// Function with named returns
	area, perimeter := calculateRectangle(5, 3)
	fmt.Printf("Rectangle 5x3: area=%.2f, perimeter=%.2f\n", area, perimeter)

	// Function that can return early
	isValid, message := validateAge(15)
	fmt.Printf("Age 15 validation: valid=%t, message='%s'\n", isValid, message)

	isValid, message = validateAge(25)
	fmt.Printf("Age 25 validation: valid=%t, message='%s'\n", isValid, message)

	// Function with complex named returns
	stats := getStringStats("Hello World!")
	fmt.Printf("String stats: %+v\n", stats)
}

// calculateRectangle demonstrates named return parameters
func calculateRectangle(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = 2 * (length + width)
	return // naked return
}

// validateAge demonstrates named returns with early return
func validateAge(age int) (isValid bool, message string) {
	if age < 0 {
		message = "age cannot be negative"
		return // isValid is false by default
	}
	
	if age < 18 {
		message = "must be 18 or older"
		return
	}
	
	isValid = true
	message = "age is valid"
	return
}

// StringStats holds statistics about a string
type StringStats struct {
	Length    int
	Words     int
	Vowels    int
	HasUpper  bool
	HasLower  bool
	HasDigits bool
}

// getStringStats demonstrates complex named returns with struct
func getStringStats(text string) (stats StringStats) {
	stats.Length = len(text)
	stats.Words = len(strings.Fields(text))
	
	for _, char := range text {
		// Count vowels
		if strings.ContainsRune("aeiouAEIOU", char) {
			stats.Vowels++
		}
		
		// Check character types
		if char >= 'A' && char <= 'Z' {
			stats.HasUpper = true
		} else if char >= 'a' && char <= 'z' {
			stats.HasLower = true
		} else if char >= '0' && char <= '9' {
			stats.HasDigits = true
		}
	}
	
	return
}

// DemoVariadicFunctions demonstrates variadic functions
func DemoVariadicFunctions() {
	fmt.Println("\n=== Variadic Functions Demo ===")

	// Function with variable number of arguments
	sum1 := sum(1, 2, 3)
	fmt.Printf("sum(1, 2, 3) = %d\n", sum1)

	sum2 := sum(1, 2, 3, 4, 5)
	fmt.Printf("sum(1, 2, 3, 4, 5) = %d\n", sum2)

	// Using slice with variadic function
	numbers := []int{10, 20, 30, 40}
	sum3 := sum(numbers...)
	fmt.Printf("sum(10, 20, 30, 40) = %d\n", sum3)

	// Variadic function with other parameters
	message1 := formatMessage("Error", "File not found", "Please check the path")
	fmt.Println(message1)

	message2 := formatMessage("Info", "Operation completed")
	fmt.Println(message2)

	// Built-in variadic functions
	fmt.Printf("Printf is also variadic: %s %d %.2f\n", "Hello", 42, 3.14)
}

// sum is a variadic function that sums all provided integers
func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// formatMessage demonstrates variadic function with other parameters
func formatMessage(level string, parts ...string) string {
	message := strings.Join(parts, " ")
	return fmt.Sprintf("[%s] %s", level, message)
}

// DemoFunctionTypes demonstrates functions as types and values
func DemoFunctionTypes() {
	fmt.Println("\n=== Function Types Demo ===")

	// Function as variable
	var mathOp func(int, int) int
	mathOp = add
	result1 := mathOp(5, 3)
	fmt.Printf("Using function variable for addition: %d\n", result1)

	// Change the function
	mathOp = multiply
	result2 := mathOp(5, 3)
	fmt.Printf("Using function variable for multiplication: %d\n", result2)

	// Function as parameter
	fmt.Println("\nUsing functions as parameters:")
	applyOperation(10, 5, add, "addition")
	applyOperation(10, 5, multiply, "multiplication")
	applyOperation(10, 5, subtract, "subtraction")

	// Anonymous functions
	fmt.Println("\nAnonymous functions:")
	square := func(x int) int {
		return x * x
	}
	fmt.Printf("Square of 7: %d\n", square(7))

	// Immediately invoked function expression
	result := func(a, b int) int {
		return a*a + b*b
	}(3, 4)
	fmt.Printf("3² + 4² = %d\n", result)

	// Function returning function
	doubler := makeMultiplier(2)
	tripler := makeMultiplier(3)
	fmt.Printf("Double 5: %d\n", doubler(5))
	fmt.Printf("Triple 5: %d\n", tripler(5))
}

// multiply multiplies two integers
func multiply(a, b int) int {
	return a * b
}

// subtract subtracts two integers
func subtract(a, b int) int {
	return a - b
}

// applyOperation demonstrates using function as parameter
func applyOperation(a, b int, op func(int, int) int, opName string) {
	result := op(a, b)
	fmt.Printf("  %d and %d with %s = %d\n", a, b, opName, result)
}

// makeMultiplier returns a function that multiplies by the given factor
func makeMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

// DemoClosures demonstrates closures in Go
func DemoClosures() {
	fmt.Println("\n=== Closures Demo ===")

	// Closure capturing variable from outer scope
	counter := makeCounter()
	fmt.Printf("Counter: %d\n", counter())
	fmt.Printf("Counter: %d\n", counter())
	fmt.Printf("Counter: %d\n", counter())

	// Another counter instance
	counter2 := makeCounter()
	fmt.Printf("Counter2: %d\n", counter2())
	fmt.Printf("Counter: %d\n", counter())

	// Closure modifying outer variable
	value := 10
	increment := func() {
		value++
		fmt.Printf("Value incremented to: %d\n", value)
	}
	
	increment()
	increment()
	fmt.Printf("Final value: %d\n", value)

	// Closure in loop (common gotcha demonstration)
	funcs := makeFunctions()
	fmt.Println("Calling functions created in loop:")
	for i, f := range funcs {
		fmt.Printf("  Function %d returns: %d\n", i, f())
	}
}

// makeCounter returns a closure that increments and returns a counter
func makeCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// makeFunctions demonstrates closure behavior in loops
func makeFunctions() []func() int {
	var funcs []func() int
	
	for i := 0; i < 3; i++ {
		// Capture i in local variable to avoid closure gotcha
		x := i
		funcs = append(funcs, func() int {
			return x * x
		})
	}
	
	return funcs
}

// DemoRecursion demonstrates recursive functions
func DemoRecursion() {
	fmt.Println("\n=== Recursion Demo ===")

	// Factorial
	fmt.Printf("Factorial of 5: %d\n", factorial(5))
	fmt.Printf("Factorial of 0: %d\n", factorial(0))

	// Fibonacci
	fmt.Printf("Fibonacci of 8: %d\n", fibonacci(8))
	
	fmt.Println("Fibonacci sequence (0-10):")
	for i := 0; i <= 10; i++ {
		fmt.Printf("  fib(%d) = %d\n", i, fibonacci(i))
	}

	// Tree-like structure traversal
	fmt.Println("\nTree sum example:")
	root := &TreeNode{
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
	
	total := sumTree(root)
	fmt.Printf("Sum of all nodes: %d\n", total)
}

// factorial calculates factorial recursively
func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

// fibonacci calculates Fibonacci number recursively
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// TreeNode represents a binary tree node
type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

// sumTree recursively sums all values in a binary tree
func sumTree(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return node.Value + sumTree(node.Left) + sumTree(node.Right)
}

// Advanced function examples

// Map applies a function to each element of a slice
func Map(slice []int, fn func(int) int) []int {
	result := make([]int, len(slice))
	for i, v := range slice {
		result[i] = fn(v)
	}
	return result
}

// Filter returns elements that satisfy the predicate
func Filter(slice []int, predicate func(int) bool) []int {
	var result []int
	for _, v := range slice {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

// Reduce reduces slice to single value using accumulator function
func Reduce(slice []int, initial int, fn func(int, int) int) int {
	acc := initial
	for _, v := range slice {
		acc = fn(acc, v)
	}
	return acc
}

// DemoHigherOrderFunctions demonstrates higher-order functions
func DemoHigherOrderFunctions() {
	fmt.Println("\n=== Higher-Order Functions Demo ===")

	numbers := []int{1, 2, 3, 4, 5}
	fmt.Printf("Original: %v\n", numbers)

	// Map: square each number
	squares := Map(numbers, func(x int) int { return x * x })
	fmt.Printf("Squares: %v\n", squares)

	// Filter: only even numbers
	evens := Filter(numbers, func(x int) bool { return x%2 == 0 })
	fmt.Printf("Evens: %v\n", evens)

	// Reduce: sum all numbers
	total := Reduce(numbers, 0, func(acc, x int) int { return acc + x })
	fmt.Printf("Sum: %d\n", total)

	// Reduce: product of all numbers
	product := Reduce(numbers, 1, func(acc, x int) int { return acc * x })
	fmt.Printf("Product: %d\n", product)
}