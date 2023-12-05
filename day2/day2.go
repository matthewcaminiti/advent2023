package day2

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

const (
	max_red   = 12
	max_green = 13
	max_blue  = 14
)

func Part1(input string) {
	lines := strings.Split(input, "\n")

	sum := 0
	for _, line := range lines {
		comps := strings.Split(line, ":")

		gameIdStrComps := strings.Split(comps[0], " ")
		gameId, _ := strconv.Atoi(gameIdStrComps[len(gameIdStrComps)-1])

		invalidGame := false
		roundStrs := strings.Split(comps[1], ";")
		for _, s := range roundStrs {
			draws := strings.Split(s, ",")

			for _, draw := range draws {
				comps := strings.Split(strings.TrimSpace(draw), " ")

				num, _ := strconv.Atoi(comps[0])
				switch comps[1] {
				case "red":
					invalidGame = num > max_red
				case "green":
					invalidGame = num > max_green
				case "blue":
					invalidGame = num > max_blue
				}

				if invalidGame {
					break
				}
			}

			if invalidGame {
				break
			}
		}

		if !invalidGame {
			sum += gameId
		}

		// fmt.Printf("%s: [%t]\n", line, invalidGame)
	}

	fmt.Println("Part1:", sum)
}

func Part2(input string) {
	lines := strings.Split(input, "\n")

	sum := 0
	for _, line := range lines {
		comps := strings.Split(line, ":")

		roundStrs := strings.Split(comps[1], ";")
		max_r, max_g, max_b := 0, 0, 0
		for _, s := range roundStrs {
			draws := strings.Split(s, ",")

			for _, draw := range draws {
				comps := strings.Split(strings.TrimSpace(draw), " ")

				num, _ := strconv.Atoi(comps[0])
				switch comps[1] {
				case "red":
					max_r = int(math.Max(float64(max_r), float64(num)))
				case "green":
					max_g = int(math.Max(float64(max_g), float64(num)))
				case "blue":
					max_b = int(math.Max(float64(max_b), float64(num)))
				}
			}
		}

		sum += max_r * max_g * max_b
		max_r, max_g, max_b = 0, 0, 0

		// fmt.Printf("%s: [%d]\n", line, max_r * max_g * max_b)
	}

	// fmt.Println("Part 2:", sum)
}
