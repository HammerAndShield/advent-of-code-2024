package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	l, err := readFile()
	if err != nil {
		log.Fatal(err)
	}

	res := 0
	for _, num := range l.leftNums {
		score := num * l.rightCounts[num]
		res += score
	}

	fmt.Printf("The result is: %d\n", res)
}

type locations struct {
	rightCounts map[int]int
	leftNums    []int
}

func readFile() (*locations, error) {
	f, err := os.Open("distances.txt")
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}

	scanner := bufio.NewScanner(f)
	res := locations{
		rightCounts: make(map[int]int),
		leftNums:    make([]int, 0),
	}
	for scanner.Scan() {
		line := scanner.Text()

		fields := strings.Fields(line)
		if len(fields) != 2 {
			return nil, fmt.Errorf("invalid line: %s", line)
		}

		l, err := strconv.Atoi(fields[0])
		if err != nil {
			return nil, fmt.Errorf("failed to parse distance: %w", err)
		}

		r, err := strconv.Atoi(fields[1])
		if err != nil {
			return nil, fmt.Errorf("failed to parse distance: %w", err)
		}

		res.rightCounts[l]++
		res.leftNums = append(res.leftNums, r)
	}

	return &res, nil
}
