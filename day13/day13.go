package day13

import (
	"fmt"
	"strings"
)

func Part1(input string) {
	blocks := strings.Split(input, "\n\n")

	colsum := 0
	rowsum := 0
	for _, block := range blocks {
		lines := strings.Split(block, "\n")

		maxRowCount := 0
		rowI := -1
		for i := 0; i+1 < len(lines); i++ {
			rowCount := 0

			l := i
			r := i + 1
			fullReflection := false
			for l >= 0 && r < len(lines) {
				if lines[l] != lines[r] {
					break
				}

				rowCount++
				l--
				r++

				if l < 0 || r >= len(lines) {
					fullReflection = true
					break
				}
			}

			if fullReflection && rowCount > maxRowCount {
				maxRowCount = rowCount
				rowI = i
			}
		}

		cols := make([]string, len(lines[0]))
		for i := 0; i < len(lines[0]); i++ {
			for j := range lines {
				cols[i] += string(lines[j][i])
			}
		}

		maxColCount := 0
		colI := -1
		for i := 0; i+1 < len(cols); i++ {
			colCount := 0

			l := i
			r := i + 1
			fullReflection := false
			for l >= 0 && r < len(cols) {
				if cols[l] != cols[r] {
					break
				}

				colCount++
				l--
				r++

				if l < 0 || r >= len(cols) {
					fullReflection = true
					break
				}
			}

			if fullReflection && colCount > maxColCount {
				colI = i
				maxColCount = colCount
			}
		}

		if maxRowCount > maxColCount {
			rowsum += rowI + 1
		} else {
			colsum += colI + 1
		}
	}

	fmt.Println("Part 1:", 100*rowsum+colsum)
}

func Part2(input string) {
	blocks := strings.Split(input, "\n\n")

	colsum := 0
	rowsum := 0
	for _, block := range blocks {
		lines := strings.Split(block, "\n")

		maxRowCount := 0
		rowI := -1
		for i := 0; i+1 < len(lines); i++ {
			rowCount := 0

			l := i
			r := i + 1
			fullReflection := false
			smudgeUsed := false
			for l >= 0 && r < len(lines) {
				if lines[l] != lines[r] {
					if smudgeUsed {
						break
					}

					// try smudge
					ndiff := 0
					for j := 0; j < len(lines[l]); j++ {
						if lines[l][j] != lines[r][j] {
							ndiff++
							if ndiff > 1 {
								break
							}
						}
					}

					if ndiff == 1 {
						smudgeUsed = true
					} else {
						break
					}
				}

				rowCount++
				l--
				r++

				if l < 0 || r >= len(lines) {
					fullReflection = true
					break
				}
			}

			if smudgeUsed && fullReflection && rowCount > maxRowCount {
				maxRowCount = rowCount
				rowI = i
			}
		}

		cols := make([]string, len(lines[0]))
		for i := 0; i < len(lines[0]); i++ {
			for j := range lines {
				cols[i] += string(lines[j][i])
			}
		}

		maxColCount := 0
		colI := -1
		for i := 0; i+1 < len(cols); i++ {
			colCount := 0
			smudgeUsed := false

			l := i
			r := i + 1
			fullReflection := false
			for l >= 0 && r < len(cols) {
				if cols[l] != cols[r] {
					if smudgeUsed {
						break
					}

					ndiff := 0
					for j := 0; j < len(cols[l]); j++ {
						if cols[l][j] != cols[r][j] {
							ndiff++
							if ndiff > 1 {
								break
							}
						}
					}

					if ndiff == 1 {
						smudgeUsed = true
					} else {
						break
					}
				}

				colCount++
				l--
				r++

				if l < 0 || r >= len(cols) {
					fullReflection = true
					break
				}
			}

			if smudgeUsed && fullReflection && colCount > maxColCount {
				colI = i
				maxColCount = colCount
			}
		}

		if maxRowCount > maxColCount {
			rowsum += rowI + 1
		} else {
			colsum += colI + 1
		}
	}

	fmt.Println("Part 2:", 100*rowsum+colsum)
}
