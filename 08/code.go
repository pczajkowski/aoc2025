package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type box struct {
	x, y, z int
	id      int
}

func readInput(file *os.File) []box {
	scanner := bufio.NewScanner(file)
	var boxes []box
	var id int

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		var x, y, z int
		n, err := fmt.Sscanf(line, "%d,%d,%d", &x, &y, &z)
		if n != 3 || err != nil {
			log.Fatalf("Bad input: %s", line)
		}

		boxes = append(boxes, box{x: x, y: y, z: z, id: id})
		id++
	}

	return boxes
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

	boxes := readInput(file)
	fmt.Println(boxes)
}
