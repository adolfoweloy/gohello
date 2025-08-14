// Package basics demonstrates fundamental Go control structures including
// if/else statements, for loops, switch statements, and range.
package basics

import (
	"fmt"
	"math/rand"
	"time"
)

// DemoIfElse demonstrates if/else statements in Go
func DemoIfElse() {
	fmt.Println("\n=== If/Else Statements Demo ===")

	// Basic if statement
	age := 18
	if age >= 18 {
		fmt.Printf("Age %d: You are an adult\n", age)
	}

	// If with else
	score := 85
	if score >= 90 {
		fmt.Printf("Score %d: Grade A\n", score)
	} else {
		fmt.Printf("Score %d: Grade B or lower\n", score)
	}

	// If with else if
	temperature := 25
	if temperature > 30 {
		fmt.Printf("Temperature %d°C: Hot\n", temperature)
	} else if temperature > 20 {
		fmt.Printf("Temperature %d°C: Warm\n", temperature)
	} else if temperature > 10 {
		fmt.Printf("Temperature %d°C: Cool\n", temperature)
	} else {
		fmt.Printf("Temperature %d°C: Cold\n", temperature)
	}

	// If with initialization statement
	if randomNum := rand.Intn(100); randomNum > 50 {
		fmt.Printf("Random number %d is greater than 50\n", randomNum)
	} else {
		fmt.Printf("Random number %d is 50 or less\n", randomNum)
	}

	// Complex conditions
	username := "alice"
	password := "secret123"
	if len(username) > 0 && len(password) >= 8 {
		fmt.Println("Login credentials are valid format")
	} else {
		fmt.Println("Invalid credentials format")
	}

	// Checking for nil or empty
	var data []string
	if data != nil && len(data) > 0 {
		fmt.Println("Data is available")
	} else {
		fmt.Println("No data available")
	}
}

// DemoForLoops demonstrates different types of for loops in Go
func DemoForLoops() {
	fmt.Println("\n=== For Loops Demo ===")

	// Classic for loop with three parts
	fmt.Println("Classic for loop (1 to 5):")
	for i := 1; i <= 5; i++ {
		fmt.Printf("  Iteration %d\n", i)
	}

	// For loop as a while loop
	fmt.Println("\nFor as while loop (countdown from 3):")
	countdown := 3
	for countdown > 0 {
		fmt.Printf("  %d...\n", countdown)
		countdown--
	}
	fmt.Println("  Blast off! 🚀")

	// Infinite loop with break
	fmt.Println("\nInfinite loop with break:")
	counter := 0
	for {
		counter++
		if counter > 3 {
			break
		}
		fmt.Printf("  Counter: %d\n", counter)
	}

	// For loop with continue
	fmt.Println("\nFor loop with continue (skip even numbers):")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue // skip even numbers
		}
		fmt.Printf("  Odd number: %d\n", i)
	}

	// Multiple variables in for loop
	fmt.Println("\nMultiple variables:")
	for i, j := 0, 10; i < 5; i, j = i+1, j-1 {
		fmt.Printf("  i=%d, j=%d\n", i, j)
	}
}

// DemoRangeLoop demonstrates range loops over different data types
func DemoRangeLoop() {
	fmt.Println("\n=== Range Loops Demo ===")

	// Range over slice
	fruits := []string{"apple", "banana", "orange"}
	fmt.Println("Range over slice:")
	for index, fruit := range fruits {
		fmt.Printf("  [%d]: %s\n", index, fruit)
	}

	// Range over slice (index only)
	fmt.Println("\nRange over slice (index only):")
	for index := range fruits {
		fmt.Printf("  Index: %d\n", index)
	}

	// Range over slice (value only)
	fmt.Println("\nRange over slice (value only):")
	for _, fruit := range fruits {
		fmt.Printf("  Fruit: %s\n", fruit)
	}

	// Range over map
	scores := map[string]int{
		"Alice":   95,
		"Bob":     87,
		"Charlie": 92,
	}
	fmt.Println("\nRange over map:")
	for name, score := range scores {
		fmt.Printf("  %s: %d\n", name, score)
	}

	// Range over string (runes)
	fmt.Println("\nRange over string (characters):")
	text := "Hello 世界"
	for index, char := range text {
		fmt.Printf("  [%d]: %c (Unicode: %d)\n", index, char, char)
	}

	// Range over array
	numbers := [4]int{10, 20, 30, 40}
	fmt.Println("\nRange over array:")
	for index, value := range numbers {
		fmt.Printf("  [%d]: %d\n", index, value)
	}

	// Range over channel (will demonstrate in concurrency section)
	fmt.Println("\nNote: Range over channels will be covered in concurrency section")
}

