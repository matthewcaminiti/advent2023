package day3

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

var lineLen int

func isDigit(r rune) bool {
	return '0' <= r && r <= '9'
}

func isSymbol(r rune) bool {
	return r != '.' && !isDigit(r)
}

func Part1(input string) {
	lines := strings.Split(input, "\n")

	lineLen = len(lines[0])

	sum := 0
	for y, line := range lines {
		lastSymX := -100
		numUsable := false
		num := ""
		for x, c := range line {
			if y > 0 && lastSymX <= x {
				if isSymbol(rune(lines[y-1][x])) {
					// symbol above
					lastSymX = x
				}

				if x < lineLen-1 && isSymbol(rune(lines[y-1][x+1])) {
					// symbol top right
					lastSymX = x + 1
				}
			}

			if lastSymX <= x && isSymbol(rune(line[x])) {
				// curr pos is symbol
				lastSymX = x
			}
			if x < lineLen-1 && isSymbol(rune(line[x+1])) {
				// symbol right
				lastSymX = x + 1
			}

			if y < len(lines)-1 && lastSymX <= x {
				if isSymbol(rune(lines[y+1][x])) {
					// symbol below
					lastSymX = x
				}
				if x < lineLen-1 && isSymbol(rune(lines[y+1][x+1])) {
					// symbol bottom right
					lastSymX = x + 1
				}
			}

			if isDigit(c) {
				num += string(c)
				if x-lastSymX <= 1 {
					numUsable = true
				}
			} else {
				if !numUsable {
					num = ""
				}

				if len(num) > 0 && numUsable {
					n, err := strconv.Atoi(num)
					if err != nil {
						panic(err)
					}

					sum += n

					num = ""
					numUsable = false
				}
			}
		}

		if len(num) > 0 && numUsable {
			n, err := strconv.Atoi(num)
			if err != nil {
				panic(err)
			}

			sum += n
		}
	}

	// fmt.Println("Part 1: ", sum)
}

func getLastSymX(lines []string, lastSymX, x, y int) int {
	if isSymbol(rune(lines[y][x])) {
		// curr pos is symbol
		return x
	}
	if x < lineLen-1 && isSymbol(rune(lines[y][x+1])) {
		// symbol right
		return x + 1
	}

	if y > 0 {
		if x < lineLen-1 && isSymbol(rune(lines[y-1][x+1])) {
			// symbol top right
			return x + 1
		}
		if isSymbol(rune(lines[y-1][x])) {
			// symbol above
			return x
		}

	}

	if y < len(lines)-1 {
		if x < lineLen-1 && isSymbol(rune(lines[y+1][x+1])) {
			// symbol bottom right
			return x + 1
		}
		if isSymbol(rune(lines[y+1][x])) {
			// symbol below
			return x
		}
	}

	return lastSymX
}

func getIsTouching(lines []string, x, y int) bool {
	for i := Max(0, x-1); i <= Min(lineLen-1, x+1); i++ {
		for j := Max(0, y-1); j <= Min(len(lines)-1, y+1); j++ {
			if isSymbol(rune(lines[j][i])) {
				return true
			}
		}
	}

	return false
}

func OptPart1(input string) {
	lines := strings.Split(input, "\n")

	lineLen = len(lines[0])

	sum := 0
	for y := 0; y < len(lines); y++ {
		numUsable := false
		num := ""
		for x := 0; x < len(lines[y]); x++ {
			if isDigit(rune(lines[y][x])) {
				if !numUsable {
					numUsable = getIsTouching(lines, x, y)
				}
				num += string(lines[y][x])
			} else {
				if !numUsable {
					num = ""
					continue
				}

				if len(num) > 0 && numUsable {
					n, _ := strconv.Atoi(num)

					sum += n

					num = ""
					numUsable = false
				}
			}
		}

		if len(num) > 0 && numUsable {
			n, _ := strconv.Atoi(num)
			sum += n
		}
	}

	// fmt.Println("Part 1: ", sum)
}

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

func isAsterisk(r byte) bool {
	return r == '*'
}

func Part2(lines []string) {
	// same as part 1, except track of valid numbers associated with '*' (record by x,y idx)
	// ensure symbol lookahead logic doesn't disregard asterisks
	// must account for scenario in which number adjacent two asterisks

	lineLen = len(lines[0])

	gears := map[string][]string{}

	sum := 0
	for y, line := range lines {
		lastGearY := -100
		lastGearX := -100

		numUsable := false
		num := ""

		for x, c := range line {
			hasGear := numUsable && len(num) > 0

			if y > 0 && lastGearX <= x {
				if isAsterisk(lines[y-1][x]) && !hasGear {
					// asterisk above
					lastGearX = x
					lastGearY = y - 1
				}

				if x < lineLen-1 && isAsterisk(lines[y-1][x+1]) && !hasGear {
					// symbol top right
					lastGearX = x + 1
					lastGearY = y - 1
				}
			}

			if lastGearX <= x && isAsterisk(line[x]) && !hasGear {
				// curr pos is symbol
				lastGearX = x
				lastGearY = y
			}

			if x < lineLen-1 && isAsterisk(line[x+1]) && !hasGear {
				// symbol right
				lastGearX = x + 1
				lastGearY = y
			}

			if y < len(lines)-1 && lastGearX <= x {
				if isAsterisk(lines[y+1][x]) && !hasGear {
					// symbol below
					lastGearX = x
					lastGearY = y + 1
				}

				if x < lineLen-1 && isAsterisk(lines[y+1][x+1]) && !hasGear {
					// symbol bottom right
					lastGearX = x + 1
					lastGearY = y + 1
				}
			}

			if isDigit(c) {
				num += string(c)
				if x-lastGearX <= 1 {
					numUsable = true
				}
			} else {
				if !numUsable {
					num = ""
					continue
				}

				// fmt.Printf("usable number \"%s\", gear: (%d, %d)\n", num, lastGearX, lastGearY)

				k := fmt.Sprintf("%d-%d", lastGearX, lastGearY)

				gear, _ := gears[k]
				gear = append(gear, num)
				gears[k] = gear

				num = ""
				numUsable = false
			}
		}

		if numUsable {
			k := fmt.Sprintf("%d-%d", lastGearX, lastGearY)

			gear, _ := gears[k]
			gear = append(gear, num)
			gears[k] = gear
		}
	}

	for _, v := range gears {
		if len(v) != 2 {
			continue
		}

		n1, _ := strconv.Atoi(v[0])
		n2, _ := strconv.Atoi(v[1])

		sum += n1 * n2
	}

	fmt.Println("Part 2: ", sum)
}
