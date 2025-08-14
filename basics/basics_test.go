package basics

import (
	"testing"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"positive numbers", 5, 3, 8},
		{"negative numbers", -2, -3, -5},
		{"mixed signs", -5, 3, -2},
		{"zero values", 0, 0, 0},
		{"with zero", 10, 0, 10},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Add(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Add(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	tests := []struct {
		name      string
		a, b      int
		wantQuot  int
		wantRem   int
		wantError bool
	}{
		{"normal division", 10, 3, 3, 1, false},
		{"exact division", 15, 5, 3, 0, false},
		{"division by zero", 10, 0, 0, 0, true},
		{"negative numbers", -15, 3, -5, 0, false},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			quot, rem, err := Divide(tt.a, tt.b)
			
			if tt.wantError {
				if err == nil {
					t.Errorf("Divide(%d, %d) expected error, got nil", tt.a, tt.b)
				}
				return
			}
			
			if err != nil {
				t.Errorf("Divide(%d, %d) unexpected error: %v", tt.a, tt.b, err)
				return
			}
			
			if quot != tt.wantQuot {
				t.Errorf("Divide(%d, %d) quotient = %d; want %d", tt.a, tt.b, quot, tt.wantQuot)
			}
			
			if rem != tt.wantRem {
				t.Errorf("Divide(%d, %d) remainder = %d; want %d", tt.a, tt.b, rem, tt.wantRem)
			}
		})
	}
}

func TestCalculateStats(t *testing.T) {
	tests := []struct {
		name    string
		numbers []int
		wantMin int
		wantMax int
		wantAvg float64
	}{
		{"normal case", []int{1, 2, 3, 4, 5}, 1, 5, 3.0},
		{"single number", []int{42}, 42, 42, 42.0},
		{"negative numbers", []int{-3, -1, -5}, -5, -1, -3.0},
		{"mixed numbers", []int{-2, 0, 5, 1}, -2, 5, 1.0},
		{"empty slice", []int{}, 0, 0, 0.0},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			min, max, avg := CalculateStats(tt.numbers)
			
			if min != tt.wantMin {
				t.Errorf("CalculateStats(%v) min = %d; want %d", tt.numbers, min, tt.wantMin)
			}
			
			if max != tt.wantMax {
				t.Errorf("CalculateStats(%v) max = %d; want %d", tt.numbers, max, tt.wantMax)
			}
			
			if avg != tt.wantAvg {
				t.Errorf("CalculateStats(%v) avg = %.2f; want %.2f", tt.numbers, avg, tt.wantAvg)
			}
		})
	}
}

func TestSum(t *testing.T) {
	tests := []struct {
		name     string
		numbers  []int
		expected int
	}{
		{"no numbers", []int{}, 0},
		{"single number", []int{5}, 5},
		{"multiple numbers", []int{1, 2, 3, 4}, 10},
		{"negative numbers", []int{-1, -2, -3}, -6},
		{"mixed numbers", []int{-5, 10, -3, 8}, 10},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Sum(tt.numbers...)
			if result != tt.expected {
				t.Errorf("Sum(%v) = %d; want %d", tt.numbers, result, tt.expected)
			}
		})
	}
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(123, 456)
	}
}

func BenchmarkCalculateStats(b *testing.B) {
	numbers := []int{1, 5, 3, 9, 2, 8, 4, 7, 6}
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		CalculateStats(numbers)
	}
}