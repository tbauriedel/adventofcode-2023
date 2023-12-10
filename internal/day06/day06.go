package day06

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type race struct {
	time               int
	record             int
	possibilitiesToWin int
}

func Execute() {
	file, err := os.ReadFile("inputs/day06.txt")
	if err != nil {
		log.Fatal(err)
	}

	content := strings.Replace(string(file), " ", "", -1)

	races := evaluateRaces(content)

	for idx, round := range races {
		races[idx] = calcWinPossibilities(round)
	}

	sum1 := 0
	for _, round := range races {
		if sum1 == 0 {
			sum1 = round.possibilitiesToWin
			continue
		}
		sum1 *= round.possibilitiesToWin
	}

	fmt.Println("Result:", sum1)
}

func evaluateRaces(file string) (races []race) {
	reg := regexp.MustCompile(`(\d+)`)
	matches := reg.FindAllString(file, -1)

	for cnt := 0; cnt < len(matches)/2; cnt++ {
		time, _ := strconv.Atoi(matches[cnt])
		dur, _ := strconv.Atoi(matches[cnt+(len(matches)/2)])

		races = append(races, race{
			time:   time,
			record: dur,
		})
	}

	return
}

func calcWinPossibilities(r race) race {
	for i := 1; i <= r.time; i++ {
		dis := i * (r.time - i)

		if dis > r.record {
			r.possibilitiesToWin += 1
		}
	}

	return r
}
