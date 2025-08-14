package basics

import "fmt"

// DemonstrateControlStructures shows if/else, for loops, and switch statements
func DemonstrateControlStructures() {
	fmt.Println("\n=== Control Structures ===")
	
	// If/else statements
	score := 85
	if score >= 90 {
		fmt.Printf("Score %d: Grade A\n", score)
	} else if score >= 80 {
		fmt.Printf("Score %d: Grade B\n", score)
	} else if score >= 70 {
		fmt.Printf("Score %d: Grade C\n", score)
	} else {
		fmt.Printf("Score %d: Grade F\n", score)
	}
	
	// If with initialization
	if remainder := score % 10; remainder == 0 {
		fmt.Printf("Score %d is divisible by 10\n", score)
	} else {
		fmt.Printf("Score %d has remainder %d when divided by 10\n", score, remainder)
	}
}

// DemonstrateLoops shows different types of for loops
func DemonstrateLoops() {
	fmt.Println("\n=== For Loops ===")
	
	// Basic for loop
	fmt.Println("Counting from 1 to 5:")
	for i := 1; i <= 5; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
	
	// For loop as while
	fmt.Println("Countdown using for as while:")
	count := 5
	for count > 0 {
		fmt.Printf("%d ", count)
		count--
	}
	fmt.Println("Blast off!")
	
	// Infinite loop with break
	fmt.Println("Finding first number divisible by 7 starting from 20:")
	num := 20
	for {
		if num%7 == 0 {
			fmt.Printf("Found: %d\n", num)
			break
		}
		num++
	}
	
	// Range over slice
	fruits := []string{"apple", "banana", "cherry", "date"}
	fmt.Println("Fruits with index:")
	for index, fruit := range fruits {
		fmt.Printf("  [%d]: %s\n", index, fruit)
	}
	
	// Range over map
	ages := map[string]int{"Alice": 30, "Bob": 25, "Charlie": 35}
	fmt.Println("Ages:")
	for name, age := range ages {
		fmt.Printf("  %s: %d years old\n", name, age)
	}
	
	// Range with underscore to ignore index
	fmt.Println("Just the fruits:")
	for _, fruit := range fruits {
		fmt.Printf("  %s\n", fruit)
	}
}

// DemonstrateSwitch shows switch statement variations
func DemonstrateSwitch() {
	fmt.Println("\n=== Switch Statements ===")
	
	// Basic switch
	day := "Wednesday"
	switch day {
	case "Monday":
		fmt.Println("Start of work week")
	case "Wednesday":
		fmt.Println("Hump day!")
	case "Friday":
		fmt.Println("TGIF!")
	case "Saturday", "Sunday":
		fmt.Println("Weekend!")
	default:
		fmt.Println("Regular day")
	}
	
	// Switch with expression
	hour := 14
	switch {
	case hour < 12:
		fmt.Println("Good morning!")
	case hour < 17:
		fmt.Println("Good afternoon!")
	case hour < 21:
		fmt.Println("Good evening!")
	default:
		fmt.Println("Good night!")
	}
	
	// Switch with initialization
	switch time := hour % 12; time {
	case 0:
		fmt.Println("It's 12 o'clock")
	case 1, 2, 3:
		fmt.Println("Early in the period")
	default:
		fmt.Printf("It's %d o'clock in 12-hour format\n", time)
	}
	
	// Type switch (we'll use interface{} here)
	demonstrateTypeSwitch()
}

// demonstrateTypeSwitch shows type switching
func demonstrateTypeSwitch() {
	fmt.Println("\nType switch examples:")
	values := []interface{}{42, "hello", 3.14, true, []int{1, 2, 3}}
	
	for i, v := range values {
		switch value := v.(type) {
		case int:
			fmt.Printf("  [%d]: integer %d\n", i, value)
		case string:
			fmt.Printf("  [%d]: string '%s'\n", i, value)
		case float64:
			fmt.Printf("  [%d]: float %.2f\n", i, value)
		case bool:
			fmt.Printf("  [%d]: boolean %t\n", i, value)
		case []int:
			fmt.Printf("  [%d]: slice of ints %v\n", i, value)
		default:
			fmt.Printf("  [%d]: unknown type %T\n", i, value)
		}
	}
}