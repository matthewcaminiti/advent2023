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

type ccoord struct {
	x, y, count, step int
	from              coord
}

func nextPossibleCCoords(grid []string, at ccoord) []ccoord {
	w := len(grid[0])
	h := len(grid)

	nexts := []ccoord{}
	for _, dir := range []coord{{1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
		next := ccoord{at.x + dir.x, at.y + dir.y, at.count, at.step + 1, coord{at.x, at.y}}
		if next.x < 0 {
			next.x = w - 1
			next.step++
		} else if next.x >= w {
			next.x = 0
			next.step++
		}

		if next.y < 0 {
			next.y = h - 1
			next.step++
		} else if next.y >= h {
			next.y = 0
			next.step++
		}

		if grid[next.y][next.x] == '#' {
			continue
		}

		nexts = append(nexts, next)
	}

	return nexts
}

func Part2(input string) {
	// do something with odd/even logic, repetitive behaviour
	// maybe get bounding box of latest search space

	// if search coord is on edge, wrap it around
	// when adding coord to search, check if the coord is already flagged for even && odd
	// if it's flagged for both, ignore
	// if it's flagged for one, mark according to the current step parity, keep in search
	lines := strings.Split(input, "\n")

	start := ccoord{}
	for y, line := range lines {
		for x, c := range line {
			if c == 'S' {
				start.x = x
				start.y = y
				start.count = 1
				start.step = 0
				break
			}
		}
	}

	searchCoords := []ccoord{start}

	// repeat := 0
	// prevSearch := []ccoord{}
	// repeatStep := 0

	// nSteps := 26501365
	nSteps := 50
	for i := 0; i < nSteps; i++ {
		nextCoords := []ccoord{}

		for _, c := range searchCoords {
			possible := nextPossibleCCoords(lines, c)

			for _, p := range possible {
				found := false
				for i, nc := range nextCoords {
					if p.x == nc.x && p.y == nc.y {
						if p.step != nc.step {
							nextCoords[i].count += p.count
						}
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

		// if len(searchCoords) == len(prevSearch) {
		//     if repeat == 2 {
		//         repeatStep = i
		//         break
		//     }
		//
		//     repeat++
		// } else {
		//     repeat = 0
		// }

		// prevSearch = searchCoords
		searchCoords = nextCoords
	}

	// fmt.Println("repeated at:", repeatStep)
	// fmt.Println(prevSearch)
	// prevTot := 0
	// for _, c := range prevSearch {
	//     prevTot += c.count
	// }
	// fmt.Println(searchCoords)
	// nextTot := 0
	// for _, c := range searchCoords {
	//     nextTot += c.count
	// }
	//
	// fmt.Println("diff:", nextTot, prevTot, nextTot - prevTot)

	tot := 0
	for _, c := range searchCoords {
		tot += c.count
	}

	fmt.Println("Part 2:", tot, len(searchCoords))
}
