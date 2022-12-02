package main

import (
	_ "embed" // if you don't use any `embed` functions you need to import it only for its sideeffect with `_`
	"fmt"
	"strings"
)

//go:embed input.txt
var inputData string

const (
	myShapeOrdStart  = 88
	oppShapeOrdStart = 65
)

var OutcomePointsMatrix = [][]int{
	{3, 6, 0},
	{0, 3, 6},
	{6, 0, 3},
}

func calculatePoints(oppShape, myShape string) int {
	oppIndex := int(oppShape[0]) % oppShapeOrdStart
	myIndex := int(myShape[0]) % myShapeOrdStart
	shapePoints := myIndex + 1
	return OutcomePointsMatrix[oppIndex][myIndex] + shapePoints
}

func calculatePoints2(oppShape, outcome string) int {
	desiredOutcomePoints := (int(outcome[0]) % myShapeOrdStart) * 3
	oppIndex := int(oppShape[0]) % oppShapeOrdStart

	var myIndex int
	for i, outcomePoints := range OutcomePointsMatrix[oppIndex] {
		if outcomePoints == desiredOutcomePoints {
			myIndex = i
			break
		}
	}

	shapePoints := myIndex + 1
	return desiredOutcomePoints + shapePoints
}

func main() {
	var points int

	// // Part 1
	// for _, line := range strings.Split(inputData, "\n") {
	// 	parts := strings.Split(line, " ")
	// 	oppShape, myShape := parts[0], parts[1]
	// 	points += calculatePoints(oppShape, myShape)
	// }

	// Part 2
	for _, line := range strings.Split(inputData, "\n") {
		parts := strings.Split(line, " ")
		oppShape, outcome := parts[0], parts[1]
		points += calculatePoints2(oppShape, outcome)
	}
	fmt.Println(points)
}
