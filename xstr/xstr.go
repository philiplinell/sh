package xstr

import (
	"sort"
	"strings"
)

// Intersection - returns the string with the runes that exists in every
// string.
func Intersection(strs ...string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	numberOfStrings := len(strs)
	var sb strings.Builder

	visited := make(map[rune]int)

	for _, str := range strs {
		str = getUnique(str)
		for _, r := range str {
			visited[r]++
		}
	}

	intersectedElements := []string{}
	for k, v := range visited {
		if v == numberOfStrings {
			intersectedElements = append(intersectedElements, string(k))
		}
	}

	sort.Strings(intersectedElements)

	for _, str := range intersectedElements {
		sb.WriteString(str)
	}

	return sb.String()
}

func getUnique(str string) string {
	var sb strings.Builder

	visited := make(map[rune]bool)

	for _, r := range str {
		if !visited[r] {
			visited[r] = true
		}
	}

	for k := range visited {
		sb.WriteRune(k)
	}

	return sb.String()
}
