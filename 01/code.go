package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Rotation struct {
	Direction byte
	Clicks    int
}

func readInput(file *os.File) []Rotation {
	scanner := bufio.NewScanner(file)
	var rotations []Rotation

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		var direction byte
		var clicks int
		n, err := fmt.Sscanf(line, "%c%d", &direction, &clicks)
		if n != 2 || err != nil {
			log.Fatalf("Bad input: %s", line)
		}

		rotations = append(rotations, Rotation{Direction: direction, Clicks: clicks})
	}

	return rotations
}

func part1(rotations []Rotation) int {
	var zeros int
	dial := 50

	for _, rotation := range rotations {
		if rotation.Direction == 'L' {
			dial -= rotation.Clicks
		} else {
			dial += rotation.Clicks
		}

		dial %= 100
		if dial < 0 {
			dial += 100
		}

		if dial == 0 {
			zeros++
		}
	}

	return zeros
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

	rotations := readInput(file)
	fmt.Println("Part1:", part1(rotations))
}
