package day3

import (
	// "fmt"
	"strconv"
	"strings"
)

type Coordinate struct {
	X int
	Y int
}

type GearRatio struct {
	Ratio1     string
	Ratio2     string
	Coordinate Coordinate
}

func David2(input string) {
	lines := strings.Split(input, "\n")

	var matrix [][]string

	for _, line := range lines {
		row := make([]string, len(line))

		for i, char := range line {
			row[i] = string(char)
		}

		matrix = append(matrix, row)
	}

	//runningSum := 0

	possiblePartNumber := ""
	gearRatiosFirstMatch := make([]GearRatio, 0)
	gearRatiosSecondMatch := make([]GearRatio, 0)
	hasStarAdjacent := false
	exists := false
	coord := Coordinate{}
	gearRatioSave := GearRatio{}

	for i, row := range matrix {
		for j, item := range row {
			if _, err := strconv.Atoi(item); err == nil {
				possiblePartNumber += item

				// Check single adjacent characters for symbols (excluding periods)
				for x := i - 1; x <= i+1; x++ {
					for y := j - 1; y <= j+1; y++ {
						if x >= 0 && x < len(matrix) && y >= 0 && y < len(row) && (x != i || y != j) {
							if matrix[x][y] == "*" {
								hasStarAdjacent = true

								// Remove from gearRatiosSecondMatch if there is a third match
								indexToRemove := -1
								for i, gr := range gearRatiosSecondMatch {
									if gr.Coordinate.X == x && gr.Coordinate.Y == y {
										indexToRemove = i
										break
									}
								}

								// If the gear ratio is found, remove it
								if indexToRemove != -1 {
									gearRatiosSecondMatch = append(gearRatiosSecondMatch[:indexToRemove], gearRatiosSecondMatch[indexToRemove+1:]...)
								}

								// Check if the coordinate already exists in the gearRatiosFirstMatch slice
								for _, gr := range gearRatiosFirstMatch {
									if gr.Coordinate.X == x && gr.Coordinate.Y == y {
										gearRatioSave = gr
										exists = true
										coord = gr.Coordinate
										break
									}
								}

								// If the coordinate doesn't exist, add a new GearRatio struct
								if !exists {
									coord = Coordinate{X: x, Y: y}
									gearRatioSave = GearRatio{
										Coordinate: coord,
									}
								}

							}
						}
					}
					if hasStarAdjacent {
						break
					}
				}
			} else {
				// not an int
				if possiblePartNumber != "" && hasStarAdjacent && exists {
					for _, gr := range gearRatiosFirstMatch {
						if gr.Coordinate.X == coord.X && gr.Coordinate.Y == coord.Y {
							gr.Ratio2 = possiblePartNumber
							gearRatiosSecondMatch = append(gearRatiosSecondMatch, gr)
							// fmt.Printf("Ratio1: %s, Ratio2: %s, Coordinate: (%d, %d)\n", gr.Ratio1, possiblePartNumber, gr.Coordinate.X, gr.Coordinate.Y)
							break
						}
					}
				} else if possiblePartNumber != "" && hasStarAdjacent {
					gearRatioSave.Ratio1 = possiblePartNumber
					gearRatiosFirstMatch = append(gearRatiosFirstMatch, gearRatioSave)
					//runningSum += partNumber
				}
				possiblePartNumber = ""
				hasStarAdjacent = false
				coord = Coordinate{}
				gearRatioSave = GearRatio{}
				exists = false
			}
		}
	}

	//fmt.Println("Summed Part Number:", runningSum)

	//fmt.Println("Gear Ratios:")
	//for _, gr := range gearRatiosFirstMatch {
	//	fmt.Printf("Ratio1: %s, Ratio2: %s, Coordinate: (%d, %d)\n", gr.Ratio1, gr.Ratio2, gr.Coordinate.X, gr.Coordinate.Y)
	//}

	runningSum := 0

	for _, gr := range gearRatiosSecondMatch {
		ratio1, err1 := strconv.Atoi(gr.Ratio1)
		ratio2, err2 := strconv.Atoi(gr.Ratio2)

		if err1 == nil && err2 == nil {
			product := ratio1 * ratio2
			runningSum += product
		}
	}

	// fmt.Println("Running Sum:", runningSum)
}
