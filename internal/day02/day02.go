package day02

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type Cubes struct {
	Reds   int
	Greens int
	Blues  int
}

func Execute() {
	file, err := os.Open("inputs/day02.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sum1 := 0
	sum2 := 0

	scanner := bufio.NewScanner(file)
	cnt := 1
	for scanner.Scan() {
		highestC := getHighestCubesOfRound(scanner.Text())
		if gameIsValid(highestC) {
			sum1 += cnt
		}

		sum2 += highestC.Blues * highestC.Reds * highestC.Greens
		cnt++
	}

	fmt.Println("Result of puzzle 1:", sum1)
	fmt.Println("Result of puzzle 2:", sum2)
}

func getHighestCubesOfRound(line string) Cubes {
	highestC := Cubes{Reds: 0, Greens: 0, Blues: 0}

	reg := regexp.MustCompile(`(\d+)\sblue`)
	matches := reg.FindAllStringSubmatch(line, -1)
	for _, match := range matches {
		if num, err := strconv.Atoi(match[1]); err == nil {
			if highestC.Blues < num {
				highestC.Blues = num
			}
		}
	}

	reg = regexp.MustCompile(`(\d+)\sred`)
	matches = reg.FindAllStringSubmatch(line, -1)
	for _, match := range matches {
		if num, err := strconv.Atoi(match[1]); err == nil {
			if highestC.Reds < num {
				highestC.Reds = num
			}
		}
	}

	reg = regexp.MustCompile(`(\d+)\sgreen`)
	matches = reg.FindAllStringSubmatch(line, -1)
	for _, match := range matches {
		if num, err := strconv.Atoi(match[1]); err == nil {
			if highestC.Greens < num {
				highestC.Greens = num
			}
		}
	}

	return highestC
}

func gameIsValid(highestC Cubes) bool {
	maxCubes := Cubes{
		Reds:   12,
		Greens: 13,
		Blues:  14,
	}

	if highestC.Blues > maxCubes.Blues {
		return false
	}

	if highestC.Reds > maxCubes.Reds {
		return false
	}

	if highestC.Greens > maxCubes.Greens {
		return false
	}

	return true
}
