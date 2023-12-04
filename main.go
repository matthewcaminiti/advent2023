package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

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
	input := strings.Join(lines, "\n")

	fmt.Println("== Day 1 ==")
	// day1.Part1(lines)
	// day1.Part2(lines)

	fmt.Println("== Day 2 ==")
	// day2.Part1(lines)
	// day2.Part2(lines)

	fmt.Println("== Day 3 ==")
	day3.Part1(input)
	day3.OptPart1(input)
	day3.KadatzPart1(input)
	day3.OptKadatzPart1(input)
	day3.Part2(lines)
}
