package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func readInput(file *os.File) ([][]int, []int) {
	scanner := bufio.NewScanner(file)
	var ranges [][]int
	var ingredients []int
	readingRanges := true

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			readingRanges = false
			continue
		}

		if readingRanges {
			var start, end int
			n, err := fmt.Sscanf(line, "%d-%d", &start, &end)
			if n != 2 || err != nil {
				log.Fatalf("Bad input: %s", line)
			}

			ranges = append(ranges, []int{start, end})
		} else {
			var ingredient int
			n, err := fmt.Sscanf(line, "%d", &ingredient)
			if n != 1 || err != nil {
				log.Fatalf("Bad input: %s", line)
			}

			ingredients = append(ingredients, ingredient)
		}
	}

	return ranges, ingredients
}

func part1(ranges [][]int, ingredients []int) int {
	var count int
	for _, ingredient := range ingredients {
		for _, r := range ranges {
			if ingredient < r[0] {
				break
			}

			if ingredient <= r[1] {
				count++
				break
			}
		}
	}

	return count
}

func part2(ranges [][]int) int {
	count := ranges[0][1] - ranges[0][0] + 1

	for i := 1; i < len(ranges); i++ {
		if ranges[i][1] <= ranges[i-1][1] {
			continue
		}

		if ranges[i][0] <= ranges[i-1][1] {
			count += ranges[i][1] - ranges[i-1][1]
		} else {
			count += ranges[i][1] - ranges[i][0] + 1
		}
	}

	return count
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("You need to specify a file!")
	}

	filePath := os.Args[1]
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open %s!\n", filePath)
	}

	ranges, ingredients := readInput(file)
	sort.Slice(ranges, func(i, j int) bool {
		if ranges[i][0] < ranges[j][0] {
			return true
		}

		if ranges[i][0] == ranges[j][0] {
			return ranges[i][1] < ranges[j][1]
		}

		return false
	})

	fmt.Println("Part1:", part1(ranges, ingredients))
	fmt.Println("Part2:", part2(ranges))
}
