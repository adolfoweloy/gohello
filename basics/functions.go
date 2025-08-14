package basics

import "fmt"

// Add returns the sum of two integers
func Add(a, b int) int {
	return a + b
}

// Divide returns the quotient and remainder of dividing a by b
// This demonstrates multiple return values
func Divide(a, b int) (int, int, error) {
	if b == 0 {
		return 0, 0, fmt.Errorf("cannot divide by zero")
	}
	quotient := a / b
	remainder := a % b
	return quotient, remainder, nil
}

// CalculateStats returns min, max, and average of a slice of integers
// This demonstrates named return values
func CalculateStats(numbers []int) (min, max int, avg float64) {
	if len(numbers) == 0 {
		return 0, 0, 0
	}
	
	min = numbers[0]
	max = numbers[0]
	sum := 0
	
	for _, num := range numbers {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
		sum += num
	}
	
	avg = float64(sum) / float64(len(numbers))
	return // naked return - returns named variables
}

// Greet is a higher-order function that takes a function as parameter
func Greet(name string, formatter func(string) string) string {
	return formatter(name)
}

// MakeMultiplier returns a function that multiplies by the given factor
// This demonstrates closures
func MakeMultiplier(factor int) func(int) int {
	return func(n int) int {
		return n * factor
	}
}

// DemonstrateFunctions shows various function concepts
func DemonstrateFunctions() {
	fmt.Println("\n=== Functions ===")
	
	// Basic function call
	result := Add(5, 3)
	fmt.Printf("5 + 3 = %d\n", result)
	
	// Multiple return values
	quotient, remainder, err := Divide(17, 5)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("17 ÷ 5 = %d remainder %d\n", quotient, remainder)
	}
	
	// Error handling
	_, _, err = Divide(10, 0)
	if err != nil {
		fmt.Printf("Division error: %v\n", err)
	}
	
	// Named return values
	numbers := []int{4, 2, 8, 1, 9, 3}
	min, max, avg := CalculateStats(numbers)
	fmt.Printf("Numbers %v: min=%d, max=%d, avg=%.2f\n", numbers, min, max, avg)
	
	// Function as parameter
	formalGreeting := func(name string) string {
		return fmt.Sprintf("Good day, %s!", name)
	}
	casualGreeting := func(name string) string {
		return fmt.Sprintf("Hey %s!", name)
	}
	
	fmt.Println(Greet("Alice", formalGreeting))
	fmt.Println(Greet("Bob", casualGreeting))
	
	// Closures
	doubler := MakeMultiplier(2)
	tripler := MakeMultiplier(3)
	
	fmt.Printf("Doubler(5) = %d\n", doubler(5))
	fmt.Printf("Tripler(4) = %d\n", tripler(4))
	
	// Anonymous function
	square := func(x int) int {
		return x * x
	}
	fmt.Printf("Square of 6 = %d\n", square(6))
	
	// Variadic function
	demonstrateVariadic()
}

// Sum is a variadic function that sums any number of integers
func Sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// PrintInfo is a variadic function with mixed parameters
func PrintInfo(prefix string, items ...string) {
	fmt.Printf("%s: ", prefix)
	for i, item := range items {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(item)
	}
	fmt.Println()
}

// demonstrateVariadic shows variadic function usage
func demonstrateVariadic() {
	fmt.Println("\n--- Variadic Functions ---")
	
	// Calling with different numbers of arguments
	fmt.Printf("Sum() = %d\n", Sum())
	fmt.Printf("Sum(1) = %d\n", Sum(1))
	fmt.Printf("Sum(1, 2, 3) = %d\n", Sum(1, 2, 3))
	fmt.Printf("Sum(1, 2, 3, 4, 5) = %d\n", Sum(1, 2, 3, 4, 5))
	
	// Expanding a slice
	nums := []int{10, 20, 30, 40}
	fmt.Printf("Sum(nums...) = %d\n", Sum(nums...))
	
	// Mixed parameters
	PrintInfo("Languages", "Go", "Python", "JavaScript")
	PrintInfo("Colors", "Red", "Green", "Blue", "Yellow")
}