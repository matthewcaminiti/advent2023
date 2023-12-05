package day2

import (
	"strconv"
	"strings"
)

type game struct {
	id int
	bC int
	rC int
	gC int
}

func Kadatz2(input string) {
	lines := strings.Split(input, "\n")
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

	// fmt.Println("Sum of valid games:", sum)
}

func OptKadatz2(input string) {
	lines := strings.Split(input, "\n")
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
