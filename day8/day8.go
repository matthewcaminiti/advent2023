package day8

import (
	"fmt"
	"strings"
)

func Part1(input string) {
	lines := strings.Split(input, "\n")

	nodes := map[string][]string{}

	for _, line := range lines[2:] {
		comps := strings.Fields(line)
		at := comps[0]
		l := comps[2][1 : len(comps[2])-1]
		r := comps[3][:len(comps[3])-1]

		nodes[at] = []string{l, r}
	}

	at := "AAA"
	instructions := lines[0]

	ii := 0
	ctr := 0
	for at != "ZZZ" {
		ctr++
		nextNodes := nodes[at]
		dir := instructions[ii]
		if dir == 'L' {
			at = nextNodes[0]
		} else {
			at = nextNodes[1]
		}
		ii++

		if ii > len(instructions)-1 {
			ii = 0
		}
	}

	fmt.Println("Part 1:", ctr)
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}

	return a
}

func lcm(ints ...int) int {
	if len(ints) < 2 {
		panic("cant lcm with less than 2 numbers")
	}

	res := ints[0] * ints[1] / gcd(ints[0], ints[1])

	for i := 2; i < len(ints); i++ {
		res = lcm(res, ints[i])
	}

	return res
}

func Part2(input string) {
	lines := strings.Split(input, "\n")

	nodes := map[string][]string{}

	at := []string{}
	for _, line := range lines[2:] {
		comps := strings.Fields(line)
		_at := comps[0]
		l := comps[2][1 : len(comps[2])-1]
		r := comps[3][:len(comps[3])-1]

		nodes[_at] = []string{l, r}
		if _at[2] == 'A' {
			at = append(at, _at)
		}
	}

	instructions := lines[0]

	zees := []int{}

	for _, _at := range at {
		ii := 0
		curr := _at
		ctr := 0

		for curr[2] != 'Z' {
			nextNodes := nodes[curr]
			dir := instructions[ii]
			if dir == 'L' {
				curr = nextNodes[0]
			} else {
				curr = nextNodes[1]
			}

			ii++
			ctr++

			if ii > len(instructions)-1 {
				ii = 0
			}
		}

		zees = append(zees, ctr)
	}

	res := lcm(zees...)
	fmt.Println("Part 2:", res)
}
