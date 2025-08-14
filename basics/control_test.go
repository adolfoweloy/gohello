package basics

import (
	"testing"
	"time"
)

func TestDemoIfElse(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("DemoIfElse panicked: %v", r)
		}
	}()
	
	DemoIfElse()
}

func TestDemoForLoops(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("DemoForLoops panicked: %v", r)
		}
	}()
	
	DemoForLoops()
}

func TestDemoRangeLoop(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("DemoRangeLoop panicked: %v", r)
		}
	}()
	
	DemoRangeLoop()
}

func TestDemoSwitchStatement(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("DemoSwitchStatement panicked: %v", r)
		}
	}()
	
	DemoSwitchStatement()
}

func TestDemoNestedLoops(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("DemoNestedLoops panicked: %v", r)
		}
	}()
	
	DemoNestedLoops()
}

func TestDemoControlFlowExamples(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("DemoControlFlowExamples panicked: %v", r)
		}
	}()
	
	DemoControlFlowExamples()
}

func TestGrade(t *testing.T) {
	testCases := []struct {
		score    int
		expected string
	}{
		{95, "A"},
		{85, "B"},
		{75, "C"},
		{65, "D"},
		{55, "F"},
		{100, "A"},
		{80, "B"},
		{70, "C"},
		{60, "D"},
		{0, "F"},
	}
	
	for _, tc := range testCases {
		result := Grade(tc.score)
		if result != tc.expected {
			t.Errorf("Grade(%d): expected %s, got %s", tc.score, tc.expected, result)
		}
	}
}

func TestIsWeekday(t *testing.T) {
	testCases := []struct {
		day      time.Weekday
		expected bool
	}{
		{time.Monday, true},
		{time.Tuesday, true},
		{time.Wednesday, true},
		{time.Thursday, true},
		{time.Friday, true},
		{time.Saturday, false},
		{time.Sunday, false},
	}
	
	for _, tc := range testCases {
		result := IsWeekday(tc.day)
		if result != tc.expected {
			t.Errorf("IsWeekday(%s): expected %t, got %t", tc.day, tc.expected, result)
		}
	}
}

func TestCountVowels(t *testing.T) {
	testCases := []struct {
		text     string
		expected int
	}{
		{"hello", 2},
		{"world", 1},
		{"aeiou", 5},
		{"AEIOU", 5},
		{"Hello World", 3},
		{"bcdfg", 0},
		{"", 0},
		{"Programming", 3},
	}
	
	for _, tc := range testCases {
		result := CountVowels(tc.text)
		if result != tc.expected {
			t.Errorf("CountVowels(%s): expected %d, got %d", tc.text, tc.expected, result)
		}
	}
}

func TestFindMax(t *testing.T) {
	// Test with normal slice
	numbers1 := []int{1, 5, 3, 9, 2}
	max1, found1 := FindMax(numbers1)
	if !found1 {
		t.Errorf("FindMax should find maximum in non-empty slice")
	}
	if max1 != 9 {
		t.Errorf("FindMax([1,5,3,9,2]): expected 9, got %d", max1)
	}
	
	// Test with single element
	numbers2 := []int{42}
	max2, found2 := FindMax(numbers2)
	if !found2 {
		t.Errorf("FindMax should find maximum in single-element slice")
	}
	if max2 != 42 {
		t.Errorf("FindMax([42]): expected 42, got %d", max2)
	}
	
	// Test with empty slice
	numbers3 := []int{}
	_, found3 := FindMax(numbers3)
	if found3 {
		t.Errorf("FindMax should return false for empty slice")
	}
	
	// Test with negative numbers
	numbers4 := []int{-1, -5, -3, -2}
	max4, found4 := FindMax(numbers4)
	if !found4 {
		t.Errorf("FindMax should find maximum in slice with negative numbers")
	}
	if max4 != -1 {
		t.Errorf("FindMax([-1,-5,-3,-2]): expected -1, got %d", max4)
	}
}

func TestSumEven(t *testing.T) {
	testCases := []struct {
		n        int
		expected int
	}{
		{0, 0},
		{1, 0},
		{2, 2},
		{3, 2},
		{4, 6},   // 2 + 4
		{6, 12},  // 2 + 4 + 6
		{10, 30}, // 2 + 4 + 6 + 8 + 10
	}
	
	for _, tc := range testCases {
		result := SumEven(tc.n)
		if result != tc.expected {
			t.Errorf("SumEven(%d): expected %d, got %d", tc.n, tc.expected, result)
		}
	}
}

