package day1

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func ReadInput() []int {
	file, err := os.Open("./inputs/day1.txt")
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	var nums []int

	for scanner.Scan() {
		line := scanner.Text()
		num, _ := strconv.Atoi(line)
		nums = append(nums, num)
	}

	return nums
}
