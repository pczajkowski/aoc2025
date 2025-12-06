package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readInput(file *os.File) ([][]int, []string) {
	scanner := bufio.NewScanner(file)
	var numbers [][]int
	var symbols []string

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		if line[0] == '+' || line[0] == '*' {
			parts := strings.Split(line, " ")
			for _, part := range parts {
				if part == "" || part == " " {
					continue
				}

				symbols = append(symbols, strings.Trim(part, " "))
			}
		} else {
			var lineNumbers []int
			parts := strings.Split(line, " ")
			for _, part := range parts {
				if part == "" || part == " " {
					continue
				}

				i, err := strconv.Atoi(strings.Trim(part, " "))
				if err != nil {
					log.Fatalf("Bad input: %s", part)
				}

				lineNumbers = append(lineNumbers, i)
			}

			numbers = append(numbers, lineNumbers)
		}
	}

	return numbers, symbols
}

func part1(numbers [][]int, symbols []string) int {
	var sum int

	for i, symbol := range symbols {
		var result int
		for row := range numbers {
			if symbol == "+" {
				result += numbers[row][i]
			} else if symbol == "*" {
				if result == 0 {
					result = 1
				}
				result *= numbers[row][i]
			}
		}

		sum += result
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

	numbers, symbols := readInput(file)
	fmt.Println("Part1:", part1(numbers, symbols))
}
