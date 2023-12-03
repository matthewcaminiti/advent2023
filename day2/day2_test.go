package day2

import (
	"bufio"
	"os"
	"testing"
)

func BenchmarkCaminitiPart2(b *testing.B) {
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

	for i := 0; i < b.N; i++ {
		Part2(lines)
	}
}

func BenchmarkKadatzPart2(b *testing.B) {
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

	for i := 0; i < b.N; i++ {
		Kadatz2(lines)
	}
}
