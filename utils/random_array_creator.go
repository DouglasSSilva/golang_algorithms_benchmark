package utils

import (
	"math/rand"
	"time"
)

// RandomInt generates a rando int in range min, max
func RandomInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min)
}

// RandomArrayCreator to be used on sort operations
func RandomArrayCreator(size, minValue, maxValue int) []int {
	arr := []int{}
	for i := 0; i < size; i++ {
		arr = append(arr, RandomInt(minValue, maxValue))
	}

	return arr
}
