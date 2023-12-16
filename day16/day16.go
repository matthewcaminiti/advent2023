package day16

import (
	"fmt"
	"strings"
)

type direction int

const (
	up direction = iota + 1
	right
	down
	left
)

type _beam struct {
	dir direction
	x   int
	y   int
}

func actBeam(b _beam, char byte) []_beam {
	return []_beam{}
}

func Part1(input string) {
	lines := strings.Split(input, "\n")

	fmt.Println("Part 1:", getEnergy(lines, _beam{right, 0, 0}))
}

func getEnergy(lines []string, start _beam) int {
	energized := make([][]int, len(lines))
	for i := range lines {
		energized[i] = make([]int, len(lines[i]))
	}

	beams := []_beam{start}

	w := len(lines[0])
	h := len(lines)

	for len(beams) > 0 {
		nextBeams := []_beam{}
		for _, beam := range beams {
			if energized[beam.y][beam.x] == int(beam.dir) {
				continue
			}

			energized[beam.y][beam.x] = int(beam.dir)

			canRight := beam.x+1 < w
			canUp := beam.y-1 >= 0
			canDown := beam.y+1 < h
			canLeft := beam.x-1 >= 0

			r := _beam{dir: right, x: beam.x + 1, y: beam.y}
			u := _beam{dir: up, x: beam.x, y: beam.y - 1}
			d := _beam{dir: down, x: beam.x, y: beam.y + 1}
			l := _beam{dir: left, x: beam.x - 1, y: beam.y}

			switch lines[beam.y][beam.x] {
			case '|':
				if beam.dir == right || beam.dir == left {
					if canUp {
						nextBeams = append(nextBeams, u)
					}
					if canDown {
						nextBeams = append(nextBeams, d)
					}
				} else if beam.dir == up && canUp {
					nextBeams = append(nextBeams, u)
				} else if beam.dir == down && canDown {
					nextBeams = append(nextBeams, d)
				}
			case '\\':
				if beam.dir == right && canDown {
					nextBeams = append(nextBeams, d)
				} else if beam.dir == up && canLeft {
					nextBeams = append(nextBeams, l)
				} else if beam.dir == left && canUp {
					nextBeams = append(nextBeams, u)
				} else if beam.dir == down && canRight {
					nextBeams = append(nextBeams, r)
				}
			case '/':
				if beam.dir == right && canUp {
					nextBeams = append(nextBeams, u)
				} else if beam.dir == up && canRight {
					nextBeams = append(nextBeams, r)
				} else if beam.dir == left && canDown {
					nextBeams = append(nextBeams, d)
				} else if beam.dir == down && canLeft {
					nextBeams = append(nextBeams, l)
				}
			case '-':
				if beam.dir == right && canRight {
					nextBeams = append(nextBeams, r)
				} else if beam.dir == up || beam.dir == down {
					if canLeft {
						nextBeams = append(nextBeams, l)
					}
					if canRight {
						nextBeams = append(nextBeams, r)
					}
				} else if beam.dir == left && canLeft {
					nextBeams = append(nextBeams, l)
				}
			default:
				if beam.dir == right && canRight {
					nextBeams = append(nextBeams, r)
				} else if beam.dir == left && canLeft {
					nextBeams = append(nextBeams, l)
				} else if beam.dir == up && canUp {
					nextBeams = append(nextBeams, u)
				} else if beam.dir == down && canDown {
					nextBeams = append(nextBeams, d)
				}
			}

			// fmt.Printf("%d %+v -> %+v\n", i, beam, nextBeams)
		}

		beams = nextBeams
	}

	tot := 0
	for _, row := range energized {
		for _, x := range row {
			if x != 0 {
				tot++
			}
		}
	}

	return tot
}

func Part2(input string) {
	lines := strings.Split(input, "\n")

	h := len(lines)
	w := len(lines[0])

	max := 0
	for x := 0; x < len(lines[0]); x++ {
		v := getEnergy(lines, _beam{dir: down, x: x, y: 0})
		if v > max {
			max = v
		}

		v = getEnergy(lines, _beam{dir: up, x: x, y: h - 1})
		if v > max {
			max = v
		}
	}

	for y := 0; y < len(lines); y++ {
		v := getEnergy(lines, _beam{dir: right, x: 0, y: y})
		if v > max {
			max = v
		}

		v = getEnergy(lines, _beam{dir: left, x: w - 1, y: y})
		if v > max {
			max = v
		}
	}

	fmt.Println("Part 2:", max)
}
