package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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
	last := len(tiles) - 1
	return abs(tiles[last].x-tiles[0].x+1) * abs(tiles[last].y-tiles[0].y+1)
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
	sort.Slice(tiles, func(i, j int) bool {
		if tiles[i].x == tiles[j].x {
			return tiles[i].y < tiles[j].y
		}

		return tiles[i].x < tiles[j].x
	})

	fmt.Println("Part1:", part1(tiles))
}
