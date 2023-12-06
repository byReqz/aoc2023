package main

import (
	"aoc2023/aoc"
	"fmt"
	"strconv"
	"strings"
)

var sample = `Time:      7  15   30
Distance:  9  40  200`

var sample2 = `Time:      71530
Distance:  940200`

func doRace(time, buttonTime int) (distance int) {
	raceTime := time - buttonTime

	return raceTime * buttonTime
}

func first(input string) {
	split := strings.Split(input, "\n")

	var times []int
	for _, t := range strings.Split(split[0], " ") {
		if n, err := strconv.Atoi(t); err == nil {
			times = append(times, n)
		}
	}

	var distances []int
	for _, d := range strings.Split(split[1], " ") {
		if n, err := strconv.Atoi(d); err == nil {
			distances = append(distances, n)
		}
	}

	var margins []int
	for n, r := range times {
		margin := 0
		for i := 0; i <= r; i++ {
			d := doRace(r, i)
			if d > distances[n] {
				margin++
			}
		}
		margins = append(margins, margin)
	}

	sum := margins[0]
	for _, m := range margins[1:] {
		sum = sum * m
	}
	fmt.Println(sum)
}

func second(input string) {
	input = strings.ReplaceAll(input, " ", "")
	input = strings.ReplaceAll(input, "Time:", "")
	input = strings.ReplaceAll(input, "Distance:", "")
	first(input)
}

func main() {
	//input := sample
	input := aoc.InputForDay("06")

	first(input)
	second(input)
}
