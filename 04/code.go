package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readInput(file *os.File) []string {
	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		lines = append(lines, line)
	}

	return lines
}

type point struct {
	x int
	y int
}

func howManyNeighbors(lines []string, x, y int) int {
	var count int

	for row := y - 1; row <= y+1; row++ {
		if row < 0 || row >= len(lines) {
			continue
		}

		for col := x - 1; col <= x+1; col++ {
			if col < 0 || col >= len(lines[row]) {
				continue
			}

			if row == y && col == x {
				continue
			}

			if lines[row][col] == '@' {
				count++
			}
		}
	}

	return count
}

func part1(lines []string) int {
	var count int

	for y := range lines {
		for x := range lines[y] {
			if lines[y][x] != '@' {
				continue
			}

			neighbors := howManyNeighbors(lines, x, y)
			if neighbors < 4 {
				count++
			}
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

	lines := readInput(file)
	fmt.Println("Part1:", part1(lines))
}
