package day_2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetInput() ([][]int, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}

	scanner := bufio.NewScanner(file)
	res := make([][]int, 0)
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			continue
		}

		fields := strings.Fields(line)
		if len(fields) <= 1 {
			return nil, fmt.Errorf("invalid line: %s", line)
		}

		nums := make([]int, len(fields))
		for i, field := range fields {
			num, err := strconv.Atoi(field)
			if err != nil {
				return nil, fmt.Errorf("failed to parse distance: %w", err)
			}

			nums[i] = num
		}

		res = append(res, nums)
	}

	return res, nil
}
