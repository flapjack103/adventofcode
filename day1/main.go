package main

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input, err := getInput("/Users/alexandra.bueno/go/src/github.com/flapjack103/adventofcode/day1/input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(input, "\n")

	var (
		calorieCounts []int
		currCalories  int
	)
	for _, line := range lines {
		if line == "" {
			calorieCounts = append(calorieCounts, currCalories)
			currCalories = 0
			continue
		}

		nCalories, err := strconv.Atoi(line)
		if err != nil {
			fmt.Printf("could not convert line '%s' to int: %s\n", line, err)
			continue
		}
		currCalories += nCalories
	}

	if currCalories > 0 {
		calorieCounts = append(calorieCounts, currCalories)
	}

	sort.Ints(calorieCounts)

	var sum int
	for _, cc := range calorieCounts[len(calorieCounts)-3:] {
		sum += cc
	}
	fmt.Println(sum)
}

func getInput(path string) (string, error) {
	fd, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer fd.Close()

	b, err := io.ReadAll(fd)
	if err != nil {
		return "", err
	}

	return string(b), err
}
