package sh

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ParseFileAsStr parses file with given filename and returns an array of
// strings. Will return an error if the file is empty or if reading file
// failed.
func ParseFileAsStr(filename string) ([]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return []string{}, err
	}

	if len(data) == 0 {
		return []string{}, errors.New("empty file")
	}

	return strings.Split(strings.TrimSpace(string((data))), "\n"), nil
}

// ParseFileAsInt parses file with given filename and returns an array of ints.
// Will return an error if the file is empty, if reading file failed or if each
// line could not be parsed to int.
func ParseFileAsInt(filename string) ([]int, error) {
	vals := []int{}
	lines, err := ParseFileAsStr(filename)
	if err != nil {
		return vals, err
	}

	for _, l := range lines {
		val, err := strconv.Atoi(l)
		if err != nil {
			return vals, fmt.Errorf("could not parse %s to int: %v", l, err)
		}
		vals = append(vals, val)
	}

	return vals, nil
}
