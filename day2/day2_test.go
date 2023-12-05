package day2

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

func BenchmarkCaminitiPart2(b *testing.B) {
	input := getFileContents()

	for i := 0; i < b.N; i++ {
		Part2(input)
	}
}

func BenchmarkKadatzPart2(b *testing.B) {
	input := getFileContents()

	for i := 0; i < b.N; i++ {
		Kadatz2(input)
	}
}

func BenchmarkOptKadatzPart2(b *testing.B) {
	input := getFileContents()

	for i := 0; i < b.N; i++ {
		OptKadatz2(input)
	}
}

func BenchmarkDavidPart2(b *testing.B) {
	input := getFileContents()

	for i := 0; i < b.N; i++ {
		DavidPart2(input)
	}
}
