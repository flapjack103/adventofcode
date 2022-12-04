package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var inputData string

func main() {
	var overlaps int
	lines := strings.Split(inputData, "\n")
	for _, line := range lines {
		parts := strings.Split(line, ",")
		a1, err := newAssignment(parts[0])
		if err != nil {
			panic(err)
		}
		a2, err := newAssignment(parts[1])
		if err != nil {
			panic(err)
		}
		if hasOverlap(a1, a2) {
			overlaps++
		}
	}
	fmt.Println(overlaps)
}

type assignment struct {
	start, end int
}

func newAssignment(a string) (*assignment, error) {
	parts := strings.Split(a, "-")
	start, err := strconv.Atoi(parts[0])
	if err != nil {
		return nil, err
	}
	end, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, err
	}
	return &assignment{start: start, end: end}, nil
}

func (a1 *assignment) contains(a2 *assignment) bool {
	return a1.start <= a2.start && a1.end >= a2.end
}

func hasOverlap(a1, a2 *assignment) bool {
	// contains
	if a1.contains(a2) || a2.contains(a1) {
		return true
	}
	// left side
	// |------ a1 ----- |
	//        |------ a2 ----- |
	if a1.end >= a2.start && a1.end <= a2.end {
		return true
	}
	// right side
	//        |------ a1 ----- |
	// |------ a2 ----- |
	if a1.start >= a2.start && a1.start <= a2.end {
		return true
	}
	return false
}
