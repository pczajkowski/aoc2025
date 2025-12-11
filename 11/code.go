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

func findPaths(entry string, devices map[string][]string, visited map[string]bool) int {
	if entry == "out" {
		if visited["dac"] && visited["fft"] {
			return 1
		}

		return 0
	}

	visited[entry] = true
	var count int
	for _, device := range devices[entry] {
		visited[device] = true
		count += findPaths(device, devices, visited)
	}

	return count
}

func part2(entry string, devices map[string][]string) int {
	var count int
	for _, device := range devices[entry] {
		visited := make(map[string]bool)
		visited[device] = true
		count += findPaths(device, devices, visited)
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

	devices := readInput(file)
	fmt.Println("Part1:", part1("you", devices))
	fmt.Println("Part2:", part2("svr", devices))
}
