package calculator_test

import (
	calc "github.com/simonnik/gb_go_level_1/hw3/calculator"
	"testing"
)

func TestCalculate(t *testing.T) {
	testCases := []struct {
		Name string
		X    float64
		Y    float64
		Op   string
		Out  float64
	}{
		{Name: "Add", X: 1, Y: 1, Op: "+", Out: 2},
		{Name: "Substract", X: 10, Y: 2, Op: "-", Out: 8},
		{Name: "Multiply", X: 12, Y: 4, Op: "*", Out: 48},
		{Name: "Divide", X: 63, Y: 7, Op: "/", Out: 9},
	}

	for _, tt := range testCases {
		t.Run(tt.Name, func(t *testing.T) {
			c := calc.Calc{X: tt.X, Y: tt.Y, Operator: tt.Op}
			if got := c.Calculate(); got != tt.Out {
				t.Fatalf("got %v, but want %v", got, tt.Out)
			}
		})
	}

}

var result float64

func BenchmarkCalculate(b *testing.B) {
	var res float64
	c := calc.Calc{X: 30, Y: 5, Operator: "*"}

	for i := 0; i < b.N; i++ {
		res = c.Calculate()
	}
	result = res
}
