package day4

import (
	"strconv"
	"strings"
)

type game struct {
	wNums   []int
	rNums   []int
	matches int
	count   int
}

func KadatzPart2(input string) {
	var games []game

	lines := strings.Split(input, "\n")

	for _, line := range lines {

		nums := strings.Split(line, ":")[1]
		wNums := strings.Split(strings.Trim(strings.Split(nums, "|")[0], " "), " ")
		rNums := strings.Split(strings.Trim(strings.Split(nums, "|")[1], " "), " ")

		var wNumsInt []int
		var rNumsInt []int

		for _, num := range wNums {
			if num == "" {
				continue
			}
			numI, _ := strconv.Atoi(num)
			wNumsInt = append(wNumsInt, numI)
		}

		for _, num := range rNums {
			numI, _ := strconv.Atoi(num)
			rNumsInt = append(rNumsInt, numI)
		}

		games = append(games, game{wNums: wNumsInt, rNums: rNumsInt, matches: 0, count: 1})
	}

	for i, game := range games {
		wNumsMap := make(map[int]bool)
		for _, num := range game.wNums {
			wNumsMap[num] = true
		}
		for _, num := range game.rNums {
			if wNumsMap[num] {
				games[i].matches++
			}
		}
	}

	for i, game := range games {
		for j := 0; j < game.matches; j++ {
			games[j+i+1].count += games[i].count
		}
	}

	sum := 0
	for _, game := range games {
		sum += game.count
	}

	// fmt.Println(sum)
}
