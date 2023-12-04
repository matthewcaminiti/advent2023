package day3

import (
	"fmt"
	"strconv"
)

var lineLen int

func isDigit(r rune) bool {
	return '0' <= r && r <= '9'
}

func isSymbol(r rune) bool {
	return r != '.' && !isDigit(r)
}

func Part1(lines []string) {
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

	fmt.Println("Part 1: ", sum)
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
