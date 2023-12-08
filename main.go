package main

import (
	"fmt"

	"advent/file"
	// "advent/day1"
	// "advent/day2"
	// "advent/day3"
	// "advent/day4"
	"advent/day5"
	// "advent/day6"
	// "advent/day7"
)

func main() {
	fmt.Println("ADVENT OF COOOOOODE")

	fmt.Println("== Day 1 ==")
	// input := file.GetFileContents("./day1/input.txt")
	// day1.Part1(input)
	// day1.Part2(input)
	// day1.KadatzPart2(input)
	// day1.DavidPart2(input)

	fmt.Println("== Day 2 ==")
	// input := file.GetFileContents("./day2/input.txt")
	// day2.Part1(input)
	// day2.Part2(input)
	// day2.Kadatz2(input)
	// day2.DavidPart2(input)

	fmt.Println("== Day 3 ==")
	// input := file.GetFileContents("./day3/input.txt")
	// day3.Part1(input)
	// day3.Part2(input)
	// day3.OptPart1(input)
	// day3.KadatzPart1(input)
	// day3.OptKadatzPart1(input)
	// day3.David2(input)

	fmt.Println("== Day 4 ==")
	// input = file.GetFileContents("./day4/input.txt")
	// day4.Part1(input)
	// day4.Part2(input)
	// day4.KadatzPart2(input)

	fmt.Println("== Day 5 ==")
	test := file.GetFileContents("./day5/test.txt")
	input := file.GetFileContents("./day5/input.txt")
	day5.Part1(input)
	// day5.Part2(input)
	day5.OptPart2(test)
	day5.OptPart2(input)

	fmt.Println("== Day 6 ==")
	// input := file.GetFileContents("./day6/input.txt")
	// day6.Part1(input)
	// day6.Part2(input)

	fmt.Println("== Day 7 ==")
	// test := file.GetFileContents("./day7/test.txt")
	// input := file.GetFileContents("./day7/input.txt")
	//    day7.Part1(test)
	//    day7.Part1(input)
	//    day7.Part2(test)
	//    day7.Part2(input)
}
