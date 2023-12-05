package day3

import (
	"strconv"
	"strings"
	"unicode"
)

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func checkIfTouching(startX int, endX int, y int, grid [][]rune) bool {
	for startRow := Max(y-1, 0); startRow <= Min(y+1, len(grid)-1); startRow++ {
		for i := Max(startX-1, 0); i <= Min(endX+1, len(grid[startRow])-1); i++ {
			if grid[startRow][i] != '.' && (grid[startRow][i] < '0' || grid[startRow][i] > '9') {
				return true
			}
		}
	}
	return false
}

func KadatzPart1(input string) {
	var grid [][]rune

	lines := strings.Split(input, "\n")

	for _, line := range lines {
		grid = append(grid, []rune(line))
	}

	sum := 0
	for i := 0; i < len(grid); i++ {
		partNum := ""
		for j := 0; j < len(grid[i]); j++ {

			if grid[i][j] >= '0' && grid[i][j] <= '9' {
				partNum += string(grid[i][j])
			}

			if !unicode.IsDigit(grid[i][j]) || j == len(grid[i])-1 {
				if partNum != "" {
					startX := Max(j-len(partNum), 0)
					endX := j - 1
					if checkIfTouching(startX, endX, i, grid) {
						num, _ := strconv.Atoi(partNum)
						sum += num
					}
					partNum = ""
				}
			}
		}
	}

	// fmt.Println("Sum: ", sum)
}

func OptKadatzPart1(input string) {
	lines := strings.Split(input, "\n")

	grid := make([][]rune, len(lines))

	for i, line := range lines {
		grid[i] = []rune(line)
	}

	sum := 0
	for i := 0; i < len(grid); i++ {
		partNum := ""
		for j := 0; j < len(grid[i]); j++ {

			if grid[i][j] >= '0' && grid[i][j] <= '9' {
				partNum += string(grid[i][j])
			}

			if !unicode.IsDigit(grid[i][j]) || j == len(grid[i])-1 {
				if partNum != "" {
					startX := Max(j-len(partNum), 0)
					endX := j - 1
					if checkIfTouching(startX, endX, i, grid) {
						num, _ := strconv.Atoi(partNum)
						sum += num
					}
					partNum = ""
				}
			}
		}
	}

	// fmt.Println("Sum: ", sum)
}
