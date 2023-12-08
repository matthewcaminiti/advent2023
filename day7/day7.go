package day7

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type hand struct {
	raw []rune
	val int
	bid int
}

func cardVal(r rune) int {
	switch r {
	case '2':
		return 2
	case '3':
		return 3
	case '4':
		return 4
	case '5':
		return 5
	case '6':
		return 6
	case '7':
		return 7
	case '8':
		return 8
	case '9':
		return 9
	case 'T':
		return 10
	case 'J':
		return 11
	case 'Q':
		return 12
	case 'K':
		return 13
	case 'A':
		return 14
	default:
		return -1
	}
}

func cardVal2(r rune) int {
	switch r {
	case '2':
		return 2
	case '3':
		return 3
	case '4':
		return 4
	case '5':
		return 5
	case '6':
		return 6
	case '7':
		return 7
	case '8':
		return 8
	case '9':
		return 9
	case 'T':
		return 10
	case 'J':
		return 1
	case 'Q':
		return 12
	case 'K':
		return 13
	case 'A':
		return 14
	default:
		return -1
	}
}

func Part1(input string) {
	lines := strings.Split(input, "\n")

	hands := make([]hand, len(lines))

	for i, line := range lines {
		comps := strings.Fields(line)

		bid, _ := strconv.Atoi(comps[1])

		chars := map[rune]int{}
		for _, c := range comps[0] {
			if _, exists := chars[c]; !exists {
				chars[c] = 0
			}

			chars[c]++
		}

		val := 0
		switch len(chars) {
		case 5:
			// high card
			val = 1
		case 4:
			// one pair
			val = 2
		case 3:
			// two pair, three of a kind
			val = 3
			for _, n := range chars {
				if n == 3 {
					val = 4
					break
				}
			}
		case 2:
			// full house, four of a kind
			val = 5
			for _, n := range chars {
				if n == 4 {
					val = 6
				}
			}
		case 1:
			// five of a kind
			val = 7
		}

		hands[i] = hand{
			raw: []rune(comps[0]),
			val: val,
			bid: bid,
		}
	}

	sort.Slice(hands, func(i, j int) bool {
		if hands[i].val != hands[j].val {
			return hands[i].val < hands[j].val
		}

		for c := 0; c < len(hands[c].raw); c++ {
			if hands[i].raw[c] != hands[j].raw[c] {
				return cardVal(hands[i].raw[c]) < cardVal(hands[j].raw[c])
			}
		}

		return true
	})

	tot := 0
	for i, hand := range hands {
		rank := i + 1
		tot += hand.bid * rank
	}

	fmt.Println("Part 1:", tot)
}

func Part2(input string) {
	lines := strings.Split(input, "\n")

	hands := make([]hand, len(lines))

	for i, line := range lines {
		comps := strings.Fields(line)

		bid, _ := strconv.Atoi(comps[1])

		chars := map[rune]int{}
		jCount := 0
		for _, c := range comps[0] {
			if c == 'J' {
				jCount++
				continue
			}

			chars[c]++
		}

		// +++best:
		// 7 five of a kind
		// 6 four of a kind
		// 5 full house (3/2)
		// 4 three of a kind
		// 3 two pair
		// 2 one pair
		// 1 high card
		val := 0
		switch len(chars) {
		case 5:
			val = 1
		case 4:
			// one pair
			val = 2
		case 3:
			// two pair, three of a kind
			val = 3
			if jCount >= 1 {
				val = 4
			} else {
				for _, n := range chars {
					if n == 3 {
						val = 4
						break
					}
				}
			}
		case 2:
			// full house, four of a kind
			val = 5
			if jCount == 3 {
				val = 6
			} else if jCount == 2 {
				val = 6
			} else if jCount == 1 {
				val = 5
				for _, n := range chars {
					if n == 3 {
						val = 6
						break
					}
				}
			} else {
				for _, n := range chars {
					if n == 4 {
						val = 6
						break
					}
				}
			}
		case 1:
			// five of a kind
			val = 7
		case 0:
			// five of a kind
			val = 7
		}

		hands[i] = hand{
			raw: []rune(comps[0]),
			val: val,
			bid: bid,
		}
	}

	sort.Slice(hands, func(i, j int) bool {
		if hands[i].val != hands[j].val {
			return hands[i].val < hands[j].val
		}

		for c := 0; c < len(hands[c].raw); c++ {
			if hands[i].raw[c] != hands[j].raw[c] {
				return cardVal2(hands[i].raw[c]) < cardVal2(hands[j].raw[c])
			}
		}

		return true
	})

	tot := 0
	for i, hand := range hands {
		rank := i + 1
		tot += hand.bid * rank
	}

	fmt.Println("Part 2:", tot)
}
