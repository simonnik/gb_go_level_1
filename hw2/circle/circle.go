package circle

import (
	"math"
)

func Diameter(radius float64) float64 {
	return 2 * radius
}
func Length(diameter float64) float64 {
	return diameter * math.Pi
}
