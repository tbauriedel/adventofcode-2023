package day03

import (
	"bufio"
	"os"
)

func Execute() {
	file, err := os.Open("inputs/day03.txt")
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

	}
}
