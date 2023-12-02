package main

import (
	"aoc2023/aoc"
	"fmt"
	"strconv"
	"strings"
)

var sample1 = `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`

var Rules = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

type Game struct {
	ID     int
	Rounds []Round
}

type Round struct {
	Sets []Draw
}

type Draw struct {
	Amount int
	Color  string
}

func getDraw(input string) Draw {
	var draw Draw

	split := strings.Split(input, " ")
	draw.Amount, _ = strconv.Atoi(split[0])
	draw.Color = split[1]

	return draw
}

func getRound(input string) Round {
	var round Round

	for _, s := range strings.Split(input, ", ") {
		round.Sets = append(round.Sets, getDraw(s))
	}

	return round
}

func newGame(input string) Game {
	var game Game

	for _, s := range strings.Split(input, "; ") {
		game.Rounds = append(game.Rounds, getRound(s))
	}

	return game
}

func isValid(game Game, rules map[string]int) bool {
	for _, r := range game.Rounds {
		for _, s := range r.Sets {
			if s.Amount > rules[s.Color] {
				return false
			}
		}
	}
	return true
}

func first(input string) {
	games := []Game{}

	for i, l := range strings.Split(input, "\n") {
		g := newGame(strings.Split(l, ": ")[1])
		g.ID = i + 1
		games = append(games, g)
	}

	sum := 0
	for _, g := range games {
		if isValid(g, Rules) {
			//fmt.Println(g.ID, "is valid")
			sum += g.ID
		}
	}
	fmt.Println("sum of valid game ids:", sum)
}

func minCubes(g Game) (red, blue, green int) {
	for _, r := range g.Rounds {
		prev := map[string]*int{
			"red":   &red,
			"blue":  &blue,
			"green": &green,
		}
		for _, s := range r.Sets {
			if s.Amount > *prev[s.Color] {
				*prev[s.Color] = s.Amount
			}
		}
	}
	return
}

func second(input string) {
	games := []Game{}

	for i, l := range strings.Split(input, "\n") {
		g := newGame(strings.Split(l, ": ")[1])
		g.ID = i + 1
		games = append(games, g)
	}

	sum := 0
	for _, g := range games {
		red, blue, green := minCubes(g)
		power := red * blue * green
		sum += power
		fmt.Printf("%d: red: %d blue: %d green: %d power: %d\n", g.ID, red, blue, green, power)
	}
	fmt.Println("sum of power of minimum cubes:", sum)
}

func main() {
	//input := sample1
	input := aoc.InputForDay("02")

	//first(input)
	second(input)
}
