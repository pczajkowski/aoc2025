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
	fmt.Println(batteries)
}
