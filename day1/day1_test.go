package day1

import (
	"bufio"
	"os"
	"strings"
	"testing"
)

func getFileContents() string {
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

func BenchmarkDay1Caminit(b *testing.B) {
	input := getFileContents()

	for i := 0; i < b.N; i++ {
		Part2(input)
	}
}

func BenchmarkDay1Kadatz(b *testing.B) {
	input := getFileContents()

	for i := 0; i < b.N; i++ {
		KadatzPart2(input)
	}
}

func BenchmarkDay1David(b *testing.B) {
	input := getFileContents()

	for i := 0; i < b.N; i++ {
		DavidPart2(input)
	}
}
