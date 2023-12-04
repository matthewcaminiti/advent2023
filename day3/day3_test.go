package day3

import (
    "os"
    "bufio"
    "strings"
    "testing"
)

func readfile() string {
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
    input := readfile()

    for i := 0; i < b.N; i++ {
        Part1(input)
    }
}

func BenchmarkOptCaminitiPart1(b *testing.B) {
    input := readfile()

    for i := 0; i < b.N; i++ {
        OptPart1(input)
    }
}

func BenchmarkKadatzPart1(b *testing.B) {
    input := readfile()

    for i := 0; i < b.N; i++ {
        KadatzPart1(input)
    }
}

func BenchmarkOptKadatzPart1(b *testing.B) {
    input := readfile()

    for i := 0; i < b.N; i++ {
        OptKadatzPart1(input)
    }
}
