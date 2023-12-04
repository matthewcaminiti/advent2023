package main

import (
	"bufio"
	"fmt"
	"os"

	// "advent/day1"
	// "advent/day2"
	"advent/day3"
)

func main() {
	fmt.Println("ADVENT OF COOOOOODE")

	f, err := os.Open("./day3/input.txt")
	if err != nil {
		panic(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)

	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	fmt.Println("== Day 1 ==")
	// day1.Part1(lines)
	// day1.Part2(lines)

	fmt.Println("== Day 2 ==")
	// day2.Part1(lines)
	// day2.Part2(lines)

	fmt.Println("== Day 3 ==")
	day3.Part1(lines)
	day3.Part2(lines)
}
