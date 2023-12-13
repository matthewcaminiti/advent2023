package main

import (
	"fmt"

	"advent/file"
	// "advent/day1"
	// "advent/day2"
	// "advent/day3"
	// "advent/day4"
	// "advent/day5"
	// "advent/day6"
	// "advent/day7"
	// "advent/day8"
	// "advent/day9"
	// "advent/day10"
	// "advent/day11"
	"advent/day12"
)

func main() {
	fmt.Println("ADVENT OF COOOOOODE")

	fmt.Println("== Day 1 ==")
	// input := file.GetFileContents("./day1/input.txt")
	// day1.Part1(input)
	// day1.Part2(input)

	fmt.Println("== Day 2 ==")
	// input := file.GetFileContents("./day2/input.txt")
	// day2.Part1(input)
	// day2.Part2(input)

	fmt.Println("== Day 3 ==")
	// input := file.GetFileContents("./day3/input.txt")
	// day3.Part1(input)
	// day3.Part2(input)

	fmt.Println("== Day 4 ==")
	// input = file.GetFileContents("./day4/input.txt")
	// day4.Part1(input)
	// day4.Part2(input)

	fmt.Println("== Day 5 ==")
	// input := file.GetFileContents("./day5/input.txt")
	// day5.Part1(input)
	// day5.OptPart2(input)

	fmt.Println("== Day 6 ==")
	// input := file.GetFileContents("./day6/input.txt")
	// day6.Part1(input)
	// day6.Part2(input)

	fmt.Println("== Day 7 ==")
	// input := file.GetFileContents("./day7/input.txt")
	// day7.Part1(input)
	// day7.Part2(input)

	fmt.Println("== Day 8 ==")
	// input := file.GetFileContents("./day8/input.txt")
	// day8.Part1(input)
	// day8.Part2(input)

	fmt.Println("== Day 9 ==")
	// input := file.GetFileContents("./day9/input.txt")
	// day9.Part1(input)
	// day9.Part2(input)

	fmt.Println("== Day 10 ==")
	// test2 := file.GetFileContents("./day10/test2.txt")
	// test3 := file.GetFileContents("./day10/test3.txt")
	// input := file.GetFileContents("./day10/input.txt")
	// // day10.Part1(test)
	// day10.Part1(input)
	// day10.Part2(test2)
	// day10.Part2(test3)
	// day10.Part2(input)
	// 582 (high)
	// 648 (high)

	fmt.Println("== Day 11 ==")
	// input := file.GetFileContents("./day11/input.txt")
	// day11.Part1(input)
	// day11.Part2(input)

	fmt.Println("== Day 12 ==")
	test := file.GetFileContents("./day12/test.txt")
	input := file.GetFileContents("./day12/input.txt")
	day12.Part1(test)
	day12.Part1(input)
	day12.Part2(test)
	day12.Part2(input)
}
