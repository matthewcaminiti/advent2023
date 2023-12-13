package day12

import (
	"advent/file"
	"testing"
)

func BenchmarkPart1(b *testing.B) {
	input := file.GetFileContents("./input.txt")

	for i := 0; i < b.N; i++ {
		Part1(input)
	}
}

func BenchmarkPart2(b *testing.B) {
	input := file.GetFileContents("./input.txt")

	for i := 0; i < b.N; i++ {
		Part2(input)
	}
}
