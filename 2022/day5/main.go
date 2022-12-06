package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	//go:embed input.txt
	inputData        string
	instructionRegex = regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)
)

func parseNumberOfStacks(input string) int {
	input = strings.ReplaceAll(input, " ", "")
	n, _ := strconv.Atoi(string(input[len(input)-1]))
	return n
}

func parseCrates(level string, rowSize int) []string {
	row := make([]string, rowSize, rowSize)
	for i := 0; i < len(level); i++ {
		if string(level[i]) == "[" {
			crateID := string(level[i+1])
			row[i/4] = crateID
			i += 2
		}
	}
	return row
}

func buildStacks(input string) []stack {
	levels := strings.Split(input, "\n")
	numStacks := parseNumberOfStacks(levels[len(levels)-1])

	var rows [][]string
	for _, level := range levels {
		rows = append(rows, parseCrates(level, numStacks))
	}

	// transform rows into stacks
	stacks := make([]stack, numStacks, numStacks)
	for i := len(rows) - 1; i >= 0; i-- {
		for j, crate := range rows[i] {
			if crate == "" {
				// wtf is with this parsing?
				continue
			}
			stacks[j].push(crate)
		}
	}
	return stacks
}

type instruction struct {
	numToMove int
	fromIdx   int
	toIdx     int
}

// ex. move 1 from 2 to 1
func parseInstruction(line string) (*instruction, error) {
	groups := instructionRegex.FindStringSubmatch(line)

	num, err := strconv.Atoi(groups[1])
	if err != nil {
		return nil, err
	}
	from, err := strconv.Atoi(groups[2])
	if err != nil {
		return nil, err
	}
	to, err := strconv.Atoi(groups[3])
	if err != nil {
		return nil, err
	}
	return &instruction{numToMove: num, fromIdx: from - 1, toIdx: to - 1}, nil
}

type stack struct {
	len      int
	contents []string
}

func (s *stack) push(obj string) {
	s.contents = append(s.contents, obj)
	s.len++
}

func (s *stack) pop() string {
	if s.len == 0 {
		return ""
	}
	obj := s.contents[s.len-1]
	s.contents = s.contents[:s.len-1]
	s.len--
	return obj
}

func moveCrates(stacks []stack, instructions string) {
	for _, line := range strings.Split(instructions, "\n") {
		if line == "" {
			continue
		}
		ins, err := parseInstruction(line)
		if err != nil {
			panic(fmt.Errorf("could not parse instruction %s: %v", line, err))
		}
		for i := 0; i < ins.numToMove; i++ {
			crate := stacks[ins.fromIdx].pop()
			if crate != "" {
				stacks[ins.toIdx].push(crate)
			}
		}
		// fmt.Printf("instruction:%s\n, stacks:%+v\n", line, stacks)
	}
}

func moveCrates9001(stacks []stack, instructions string) {
	for _, line := range strings.Split(instructions, "\n") {
		if line == "" {
			continue
		}
		ins, err := parseInstruction(line)
		if err != nil {
			panic(fmt.Errorf("could not parse instruction %s: %v", line, err))
		}
		// getting a little messy :/
		// pop off all the crates then push them on in reverse of pop order
		var crates []string
		for i := 0; i < ins.numToMove; i++ {
			crate := stacks[ins.fromIdx].pop()
			if crate == "" {
				break
			}
			crates = append(crates, crate)
		}
		for i := len(crates) - 1; i >= 0; i-- {
			stacks[ins.toIdx].push(crates[i])
		}
	}
}

func main() {
	div := strings.Index(inputData, "\n\n")
	stacks := buildStacks(inputData[:div])
	moveCrates9001(stacks, inputData[div:])

	var topCrates []string
	for _, stack := range stacks {
		topCrates = append(topCrates, stack.pop())
	}
	fmt.Println(strings.Join(topCrates, ""))
}
