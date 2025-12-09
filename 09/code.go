package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type tile struct {
	x, y int
}

func readInput(file *os.File) []tile {
	scanner := bufio.NewScanner(file)
	var tiles []tile

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		var x, y int
		n, err := fmt.Sscanf(line, "%d,%d", &x, &y)
		if n != 2 || err != nil {
			log.Fatalf("Bad input: %s", line)
		}

		tiles = append(tiles, tile{x: x, y: y})
	}

	return tiles
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}

func part1(tiles []tile) int {
	end := len(tiles)
	var maxArea int

	for i := range tiles {
		for j := i + 1; j < end; j++ {
			area := (abs(tiles[j].x-tiles[i].x) + 1) * (abs(tiles[j].y-tiles[i].y) + 1)
			if area > maxArea {
				maxArea = area
			}
		}
	}

	return maxArea
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

	tiles := readInput(file)
	fmt.Println("Part1:", part1(tiles))
}
