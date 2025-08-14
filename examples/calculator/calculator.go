// Package calculator provides a simple command-line calculator
package calculator

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Operation represents a mathematical operation
type Operation struct {
	Left     float64
	Right    float64
	Operator string
	Result   float64
}

// Calculator performs mathematical operations
type Calculator struct {
	History []Operation
}

// NewCalculator creates a new calculator instance
func NewCalculator() *Calculator {
	return &Calculator{
		History: make([]Operation, 0),
	}
}

// Calculate performs a calculation and adds it to history
func (c *Calculator) Calculate(left, right float64, operator string) (float64, error) {
	var result float64
	var err error
	
	switch operator {
	case "+":
		result = left + right
	case "-":
		result = left - right
	case "*":
		result = left * right
	case "/":
		if right == 0 {
			return 0, errors.New("division by zero")
		}
		result = left / right
	case "^", "**":
		result = power(left, right)
	case "%":
		if right == 0 {
			return 0, errors.New("modulo by zero")
		}
		result = float64(int(left) % int(right))
	default:
		return 0, fmt.Errorf("unsupported operator: %s", operator)
	}
	
	// Add to history
	operation := Operation{
		Left:     left,
		Right:    right,
		Operator: operator,
		Result:   result,
	}
	c.History = append(c.History, operation)
	
	return result, err
}

// power calculates left raised to the power of right
func power(base, exponent float64) float64 {
	if exponent == 0 {
		return 1
	}
	if exponent == 1 {
		return base
	}
	
	result := base
	for i := 1; i < int(exponent); i++ {
		result *= base
	}
	return result
}

// GetHistory returns the calculation history
func (c *Calculator) GetHistory() []Operation {
	return c.History
}

// ClearHistory clears the calculation history
func (c *Calculator) ClearHistory() {
	c.History = make([]Operation, 0)
}

// PrintHistory prints the calculation history
func (c *Calculator) PrintHistory() {
	if len(c.History) == 0 {
		fmt.Println("No calculations in history")
		return
	}
	
	fmt.Println("Calculation History:")
	fmt.Println("-------------------")
	for i, op := range c.History {
		fmt.Printf("%d. %.2f %s %.2f = %.2f\n", 
			i+1, op.Left, op.Operator, op.Right, op.Result)
	}
}

// RunInteractive starts an interactive calculator session
func (c *Calculator) RunInteractive() {
	reader := bufio.NewReader(os.Stdin)
	
	fmt.Println("=== Go Calculator ===")
	fmt.Println("Supported operations: +, -, *, /, ^, %")
	fmt.Println("Commands: history, clear, help, quit")
	fmt.Println("Enter expressions like: 5 + 3")
	fmt.Println("=====================")
	
	for {
		fmt.Print("calc> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Error reading input: %v\n", err)
			continue
		}
		
		input = strings.TrimSpace(input)
		if input == "" {
			continue
		}
		
		// Handle commands
		switch strings.ToLower(input) {
		case "quit", "exit", "q":
			fmt.Println("Goodbye!")
			return
		case "history", "h":
			c.PrintHistory()
			continue
		case "clear", "c":
			c.ClearHistory()
			fmt.Println("History cleared")
			continue
		case "help":
			c.printHelp()
			continue
		}
		
		// Parse and evaluate expression
		if err := c.processExpression(input); err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}
}

// processExpression parses and evaluates a mathematical expression
func (c *Calculator) processExpression(input string) error {
	// Simple parsing: "left operator right"
	parts := strings.Fields(input)
	if len(parts) != 3 {
		return errors.New("invalid expression format. Use: number operator number")
	}
	
	left, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		return fmt.Errorf("invalid left operand: %s", parts[0])
	}
	
	operator := parts[1]
	
	right, err := strconv.ParseFloat(parts[2], 64)
	if err != nil {
		return fmt.Errorf("invalid right operand: %s", parts[2])
	}
	
	result, err := c.Calculate(left, right, operator)
	if err != nil {
		return err
	}
	
	fmt.Printf("%.2f %s %.2f = %.2f\n", left, operator, right, result)
	return nil
}

// printHelp prints help information
func (c *Calculator) printHelp() {
	fmt.Println("Calculator Help:")
	fmt.Println("  Expressions: number operator number")
	fmt.Println("  Examples:")
	fmt.Println("    5 + 3")
	fmt.Println("    10.5 - 2.3")
	fmt.Println("    4 * 7")
	fmt.Println("    15 / 3")
	fmt.Println("    2 ^ 8")
	fmt.Println("    17 % 5")
	fmt.Println("  Commands:")
	fmt.Println("    history (h) - show calculation history")
	fmt.Println("    clear (c)   - clear history")
	fmt.Println("    help        - show this help")
	fmt.Println("    quit (q)    - exit calculator")
}

// DemonstrateCalculator shows calculator functionality
func DemonstrateCalculator() {
	fmt.Println("=== Simple Calculator Demo ===")
	
	calc := NewCalculator()
	
	// Perform some calculations
	operations := []struct {
		left, right float64
		operator    string
	}{
		{10, 5, "+"},
		{20, 8, "-"},
		{7, 6, "*"},
		{15, 3, "/"},
		{2, 8, "^"},
		{17, 5, "%"},
	}
	
	fmt.Println("Performing calculations:")
	for _, op := range operations {
		result, err := calc.Calculate(op.left, op.right, op.operator)
		if err != nil {
			fmt.Printf("Error: %.2f %s %.2f -> %v\n", 
				op.left, op.operator, op.right, err)
		} else {
			fmt.Printf("%.2f %s %.2f = %.2f\n", 
				op.left, op.operator, op.right, result)
		}
	}
	
	fmt.Println()
	calc.PrintHistory()
	
	// Test error cases
	fmt.Println("\nTesting error cases:")
	_, err := calc.Calculate(10, 0, "/")
	if err != nil {
		fmt.Printf("Division by zero: %v\n", err)
	}
	
	_, err = calc.Calculate(5, 3, "@")
	if err != nil {
		fmt.Printf("Invalid operator: %v\n", err)
	}
}

// RunCalculatorCLI starts the interactive calculator
func RunCalculatorCLI() {
	calc := NewCalculator()
	calc.RunInteractive()
}