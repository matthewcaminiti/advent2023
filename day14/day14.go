package day14

import (
	"fmt"
	"strings"
)

func Part1(input string) {
	lines := strings.Split(input, "\n")

	cols := make([]string, len(lines[0]))
	for i := 0; i < len(lines[0]); i++ {
		for j := range lines {
			cols[i] += string(lines[j][i])
		}
	}

	load := 0
	for _, col := range cols {
		prevObstacle := -1
		for i := 0; i < len(col); i++ {
			if col[i] == '#' {
				prevObstacle = i
				continue
			}

			if col[i] == 'O' {
				load += len(col) - prevObstacle - 1
				prevObstacle = prevObstacle + 1
			}
		}
	}

	fmt.Println("Part 1:", load)
}

func compute4Cycle(_grid [][]byte) [][]byte {
	grid := [][]byte{}
	for y := 0; y < len(_grid); y++ {
		tmp := make([]byte, len(_grid[0]))
		for x := 0; x < len(_grid[0]); x++ {
			tmp[x] = _grid[y][x]
		}
		grid = append(grid, tmp)
	}

	// north, west, south, east
	// north
	for x := 0; x < len(grid[0]); x++ {
		prevObstacle := -1
		for y := 0; y < len(grid); y++ {
			if grid[y][x] == '#' {
				prevObstacle = y
				continue
			}

			if grid[y][x] == 'O' {
				grid[y][x] = '.'
				prevObstacle++
				grid[prevObstacle][x] = 'O'
			}
		}
	}

	// west
	for y := 0; y < len(grid); y++ {
		prevObstacle := -1
		for x := 0; x < len(grid[0]); x++ {
			if grid[y][x] == '#' {
				prevObstacle = x
				continue
			}

			if grid[y][x] == 'O' {
				grid[y][x] = '.'
				prevObstacle++
				grid[y][prevObstacle] = 'O'
			}
		}
	}

	// south
	for x := 0; x < len(grid[0]); x++ {
		prevObstacle := len(grid)
		for y := len(grid) - 1; y >= 0; y-- {
			if grid[y][x] == '#' {
				prevObstacle = y
				continue
			}

			if grid[y][x] == 'O' {
				grid[y][x] = '.'
				prevObstacle--
				grid[prevObstacle][x] = 'O'
			}
		}
	}

	// east
	for y := 0; y < len(grid); y++ {
		prevObstacle := len(grid[0])
		for x := len(grid[0]) - 1; x >= 0; x-- {
			if grid[y][x] == '#' {
				prevObstacle = x
				continue
			}

			if grid[y][x] == 'O' {
				grid[y][x] = '.'
				prevObstacle--
				grid[y][prevObstacle] = 'O'
			}
		}
	}

	return grid
}

func equal(g1 [][]byte, g2 [][]byte) bool {
	for y := 0; y < len(g1); y++ {
		for x := 0; x < len(g1[0]); x++ {
			if g1[y][x] != g2[y][x] {
				return false
			}
		}
	}

	return true
}

func calcWeight(grid [][]byte) int {
	tot := 0
	for x := 0; x < len(grid[0]); x++ {
		for y := 0; y < len(grid); y++ {
			if grid[y][x] == 'O' {
				tot += len(grid) - y
			}
		}
	}

	return tot
}

func Part2(input string) {
	lines := strings.Split(input, "\n")

	grid := [][]byte{}

	for _, line := range lines {
		tmp := make([]byte, len(line))
		for i, c := range line {
			tmp[i] = byte(c)
		}
		grid = append(grid, tmp)
	}

	grids := [][][]byte{grid}
	iterations := 1000000000
	for i := 0; i < iterations; i++ {
		nextGrid := compute4Cycle(grid)

		for j := 0; j < len(grids); j++ {
			if equal(grids[j], nextGrid) {
				flip := (iterations - j) % (i - j + 1)
				for k := 0; k < flip; k++ {
					nextGrid = compute4Cycle(nextGrid)
				}

				fmt.Println("Part 2:", calcWeight(nextGrid))
				return
			}
		}

		grids = append(grids, nextGrid)
		grid = nextGrid
	}
}
