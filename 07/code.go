package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readInput(file *os.File) ([]string, int) {
	scanner := bufio.NewScanner(file)
	var maze []string
	start := -1

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		if start < 0 {
			for i := range line {
				if line[i] == 'S' {
					start = i
					break
				}
			}
		}

		maze = append(maze, line)
	}

	return maze, start
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

	maze, start := readInput(file)
	fmt.Println(maze, start)
}
