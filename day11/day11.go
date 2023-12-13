package day11

import (
	"fmt"
	"math"
	"strings"
)

func manhattanDist(a, b []int) int {
	return int(math.Abs(float64(a[0]-b[0])) + math.Abs(float64(a[1]-b[1])))
}

func Part1(input string) {
	lines := strings.Split(input, "\n")

	colNonEmpty := make([]bool, len(lines[0]))
	rowNonEmpty := make([]bool, len(lines))

	for y, line := range lines {
		for x, char := range line {
			if char != '.' {
				colNonEmpty[x] = true
				rowNonEmpty[y] = true
			}
		}
	}

	newSpace := []string{}
	for y, nonEmpty := range rowNonEmpty {
		row := ""
		for x := range lines[y] {
			if !colNonEmpty[x] {
				// empty
				row += ".."
			} else {
				row += string(lines[y][x])
			}
		}

		if !nonEmpty {
			newSpace = append(newSpace, row, row)
			// is empty
		} else {
			newSpace = append(newSpace, row)
		}
	}

	galaxyCoords := [][]int{}
	for y := range newSpace {
		for x := range newSpace[y] {
			if newSpace[y][x] == '#' {
				galaxyCoords = append(galaxyCoords, []int{x, y})
			}
		}
	}

	sum := 0
	for i, coord := range galaxyCoords {
		shortestPath := 1000000000000
		for j, coord2 := range galaxyCoords {
			if i == j {
				continue
			}

			dist := manhattanDist(coord, coord2)
			sum += dist
			if dist < shortestPath {
				shortestPath = dist
			}
		}
	}

	fmt.Println("Part 1:", sum/2)
}

func Part2(input string) {
	lines := strings.Split(input, "\n")

	colNonEmpty := make([]bool, len(lines[0]))
	rowNonEmpty := make([]bool, len(lines))

	for y, line := range lines {
		for x, char := range line {
			if char != '.' {
				colNonEmpty[x] = true
				rowNonEmpty[y] = true
			}
		}
	}

	factor := 1000000
	galaxyCoords := [][]int{}
	for y := range lines {
		for x := range lines[y] {
			if lines[y][x] == '#' {
				xLeaps, yLeaps := 0, 0
				for i := 0; i < x; i++ {
					if !colNonEmpty[i] {
						xLeaps++
					}
				}
				for i := 0; i < y; i++ {
					if !rowNonEmpty[i] {
						yLeaps++
					}
				}
				galaxyCoords = append(galaxyCoords, []int{
					x - xLeaps + (xLeaps * factor),
					y - yLeaps + (yLeaps * factor),
				})
			}
		}
	}

	sum := 0
	for i, coord := range galaxyCoords {
		shortestPath := 1000000000000
		for j, coord2 := range galaxyCoords {
			if i == j {
				continue
			}

			dist := manhattanDist(coord, coord2)
			sum += dist
			if dist < shortestPath {
				shortestPath = dist
			}
		}
	}

	fmt.Println("Part 2:", sum/2)
}
