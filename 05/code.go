package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
	fmt.Println(ranges, ingredients)
}
