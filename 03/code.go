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

func part1(batteries [][]int) int {
	var sum int

	for row := range batteries {
		var maxLeft int
		for col := 0; col < len(batteries[row])-1; col++ {
			if batteries[row][col] > batteries[row][maxLeft] {
				maxLeft = col
				if batteries[row][maxLeft] == 9 {
					break
				}
			}
		}

		maxRight := maxLeft + 1
		for col := maxRight; col < len(batteries[row]); col++ {
			if batteries[row][col] > batteries[row][maxRight] {
				maxRight = col
				if batteries[row][maxRight] == 9 {
					break
				}
			}
		}

		sum += batteries[row][maxLeft]*10 + batteries[row][maxRight]
	}

	return sum
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
	fmt.Println("Part1:", part1(batteries))
}
