package main

import (
	"bufio"
	"os"

	// "advent/day1"
	"advent/day2"
)

func main() {
	f, err := os.Open("./day2/input.txt")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// day1.Part1(lines)
	// day1.Part2(lines)

	day2.Part1(lines)
	day2.Part2(lines)
    day2.Kadatz2(lines)
}
