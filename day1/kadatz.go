package day1

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func KadatzPart2(input string) {
	lines := strings.Split(input, "\n")

	stringNumbers := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	results := []string{}

	//read file line by line
	for _, line := range lines {
		numA := 0
		numB := 0

		runes := []rune(line)
		for i := 0; i < len(runes); i++ {
			if unicode.IsDigit(runes[i]) {
				if numA == 0 {
					numA = int(runes[i]) - 48
				}
				numB = int(runes[i]) - 48
			} else {
				// look ahead and see if chars ahead match any of the stringNumbers
				for j := 0; j < len(stringNumbers); j++ {
					if i+len(stringNumbers[j]) <= len(runes) {
						if string(runes[i:i+len(stringNumbers[j])]) == stringNumbers[j] {
							// fmt.Println(stringNumbers[j])
							if numA == 0 {
								numA = j + 1
							}
							numB = j + 1
							break
						}
					}
				}
			}
		}

		result := fmt.Sprintf("%d%d", numA, numB)
		results = append(results, result)
	}

	//sum results
	sum := 0
	for i := 0; i < len(results); i++ {
		// fmt.Println("row", i+1)
		// fmt.Println(results[i])
		number, _ := strconv.Atoi(results[i])
		sum += number
	}

	// fmt.Println(sum)
}
