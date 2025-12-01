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

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func parts(rotations []Rotation) (int, int) {
	var zeros int
	var passedZeros int
	dial := 50

	for _, rotation := range rotations {
		was := dial
		if rotation.Direction == 'L' {
			dial -= rotation.Clicks
		} else {
			dial += rotation.Clicks
		}

		if dial > 100 {
			passedZeros += dial / 100
		} else if dial < 0 {
			passedZeros += abs(dial) / 100
			if was != 0 {
				passedZeros++
			}
		}

		dial %= 100
		if dial < 0 {
			dial += 100
		}

		if dial == 0 {
			zeros++
		}
	}

	return zeros, passedZeros
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
	zeros, passedZeros := parts(rotations)
	fmt.Println("Part1:", zeros)
	fmt.Println("Part2:", zeros+passedZeros)
}
