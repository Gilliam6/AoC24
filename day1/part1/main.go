package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type pair struct {
	Value int
	Index int
}

func main() {
	var input = os.Args[1]

	data, err := os.ReadFile(input)
	if err != nil {
		log.Fatal(err)
	}

	var left, right []pair
	lines := strings.Split(string(data), "\n")
	for i, line := range lines {
		nums := strings.Fields(line)
		l, err := strconv.Atoi(nums[0])
		if err != nil {
			log.Fatal(err)
		}
		r, err := strconv.Atoi(nums[1])
		if err != nil {
			log.Fatal(err)
		}
		left = append(left, pair{Value: l, Index: i})
		right = append(right, pair{Value: r, Index: i})
	}

	sort.Slice(left, func(i, j int) bool {
		return left[i].Value < left[j].Value
	})
	sort.Slice(right, func(i, j int) bool {
		return right[i].Value < right[j].Value
	})

	var distances int
	for i := range left {
		dif := int(math.Abs(float64(left[i].Value - right[i].Value)))
		distances += dif
	}
	fmt.Println("Distances:", distances)
}
