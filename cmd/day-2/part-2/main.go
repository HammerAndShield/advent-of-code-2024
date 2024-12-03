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

	safe := 0
	for _, report := range reports {
		if isValid(report) {
			safe += 1
		}
	}

	fmt.Printf("The result is: %d\n", safe)
}

func isValid(report []int) bool {
	if isValidSeq(report) {
		return true
	}

	for i := range report {
		newReport := make([]int, 0, len(report)-1)
		newReport = append(newReport, report[:i]...)
		newReport = append(newReport, report[i+1:]...)
		if isValidSeq(newReport) {
			return true
		}
	}

	return false
}

func isValidSeq(report []int) bool {
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