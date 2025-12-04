package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readInput(file *os.File) [][]byte {
	scanner := bufio.NewScanner(file)
	var lines [][]byte

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		items := make([]byte, len(line))
		for i := range line {
			items[i] = line[i]
		}

		lines = append(lines, items)
	}

	return lines
}

type point struct {
	x int
	y int
}

func howManyNeighbors(lines [][]byte, x, y int) int {
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

func removeRolls(lines [][]byte) int {
	var count int
	var toRemove []point

	for y := range lines {
		for x := range lines[y] {
			if lines[y][x] != '@' {
				continue
			}

			neighbors := howManyNeighbors(lines, x, y)
			if neighbors < 4 {
				toRemove = append(toRemove, point{x: x, y: y})
				count++
			}
		}
	}

	for _, p := range toRemove {
		lines[p.y][p.x] = '.'
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
	part1 := removeRolls(lines)
	fmt.Println("Part1:", part1)

	part2 := part1
	for {
		removed := removeRolls(lines)
		if removed == 0 {
			break
		}

		part2 += removed
	}

	fmt.Println("Part2:", part2)
}
