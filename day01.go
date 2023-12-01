package main

import (
	"aoc2023/aoc"
	"fmt"
	"strconv"
	"strings"
)

var sample1 = `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`

var sample2 = `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`

func first(input string) {
	var values []string
	for _, l := range strings.Split(input, "\n") {
		var (
			lf     string
			lfdone bool
			ll     string
		)
		for _, b := range strings.Split(l, "") {
			if _, err := strconv.Atoi(b); err == nil {
				if !lfdone {
					lf = b
					lfdone = true
				}
				ll = b
			}
		}
		fmt.Println("line first:", lf, "line last:", ll, "result:", lf+ll)
		values = append(values, lf+ll)
	}

	var sum int
	for _, l := range values {
		i, _ := strconv.Atoi(l)
		sum += i
	}

	fmt.Println("sum of calibration values:", sum)
}

var nummap = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func isNum(n string) string {
	if _, err := strconv.Atoi(n); err == nil {
		return n
	}

	return nummap[n]
}

func repLine(input *string) {
	/*for str, i := range nummap {
		*input = strings.ReplaceAll(*input, str, i)
	}*/
	*input = strings.ReplaceAll(*input, "one", "1")
	*input = strings.ReplaceAll(*input, "two", "2")
	*input = strings.ReplaceAll(*input, "three", "3")
	*input = strings.ReplaceAll(*input, "four", "4")
	*input = strings.ReplaceAll(*input, "five", "5")
	*input = strings.ReplaceAll(*input, "six", "6")
	*input = strings.ReplaceAll(*input, "seven", "7")
	*input = strings.ReplaceAll(*input, "eight", "8")
	*input = strings.ReplaceAll(*input, "nine", "9")

}

func numFirstFromLine(input string) string {
	for i := 0; i <= len(input); i++ {
		firstchar := strings.Split(input, "")[i]
		if _, err := strconv.Atoi(firstchar); err == nil {
			return firstchar
		}

		trimline := strings.Join(strings.Split(input, "")[i:], "")
		for str, num := range nummap {
			if strings.HasPrefix(trimline, str) {
				return num
			}
		}
	}

	return ""
}

func numLastFromLine(input string) string {
	for i := 0; i < len(input); i++ {
		lastchar := strings.Split(input, "")[len(input)-1-i]
		if _, err := strconv.Atoi(lastchar); err == nil {
			return lastchar
		}

		start := (len(input) - 1) - i
		trimline := strings.Join(strings.Split(input, "")[start:], "")
		for str, num := range nummap {
			if strings.HasPrefix(trimline, str) {
				return num
			}
		}
	}

	return ""
}

func second(input string) {
	var values []string
	for i, l := range strings.Split(input, "\n") {

		lf := numFirstFromLine(l)
		ll := numLastFromLine(l)
		fmt.Println("line", i, "first:", lf, "line last:", ll, "result:", lf+ll)
		values = append(values, lf+ll)
	}

	var sum int
	for _, l := range values {
		i, _ := strconv.Atoi(l)
		sum += i
	}

	fmt.Println("sum of calibration values:", sum)
}

func main() {
	input := aoc.InputForDay("01")
	//input := sample2

	first(input)
	second(input)
}
