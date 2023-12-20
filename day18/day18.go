package day18

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type coord struct {
	x, y int
}

func abs(x int) int {
	if x < 0 {
		return x * -1
	}

	return x
}

func Part1(input string) {
	lines := strings.Split(input, "\n")

	pos := coord{0, 0}
	minCo, maxCo := coord{math.MaxInt64, math.MaxInt64}, coord{0, 0}
	l, r := 0, 0
	prev := ""
	for _, line := range lines {
		comps := strings.Fields(line)

		v, _ := strconv.Atoi(comps[1])

		switch comps[0] {
		case "R":
			pos.x += v
			if prev == "U" {
				r++
			} else if prev == "D" {
				l++
			}
		case "D":
			pos.y += v
			if prev == "R" {
				r++
			} else if prev == "L" {
				l++
			}
		case "L":
			pos.x -= v
			if prev == "U" {
				l++
			} else if prev == "D" {
				r++
			}
		case "U":
			pos.y -= v
			if prev == "R" {
				l++
			} else if prev == "L" {
				r++
			}
		}

		prev = comps[0]

		if pos.x > maxCo.x {
			maxCo.x = pos.x
		}
		if pos.x < minCo.x {
			minCo.x = pos.x
		}

		if pos.y > maxCo.y {
			maxCo.y = pos.y
		}
		if pos.y < minCo.y {
			minCo.y = pos.y
		}
	}

	rDominant := r > l
	// fmt.Println(minCo, maxCo, rDominant)

	grid := [][]int{}
	for y := minCo.y; y <= maxCo.y-minCo.y; y++ {
		grid = append(grid, make([]int, maxCo.x-minCo.x+1))
	}

	pos = coord{}
	seedCoord := coord{}
	for _, line := range lines {
		comps := strings.Fields(line)

		v, _ := strconv.Atoi(comps[1])

		dir := coord{}
		switch comps[0] {
		case "R":
			dir.x = 1
		case "D":
			dir.y = 1
		case "L":
			dir.x = -1
		case "U":
			dir.y = -1
		}

		for i := 0; i < v; i++ {
			pos.x += dir.x
			pos.y += dir.y
			if i == 1 && comps[0] == "R" && seedCoord.x == 0 && seedCoord.y == 0 {
				seedCoord.x = pos.x - minCo.x
				if rDominant {
					seedCoord.y = pos.y - minCo.y + 1
				} else {
					seedCoord.y = pos.y - minCo.y - 1
				}
			}
			grid[pos.y-minCo.y][pos.x-minCo.x] = 1
		}
	}

	h := len(grid)
	w := len(grid[0])

	var fillIn func(at coord)
	fillIn = func(at coord) {
		grid[at.y][at.x] = 1
		for _, dir := range []coord{{1, 0}, {0, -1}, {-1, 0}, {0, 1}} {
			next := coord{at.x + dir.x, at.y + dir.y}
			if next.x < w && next.x >= 0 && next.y < h && next.y >= 0 && grid[next.y][next.x] != 1 {
				fillIn(next)
			}
		}
	}

	fillIn(seedCoord)

	tot := 0
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if grid[y][x] == 1 {
				tot++
			}
		}
	}

	fmt.Println("Part 1:", tot)
}

func det(a, b coord) int {
	return a.x*b.y - a.y*b.x
}

func Part2(input string) {
	lines := strings.Split(input, "\n")

	pos := coord{0, 0}
	vertices := []coord{pos}
	for _, line := range lines {
		comps := strings.Fields(line)

		x := comps[2][2 : len(comps[2])-1]
		_v, _ := strconv.ParseInt(x[:5], 16, 0)
		v := int(_v)

		switch x[5] {
		case '0': // "R"
			pos.x += v
		case '1': // "D"
			pos.y += v
		case '2': // "L"
			pos.x -= v
		case '3': // "U"
			pos.y -= v
		}

		vertices = append(vertices, coord{pos.x, pos.y})
	}

	// vertices = vertices[:len(vertices)-1]
	fmt.Println(vertices[0], vertices[len(vertices)-1])

	tot := 0
	for i := 0; i < len(vertices)-1; i++ {
		// fmt.Println(vertices[i])
		tot += det(vertices[i], vertices[i+1])
		tot += abs(vertices[i].x - vertices[i+1].x)
		tot += abs(vertices[i].y - vertices[i+1].y)
	}
	tot += det(vertices[len(vertices)-1], vertices[0])

	fmt.Println("Part 2:", (tot/2)+1)
}
