package sh_test

import (
	"strings"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/philiplinell/sh"
)

func TestIntersectionStr(t *testing.T) {
	str1 := "ABCDEFG"
	str2 := "ABXYZ"
	str3 := "ABC"

	want := []string{"A", "B"}

	got := sh.Intersection(
		strings.Split(str1, ""),
		strings.Split(str2, ""),
		strings.Split(str3, ""),
	)
	assert.Equal(t, want, got)
}

func TestIntersectionInt(t *testing.T) {
	num1 := []int{1, 2, 3, 4}
	num2 := []int{5, 4, 3, 2}
	num3 := []int{1, 2, 3}

	want := []int{2, 3}

	got := sh.Intersection(num1, num2, num3)
	assert.Equal(t, want, got)
}

func TestIntersectionSingleElement(t *testing.T) {
	want := []int{1, 2, 3}

	got := sh.Intersection(want)
	assert.Equal(t, want, got)
}

func TestGetUnique(t *testing.T) {
	input := strings.Split("ABBA", "")

	want := []string{"A", "B"}

	got := sh.GetUnique(input)
	assert.Equal(t, want, got)
}
