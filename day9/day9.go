package day9

import (
	"fmt"
	"strconv"
	"strings"
)

func allZeros(x []int) bool {
	for i := range x {
		if x[i] != 0 {
			return false
		}
	}

	return true
}

func Part1(input string) {
	lines := strings.Split(input, "\n")

	sum := 0
	for _, line := range lines {
		_nums := strings.Fields(line)

		nums := make([]int, len(_nums))
		for i := range _nums {
			nums[i], _ = strconv.Atoi(_nums[i])
		}

		seqs := [][]int{nums}

		for true {
			diffs := []int{}
			curr := seqs[len(seqs)-1]

			allZeros := true
			for i := range curr {
				if i < len(curr)-1 {
					diff := curr[i+1] - curr[i]
					if diff != 0 {
						allZeros = false
					}
					diffs = append(diffs, diff)
				}
			}

			seqs = append(seqs, diffs)

			if allZeros {
				break
			}
		}

		// rollup
		curr := 0
		for i := len(seqs) - 1; i >= 0; i-- {
			curr += seqs[i][len(seqs[i])-1]
		}

		sum += curr

	}

	fmt.Println("Part 1:", sum)
}

func Part2(input string) {
	lines := strings.Split(input, "\n")

	sum := 0
	for _, line := range lines {
		_nums := strings.Fields(line)

		nums := make([]int, len(_nums))
		for i := range _nums {
			nums[i], _ = strconv.Atoi(_nums[i])
		}

		seqs := [][]int{nums}

		for true {
			diffs := []int{}
			curr := seqs[len(seqs)-1]

			allZeros := true
			for i := range curr {
				if i < len(curr)-1 {
					diff := curr[i+1] - curr[i]
					if diff != 0 {
						allZeros = false
					}
					diffs = append(diffs, diff)
				}
			}

			seqs = append(seqs, diffs)

			if allZeros {
				break
			}
		}

		// rollup
		curr := 0
		for i := len(seqs) - 1; i >= 0; i-- {
			curr = seqs[i][0] - curr
		}

		sum += curr
	}

	fmt.Println("Part 2:", sum)
}
