package calculator

import (
	"testing"
)

func TestCalculator_Calculate(t *testing.T) {
	calc := NewCalculator()
	
	tests := []struct {
		name      string
		left      float64
		right     float64
		operator  string
		expected  float64
		wantError bool
	}{
		{"addition", 5.0, 3.0, "+", 8.0, false},
		{"subtraction", 10.0, 4.0, "-", 6.0, false},
		{"multiplication", 6.0, 7.0, "*", 42.0, false},
		{"division", 15.0, 3.0, "/", 5.0, false},
		{"power", 2.0, 3.0, "^", 8.0, false},
		{"modulo", 17.0, 5.0, "%", 2.0, false},
		{"division by zero", 10.0, 0.0, "/", 0.0, true},
		{"modulo by zero", 10.0, 0.0, "%", 0.0, true},
		{"invalid operator", 5.0, 3.0, "@", 0.0, true},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := calc.Calculate(tt.left, tt.right, tt.operator)
			
			if tt.wantError {
				if err == nil {
					t.Errorf("Calculate(%.2f, %.2f, %s) expected error, got nil", 
						tt.left, tt.right, tt.operator)
				}
				return
			}
			
			if err != nil {
				t.Errorf("Calculate(%.2f, %.2f, %s) unexpected error: %v", 
					tt.left, tt.right, tt.operator, err)
				return
			}
			
			if result != tt.expected {
				t.Errorf("Calculate(%.2f, %.2f, %s) = %.2f; want %.2f", 
					tt.left, tt.right, tt.operator, result, tt.expected)
			}
		})
	}
}

func TestCalculator_History(t *testing.T) {
	calc := NewCalculator()
	
	// Initially history should be empty
	if len(calc.GetHistory()) != 0 {
		t.Errorf("New calculator should have empty history, got %d items", len(calc.GetHistory()))
	}
	
	// Perform some calculations
	calc.Calculate(5, 3, "+")
	calc.Calculate(10, 2, "*")
	calc.Calculate(15, 3, "/")
	
	history := calc.GetHistory()
	if len(history) != 3 {
		t.Errorf("Expected 3 items in history, got %d", len(history))
	}
	
	// Check first calculation
	if history[0].Left != 5 || history[0].Right != 3 || history[0].Operator != "+" || history[0].Result != 8 {
		t.Errorf("First calculation incorrect: got %+v", history[0])
	}
	
	// Clear history
	calc.ClearHistory()
	if len(calc.GetHistory()) != 0 {
		t.Errorf("History should be empty after clear, got %d items", len(calc.GetHistory()))
	}
}

func BenchmarkCalculator_Addition(b *testing.B) {
	calc := NewCalculator()
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		calc.Calculate(123.456, 789.123, "+")
	}
}

func BenchmarkCalculator_Division(b *testing.B) {
	calc := NewCalculator()
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		calc.Calculate(1000.0, 7.0, "/")
	}
}