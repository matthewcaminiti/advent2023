package day4

import (
	"bufio"
	"os"
	"strings"
	"testing"
)

func readFile() string {
	f, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)

	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return strings.Join(lines, "\n")
}

func BenchmarkCaminitiPart1(b *testing.B) {
	input := readFile()

	for i := 0; i < b.N; i++ {
		Part1(input)
	}
}

func BenchmarkCaminitiPart2(b *testing.B) {
	input := readFile()

	for i := 0; i < b.N; i++ {
		Part2(input)
	}
}
