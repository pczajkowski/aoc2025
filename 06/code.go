package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readInput(file *os.File) ([][]string, []string) {
	scanner := bufio.NewScanner(file)
	var numbers [][]string
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
			var lineNumbers []string
			var number string
			for i := range line {
				if line[i] == ' ' {
					if i > 0 && line[i-1] >= '0' && line[i-1] <= '9' {
						lineNumbers = append(lineNumbers, number)
						number = ""
						continue
					}
				}

				number += string(line[i])
			}

			lineNumbers = append(lineNumbers, number)
			numbers = append(numbers, lineNumbers)
		}
	}

	return numbers, symbols
}

func part1(numbers [][]string, symbols []string) int {
	var sum int

	for i, symbol := range symbols {
		var result int
		for row := range numbers {
			number, err := strconv.Atoi(strings.Trim(numbers[row][i], " "))
			if err != nil {
				log.Fatalf("Failed to convert %s to int!\n", numbers[row][i])
			}

			if symbol == "+" {
				result += number
			} else if symbol == "*" {
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

	numbers, symbols := readInput(file)
	fmt.Println("Part1:", part1(numbers, symbols))
}
