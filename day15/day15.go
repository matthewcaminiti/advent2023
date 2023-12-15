package day15

import (
	"fmt"
	"strconv"
	"strings"
)

func hash(s string) int {
	val := 0
	for _, c := range s {
		val = ((val + int(c)) * 17) % 256
	}
	return val
}

func Part1(input string) {
	runs := strings.Split(input, ",")

	tot := 0
	for _, run := range runs {
		tot += hash(run)
	}

	fmt.Println("Part 1:", tot)
}

type lens struct {
	name string
	val  int
}

func Part2(input string) {
	runs := strings.Split(input, ",")

	boxes := map[int][]lens{}
	for _, run := range runs {
		comps := strings.Split(run, "=")
		s := comps[0]
		v := -1
		if len(comps) == 2 {
			v, _ = strconv.Atoi(comps[1])
		} else {
			s = s[:len(s)-1]
		}

		k := hash(s)

		lenses, exists := boxes[k]
		if v == -1 {
			if len(lenses) == 0 {
				continue
			}

			tmp := []lens{}
			for _, l := range lenses {
				if l.name != s {
					tmp = append(tmp, l)
				}
			}
			boxes[k] = tmp
			continue
		}

		if !exists {
			boxes[k] = []lens{{s, v}}
			continue
		}

		found := false
		for i := range lenses {
			if lenses[i].name == s {
				found = true
				lenses[i].val = v
				break
			}
		}

		if !found {
			lenses = append(lenses, lens{s, v})
		}

		boxes[k] = lenses
	}

	tot := 0
	for k, v := range boxes {
		for i, l := range v {
			tot += (k + 1) * (i + 1) * l.val
		}
	}

	fmt.Println("Part 2:", tot)
}
