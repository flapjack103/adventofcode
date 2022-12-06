package main

import (
	_ "embed"
	"fmt"
)

//go:embed input.txt
var inputData string

const (
	startPacketMarkerSize  = 4
	startMessageMarkerSize = 14
)

func main() {
	charMap := make(map[string]int)
	for i, c := range inputData {
		// Add the new char
		newChar := string(c)
		if _, ok := charMap[newChar]; !ok {
			charMap[newChar] = 0
		}
		charMap[newChar]++

		// Pop off the old char
		if i >= startMessageMarkerSize {
			oldChar := string(inputData[i-startMessageMarkerSize])
			charMap[oldChar]--
			if charMap[oldChar] == 0 {
				delete(charMap, oldChar)
			}
		}

		if i >= startMessageMarkerSize-1 && len(charMap) == startMessageMarkerSize {
			// Done: all unique chars
			fmt.Println(i + 1)
			return
		}
	}
}
