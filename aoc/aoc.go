package aoc

import (
	"log"
	"os"
	"strings"
)

func InputForDay(day string) string {
	f, err := os.ReadFile("./day" + day + ".input")
	if err != nil {
		log.Fatal("couldnt read input file:", err)
	}
	return strings.TrimSpace(string(f))
}
