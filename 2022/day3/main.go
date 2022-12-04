package main

import (
	_ "embed" // if you don't use any `embed` functions you need to import it only for its sideeffect with `_`
	"fmt"
	"strings"
)

//go:embed input.txt
var inputData string

func main() {
	var prioritySum int
	lines := strings.Split(inputData, "\n")
	for i := range lines {
		// comp1, comp2 := line[:len(line)/2], line[len(line)/2:]
		// for item := range getSharedItems(comp1, comp2) {
		// 	prioritySum += getPriority(item)
		// }

		if i > 0 && i%3 == 0 {
			group := lines[i-3 : i]
			shared := getSharedItems(group[0], group[1])
			for _, c := range group[2] {
				if _, ok := shared[c]; ok {
					prioritySum += getPriority(c)
					break
				}
			}
		}
	}

	group := lines[len(lines)-3:]
	shared := getSharedItems(group[0], group[1])
	for _, c := range group[2] {
		if _, ok := shared[c]; ok {
			prioritySum += getPriority(c)
			break
		}
	}

	fmt.Println(prioritySum)
}

// messy set operations in go T_T
func getSharedItems(comp1, comp2 string) map[rune]struct{} {
	shared := make(map[rune]struct{})
	comp1Map := make(map[rune]struct{})
	for _, c := range comp1 {
		comp1Map[c] = struct{}{}
	}
	for _, c := range comp2 {
		if _, ok := comp1Map[c]; ok {
			shared[c] = struct{}{}
		}
	}
	return shared
}

func getPriority(r rune) int {
	if r >= 'A' && r <= 'Z' {
		return int(r-'A') + 27
	}
	return int(r-'a') + 1
}
