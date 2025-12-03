package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readInput(file *os.File) [][]int {
	scanner := bufio.NewScanner(file)
	var batteries [][]int

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		batteriesRow := make([]int, len(line))
		for i, char := range line {
			batteriesRow[i] = int(char - '0')
		}

		batteries = append(batteries, batteriesRow)
	}

	return batteries
}

func findMaxIndex(slice []int, start, end int) int {
	maxIndex := start
	for i := start; i < end; i++ {
		if slice[i] > slice[maxIndex] {
			maxIndex = i
		}
	}

	return maxIndex
}

func part1(batteries [][]int) int {
	var sum int

	for row := range batteries {
		maxLeft := findMaxIndex(batteries[row], 0, len(batteries[row])-1)
		maxRight := findMaxIndex(batteries[row], maxLeft+1, len(batteries[row]))

		sum += batteries[row][maxLeft]*10 + batteries[row][maxRight]
	}

	return sum
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

	batteries := readInput(file)
	fmt.Println("Part1:", part1(batteries))
}
