package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	// "advent/day1"
	// "advent/day2"
	"advent/day3"
	"advent/day4"
)

func getFileContents(path string) string {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)

	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return strings.Join(lines, "\n")
}

func main() {
	fmt.Println("ADVENT OF COOOOOODE")

	fmt.Println("== Day 1 ==")
	// input := getFileContents("./day1/input.txt")
	// day1.Part1(input)
	// day1.Part2(input)
	// day1.KadatzPart2(input)
	// day1.DavidPart2(input)

	fmt.Println("== Day 2 ==")
	// input := getFileContents("./day2/input.txt")
	// day2.Part1(input)
	// day2.Part2(input)
	//    day2.Kadatz2(input)
	//    day2.DavidPart2(input)

	fmt.Println("== Day 3 ==")
	input := getFileContents("./day3/input.txt")
	day3.Part1(input)
	day3.Part2(input)
	day3.OptPart1(input)
	day3.KadatzPart1(input)
	day3.OptKadatzPart1(input)
	day3.David2(input)

	fmt.Println("== Day 4 ==")
	input = getFileContents("./day4/input.txt")
	day4.Part1(input)
	day4.Part2(input)
	day4.KadatzPart2(input)
}
