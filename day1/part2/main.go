package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var input = os.Args[1]

	data, err := os.ReadFile(input)
	if err != nil {
		log.Fatal(err)
	}

	var left []int
	var right = make(map[int]int)
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		nums := strings.Fields(line)
		l, err := strconv.Atoi(nums[0])
		if err != nil {
			log.Fatal(err)
		}
		r, err := strconv.Atoi(nums[1])
		if err != nil {
			log.Fatal(err)
		}
		left = append(left, l)
		right[r] += 1
	}

	var score int
	for _, i := range left {
		score += i * right[i]
	}
	fmt.Println("Score:", score)
}
