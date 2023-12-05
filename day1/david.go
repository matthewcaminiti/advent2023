package day1

import (
	// "fmt"
	"strconv"
	"strings"
)

func DavidPart2(input string) {
	lines := strings.Split(input, "\n")

	runningSum := 0
	for _, line := range lines {
		sumString := extractLineConcat(line)
		sumInt, _ := strconv.Atoi(sumString)
		runningSum += sumInt
		//fmt.Println("running sum:", runningSum)
	}
	// fmt.Println("final sum: ", runningSum)
}

func extractLineConcat(line string) string {

	// fmt.Println("Line: ", line)

	intA := ""
	intB := ""

	for i := 0; i < len(line); i++ {
		match := getMatchFromString(line[i:])

		if match != "" {
			if intA == "" {
				intA = match
			} else {
				intB = match
			}
			match = ""
		}
	}

	if intB == "" {
		intB = intA
	}
	// fmt.Println("intA: ", intA)
	// fmt.Println("intB: ", intB)
	sum := intA + intB
	// fmt.Println("concat sum: ", sum)
	// fmt.Println()
	return sum
}

func getMatchFromString(line string) string {
	// fmt.Println("TEST LINE: ", line)
	numbers := []string{
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
	}

	word := ""
	for _, ch := range line {
		if digit, err := strconv.Atoi(string(ch)); err == nil {
			return strconv.Itoa(digit)
		} else {
			word += string(ch)
			match := containsSubstring(word, numbers)
			if match != "" {
				return match
			}
		}
	}
	return ""
}

func containsSubstring(s string, substrings []string) string {
	for i, substr := range substrings {
		if strings.Contains(s, substr) {
			return strconv.Itoa(i + 1)
		}
	}
	return ""
}
