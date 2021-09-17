package main

import (
	"fmt"
	"github.com/simonnik/gb_go_level_1/hw3/calculator"
)

func main() {
	fmt.Println("Введите число x:")
	var x float64
	_, err := fmt.Scan(&x)

	if err != nil {
		fmt.Println("Введите число")
		return
	}

	fmt.Println("Введите число y:")
	var y float64
	_, err = fmt.Scan(&y)

	if err != nil {
		fmt.Println("Введите число")
		return
	}

	fmt.Println("Выберите операцию: +, -, *, /")

	var operation string

	_, err = fmt.Scanln(&operation)

	if err != nil {
		fmt.Println("Выберите операцию")
		return
	}

	c := calculator.Calc{X: x, Y: y, Operator: operation}
	result := c.Calculate()
	fmt.Printf("%v %s %v = %v\n", c.X, operation, c.Y, result)
}
