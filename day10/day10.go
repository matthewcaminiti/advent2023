package day10

import (
	"fmt"
	// "math"
	"sort"
	"strings"
)

type coord struct {
	x, y int
}

type pipe struct {
	from  coord
	at    coord
	depth int
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

func canGoTo(toC byte, from coord, to coord) bool {
	// top
	if to.x == from.x && to.y < from.y {
		return !(toC != '|' && toC != '7' && toC != 'F')
	}
	// right
	if to.x > from.x && to.y == from.y {
		return !(toC != '-' && toC != '7' && toC != 'J')
	}
	// bottom
	if to.x == from.x && to.y > from.y {
		return !(toC != '|' && toC != 'L' && toC != 'J')
	}
	// left
	if to.x < from.x && to.y == from.y {
		return !(toC != '-' && toC != 'L' && toC != 'F')
	}

	panic("why you go diagonal?")
}

func Part1(input string) {
	lines := strings.Split(input, "\n")

	h, w := len(lines), len(lines[0])

	// put pipes into 2d array and find S coord
	// from S, call "traverse" each connected pipe BFS, keeping track of depth
	// if two coords match, mark distance and eliminate from search queue

	sCoord := coord{}
	for y := 0; y < h; y++ {
		found := false
		for x := 0; x < w; x++ {
			if lines[y][x] == 'S' {
				sCoord.x = x
				sCoord.y = y
				found = true
				break
			}
		}
		if found {
			break
		}
	}

	charAtCoord := func(c coord) byte {
		if c.y < 0 || c.y > h {
			return '.'
		}

		if c.x < 0 || c.x > w {
			return '.'
		}

		return lines[c.y][c.x]
	}

	pipes := []pipe{{from: sCoord, at: sCoord, depth: 0}}
	maxDepth := 0
	for len(pipes) != 0 {
		newPipes := []pipe{}

		for i, _pipe := range pipes {
			atChar := lines[_pipe.at.y][_pipe.at.x]

			if _pipe.from == _pipe.at {
				// start coord
				// top
				to := coord{x: _pipe.at.x, y: _pipe.at.y - 1}
				if canGoTo(charAtCoord(to), _pipe.at, to) {
					newPipes = append(newPipes, pipe{from: _pipe.at, at: to, depth: _pipe.depth + 1})
				}
				// right
				to = coord{x: _pipe.at.x + 1, y: _pipe.at.y}
				if canGoTo(charAtCoord(to), _pipe.at, to) {
					newPipes = append(newPipes, pipe{from: _pipe.at, at: to, depth: _pipe.depth + 1})
				}
				// bottom
				to = coord{x: _pipe.at.x, y: _pipe.at.y + 1}
				if canGoTo(charAtCoord(to), _pipe.at, to) {
					newPipes = append(newPipes, pipe{from: _pipe.at, at: to, depth: _pipe.depth + 1})
				}
				// left
				to = coord{x: _pipe.at.x - 1, y: _pipe.at.y}
				if canGoTo(charAtCoord(to), _pipe.at, to) {
					newPipes = append(newPipes, pipe{from: _pipe.at, at: to, depth: _pipe.depth + 1})
				}
				break
			}

			// check that pipe is not also in list
			matchFound := false
			for j, __pipe := range pipes {
				if _pipe.at == __pipe.at && i != j {
					matchFound = true
					if _pipe.depth > maxDepth {
						maxDepth = _pipe.depth
					}
					break
				}
			}
			if matchFound {
				continue
			}

			diry, dirx := _pipe.at.y-_pipe.from.y, _pipe.at.x-_pipe.from.x

			fromTop := diry > 0
			fromBottom := diry < 0
			fromLeft := dirx > 0
			fromRight := dirx < 0

			var to coord

			switch atChar {
			case '|':
				if fromBottom {
					to = coord{x: _pipe.at.x, y: _pipe.at.y - 1}
				} else if fromTop {
					to = coord{x: _pipe.at.x, y: _pipe.at.y + 1}
				}
			case '-':
				if fromLeft {
					to = coord{x: _pipe.at.x + 1, y: _pipe.at.y}
				} else if fromRight {
					to = coord{x: _pipe.at.x - 1, y: _pipe.at.y}
				}
			case 'L':
				if fromTop {
					to = coord{x: _pipe.at.x + 1, y: _pipe.at.y}
				} else if fromRight {
					to = coord{x: _pipe.at.x, y: _pipe.at.y - 1}
				}
			case 'J':
				if fromTop {
					to = coord{x: _pipe.at.x - 1, y: _pipe.at.y}
				} else if fromLeft {
					to = coord{x: _pipe.at.x, y: _pipe.at.y - 1}
				}
			case '7':
				if fromBottom {
					to = coord{x: _pipe.at.x - 1, y: _pipe.at.y}
				} else if fromLeft {
					to = coord{x: _pipe.at.x, y: _pipe.at.y + 1}
				}
			case 'F':
				if fromBottom {
					to = coord{x: _pipe.at.x + 1, y: _pipe.at.y}
				} else if fromRight {
					to = coord{x: _pipe.at.x, y: _pipe.at.y + 1}
				}
			}

			if canGoTo(charAtCoord(to), _pipe.at, to) {
				newPipes = append(newPipes, pipe{from: _pipe.at, at: to, depth: _pipe.depth + 1})
			}
		}

		pipes = newPipes
	}

	fmt.Println("Part 1:", maxDepth)
}

func Part2(input string) {
	lines := strings.Split(input, "\n")

	h, w := len(lines), len(lines[0])

	// put pipes into 2d array and find S coord
	// from S, call "traverse" each connected pipe BFS, keeping track of depth
	// keep track of pipes traversed to keep coordinates of loop
	// if two coords match, mark distance and eliminate from search queue

	sCoord := coord{}
	for y := 0; y < h; y++ {
		found := false
		for x := 0; x < w; x++ {
			if lines[y][x] == 'S' {
				sCoord.x = x
				sCoord.y = y
				found = true
				break
			}
		}
		if found {
			break
		}
	}

	charAtCoord := func(c coord) byte {
		if c.y < 0 || c.y > h {
			return '.'
		}

		if c.x < 0 || c.x > w {
			return '.'
		}

		return lines[c.y][c.x]
	}

	type step struct {
		group int
		from  coord
		at    coord
		depth int
	}

	historicalSteps := []step{}
	steps := []step{{group: 0, from: sCoord, at: sCoord, depth: 0}}
	maxDepth := 0
	maxDepthGroupA, maxDepthGroupB := 0, 0
	for len(steps) != 0 {
		newSteps := []step{}

		for i, _pipe := range steps {
			atChar := lines[_pipe.at.y][_pipe.at.x]

			if _pipe.from == _pipe.at {
				// start coord
				// top
				to := coord{x: _pipe.at.x, y: _pipe.at.y - 1}
				if canGoTo(charAtCoord(to), _pipe.at, to) {
					newSteps = append(newSteps, step{
						group: 1,
						from:  _pipe.at,
						at:    to,
						depth: _pipe.depth + 1,
					})
				}
				// right
				to = coord{x: _pipe.at.x + 1, y: _pipe.at.y}
				if canGoTo(charAtCoord(to), _pipe.at, to) {
					newSteps = append(newSteps, step{
						group: 2,
						from:  _pipe.at,
						at:    to,
						depth: _pipe.depth + 1,
					})
				}
				// bottom
				to = coord{x: _pipe.at.x, y: _pipe.at.y + 1}
				if canGoTo(charAtCoord(to), _pipe.at, to) {
					newSteps = append(newSteps, step{
						group: 3,
						from:  _pipe.at,
						at:    to,
						depth: _pipe.depth + 1,
					})
				}
				// left
				to = coord{x: _pipe.at.x - 1, y: _pipe.at.y}
				if canGoTo(charAtCoord(to), _pipe.at, to) {
					newSteps = append(newSteps, step{
						group: 4,
						from:  _pipe.at,
						at:    to,
						depth: _pipe.depth + 1,
					})
				}
				break
			}

			// check that pipe is not also in list
			matchFound := false
			for j, __pipe := range steps {
				if _pipe.at == __pipe.at && i != j {
					matchFound = true
					if _pipe.depth > maxDepth {
						maxDepth = _pipe.depth
						maxDepthGroupA = steps[i].group
						maxDepthGroupB = steps[j].group
					}
					break
				}
			}
			if matchFound {
				continue
			}

			diry, dirx := _pipe.at.y-_pipe.from.y, _pipe.at.x-_pipe.from.x

			fromTop := diry > 0
			fromBottom := diry < 0
			fromLeft := dirx > 0
			fromRight := dirx < 0

			var to coord

			switch atChar {
			case '|':
				if fromBottom {
					to = coord{x: _pipe.at.x, y: _pipe.at.y - 1}
				} else if fromTop {
					to = coord{x: _pipe.at.x, y: _pipe.at.y + 1}
				}
			case '-':
				if fromLeft {
					to = coord{x: _pipe.at.x + 1, y: _pipe.at.y}
				} else if fromRight {
					to = coord{x: _pipe.at.x - 1, y: _pipe.at.y}
				}
			case 'L':
				if fromTop {
					to = coord{x: _pipe.at.x + 1, y: _pipe.at.y}
				} else if fromRight {
					to = coord{x: _pipe.at.x, y: _pipe.at.y - 1}
				}
			case 'J':
				if fromTop {
					to = coord{x: _pipe.at.x - 1, y: _pipe.at.y}
				} else if fromLeft {
					to = coord{x: _pipe.at.x, y: _pipe.at.y - 1}
				}
			case '7':
				if fromBottom {
					to = coord{x: _pipe.at.x - 1, y: _pipe.at.y}
				} else if fromLeft {
					to = coord{x: _pipe.at.x, y: _pipe.at.y + 1}
				}
			case 'F':
				if fromBottom {
					to = coord{x: _pipe.at.x + 1, y: _pipe.at.y}
				} else if fromRight {
					to = coord{x: _pipe.at.x, y: _pipe.at.y + 1}
				}
			}

			if canGoTo(charAtCoord(to), _pipe.at, to) {
				newSteps = append(newSteps, step{
					group: _pipe.group,
					from:  _pipe.at,
					at:    to,
					depth: _pipe.depth + 1,
				})
			}
		}

		historicalSteps = append(historicalSteps, steps...)
		steps = newSteps
	}

	loopSteps := []step{}
	for _, _step := range historicalSteps {
		if _step.group == maxDepthGroupA || _step.group == maxDepthGroupB || _step.group == 0 {
			if _step.group == maxDepthGroupB {
				loopSteps = append(loopSteps, step{
					group: _step.group,
					from:  _step.at,
					at:    _step.from,
					depth: _step.depth + maxDepth,
				})
			} else {
				loopSteps = append(loopSteps, _step)
			}
		}
	}

	sort.Slice(loopSteps, func(i, j int) bool {
		return loopSteps[i].depth < loopSteps[j].depth
		// if loopSteps[i].at.y != loopSteps[j].at.y {
		//     return loopSteps[i].at.y < loopSteps[j].at.y
		// }
		//
		// return loopSteps[i].at.x < loopSteps[j].at.x
	})

	minX, maxX := 100000000000, 0
	minY, maxY := 100000000000, 0
	l, r := 0, 0
	for _, step := range loopSteps {
		minX = min(minX, step.at.x)
		maxX = max(maxX, step.at.x)
		minY = min(minY, step.at.y)
		maxY = max(maxY, step.at.y)

		dx := step.at.x - step.from.x
		dy := step.at.y - step.from.y

		fromTop := dy > 0
		fromBottom := dy < 0
		fromLeft := dx > 0

		c := charAtCoord(step.at)
		switch c {
		case 'F':
			if fromBottom {
				r++
			} else {
				l++
			}
		case '7':
			if fromBottom {
				l++
			} else {
				r++
			}
		case 'J':
			if fromLeft {
				l++
			} else {
				r++
			}
		case 'L':
			if fromTop {
				l++
			} else {
				r++
			}
		}
	}

	leftInside := l > r
	w = maxX - minX + 1
	h = maxY - minY + 1
	fmt.Printf("dim: (%d, %d), left inside: %t\n", w, h, leftInside)

	area := 0

	// create map, 0 is empty, 1 is path, 2 is area
	space := make([][]int, h)
	for i := 0; i < h; i++ {
		space[i] = make([]int, w)
	}

	for i := 1; i < len(loopSteps); i++ {
		_step := loopSteps[i]

		space[_step.at.y-minY][_step.at.x-minX] = 1

		dx := _step.at.x - _step.from.x
		dy := _step.at.y - _step.from.y

		fromTop := dy > 0
		fromBottom := dy < 0
		fromLeft := dx > 0
		fromRight := dx < 0

		x, y := _step.at.x, _step.at.y
		if fromTop {
			if leftInside {
				x++
			} else {
				x--
			}
		} else if fromBottom {
			if leftInside {
				x++
			} else {
				x--
			}
		} else if fromRight {
			if leftInside {
				y++
			} else {
				y--
			}
		} else if fromLeft {
			if leftInside {
				y--
			} else {
				y++
			}
		}

		if x < 0 || x >= w || y < 0 || y >= h {
			continue
		}

		s := space[y][x]
		if s == 0 {
			space[y][x] = 2
		}
	}

	for y, r := range space {
		for x := range r {
			if y == 0 || y == h-1 || x == 0 || x == w-1 {
				if space[y][x] != 1 {
					space[y][x] = 3
				}
			}
		}
	}

	var recurseArea func(x, y, val int)
	recurseArea = func(x, y, val int) {
		// left
		if x-1 >= 0 && space[y][x-1] != 1 && space[y][x-1] != val && space[y][x-1] != 3 {
			space[y][x-1] = val
			recurseArea(x-1, y, val)
		}

		// right
		if x+1 < w && space[y][x+1] != 1 && space[y][x+1] != val && space[y][x+1] != 3 {
			space[y][x+1] = val
			recurseArea(x+1, y, val)
		}

		// top
		if y-1 >= 0 && space[y-1][x] != 1 && space[y-1][x] != val && space[y-1][x] != 3 {
			space[y-1][x] = val
			recurseArea(x, y-1, val)
		}

		// bottom
		if y+1 < h && space[y+1][x] != 1 && space[y+1][x] != val && space[y+1][x] != 3 {
			space[y+1][x] = val
			recurseArea(x, y+1, val)
		}

		// do diagonals for void
		// TODO: need to ensure that every 2 is correctly surrounded
		// i.e a 2 at the corner of a lake must have the 5 adjacent tiles
		// (adjacent including diagonal) having no difference of depth > 1
		if val == 3 {
			if y+1 < h && x+1 < w {
				next := space[y+1][x+1]
				if next != 1 && next != val {
					space[y+1][x+1] = val
					recurseArea(x+1, y+1, val)
				}
			}
			if y+1 < h && x-1 >= 0 {
				next := space[y+1][x-1]
				if next != 1 && next != val {
					space[y+1][x-1] = val
					recurseArea(x-1, y+1, val)
				}
			}
			if y-1 >= 0 && x-1 >= 0 {
				next := space[y-1][x-1]
				if next != 1 && next != val {
					space[y-1][x-1] = val
					recurseArea(x-1, y-1, val)
				}
			}
			if y-1 >= 0 && x+1 < w {
				next := space[y-1][x+1]
				if next != 1 && next != val {
					space[y-1][x+1] = val
					recurseArea(x+1, y-1, val)
				}
			}
		}
	}

	for y := range space {
		for x := range space[y] {
			v := space[y][x]
			if v == 1 || v == 0 {
				continue
			}

			if v == 3 || v == 2 {
				recurseArea(x, y, v)
			}
		}
	}

	for y := range space {
		for x := range space[y] {
			v := space[y][x]
			if v == 2 {
				area++
			}
			if v == 3 {
				fmt.Print(".")
			} else {
				fmt.Printf("%d", v)
			}
		}
		fmt.Println()
	}

	fmt.Println("Part 2:", area)
}
