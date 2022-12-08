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
	var treeGrid [][]int
	for _, line := range strings.Split(inputData, "\n") {
		var row []int
		for _, c := range line {
			height, err := strconv.Atoi(string(c))
			if err != nil {
				panic(fmt.Errorf("could not parse tree height %s: %v", string(c), err))
			}
			row = append(row, height)
		}
		treeGrid = append(treeGrid, row)
	}

	// count all the outer trees
	// visibleTreeCount := len(treeGrid)*2 + len(treeGrid[0])*2 - 4

	// find the tree with the max scenic score
	var maxScenicScore int
	for i := 1; i < len(treeGrid)-1; i++ {
		row := treeGrid[i]
		for j := 1; j < len(row)-1; j++ {
			scenicScore := visibleDown(i, j, treeGrid)
			scenicScore *= visibleUp(i, j, treeGrid)
			scenicScore *= visibleLeft(i, j, treeGrid)
			scenicScore *= visibleRight(i, j, treeGrid)
			if scenicScore > maxScenicScore {
				maxScenicScore = scenicScore
			}
		}
	}
	// fmt.Println(visibleTreeCount)
	fmt.Println(maxScenicScore)
}

func visibleDown(startRow, col int, treeGrid [][]int) int {
	var treeCount int
	treeSize := treeGrid[startRow][col]
	for i := startRow + 1; i < len(treeGrid); i++ {
		treeCount++
		if treeGrid[i][col] >= treeSize {
			break
		}
	}
	return treeCount
}

func visibleUp(startRow, col int, treeGrid [][]int) int {
	var treeCount int
	treeSize := treeGrid[startRow][col]
	for i := startRow - 1; i >= 0; i-- {
		treeCount++
		if treeGrid[i][col] >= treeSize {
			break
		}
	}
	return treeCount
}

func visibleLeft(row, startCol int, treeGrid [][]int) int {
	var treeCount int
	treeSize := treeGrid[row][startCol]
	for i := startCol - 1; i >= 0; i-- {
		treeCount++
		if treeGrid[row][i] >= treeSize {
			break
		}
	}
	return treeCount
}

func visibleRight(row, startCol int, treeGrid [][]int) int {
	var treeCount int
	treeSize := treeGrid[row][startCol]
	for i := startCol + 1; i < len(treeGrid[0]); i++ {
		treeCount++
		if treeGrid[row][i] >= treeSize {
			break
		}
	}
	return treeCount
}
