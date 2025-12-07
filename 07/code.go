package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readInput(file *os.File) []string {
	scanner := bufio.NewScanner(file)
	var maze []string

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		maze = append(maze, line)
	}

	return maze
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

	maze := readInput(file)
	fmt.Println(maze)
}
