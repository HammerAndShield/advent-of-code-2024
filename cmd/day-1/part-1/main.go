package main

import (
	"bufio"
	"fmt"
	"github.com/emirpasic/gods/v2/trees/binaryheap"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	ds, err := readDistances()
	if err != nil {
		log.Fatal(err)
	}

	res := 0
	for ds.RightD.Size() > 0 && ds.LeftD.Size() > 0 {
		res += ds.MinDistSum()
	}

	fmt.Printf("The result is: %d\n", res)
}

func readDistances() (*Distances, error) {
	file, err := os.Open("distances.txt")
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer func() {
		err = file.Close()
		if err != nil {
			panic(err)
		}
	}()

	scanner := bufio.NewScanner(file)
	res := Distances{
		LeftD:  binaryheap.New[int](),
		RightD: binaryheap.New[int](),
	}

	for scanner.Scan() {
		line := scanner.Text()

		fields := strings.Fields(line)
		if len(fields) != 2 {
			return nil, fmt.Errorf("invalid line: %s", line)
		}

		a, err := strconv.Atoi(fields[0])
		if err != nil {
			return nil, fmt.Errorf("failed to parse distance: %w", err)
		}
		b, err := strconv.Atoi(fields[1])
		if err != nil {
			return nil, fmt.Errorf("failed to parse distance: %w", err)
		}

		res.DPush(a, b)
	}

	return &res, nil
}
