package day8

import (
	"advent/file"
	"testing"
)

func BenchmarkOk(b *testing.B) {
	input := file.GetFileContents("./input.txt")

	for i := 0; i < b.N; i++ {
		Part2(input)
	}
}