func TestFizzBuzz(t *testing.T) {
	// Test FizzBuzz for 15
	result := FizzBuzz(15)
	
	if len(result) != 15 {
		t.Errorf("FizzBuzz(15): expected length 15, got %d", len(result))
	}
	
	// Test specific values
	expected := []string{
		"1", "2", "Fizz", "4", "Buzz",
		"Fizz", "7", "8", "Fizz", "Buzz",
		"11", "Fizz", "13", "14", "FizzBuzz",
	}
	
	for i, exp := range expected {
		if result[i] != exp {
			t.Errorf("FizzBuzz(15)[%d]: expected %s, got %s", i, exp, result[i])
		}
	}
	
	// Test smaller case
	result5 := FizzBuzz(5)
	expected5 := []string{"1", "2", "Fizz", "4", "Buzz"}
	
	for i, exp := range expected5 {
		if result5[i] != exp {
			t.Errorf("FizzBuzz(5)[%d]: expected %s, got %s", i, exp, result5[i])
		}
	}
}

func TestControlStructureBehavior(t *testing.T) {
	// Test if statement behavior
	value := 10
	var result string
	
	if value > 5 {
		result = "greater"
	} else {
		result = "lesser"
	}
	
	if result != "greater" {
		t.Errorf("If statement: expected 'greater', got '%s'", result)
	}
	
	// Test for loop behavior
	sum := 0
	for i := 1; i <= 5; i++ {
		sum += i
	}
	
	expectedSum := 15 // 1+2+3+4+5
	if sum != expectedSum {
		t.Errorf("For loop sum: expected %d, got %d", expectedSum, sum)
	}
	
	// Test range behavior
	numbers := []int{1, 2, 3}
	total := 0
	
	for _, num := range numbers {
		total += num
	}
	
	expectedTotal := 6
	if total != expectedTotal {
		t.Errorf("Range sum: expected %d, got %d", expectedTotal, total)
	}
}

func TestSwitchBehavior(t *testing.T) {
	// Test switch with multiple cases
	getValue := func(input int) string {
		switch input {
		case 1, 2, 3:
			return "low"
		case 4, 5, 6:
			return "medium"
		case 7, 8, 9:
			return "high"
		default:
			return "unknown"
		}
	}
	
	testCases := []struct {
		input    int
		expected string
	}{
		{1, "low"},
		{2, "low"},
		{3, "low"},
		{4, "medium"},
		{5, "medium"},
		{6, "medium"},
		{7, "high"},
		{8, "high"},
		{9, "high"},
		{10, "unknown"},
		{0, "unknown"},
	}
	
	for _, tc := range testCases {
		result := getValue(tc.input)
		if result != tc.expected {
			t.Errorf("Switch test for %d: expected %s, got %s", tc.input, tc.expected, result)
		}
	}
}

func TestNestedLoopBreak(t *testing.T) {
	// Test that we can break out of nested loops
	found := false
	target := 6 // 2*3 = 6, which is achievable with i=2, j=3
	
	outer:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			if i*j == target {
				found = true
				break outer
			}
		}
	}
	
	if !found {
		t.Errorf("Should have found target %d in nested loops", target)
	}
}

func TestRangeOverDifferentTypes(t *testing.T) {
	// Test range over string
	text := "Go"
	chars := []rune{}
	
	for _, char := range text {
		chars = append(chars, char)
	}
	
	if len(chars) != 2 {
		t.Errorf("Range over string 'Go': expected 2 characters, got %d", len(chars))
	}
	
	// Test range over map
	m := map[string]int{"a": 1, "b": 2}
	count := 0
	
	for range m {
		count++
	}
	
	if count != 2 {
		t.Errorf("Range over map: expected 2 iterations, got %d", count)
	}
	
	// Test range over array
	arr := [3]int{1, 2, 3}
	sum := 0
	
	for _, value := range arr {
		sum += value
	}
	
	if sum != 6 {
		t.Errorf("Range over array: expected sum 6, got %d", sum)
	}
}