package day1

import (
	"bufio"
	"os"
	"testing"
)

func BenchmarkDay1Caminit(b *testing.B) {
	f, err := os.Open("./day1.txt")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	for i := 0; i < b.N; i++ {
		Part2(lines)
	}
}

func BenchmarkDay1Kadatz(b *testing.B) {
	f, err := os.Open("./day1.txt")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	for i := 0; i < b.N; i++ {
		KadatzPart2(lines)
	}
}
