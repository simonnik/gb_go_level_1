package calculator

import (
	"math"
)

type Calc struct {
	X        float64
	Y        float64
	Operator string
}

func (c *Calc) Add() float64 {
	return c.X + c.Y
}

func (c *Calc) Substract() float64 {
	return c.X - c.Y
}

func (c *Calc) Multiply() float64 {
	return c.X * c.Y
}

func (c *Calc) Divide() float64 {
	if c.Y == 0 {
		return math.Inf(int(c.X))
	} else {
		return c.X / c.Y
	}
}

func (c *Calc) Calculate() float64 {
	var result float64
	switch c.Operator {
	case "+":
		result = c.Add()
	case "-":
		result = c.Substract()
	case "*":
		result = c.Multiply()
	case "/":
		result = c.Divide()
	}
	return result
}
