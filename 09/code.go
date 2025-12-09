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
	fmt.Println(tiles)
}
