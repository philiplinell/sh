package sh_test

import (
	"testing"

	"github.com/philiplinell/sh"
)

func TestParseFileAsInt(t *testing.T) {
	filename := "./testdata/numbers.txt"

	_, err := sh.ParseFileAsInt(filename)
	if err != nil {
		t.Errorf("expected nil error, got: %v", err)
	}
}

func TestParseFileAsStr(t *testing.T) {
	filename := "./testdata/strs.txt"

	_, err := sh.ParseFileAsStr(filename)
	if err != nil {
		t.Errorf("expected nil error, got: %v", err)
	}
}
