package day05

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type seedData struct {
	seed        int
	soil        int
	fertilizer  int
	water       int
	light       int
	temperature int
	humidity    int
	location    int
}

type rangeObject struct {
	destStart   int
	sourceStart int
	length      int
}

var (
	seeds []seedData
)

func Execute() {
	file, err := os.ReadFile("inputs/day05.txt")
	if err != nil {
		log.Fatal(err)
	}

	parts := strings.Split(string(file), "\n\n")

	reg := regexp.MustCompile(`(\d+)`)
	resSeeds := reg.FindAllString(parts[0], -1)

	for _, match := range resSeeds {
		seedID, _ := strconv.Atoi(match)
		seed := seedData{seed: seedID}
		seed.getCorrespondence(parts[1], "soil")
		seed.getCorrespondence(parts[2], "fertilizer")
		seed.getCorrespondence(parts[3], "water")
		seed.getCorrespondence(parts[4], "light")
		seed.getCorrespondence(parts[5], "temperature")
		seed.getCorrespondence(parts[6], "humidity")
		seed.getCorrespondence(parts[7], "location")

		seeds = append(seeds, seed)
	}

	sum1 := 0
	for _, seed := range seeds {
		if sum1 == 0 {
			sum1 = seed.location
			continue
		}
		if seed.location < sum1 {
			sum1 = seed.location
		}
	}

	fmt.Println("Result of puzzle 1:", sum1)
}

func (s *seedData) getCorrespondence(content string, mode string) {
	var source *int
	var dest *int

	switch mode {
	case "soil":
		source = &s.seed
		dest = &s.soil
		break
	case "fertilizer":
		source = &s.soil
		dest = &s.fertilizer
		break
	case "water":
		source = &s.fertilizer
		dest = &s.water
		break
	case "light":
		source = &s.water
		dest = &s.light
		break
	case "temperature":
		source = &s.light
		dest = &s.temperature
		break
	case "humidity":
		source = &s.temperature
		dest = &s.humidity
		break
	case "location":
		source = &s.humidity
		dest = &s.location
		break
	}

	possibleRanges := strings.Split(content, "\n")
	matchFound := false

	for idx, part := range possibleRanges {
		if idx == 0 {
			continue
		}

		r := parseRanges(part)

		if !matchFound {
			if *source >= r.sourceStart && *source < r.sourceStart+r.length {
				for cnt := 0; cnt <= r.length; cnt++ {
					if *source == (r.sourceStart + cnt) {
						*dest = r.destStart + cnt
						matchFound = true
						break
					}
				}
			}
		}
		if *dest == 0 {
			*dest = *source
		}
	}
}

func parseRanges(line string) (r rangeObject) {
	reg := regexp.MustCompile(`(\d+)`)
	res := reg.FindAllString(line, -1)
	r.destStart, _ = strconv.Atoi(res[0])
	r.sourceStart, _ = strconv.Atoi(res[1])
	r.length, _ = strconv.Atoi(res[2])

	return
}
