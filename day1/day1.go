package day1

import (
	"fmt"
	"strconv"
	"strings"
)

var zero = int('0')
var nine = int('9')

func Part1(input string) {
	lines := strings.Split(input, "\n")
	sum := int64(0)
	for _, code := range lines {
		li, ri := 0, len(code)-1
		l, r := "", ""

		for li <= len(code)-1 {
			if zero <= int(code[li]) && int(code[li]) <= nine {
				l = string(code[li])
				break
			}
			li++
		}

		for ri >= 0 {
			if zero <= int(code[ri]) && int(code[ri]) <= nine {
				r = string(code[ri])
				break
			}

			ri--
		}

		x, err := strconv.ParseInt(l+r, 10, 64)
		if err != nil {
			fmt.Println(err)
		}
		sum += x
	}

	// fmt.Printf("Part 1 final sum: %d\n", sum)
}

func isDigit(x byte) bool {
	return zero <= int(x) && int(x) <= nine
}

func getDigit(x string) string {
	switch x {
	case "one":
		return "1"
	case "two":
		return "2"
	case "three":
		return "3"
	case "four":
		return "4"
	case "five":
		return "5"
	case "six":
		return "6"
	case "seven":
		return "7"
	case "eight":
		return "8"
	case "nine":
		return "9"
	}

	return ""
}

func canBeLDigit(x string) bool {
	digits := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	for _, digit := range digits {
		if len(x) > len(digit) {
			continue
		}

		if x == digit[0:len(x)] {
			return true
		}
	}

	return false
}

func canBeRDigit(x string) bool {
	digits := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	for _, digit := range digits {
		if len(x) > len(digit) {
			continue
		}

		if x == digit[len(digit)-len(x):] {
			return true
		}
	}

	return false
}

func Part2(input string) {
	lines := strings.Split(input, "\n")
	sum := int64(0)
	for _, code := range lines {
		li, ri := 0, len(code)-1
		l, r := "", ""

		lstr := ""
		for li <= len(code)-1 {
			digit := getDigit(lstr)
			if digit != "" {
				l = digit
				break
			}

			lstr = lstr + string(code[li])

			for !canBeLDigit(lstr) {
				lstr = lstr[1:]
			}

			if isDigit(code[li]) {
				l = string(code[li])
				break
			}

			li++
		}

		if digit := getDigit(lstr); digit != "" {
			l = digit
		}

		rstr := ""
		for ri >= 0 {
			digit := getDigit(rstr)
			if digit != "" {
				r = digit
				break
			}

			rstr = string(code[ri]) + rstr

			for !canBeRDigit(rstr) {
				rstr = rstr[:len(rstr)-1]
			}

			if isDigit(code[ri]) {
				r = string(code[ri])
				break
			}

			ri--
		}

		if digit := getDigit(rstr); digit != "" {
			r = digit
		}

		x, err := strconv.ParseInt(l+r, 10, 64)
		if err != nil {
			fmt.Println(err)
		}
		sum += x
	}

	// fmt.Printf("Part 2 final sum: %d\n", sum)
}
