package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func readInput(file *os.File) map[string][]string {
	scanner := bufio.NewScanner(file)
	devices := make(map[string][]string)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		parts := strings.Split(line, ": ")
		if len(parts) < 2 {
			log.Fatalf("Bad input: %s", line)
		}

		name := parts[0]
		connections := strings.Split(parts[1], " ")
		devices[name] = connections
	}

	return devices
}

func part1(entry string, devices map[string][]string) int {
	if entry == "out" {
		return 1
	}

	var count int
	for _, device := range devices[entry] {
		count += part1(device, devices)
	}

	return count
}

func passedDACandFFT(path []string) bool {
	var dac, fft bool
	for i := range path {
		if path[i] == "dac" {
			dac = true
		}

		if path[i] == "fft" {
			fft = true
		}
	}

	return dac && fft
}

func part2(entry string, devices map[string][]string) int {
	var count int
	visited := make(map[string]bool)
	path := []string{}

	var dfs func(current string)
	dfs = func(current string) {
		path = append(path, current)
		visited[current] = true
		defer func() {
			path = path[:len(path)-1]
			visited[current] = false
		}()

		if current == "out" {
			if passedDACandFFT(path) {
				count++
			}

			return
		} else {
			for _, neighbor := range devices[current] {
				if !visited[neighbor] {
					dfs(neighbor)
				}
			}
		}
	}

	dfs(entry)
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

	devices := readInput(file)
	fmt.Println("Part1:", part1("you", devices))
	fmt.Println("Part2:", part2("svr", devices))
}
