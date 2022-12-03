package sh

import (
	"log"

	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
)

// FailOnErr will log with a log.Fatal if there is any errors (meaning it will
// directly exit the routine with a exit code 1).
func FailOnErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Intersection - returns the set of elements that exists in all inputs.
func Intersection[T constraints.Ordered](inputs ...[]T) []T {
	if len(inputs) == 1 {
		return inputs[0]
	}

	numberOfElements := len(inputs)
	visited := make(map[T]int)

	for _, input := range inputs {
		input = GetUnique(input)
		for _, v := range input {
			visited[v]++
		}
	}

	result := []T{}
	for k, v := range visited {
		if v == numberOfElements {
			result = append(result, k)
		}
	}

	slices.Sort(result)

	return result
}

// GetUnique - returns unique elements from input.
func GetUnique[T constraints.Ordered](input []T) []T {
	visited := make(map[T]bool)
	result := make([]T, 0)

	for _, r := range input {
		if !visited[r] {
			visited[r] = true
		}
	}

	for k := range visited {
		result = append(result, k)
	}

	return result
}
