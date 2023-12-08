package day04

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

type Scratchcard struct {
	id       int
	actual   []string
	winnings []string
	points   int
	hits     int
	amount   int
}

var cards []Scratchcard

func Execute() {
	file, err := os.Open("inputs/day04.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sum1 := 0

	scanner := bufio.NewScanner(file)
	index := 1
	for scanner.Scan() {
		winnings, actual := evaScratchboard(scanner.Text())
		card := Scratchcard{
			id:       index,
			actual:   actual,
			winnings: winnings,
			points:   ActualInWinning(actual, winnings),
			hits:     hitsInScratchcard(actual, winnings),
			amount:   1,
		}
		cards = append(cards, card)

		sum1 += card.points
		index += 1
	}

	sum2 := 0
	for _, card := range cards {
		if card.hits > 0 {
			for a := 1; a <= card.amount; a++ {
				for cnt := 1; cnt <= card.hits; cnt++ {
					cards[card.id+cnt-1].amount += 1
				}
			}
		}

		sum2 += card.amount
	}

	fmt.Println("Result of puzzle 1:", sum1)
	fmt.Println("Result of puzzle 2:", sum2)
}

func evaScratchboard(line string) ([]string, []string) {
	var winnings, actuals []string
	reg := regexp.MustCompile(`Card\s+\d+:\s`)
	res := reg.FindAllString(line, -1)
	line = strings.Replace(line, res[0], "", 1)

	split := strings.Split(line, "|")
	reg = regexp.MustCompile(`(\d+)`)
	w := reg.FindAllString(split[0], -1)
	a := reg.FindAllString(split[1], -1)

	for _, match := range w {
		winnings = append(winnings, match)
	}
	for _, match := range a {
		actuals = append(actuals, match)
	}

	return winnings, actuals
}

func ActualInWinning(actuals, winning []string) int {
	matches := 0
	for _, a := range actuals {
		for _, w := range winning {
			if a == w {
				if matches == 0 {
					matches = 1
				} else {
					matches *= 2
				}
			}
		}
	}

	return matches
}

func hitsInScratchcard(actuals, winning []string) int {
	hits := 0
	for _, a := range actuals {
		for _, w := range winning {
			if a == w {
				hits += 1
			}
		}
	}
	return hits
}
