package day6

import (
	"fmt"
	"strconv"
	"strings"
)

func Part1(input string) {
	lines := strings.Split(input, "\n")

	times := strings.Fields(lines[0])[1:]
	distances := strings.Fields(lines[1])[1:]

	tot := 0
	for i := 0; i < len(times); i++ {
		numWays := 0

		time, _ := strconv.Atoi(times[i])
		distance, _ := strconv.Atoi(distances[i])

		for j := 0; j <= time; j++ {
			if j*(time-j) > distance {
				numWays++
			}
		}

		if tot == 0 {
			tot = numWays
		} else {
			tot *= numWays
		}
	}

	fmt.Println("Part 1:", tot)
}

func Part2(input string) {
	lines := strings.Split(input, "\n")

	time, _ := strconv.Atoi(strings.Join(strings.Fields(lines[0])[1:], ""))
	distance, _ := strconv.Atoi(strings.Join(strings.Fields(lines[1])[1:], ""))

	numWay := 0
	for i := 0; i <= time; i++ {
		if i*(time-i) > distance {
			numWay++
		}
	}

	fmt.Println("Part 2:", numWay)
}
