package day21

import (
	"fmt"
	"strings"
)

type coord struct {
	x, y int
}

func nextPossibleCoords(grid []string, at coord) []coord {
	w := len(grid[0])
	h := len(grid)

	nexts := []coord{}
	for _, dir := range []coord{{1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
		next := coord{at.x + dir.x, at.y + dir.y}
		if next.x < 0 || next.x >= w || next.y < 0 || next.y >= h || grid[next.y][next.x] == '#' {
			continue
		}

		nexts = append(nexts, next)
	}

	return nexts
}

func Part1(input string) {
	lines := strings.Split(input, "\n")

	start := coord{}
	for y, line := range lines {
		for x, c := range line {
			if c == 'S' {
				start.x = x
				start.y = y
				break
			}
		}

		if start.x != 0 || start.y != 0 {
			break
		}
	}

	searchCoords := []coord{start}
	nSteps := 64
	for i := 0; i < nSteps; i++ {
		nextCoords := []coord{}

		for _, c := range searchCoords {
			possible := nextPossibleCoords(lines, c)

			for _, p := range possible {
				found := false
				for _, nc := range nextCoords {
					if p.x == nc.x && p.y == nc.y {
						found = true
						break
					}
				}

				if found {
					continue
				}

				nextCoords = append(nextCoords, p)
			}
		}

		searchCoords = nextCoords
	}

	fmt.Println("Part 1:", len(searchCoords))
}

func Part2(input string) {
	// do something with odd/even logic, repetitive behaviour
	// maybe get bounding box of latest search space
}
