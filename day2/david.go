package day2

import (
	// "fmt"
	"strconv"
	"strings"
)

func DavidPart2(input string) {
	lines := strings.Split(input, "\n")

	powerSum := 0

	for _, line := range lines {
		power := parseGame(line)
		powerSum = powerSum + power
	}

	// fmt.Println(powerSum)
}

func parseGame(line string) int {

	gameParts := strings.Split(line, ": ")

	//successes := 0
	//gameId := strings.Split(gameParts[0], " ")
	//gameIdInt, _ := strconv.Atoi(gameId[1])

	redMax := 0
	greenMax := 0
	blueMax := 0

	subGames := strings.Split(gameParts[1], "; ")
	for _, subGame := range subGames {
		red, green, blue := extractColorInts(subGame)
		// fmt.Printf("ANSWER: Red: %d, Green: %d, Blue: %d\n", red, green, blue)

		if red > redMax {
			redMax = red
		}

		if green > greenMax {
			greenMax = green
		}

		if blue > blueMax {
			blueMax = blue
		}

		//if red <= redTest && green <= greenTest && blue <= blueTest {
		//	successes++
		//}
	}

	power := redMax * greenMax * blueMax
	return power

	//if successes == len(subGames) {
	//	return gameIdInt
	//} else {
	//	return 0
	//}
}

func extractColorInts(input string) (int, int, int) {

	colourMap := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}

	colourCounts := strings.Split(input, ", ")

	for _, colourCount := range colourCounts {
		x := strings.Split(colourCount, " ")

		numberStr := x[0]
		colour := x[1]

		number, err := strconv.Atoi(numberStr)
		if err == nil {
			// Check if the colour is in the colourMap
			if _, ok := colourMap[colour]; ok {
				// Assign the number to the colour in colourMap
				colourMap[colour] = number
			}
		}
	}

	return colourMap["red"], colourMap["green"], colourMap["blue"]
}
