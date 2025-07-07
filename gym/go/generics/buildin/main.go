package main

import (
	"fmt"
	// "golang.org/x/exp/constraints" // Importing constraints for generic type constraints
)

func MapValueWithoutGenerics(values []int, mapFunc func(int) int) []int {
	var newValues []int
	for _, v := range values {
		newValues = append(newValues, mapFunc(v))
	}
	return newValues
}

func MapValueWithGenerics[T any](values []T, mapFunc func(T) T) []T {
	var newValues []T
	for _, v := range values {
		newValues = append(newValues, mapFunc(v))
	}
	return newValues
}

func main() {
	values := []int{1, 2, 3, 4, 5}
	result := MapValueWithoutGenerics(values, func(v int) int {
		return v * 2
	})

	fmt.Println("Mapped values without generics:", result)

	newValues := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	resultWithGenerics := MapValueWithGenerics(newValues, func(v float64) float64 {
		return v * 2.0
	})
	fmt.Println("Mapped values with generics:", resultWithGenerics)
}