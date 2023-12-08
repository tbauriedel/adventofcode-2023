package main

import (
	"flag"
	"github.com/tbauriedel/adventofcode-2023/internal/day01"
	"github.com/tbauriedel/adventofcode-2023/internal/day02"
	"github.com/tbauriedel/adventofcode-2023/internal/day03"
	"github.com/tbauriedel/adventofcode-2023/internal/day04"
	"github.com/tbauriedel/adventofcode-2023/internal/day05"
)

var (
	possibleDays = map[string]func(){
		"01": day01.Execute,
		"02": day02.Execute,
		"03": day03.Execute,
		"04": day04.Execute,
		"05": day05.Execute,
	}
)

func main() {
	day := flag.String("day", "", "Day to execute (format 'dayXX'")
	flag.Parse()

	// Call Execute() function of day
	possibleDays[*day]()
}
