package main

import (
	"fmt"
)

func main() {
	fmt.Println("Введите число a:")
	var a float64
	_, err := fmt.Scan(&a)

	if err != nil {
		fmt.Println("Введите число")
		return
	}

	fmt.Println("Введите число b:")
	var b float64
	_, err = fmt.Scan(&b)

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

	switch operation {
	case "+":
		fmt.Println("Результат:", a+b)
	case "-":
		fmt.Println("Результат:", a-b)
	case "*":
		fmt.Println("Результат:", a*b)
	case "/":
		if b != 0 {
			fmt.Println("Результат:", a/b)
		} else {
			fmt.Println("Деление на 0")
		}
	default:
		fmt.Println("Неизвестная операция")
	}
}
