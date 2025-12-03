package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readInput(file *os.File) [][]byte {
	scanner := bufio.NewScanner(file)
	var batteries [][]byte

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		batteriesRow := make([]byte, len(line))
		for i := range line {
			batteriesRow[i] = line[i]
		}

		batteries = append(batteries, batteriesRow)
	}

	return batteries
}

func findMaxIndex(slice []byte, start, end int) int {
	maxIndex := start
	for i := start; i < end; i++ {
		if slice[i] > slice[maxIndex] {
			maxIndex = i
		}
	}

	return maxIndex
}

func getANumber(text string) int {
	num, err := strconv.Atoi(text)
	if err != nil {
		log.Fatalf("Failed to convert %s to int!\n", text)
	}

	return num
}

func part1(batteries [][]byte) int {
	var sum int

	for row := range batteries {
		maxLeft := findMaxIndex(batteries[row], 0, len(batteries[row])-1)
		maxRight := findMaxIndex(batteries[row], maxLeft+1, len(batteries[row]))

		sum += getANumber(string(batteries[row][maxLeft]) + string(batteries[row][maxRight]))
	}

	return sum
}

func part2(batteries [][]byte) int {
	var sum int

	for row := range batteries {
		digits := make([]byte, 12)
		var index int

		for i := 12; i > 0; i-- {
			index = findMaxIndex(batteries[row], index, len(batteries[row])-i+1)
			digits[12-i] = batteries[row][index]
			index++
		}

		sum += getANumber(string(digits))
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
	fmt.Println("Part2:", part2(batteries))
}
