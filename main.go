package main

import (
	"flag"
	"github.com/tbauriedel/adventofcode-2023/internal/day01"
	"github.com/tbauriedel/adventofcode-2023/internal/day02"
)

var (
	possibleDays = map[string]func(){
		"01": day01.Execute,
		"02": day02.Execute,
	}
)

func main() {
	day := flag.String("day", "", "Day to execute (format 'dayXX'")
	flag.Parse()

	possibleDays[*day]()
}
