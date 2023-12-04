package day2

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// type game struct {
//     id int
//     rounds []round
// }
//
// type round struct {
//     red int
//     green int
//     blue int
// }

const (
	max_red   = 12
	max_green = 13
	max_blue  = 14
)

func Part1(lines []string) {
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

func Part2(lines []string) {
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

type game struct {
	id int
	bC int
	rC int
	gC int
}

func Kadatz2(lines []string) {
	var games []game

	for _, line := range lines {
		name := strings.Split(line, ":")[0]
		id, _ := strconv.Atoi(strings.Split(name, " ")[1])

		game := game{
			id: id,
			bC: 0,
			rC: 0,
			gC: 0,
		}

		hints := strings.Split(strings.Split(line, ":")[1], ";")

		for _, hint := range hints {
			pulls := strings.Split(hint, ",")
			for _, pull := range pulls {
				pull = strings.TrimSpace(pull)
				num, _ := strconv.Atoi(strings.Split(pull, " ")[0])
				if strings.Contains(pull, "red") {
					if game.rC < num {
						game.rC = num
					}
				} else if strings.Contains(pull, "green") {
					if game.gC < num {
						game.gC = num
					}
				} else if strings.Contains(pull, "blue") {
					if game.bC < num {
						game.bC = num
					}
				}
			}
		}

		games = append(games, game)
	}

	sum := 0
	for _, game := range games {
		sum += game.bC * game.rC * game.gC
	}

	// fmt.Println("Sum of valid games:")
	// fmt.Println(sum)
}

func OptKadatz2(lines []string) {
	sum := 0
	for _, line := range lines {
		game := game{
			bC: 0,
			rC: 0,
			gC: 0,
		}

		hints := strings.Split(strings.Split(line, ":")[1], ";")

		for _, hint := range hints {
			pulls := strings.Split(hint, ",")
			for _, pull := range pulls {
				comps := strings.Split(strings.TrimSpace(pull), " ")

				pull = comps[1]
				num, _ := strconv.Atoi(comps[0])

				if pull == "red" && game.rC < num {
					game.rC = num
				} else if pull == "green" && game.gC < num {
					game.gC = num
				} else if pull == "blue" && game.bC < num {
					game.bC = num
				}
			}
		}

		sum += game.bC * game.rC * game.gC
	}

	// fmt.Println("Sum of valid games:", sum)
}
