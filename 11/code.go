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
	fmt.Println(devices)
}
