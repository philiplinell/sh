package sh_test

import (
	"fmt"

	"github.com/philiplinell/sh"
)

func ExampleIntersection() {
	str1 := []string{"a", "b", "c", "d"}
	str2 := []string{"b", "c", "d", "e"}
	str3 := []string{"c", "d", "e", "f"}

	fmt.Println(sh.Intersection(str1, str2, str3))
	// Output: [c d]
}

func ExampleGetUnique() {
	str1 := []string{"a", "b", "b", "a"}

	fmt.Println(sh.GetUnique(str1))
	// Output: [a b]
}
