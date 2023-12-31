package day23

import (
	"fmt"
	"strings"
)

type coord struct {
	x, y int
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func maxPath(grid []string, from, at coord) int {
	w := len(grid[0])
	h := len(grid)

	// fmt.Printf("(%d, %d)\n", at.x, at.y)

	nextSteps := []coord{}
	for _, dir := range []coord{{1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
		next := coord{at.x + dir.x, at.y + dir.y}
		if next.x == from.x && next.y == from.y {
			continue
		}

		if next.x < 0 || next.x >= w || next.y < 0 || next.y >= h {
			continue
		}

		c := grid[next.y][next.x]
		if c == '#' {
			continue
		}

		if c != '.' {
			switch c {
			case 'v':
				// if going up, continue
				if dir.y == -1 {
					continue
				}
			case '<':
				// if going right, continue
				if dir.x == 1 {
					continue
				}
			case '>':
				// if going left, continue
				if dir.x == -1 {
					continue
				}
			case '^':
				// if going down, continue
				if dir.y == 1 {
					continue
				}
			}
		}

		nextSteps = append(nextSteps, next)
	}

	maxDist := 0
	for _, next := range nextSteps {
		maxDist = max(maxDist, 1+maxPath(grid, at, next))
	}

	return maxDist
}

func Part1(input string) {
	lines := strings.Split(input, "\n")

	start := coord{}
	for x, c := range lines[0] {
		if c == '.' {
			start = coord{x: x, y: 0}
		}
	}

	fmt.Println("Part 1:", maxPath(lines, coord{x: start.x, y: start.y - 1}, start))
}

func maxPath2(grid []string, end, from, at coord, visited []coord) int {
	w := len(grid[0])
	h := len(grid)

	nextSteps := []coord{}
	for _, dir := range []coord{{1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
		next := coord{at.x + dir.x, at.y + dir.y}
		if next.x < 0 || next.x >= w || next.y < 0 || next.y >= h {
			continue
		}

		c := grid[next.y][next.x]
		if c == '#' {
			continue
		}

		seen := false
		for _, v := range visited {
			if next.x == v.x && next.y == v.y {
				seen = true
				break
			}
		}

		if seen {
			continue
		}

		nextSteps = append(nextSteps, next)
	}

	if len(nextSteps) == 0 && !(at.x == end.x && at.y == end.y) {
		return -100
	}

	maxDist := 0
	for _, next := range nextSteps {
		nextVisited := []coord{next}
		for _, v := range visited {
			nextVisited = append(nextVisited, v)
		}

		maxDist = max(maxDist, 1+maxPath2(grid, end, at, next, nextVisited))
	}

	return maxDist
}

func Part2(input string) {
	lines := strings.Split(input, "\n")

	start := coord{}
	for x, c := range lines[0] {
		if c == '.' {
			start = coord{x: x, y: 0}
		}
	}

	end := coord{}
	for x, c := range lines[len(lines)-1] {
		if c == '.' {
			end = coord{x: x, y: len(lines) - 1}
		}
	}

	visited := []coord{start}
	fmt.Println("Part 2:", maxPath2(lines, end, coord{x: start.x, y: start.y - 1}, start, visited))
}
