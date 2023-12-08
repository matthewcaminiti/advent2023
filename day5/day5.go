package day5

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

type transition struct {
	destStart   int
	sourceStart int
	length      int
}

func resolveTransitions(tx [][]transition, src int) int {
	for i, transitions := range tx {
		for _, transition := range transitions {
			d := src - transition.sourceStart
			if d > 0 && d < transition.length {
				dest := transition.destStart + d
				// matches this transition
				if i == len(tx)-1 {
					// is last transition block
					return dest
				}

				return resolveTransitions(tx[i+1:], dest)
			}
		}

		return resolveTransitions(tx[i+1:], src)
	}

	return src
}

func Part1(input string) {
	// source-to-dest
	// dest_range_start source_range_start range_len

	// starting with seed
	// find it's corresponding soil number
	// with soil, find its correspodning fertilizer number
	// with fertilizer, find it's corresponding water
	groups := strings.Split(input, "\n\n")

	transs := [][]transition{}
	for i := 1; i < len(groups); i++ {
		lines := strings.Split(groups[i], "\n")

		transs = append(transs, make([]transition, len(lines)))

		for i, line := range lines {
			if strings.Contains(line, ":") {
				continue
			}
			comps := strings.Fields(line)

			destStart, _ := strconv.Atoi(comps[0])
			sourceStart, _ := strconv.Atoi(comps[1])
			length, _ := strconv.Atoi(comps[2])

			transs[len(transs)-1][i] = transition{
				destStart:   destStart,
				sourceStart: sourceStart,
				length:      length,
			}
		}
	}

	seeds := strings.Fields(groups[0][7:])
	lowest := 1000000000000000000
	for _, _seed := range seeds {
		src, _ := strconv.Atoi(_seed)

		if loc := resolveTransitions(transs, src); loc < lowest {
			lowest = loc
		}
	}

	fmt.Println("Part 1:", lowest)
}

func Part2(input string) {
	// source-to-dest
	// dest_range_start source_range_start range_len

	// starting with seed
	// find it's corresponding soil number
	// with soil, find its correspodning fertilizer number
	// with fertilizer, find it's corresponding water
	groups := strings.Split(input, "\n\n")

	transs := [][]transition{}
	for i := 1; i < len(groups); i++ {
		lines := strings.Split(groups[i], "\n")

		transs = append(transs, make([]transition, len(lines)))

		for i, line := range lines {
			if strings.Contains(line, ":") {
				continue
			}
			comps := strings.Fields(line)

			destStart, _ := strconv.Atoi(comps[0])
			sourceStart, _ := strconv.Atoi(comps[1])
			length, _ := strconv.Atoi(comps[2])

			transs[len(transs)-1][i] = transition{
				destStart:   destStart,
				sourceStart: sourceStart,
				length:      length,
			}
		}
	}

	var wg sync.WaitGroup
	seeds := strings.Fields(groups[0][7:])
	lowest := 1000000000000000000
	for i := 0; i < len(seeds)-1; i += 2 {
		fmt.Println(i, len(seeds)-1)
		seedStart, _ := strconv.Atoi(seeds[i])
		seedLen, _ := strconv.Atoi(seeds[i+1])

		// for _, transition := range transs[0] {
		// }

		wg.Add(1)
		ret := func() {
			defer wg.Done()
			for src := seedStart; src <= seedStart+seedLen; src++ {
				loc := resolveTransitions(transs, src)
				if loc < lowest {
					lowest = loc
				}
				// if loc := resolveTransitions(transs, src); loc < lowest {
				//     lowest = loc
				// }
			}
		}

		go ret()
	}

	wg.Wait()
	fmt.Println("Part 2:", lowest)
}

func max(x, y int) int {
	if x >= y {
		return x
	}

	return y
}

func min(x, y int) int {
	if x <= y {
		return x
	}

	return y
}

func resolveTransitionsRange(tx [][]transition, src int, n int) int {
	srcR := src + n
	if len(tx) == 1 {
		for _, transition := range tx[0] {
			tl, tr := transition.sourceStart, transition.sourceStart+transition.length

			if src < tl {
				if srcR < tl {
					continue
				}

				return min(src, transition.destStart)
			} else if src <= tr {
				if srcR <= tr {
					return transition.destStart + src - tl
				}

				return min(transition.destStart+src-tl, tr)
			}
		}

		return src
	}

	ranges := [][]int{{src, n}}
	lowest := 1000000000000000000
	for _, transition := range tx[0] {
		newRanges := [][]int{}
		for _, _range := range ranges {
			tl, tr := transition.sourceStart, transition.sourceStart+transition.length
			rl, rr := _range[0], _range[0]+_range[1]

			if rl < tl {
				if rr < tl {
					newRanges = append(newRanges, []int{_range[0], _range[1]})
					continue
				}

				if rr <= tr {
					newRanges = append(newRanges, []int{rl, tl - rl})
					val := resolveTransitionsRange(tx[1:], transition.destStart, rr-tl)
					lowest = min(lowest, val)
					continue
				}

				if rr > tr {
					newRanges = append(newRanges, []int{rl, tl - rl}, []int{tr, rr - tr})
					val := resolveTransitionsRange(tx[1:], transition.destStart, transition.length)
					lowest = min(lowest, val)
				}
			} else if rl <= tr {
				if rr <= tr {
					val := resolveTransitionsRange(tx[1:], transition.destStart+rl-tl, _range[1])
					lowest = min(lowest, val)
					continue
				}

				val := resolveTransitionsRange(tx[1:], transition.destStart+rl-tl, tr-rl)
				lowest = min(lowest, val)
				newRanges = append(newRanges, []int{tr, rr - tr})
			} else {
				newRanges = append(newRanges, []int{_range[0], _range[1]})
			}
		}

		ranges = newRanges
	}

	// this code unnecessary for run against input data! (but is logically required)
	for _, _range := range ranges {
		val := resolveTransitionsRange(tx[1:], _range[0], _range[1])
		lowest = min(lowest, val)
	}

	return lowest
}

func OptPart2(input string) {
	groups := strings.Split(input, "\n\n")

	transs := make([][]transition, 0)
	for i := 1; i < len(groups); i++ {
		lines := strings.Split(groups[i], "\n")

		transs = append(transs, make([]transition, len(lines)-1))

		for j, line := range lines {
			if strings.Contains(line, ":") {
				continue
			}
			comps := strings.Fields(line)

			destStart, _ := strconv.Atoi(comps[0])
			sourceStart, _ := strconv.Atoi(comps[1])
			length, _ := strconv.Atoi(comps[2])

			transs[len(transs)-1][j-1] = transition{
				destStart:   destStart,
				sourceStart: sourceStart,
				length:      length,
			}
		}
	}

	seeds := strings.Fields(groups[0][7:])

	lowest := 1000000000000000000
	for i := 0; i < len(seeds)-1; i += 2 {
		seedStart, _ := strconv.Atoi(seeds[i])
		seedLen, _ := strconv.Atoi(seeds[i+1])

		loc := resolveTransitionsRange(transs, seedStart, seedLen)
		if loc < lowest {
			lowest = loc
		}
	}

	fmt.Println("Part 2 Opt:", lowest)
}
