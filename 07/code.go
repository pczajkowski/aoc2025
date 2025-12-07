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

func part1(maze []string, beams []int) int {
	var count int

	for row := 1; row < len(maze); row++ {
		for i := range beams {
			if beams[i] == 0 {
				continue
			}

			if maze[row][i] == '^' {
				if i > 0 {
					beams[i-1]++
				}

				if i < len(beams)-1 {
					beams[i+1]++
				}

				beams[i] = 0

				count++
			}
		}
	}

	return count
}

func part2(maze []string, x, y int) int {
	if y >= len(maze) {
		return 1
	}

	var count int
	var left, right int
	if maze[y][x] == '.' {
		count += part2(maze, x, y+1)
	} else if maze[y][x] == '^' {
		if x > 0 {
			left += part2(maze, x-1, y+1)
		}

		if x < len(maze[0])-1 {
			right += part2(maze, x+1, y+1)
		}

		count += left + right
	}

	return count
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
	beams := make([]int, len(maze[0]))
	beams[start] = 1
	fmt.Println("Part1:", part1(maze, beams))
	fmt.Println("Part2:", part2(maze, start, 1))
}
