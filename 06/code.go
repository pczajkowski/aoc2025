package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readInput(file *os.File) ([]string, []string) {
	scanner := bufio.NewScanner(file)
	var numberLines []string
	var symbols []string

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		if line[0] == '+' || line[0] == '*' {
			var symbol []byte
			for i := range line {
				if i > 0 && line[i] == '+' || line[i] == '*' {
					symbols = append(symbols, string(symbol))
					symbol = []byte{}
				}

				symbol = append(symbol, line[i])
			}

			symbols = append(symbols, string(symbol))
		} else {
			numberLines = append(numberLines, line)
		}
	}

	return numberLines, symbols
}

func parseNumbers(numberLines, symbols []string) [][]string {
	var numbers [][]string

	for _, line := range numberLines {
		var start, end int
		var lineNumbers []string
		for _, symbol := range symbols {
			end += len(symbol)
			var number string
			for i := start; i < end; i++ {
				number += string(line[i])
			}

			lineNumbers = append(lineNumbers, number)
			start = end
		}

		numbers = append(numbers, lineNumbers)
	}

	return numbers
}

func part1(numbers [][]string, symbols []string) int {
	var sum int

	for i, symbol := range symbols {
		var result int
		for row := range numbers {
			if numbers[row][i] == "" {
				continue
			}

			number, err := strconv.Atoi(strings.Trim(numbers[row][i], " "))
			if err != nil {
				log.Fatalf("Failed to convert %s to int (row: %d, col: %d)!\n", numbers[row][i], row, i)
			}

			if symbol[0] == '+' {
				result += number
			} else if symbol[0] == '*' {
				if result == 0 {
					result = 1
				}
				result *= number
			}
		}

		sum += result
	}

	return sum
}

func part2(numbers [][]string, symbols []string) int {
	var sum int

	for i, symbol := range symbols {
		var result int

		for j := len(symbol) - 1; j >= 0; j-- {
			var digits []byte
			for row := range numbers {
				if numbers[row][i] == "" {
					continue
				}

				if numbers[row][i][j] >= '0' && numbers[row][i][j] <= '9' {
					digits = append(digits, numbers[row][i][j])
				}
			}

			numberString := string(digits)
			if numberString == "" {
				continue
			}

			number, err := strconv.Atoi(numberString)
			if err != nil {
				log.Fatalf("Failed to convert %s to int!\n", numberString)
			}

			if symbol[0] == '+' {
				result += number
			} else if symbol[0] == '*' {
				if result == 0 {
					result = 1
				}
				result *= number
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

	numberLines, symbols := readInput(file)
	numbers := parseNumbers(numberLines, symbols)
	fmt.Println("Part1:", part1(numbers, symbols))
	fmt.Println("Part2:", part2(numbers, symbols))
}
