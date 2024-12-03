package main

import (
	day_2 "advent-of-code-2024/cmd/day-2"
	"fmt"
	"log"
	"math"
)

func main() {
	reports, err := day_2.GetInput()
	if err != nil {
		log.Fatal(err)
	}

	res := 0
	for _, report := range reports {
		if isValid(report) {
			res += 1
		}
	}

	fmt.Printf("The result is: %d\n", res)
}

func isValid(report []int) bool {
	increasing, decreasing := true, true

	for i := 1; i < len(report); i++ {
		if report[i] <= report[i-1] {
			increasing = false
		}
		if report[i] >= report[i-1] {
			decreasing = false
		}
		difference := math.Abs(float64(report[i] - report[i-1]))
		if difference > 3 {
			return false
		}
	}

	return increasing || decreasing
}
