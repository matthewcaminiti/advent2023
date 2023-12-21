package day20

import (
	"fmt"
	"strings"
)

const (
	BROADCAST = 1
	FLIP      = 2
	CONJ      = 3

	LOW  = 1
	HIGH = 2
)

type flipFlop struct {
	id       string
	on       bool
	children []string
}

func (m *flipFlop) recv(from string, pulse int) *signal {
	if pulse == HIGH {
		return nil
	}

	if m.on {
		m.on = false
		return &signal{
			from:  m.id,
			pulse: LOW,
			to:    m.children,
		}
	}

	m.on = true
	return &signal{
		from:  m.id,
		pulse: HIGH,
		to:    m.children,
	}
}

var _ module = &flipFlop{}

type conjunction struct {
	id       string
	memory   map[string]int
	children []string
}

func (m *conjunction) recv(from string, pulse int) *signal {
	m.memory[from] = pulse

	for _, p := range m.memory {
		if p == LOW {
			return &signal{
				from:  m.id,
				pulse: HIGH,
				to:    m.children,
			}
		}
	}

	return &signal{
		from:  m.id,
		pulse: LOW,
		to:    m.children,
	}
}

type broadcast struct {
	id       string
	children []string
}

func (m broadcast) recv(from string, pulse int) *signal {
	return &signal{
		from:  "broadcaster",
		pulse: pulse,
		to:    m.children,
	}
}

var _ module = &conjunction{}

type signal struct {
	from  string
	pulse int
	to    []string
}

type module interface {
	recv(from string, pulse int) *signal
}

func Part1(input string) {
	modules := map[string]module{}

	rawConjunctions := map[string][]string{}

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		comps := strings.Split(line, "->")
		in, outs := strings.TrimSpace(comps[0]), strings.Split(strings.TrimSpace(comps[1]), ", ")

		if in[0] == '%' {
			modules[in[1:]] = &flipFlop{
				id:       in[1:],
				on:       false,
				children: outs,
			}
		} else if in[0] == '&' {
			rawConjunctions[in[1:]] = outs
		} else {
			modules[in] = broadcast{
				id:       in,
				children: outs,
			}
		}
	}

	for id, children := range rawConjunctions {
		memory := map[string]int{}

		for _, line := range lines {
			comps := strings.Split(line, "->")
			in, outs := strings.TrimSpace(comps[0]), strings.Split(strings.TrimSpace(comps[1]), ", ")

			if in[0] == '%' || in[0] == '&' {
				in = in[1:]
			}

			for _, out := range outs {
				if out == id {
					memory[in] = LOW
					break
				}
			}
		}

		modules[id] = &conjunction{
			id:       id,
			memory:   memory,
			children: children,
		}
	}

	nLow, nHigh := 0, 0
	for i := 0; i < 1000; i++ {
		signals := []signal{
			{
				from:  "button",
				pulse: LOW,
				to:    []string{"broadcaster"},
			},
		}

		for len(signals) > 0 {
			nextSignals := []signal{}

			for _, s := range signals {
				if s.pulse == LOW {
					nLow += len(s.to)
				} else {
					nHigh += len(s.to)
				}

				for _, dest := range s.to {
					mod, ok := modules[dest]
					if !ok {
						continue
					}

					newSignal := mod.recv(s.from, s.pulse)
					if newSignal == nil {
						continue
					}

					nextSignals = append(nextSignals, *newSignal)
				}
			}

			signals = nextSignals
		}
	}

	fmt.Println("Part 1:", nLow*nHigh)
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
	modules := map[string]module{}

	rawConjunctions := map[string][]string{}

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		comps := strings.Split(line, "->")
		in, outs := strings.TrimSpace(comps[0]), strings.Split(strings.TrimSpace(comps[1]), ", ")

		if in[0] == '%' {
			modules[in[1:]] = &flipFlop{
				id:       in[1:],
				on:       false,
				children: outs,
			}
		} else if in[0] == '&' {
			rawConjunctions[in[1:]] = outs
		} else {
			modules[in] = broadcast{
				id:       in,
				children: outs,
			}
		}
	}

	for id, children := range rawConjunctions {
		memory := map[string]int{}

		for _, line := range lines {
			comps := strings.Split(line, "->")
			in, outs := strings.TrimSpace(comps[0]), strings.Split(strings.TrimSpace(comps[1]), ", ")

			if in[0] == '%' || in[0] == '&' {
				in = in[1:]
			}

			for _, out := range outs {
				if out == id {
					memory[in] = LOW
					break
				}
			}
		}

		modules[id] = &conjunction{
			id:       id,
			memory:   memory,
			children: children,
		}
	}

	nPresses := 0

	parents := map[string][]int{}
	for {
		nPresses++
		signals := []signal{
			{
				from:  "button",
				pulse: LOW,
				to:    []string{"broadcaster"},
			},
		}

		for len(signals) > 0 {
			nextSignals := []signal{}

			for _, s := range signals {
				for _, dest := range s.to {
					if dest == "qb" {
						if s.from == "kv" || s.from == "jg" || s.from == "rz" || s.from == "mr" {
							if s.pulse == HIGH {
								if _, exists := parents[s.from]; !exists {
									parents[s.from] = []int{}
								}

								parents[s.from] = append(parents[s.from], nPresses)
								// fmt.Printf("%s -[%d]-> %s\n", s.from, s.pulse, s.to)
							}
						}
					}

					mod, ok := modules[dest]
					if !ok {
						continue
					}

					newSignal := mod.recv(s.from, s.pulse)
					if newSignal == nil {
						continue
					}

					nextSignals = append(nextSignals, *newSignal)
				}
			}

			signals = nextSignals
		}

		if len(parents) == 4 {
			jah := []int{}
			for _, v := range parents {
				jah = append(jah, v[0])
			}

			fmt.Println("Part 2:", lcm(jah...))
			return
		}
	}
}
