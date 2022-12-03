package xstr_test

import (
	"testing"

	"github.com/philiplinell/sh/xstr"
)

func TestIntersection(t *testing.T) {
	str1 := "ABCDEFG"
	str2 := "ABXYZ"
	str3 := "ABC"

	want := "AB"

	got := xstr.Intersection(str1, str2, str3)
	if got != want {
		t.Errorf("Expected %q got %q", want, got)
	}
}
