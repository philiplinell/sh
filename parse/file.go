package parse

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// FileAsStr parses file with given filename and returns an array of
// strings. Will return an error if the file is empty or if reading file
// failed.
func FileAsStr(filename string) ([]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return []string{}, err
	}

	if len(data) == 0 {
		return []string{}, errors.New("empty file")
	}

	return strings.Split(string(data), "\n"), nil
}

// FileAsInt parses file with given filename and returns an array of ints.
// Will return an error if the file is empty, if reading file failed or if each
// line could not be parsed to int.
func FileAsInt(filename string) ([]int, error) {
	vals := []int{}
	lines, err := FileAsStr(filename)
	if err != nil {
		return vals, err
	}

	for _, l := range lines {
		if len(l) == 0 {
			continue
		}
		val, err := strconv.Atoi(l)
		if err != nil {
			return vals, fmt.Errorf("could not parse %s to int: %v", l, err)
		}
		vals = append(vals, val)
	}

	return vals, nil
}
