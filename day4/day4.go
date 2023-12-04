package day4

import (
	"fmt"
	"strconv"
	"strings"
)

func Part1(input string) {
	lines := strings.Split(input, "\n")

	sum := 0
	for _, line := range lines {
		card := strings.Split(strings.Split(line, ":")[1], "|")

		_winningNums := make([]bool, 100)
		winningNums := strings.Fields(card[0])

		for _, wn := range winningNums {
			n, _ := strconv.Atoi(wn)
			_winningNums[n] = true
		}

		win := 0
		havingNums := strings.Fields(card[1])
		for _, hn := range havingNums {
			n, _ := strconv.Atoi(hn)
			if _winningNums[n] {
				if win == 0 {
					win = 1
				} else {
					win *= 2
				}
			}
		}

		sum += win
	}

	fmt.Println("Part 1:", sum)
}

func Part2(input string) {
	lines := strings.Split(input, "\n")

	numlines := len(lines)

	copies := make([]int, numlines)
	for i := range copies {
		copies[i] = 1
	}

	total := 0
	for i, line := range lines {
		card := strings.Split(line[9:], " | ")

		_winningNums := make([]bool, 100)
		winningNums := strings.Fields(card[0])

		for _, wn := range winningNums {
			n, _ := strconv.Atoi(wn)
			_winningNums[n] = true
		}

		j := 1
		havingNums := strings.Fields(card[1])
		for _, hn := range havingNums {
			n, _ := strconv.Atoi(hn)
			if _winningNums[n] {
				if i+j <= numlines-1 {
					copies[i+j] += copies[i]
				}
				j++
			}
		}

		if i > 0 {
			total += copies[i-1]
		}
	}

	total += copies[len(copies)-1]

	fmt.Println("Part 2:", total)
}
