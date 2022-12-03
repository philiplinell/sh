package parse_test

import (
	"testing"

	"github.com/philiplinell/sh/parse"
)

func TestParseFileAsInt(t *testing.T) {
	filename := "./testdata/numbers.txt"

	_, err := parse.FileAsInt(filename)
	if err != nil {
		t.Errorf("expected nil error, got: %v", err)
	}
}

func TestParseFileAsStr(t *testing.T) {
	filename := "./testdata/strs.txt"

	_, err := parse.FileAsStr(filename)
	if err != nil {
		t.Errorf("expected nil error, got: %v", err)
	}
}
