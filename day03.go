package main

import (
	"fmt"
	"strconv"
	"strings"
)

var sample = `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

var bettersample = `11....11
..$..$..
11....11`

func isSymbol(char string) bool {
	if char == "." {
		return false
	} else if _, err := strconv.Atoi(char); err == nil {
		return false
	}
	return true
}

func checkSurroundings(grid [][]string, row, char int) bool {
	// left
	if row != 0 && char != 0 && isSymbol(grid[row-1][char-1]) {
		return true
	} else if char != 0 && isSymbol(grid[row][char-1]) {
		return true
	} else if row != (len(grid)-1) && char != 0 && isSymbol(grid[row+1][char-1]) {
		return true
	}

	// right
	if row != 0 && char != (len(grid[row])-1) && isSymbol(grid[row-1][char+1]) {
		return true
	} else if char != (len(grid[row])-1) && isSymbol(grid[row][char+1]) {
		return true
	} else if row != (len(grid)-1) && char != (len(grid[row])-1) && isSymbol(grid[row+1][char+1]) {
		return true
	}

	// top
	if row != 0 && isSymbol(grid[row-1][char]) {
		return true
	}

	// bottom
	if row != (len(grid)-1) && isSymbol(grid[row+1][char]) {
		return true
	}

	return false
}

func first(input string) {
	split := strings.Split(input, "\n")

	scheme := make([][]string, len(split))
	for i, l := range split {
		scheme[i] = strings.Split(l, "")
	}

	var nums []string
	var nonnums []string
	for irow, row := range scheme {
		var tempword string
		var hasSurrounding bool
		for ichar, char := range row {
			if _, err := strconv.Atoi(char); err == nil {
				tempword += char
				if checkSurroundings(scheme, irow, ichar) {
					hasSurrounding = true
				}
			} else if tempword != "" {
				if hasSurrounding {
					nums = append(nums, tempword)
				} else {
					nonnums = append(nonnums, tempword)
				}
				tempword = ""
				hasSurrounding = false
			}

			if ichar == (len(row)-1) && tempword != "" {
				if hasSurrounding {
					nums = append(nums, tempword)
				} else {
					nonnums = append(nonnums, tempword)
				}
				tempword = ""
				hasSurrounding = false
			}
		}
	}

	sum := 0
	for _, n := range nums {
		num, _ := strconv.Atoi(n)
		sum += num
	}

	fmt.Println("valid part numbers:", nums)
	fmt.Println("invalid part numbers:", nonnums)
	fmt.Println("sum of valid part numbers:", sum)

}

/*func first2(input string) {
	split := strings.Split(input, "\n")

	scheme := make([][]string, len(split))
	for i, l := range split {
		scheme[i] = strings.Split(l, ".")
	}

	fmt.Println(len(scheme[0]), len(scheme[1]))

	var nums []string
	var nonnums []string
	for irow, row := range scheme {
		var tempword string
		var hasSurrounding bool
		for ichar, char := range row {
			if _, err := strconv.Atoi(char); err == nil {
				tempword += char
				if checkSurroundings(scheme, irow, ichar) {
					hasSurrounding = true
				}
			} else if tempword != "" {
				if hasSurrounding {
					nums = append(nums, tempword)
				} else {
					nonnums = append(nonnums, tempword)
				}
				tempword = ""
				hasSurrounding = false
			}
		}
	}

	sum := 0
	for _, n := range nums {
		num, _ := strconv.Atoi(n)
		sum += num
	}

	fmt.Println("valid part numbers:", nums)
	fmt.Println("invalid part numbers:", nonnums)
	fmt.Println("sum of valid part numbers:", sum)

}*/

func checkGearSurroundings(grid [][]string, row, char int) []string {

	surr := make([]string, 4)

	// top
	if row != 0 && char > 2 {
		surr[0] += grid[row-1][char-3]
	}
	if row != 0 && char > 1 {
		surr[0] += grid[row-1][char-2]
	}
	if row != 0 && char != 0 {
		surr[0] += grid[row-1][char-1]
	}
	if row != 0 {
		surr[0] += grid[row-1][char]
	}
	if row != 0 && char < (len(grid[row])-1) {
		surr[0] += grid[row-1][char+1]
	}
	if row != 0 && char < (len(grid[row])-2) {
		surr[0] += grid[row-1][char+2]
	}
	if row != 0 && char < (len(grid[row])-3) {
		surr[0] += grid[row-1][char+3]
	}

	// bottom
	if row != (len(grid)-1) && char > 0 && (_, err := strconv.Atoi(grid[row+1][char-2]); err == nil) {
		if row != (len(grid)-1) && char > 2 {
			surr[1] += grid[row+1][char-3]
		}
		if row != (len(grid)-1) && char > 1 {
			surr[1] += grid[row+1][char-2]
		}
	}
	if row != (len(grid)-1) && char > 0 {
		surr[1] += grid[row+1][char-1]
	}
	if row != (len(grid) - 1) {
		surr[1] += grid[row+1][char]
	}
	if row != (len(grid)-1) && char < (len(grid[row])-1) {
		surr[1] += grid[row+1][char+1]
	}
	if row != (len(grid)-1) && char < (len(grid[row])-2) {
		surr[1] += grid[row+1][char+2]
	}
	if row != (len(grid)-1) && char < (len(grid[row])-3) {
		surr[1] += grid[row+1][char+3]
	}

	// left
	if char > 2 {
		surr[2] += grid[row][char-3]
	}
	if char > 1 {
		surr[2] += grid[row][char-2]
	}
	if char > 0 {
		surr[2] += grid[row][char-1]
	}

	// right
	if char != (len(grid[row]) - 1) {
		surr[3] += grid[row][char+1]
	}
	if char < (len(grid[row]) - 2) {
		surr[3] += grid[row][char+2]
	}
	if char < (len(grid[row]) - 3) {
		surr[3] += grid[row][char+3]
	}

	return surr
}

func second(input string) {
	split := strings.Split(input, "\n")

	scheme := make([][]string, len(split))
	for i, l := range split {
		scheme[i] = strings.Split(l, "")
	}

	var nums []string
	var nonnums []string
	for irow, row := range scheme {
		var tempword string
		var hasSurrounding bool
		for ichar, char := range row {
			if char == "*" {
				surr := checkGearSurroundings(scheme, irow, ichar)
				fmt.Println(surr)
				if false {
					hasSurrounding = true
				}
			} else if tempword != "" {
				if hasSurrounding {
					nums = append(nums, tempword)
				} else {
					nonnums = append(nonnums, tempword)
				}
				tempword = ""
				hasSurrounding = false
			}

			if ichar == (len(row)-1) && tempword != "" {
				if hasSurrounding {
					nums = append(nums, tempword)
				} else {
					nonnums = append(nonnums, tempword)
				}
				tempword = ""
				hasSurrounding = false
			}
		}
	}

	sum := 0
	for _, n := range nums {
		num, _ := strconv.Atoi(n)
		sum += num
	}

	fmt.Println("valid part numbers:", nums)
	fmt.Println("invalid part numbers:", nonnums)
	fmt.Println("sum of valid part numbers:", sum)
}

func main() {
	input := sample
	//input := aoc.InputForDay("03")
	//input := bettersample

	//first(input)
	second(input)
}
