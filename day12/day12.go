package day12

import (
	"fmt"
	"strconv"
	"strings"
)

func isValid(n int, s string, next byte) bool {
	ctr := 0
	if next == '#' {
		return false
	}

	for _, c := range s {
		if c == '#' || c == '?' {
			ctr++
		}
	}

	return ctr == n
}

func getPermutations(s string, nums []int) int {
	if len(nums) == 0 {
		for _, c := range s {
			if c == '#' {
				return 0
			}
		}

		return 1
	}

	n := nums[0]
	ctr := 0
	for i := 0; i+n <= len(s); i++ {
		var nextChar byte
		if i+n <= len(s)-1 {
			nextChar = s[i+n]
		}

		ss := s[i : i+n]
		valid := isValid(n, ss, nextChar)
		// fmt.Printf("%d %s(%c) %t\n", n, ss, nextChar, valid)

		if valid {
			if i+n+1 > len(s) {
				if len(nums) == 1 {
					ctr++
				} else {
					break // performance opt
				}
			} else {
				// +1 to skip over padding char #x# (i.e x cannot be used)
				ctr += getPermutations(s[i+n+1:], nums[1:])
			}
		}

		// if first char is '#', we must stop this window
		if ss[0] == '#' {
			break
		}
	}

	return ctr
}

func Part1(input string) {
	lines := strings.Split(input, "\n")

	puzzles := make([]string, len(lines))
	inputs := make([][]int, len(lines))
	for i, line := range lines {
		comps := strings.Fields(line)
		puzzle := strings.Trim(comps[0], ".")
		_tmp := strings.ReplaceAll(puzzle, "..", ".")
		for _tmp != puzzle {
			puzzle = _tmp
			_tmp = strings.ReplaceAll(_tmp, "..", ".")
		}
		puzzles[i] = _tmp

		nums := strings.Split(comps[1], ",")
		tmp := make([]int, len(nums))
		for j, c := range nums {
			n, _ := strconv.Atoi(c)
			tmp[j] = n
		}

		inputs[i] = tmp
	}

	tot := 0
	for i := 0; i < len(puzzles); i++ {
		n := getPermutations(puzzles[i], inputs[i])
		tot += n
	}

	fmt.Println("Part 1:", tot)
}

type permHit struct {
	nums  []int
	perms int
}

func getMemoizedPermutations(history *map[string][]permHit, s string, nums []int) int {
	if len(nums) == 0 {
		for _, c := range s {
			if c == '#' {
				return 0
			}
		}

		return 1
	}

	hits, exists := (*history)[s]
	if exists {
		for _, hit := range hits {
			if len(hit.nums) != len(nums) {
				continue
			}

			same := true
			for i := 0; i < len(hit.nums); i++ {
				if hit.nums[i] != nums[i] {
					same = false
					break
				}
			}

			if same {
				return hit.perms
			}
		}
	}

	n := nums[0]
	ctr := 0
	for i := 0; i+n <= len(s); i++ {
		var nextChar byte
		if i+n <= len(s)-1 {
			nextChar = s[i+n]
		}

		ss := s[i : i+n]
		valid := isValid(n, ss, nextChar)

		if valid {
			if i+n+1 > len(s) {
				if len(nums) == 1 {
					ctr++
				} else {
					break // performance opt
				}
			} else {
				// +1 to skip over padding char #x# (i.e x cannot be used)
				subs := s[i+n+1:]
				subnums := nums[1:]
				val := getMemoizedPermutations(history, subs, subnums)
				if _, exists := (*history)[subs]; !exists {
					(*history)[subs] = []permHit{
						{
							nums:  subnums,
							perms: val,
						},
					}
				} else {
					(*history)[subs] = append((*history)[subs],
						permHit{
							nums:  subnums,
							perms: val,
						},
					)
				}

				ctr += val
			}
		}

		// if first char is '#', we must stop this window
		if ss[0] == '#' {
			break
		}
	}

	return ctr
}

func Part2(input string) {
	lines := strings.Split(input, "\n")

	puzzles := make([]string, 0)
	inputs := make([][]int, 0)
	for _, line := range lines {
		comps := strings.Fields(line)

		for i := 0; i < 5; i++ {
			raws := []string{}
			for j := 0; j < i+1; j++ {
				raws = append(raws, comps[0])
			}
			puzzle := strings.Trim(strings.Join(raws, "?"), ".")
			_tmp := strings.ReplaceAll(puzzle, "..", ".")
			for _tmp != puzzle {
				puzzle = _tmp
				_tmp = strings.ReplaceAll(_tmp, "..", ".")
			}
			puzzles = append(puzzles, _tmp)
		}

		nums := strings.Split(comps[1], ",")
		for i := 0; i < 5; i++ {
			tmp := make([]int, 0)
			for j := 0; j < i+1; j++ {
				for _, c := range nums {
					n, _ := strconv.Atoi(c)
					tmp = append(tmp, n)
				}
			}

			inputs = append(inputs, tmp)
		}
	}

	tot := 0
	for i := 0; i+5 <= len(puzzles); i += 5 {
		history := map[string][]permHit{}
		v := getMemoizedPermutations(&history, puzzles[i+4], inputs[i+4])
		tot += v
	}

	fmt.Printf("Part 2: %d\n", tot)
}
