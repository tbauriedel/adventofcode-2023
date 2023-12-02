package day01

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	replacements = map[string]string{
		"one":   "o1e",
		"two":   "t2w",
		"three": "t3e",
		"four":  "f4r",
		"five":  "f5e",
		"six":   "s6x",
		"seven": "s7n",
		"eight": "e8t",
		"nine":  "n9e",
	}
)

func Execute() {
	file, err := os.Open("inputs/day01.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sum1 := 0
	sum2 := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		firstD, _ := strconv.Atoi(getFirstDigitInString(line))
		lastD, _ := strconv.Atoi(getLastDigitInString(line))
		sum1 += (firstD * 10) + lastD

		line = replaceDigits(line)
		firstD, _ = strconv.Atoi(getFirstDigitInString(line))
		lastD, _ = strconv.Atoi(getLastDigitInString(line))
		sum2 += (firstD * 10) + lastD
	}

	fmt.Println("Result of puzzle 1:", sum1)
	fmt.Println("Result of puzzle 2:", sum2)
}

func getFirstDigitInString(line string) string {
	for _, char := range line {
		if char > 47 && char < 58 {
			return string(char)
		}
	}

	return ""
}

func getLastDigitInString(line string) string {
	for cnt := len(line) - 1; cnt >= 0; cnt-- {
		if line[cnt] > 47 && line[cnt] < 58 {
			return string(line[cnt])
		}
	}
	return ""
}

func replaceDigits(line string) string {
	for digitString, digit := range replacements {
		line = strings.Replace(line, digitString, digit, -1)
	}
	return line
}