// DemoSwitchStatement demonstrates switch statements in Go
func DemoSwitchStatement() {
	fmt.Println("\n=== Switch Statements Demo ===")

	// Basic switch
	day := time.Now().Weekday()
	fmt.Printf("Today is %s: ", day)
	switch day {
	case time.Monday:
		fmt.Println("Start of the work week")
	case time.Tuesday, time.Wednesday, time.Thursday:
		fmt.Println("Midweek")
	case time.Friday:
		fmt.Println("TGIF!")
	case time.Saturday, time.Sunday:
		fmt.Println("Weekend!")
	default:
		fmt.Println("Unknown day")
	}

	// Switch with expressions
	score := 85
	fmt.Printf("\nScore %d: ", score)
	switch {
	case score >= 90:
		fmt.Println("Grade A - Excellent!")
	case score >= 80:
		fmt.Println("Grade B - Good!")
	case score >= 70:
		fmt.Println("Grade C - Average")
	case score >= 60:
		fmt.Println("Grade D - Below Average")
	default:
		fmt.Println("Grade F - Failing")
	}

	// Switch with initialization
	switch age := 25; {
	case age < 13:
		fmt.Printf("Age %d: Child\n", age)
	case age < 20:
		fmt.Printf("Age %d: Teenager\n", age)
	case age < 65:
		fmt.Printf("Age %d: Adult\n", age)
	default:
		fmt.Printf("Age %d: Senior\n", age)
	}

	// Type switch (will be covered more in interfaces section)
	var value interface{} = "hello"
	fmt.Printf("\nType switch for value '%v': ", value)
	switch v := value.(type) {
	case string:
		fmt.Printf("String with length %d\n", len(v))
	case int:
		fmt.Printf("Integer with value %d\n", v)
	case bool:
		fmt.Printf("Boolean with value %t\n", v)
	default:
		fmt.Printf("Unknown type %T\n", v)
	}

	// Switch with fallthrough
	number := 2
	fmt.Printf("\nSwitch with fallthrough for number %d:\n", number)
	switch number {
	case 1:
		fmt.Println("  One")
		fallthrough
	case 2:
		fmt.Println("  Two")
		fallthrough
	case 3:
		fmt.Println("  Three or less")
	default:
		fmt.Println("  Greater than three")
	}
}

// DemoNestedLoops demonstrates nested loops and control flow
func DemoNestedLoops() {
	fmt.Println("\n=== Nested Loops Demo ===")

	// Multiplication table
	fmt.Println("Multiplication table (3x3):")
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Printf("%2d ", i*j)
		}
		fmt.Println()
	}

	// Pattern printing
	fmt.Println("\nTriangle pattern:")
	for i := 1; i <= 5; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}

	// Breaking out of nested loops with labels
	fmt.Println("\nBreaking out of nested loops with labels:")
	outer:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Printf("  i=%d, j=%d\n", i, j)
			if i == 2 && j == 2 {
				fmt.Println("  Breaking out of both loops")
				break outer
			}
		}
	}
	fmt.Println("  Exited nested loops")
}

// Grade calculates letter grade based on numeric score
func Grade(score int) string {
	switch {
	case score >= 90:
		return "A"
	case score >= 80:
		return "B"
	case score >= 70:
		return "C"
	case score >= 60:
		return "D"
	default:
		return "F"
	}
}

// IsWeekday checks if a time.Weekday is a weekday
func IsWeekday(day time.Weekday) bool {
	switch day {
	case time.Saturday, time.Sunday:
		return false
	default:
		return true
	}
}

// CountVowels counts vowels in a string using range and switch
func CountVowels(text string) int {
	count := 0
	for _, char := range text {
		switch char {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			count++
		}
	}
	return count
}

// FindMax finds the maximum value in a slice using for loop
func FindMax(numbers []int) (int, bool) {
	if len(numbers) == 0 {
		return 0, false
	}

	max := numbers[0]
	for _, num := range numbers {
		if num > max {
			max = num
		}
	}
	return max, true
}

// SumEven calculates the sum of even numbers up to n
func SumEven(n int) int {
	sum := 0
	for i := 2; i <= n; i += 2 {
		sum += i
	}
	return sum
}

// FizzBuzz implements the classic FizzBuzz problem
func FizzBuzz(n int) []string {
	result := make([]string, n)
	for i := 1; i <= n; i++ {
		switch {
		case i%15 == 0:
			result[i-1] = "FizzBuzz"
		case i%3 == 0:
			result[i-1] = "Fizz"
		case i%5 == 0:
			result[i-1] = "Buzz"
		default:
			result[i-1] = fmt.Sprintf("%d", i)
		}
	}
	return result
}

// DemoControlFlowExamples demonstrates practical examples
func DemoControlFlowExamples() {
	fmt.Println("\n=== Practical Control Flow Examples ===")

	// Grade calculation
	scores := []int{95, 87, 76, 62, 45}
	fmt.Println("Grade calculations:")
	for i, score := range scores {
		grade := Grade(score)
		fmt.Printf("  Student %d: Score %d -> Grade %s\n", i+1, score, grade)
	}

	// Vowel counting
	text := "Hello World"
	vowelCount := CountVowels(text)
	fmt.Printf("\nVowels in '%s': %d\n", text, vowelCount)

	// Find maximum
	numbers := []int{23, 45, 12, 67, 34, 89, 56}
	if max, found := FindMax(numbers); found {
		fmt.Printf("Maximum in %v: %d\n", numbers, max)
	}

	// Sum of even numbers
	n := 10
	evenSum := SumEven(n)
	fmt.Printf("Sum of even numbers up to %d: %d\n", n, evenSum)

	// FizzBuzz
	fmt.Println("\nFizzBuzz (1-15):")
	fizzBuzz := FizzBuzz(15)
	for i, value := range fizzBuzz {
		fmt.Printf("  %2d: %s\n", i+1, value)
	}
}