package day3

import (
	"bufio"
	"log"
	"os"
)

func ReadInput() []string {
	file, err := os.Open("./inputs/day3.txt")
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	return lines
}
