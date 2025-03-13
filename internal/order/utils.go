package order

import (
	"math/rand"
)

func SimulatePrice(basePrice float64) float64 {
	// Random value between -1.00 and +1.00
	change := (rand.Float64() * 2) - 1
	return basePrice + change
}