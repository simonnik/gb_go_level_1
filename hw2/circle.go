package main

import (
	"fmt"
	"github.com/simonnik/gb_go_level_1/hw2/circle"
	"math"
)

func main() {
	fmt.Println("Введите площадь круга:")
	var s float64
	_, errScan := fmt.Scan(&s)

	if errScan != nil {
		fmt.Println("Введите число")
		return
	}

	radius := math.Sqrt(s / math.Pi)
	diameter := circle.Diameter(radius)
	length := circle.Length(diameter)

	fmt.Println("Длина окружности:", length)
	fmt.Println("Диаметр окружности:", diameter)
}
