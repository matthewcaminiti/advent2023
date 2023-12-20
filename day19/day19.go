package day19

import (
	"fmt"
	"strconv"
	"strings"
)

type part struct {
	x, m, a, s int
}

type stage func(part) (bool, string)

func Part1(input string) {
	groups := strings.Split(input, "\n\n")

	pipelines := map[string][]stage{}
	for _, line := range strings.Split(groups[0], "\n") {
		outer := strings.Split(line, "{")

		pipelineKey := outer[0]

		rawStages := strings.Split(outer[1][:len(outer[1])-1], ",")
		stages := []stage{}
		for _, rw := range rawStages {
			comps := strings.Split(rw, ":")
			if len(comps) == 1 {
				// just a key or A/R
				stages = append(stages, func(x part) (bool, string) {
					return true, comps[0]
				})
			} else {
				op := strings.Split(comps[0], ">")
				gt := true

				if len(op) == 1 {
					gt = false
					op = strings.Split(comps[0], "<")
				}

				control, _ := strconv.Atoi(op[1])

				stages = append(stages, func(x part) (bool, string) {
					v := 0
					switch op[0] {
					case "x":
						v = x.x
					case "m":
						v = x.m
					case "a":
						v = x.a
					case "s":
						v = x.s
					}

					if gt {
						return v > control, comps[1]
					}

					return v < control, comps[1]
				})
			}
		}

		pipelines[pipelineKey] = stages
	}

	parts := []part{}
	for _, line := range strings.Split(groups[1], "\n") {
		subComps := strings.Split(line[1:len(line)-1], ",")
		p := part{}
		for _, c := range subComps {
			zz := strings.Split(c, "=")
			v, _ := strconv.Atoi(zz[1])
			switch zz[0] {
			case "x":
				p.x = v
			case "m":
				p.m = v
			case "a":
				p.a = v
			case "s":
				p.s = v
			}
		}

		parts = append(parts, p)
	}

	tot := 0
	for _, p := range parts {
		at := "in"
		for at != "A" && at != "R" {
			stages := pipelines[at]

			for _, stage := range stages {
				good, next := stage(p)
				if good {
					at = next
					break
				}
			}
		}

		if at == "A" {
			tot += p.x + p.m + p.a + p.s
		}
	}

	fmt.Println("Part 1:", tot)
}

type pair struct {
	l, r int
}

type partRange struct {
	x pair
	m pair
	a pair
	s pair
}

// if len == 1, pair is valid if populated
// if len == 2, 0th is valid pair, 1st is invalid pair
func splitPair(x pair, v int, lt bool) []pair {
	if lt {
		if x.r < v {
			return []pair{x}
		}
		if x.l >= v {
			return []pair{{}}
		}

		return []pair{
			{x.l, v - 1},
			{v, x.r},
		}
	}

	// gt
	if x.r <= v {
		return []pair{{}}
	}

	if x.l > v {
		return []pair{x}
	}

	return []pair{
		{v + 1, x.r},
		{x.l, v},
	}
}

// given a range, apply a stage to it, determine if stage impacts the provided partRange
// if so, return the new part ranges, with their corresponding next arg
type rangeSplitter func(x partRange) (bool, []partRange, string)

func Part2(input string) {
	groups := strings.Split(input, "\n\n")

	// go through every pipeline and maintain set of part ranges that can yield A
	pipelines := map[string][]rangeSplitter{}
	for _, line := range strings.Split(groups[0], "\n") {
		outer := strings.Split(line[:len(line)-1], "{")

		pipelineKey := outer[0]

		rawStages := strings.Split(outer[1], ",")
		stages := []rangeSplitter{}
		for _, rw := range rawStages {
			comps := strings.Split(rw, ":")
			if len(comps) == 1 {
				// just a key or A/R
				stages = append(stages, func(pr partRange) (bool, []partRange, string) {
					return true, []partRange{pr}, comps[0]
				})
			} else {
				op := strings.Split(comps[0], ">")
				lt := false

				if len(op) == 1 {
					lt = true
					op = strings.Split(comps[0], "<")
				}

				control, _ := strconv.Atoi(op[1])
				letter := op[0]

				stages = append(stages, func(pr partRange) (bool, []partRange, string) {
					newPairs := []pair{}
					switch letter {
					case "x":
						newPairs = splitPair(pr.x, control, lt)
					case "m":
						newPairs = splitPair(pr.m, control, lt)
					case "a":
						newPairs = splitPair(pr.a, control, lt)
					case "s":
						newPairs = splitPair(pr.s, control, lt)
					}

					if len(newPairs) == 1 {
						return newPairs[0].l != 0 || newPairs[0].r != 0,
							[]partRange{pr},
							comps[1]
					}

					newPartRanges := []partRange{}

					for _, newPair := range newPairs {
						switch letter {
						case "x":
							newPartRanges = append(newPartRanges, partRange{
								x: pair{newPair.l, newPair.r},
								m: pr.m,
								a: pr.a,
								s: pr.s,
							})
						case "m":
							newPartRanges = append(newPartRanges, partRange{
								x: pr.x,
								m: pair{newPair.l, newPair.r},
								a: pr.a,
								s: pr.s,
							})
						case "a":
							newPartRanges = append(newPartRanges, partRange{
								x: pr.x,
								m: pr.m,
								a: pair{newPair.l, newPair.r},
								s: pr.s,
							})
						case "s":
							newPartRanges = append(newPartRanges, partRange{
								x: pr.x,
								m: pr.m,
								a: pr.a,
								s: pair{newPair.l, newPair.r},
							})
						}
					}

					return true, newPartRanges, comps[1]
				})
			}
		}

		pipelines[pipelineKey] = stages
	}

	partRanges := []partRange{
		{
			x: pair{1, 4000},
			m: pair{1, 4000},
			a: pair{1, 4000},
			s: pair{1, 4000},
		},
	}

	ats := []string{
		"in",
	}

	goodRanges := []partRange{}
	for len(partRanges) > 0 {
		// for i, pr := range partRanges {
		//     fmt.Printf("[%3s] x: (%4d, %4d) m: (%4d, %4d) a: (%4d, %4d) s: (%4d, %4d)\n", ats[i], pr.x.l, pr.x.r, pr.m.l, pr.m.r, pr.a.l, pr.a.r, pr.s.l, pr.s.r)
		// }
		// fmt.Println()
		newPartRanges := []partRange{}
		newAts := []string{}

		for i := 0; i < len(partRanges); i++ {
			at := ats[i]
			pr := partRanges[i]
			if at == "A" {
				goodRanges = append(goodRanges, pr)
				continue
			} else if at == "R" {
				continue
			}

			rangeSplitters := pipelines[at]
			for _, splitter := range rangeSplitters {
				ok, newRanges, next := splitter(pr)
				if !ok {
					continue
				}

				newPartRanges = append(newPartRanges, newRanges[0])
				newAts = append(newAts, next)

				if len(newRanges) == 1 {
					break
				}

				pr = newRanges[1]
			}
		}

		partRanges = newPartRanges
		ats = newAts
	}

	d := func(x pair) int {
		return x.r - x.l + 1
	}

	tot := 0
	for _, pr := range goodRanges {
		// fmt.Printf("x: (%4d, %4d) m: (%4d, %4d) a: (%4d, %4d) s: (%4d, %4d)\n", pr.x.l, pr.x.r, pr.m.l, pr.m.r, pr.a.l, pr.a.r, pr.s.l, pr.s.r)
		tot += d(pr.x) * d(pr.m) * d(pr.a) * d(pr.s)
	}

	fmt.Println("Part 2:", tot)
}
