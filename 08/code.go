package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
)

type box struct {
	x, y, z int
	id      int
}

func (p box) DistanceTo(other box) float64 {
	dx := other.x - p.x
	dy := other.y - p.y
	dz := other.z - p.z
	return math.Sqrt(float64(dx*dx + dy*dy + dz*dz))
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

	markerBox := box{x: 0, y: 0, z: 0}
	boxes := readInput(file)
	fmt.Println(boxes[0], boxes[len(boxes)-1], boxes[0].DistanceTo(boxes[len(boxes)-1]))
	sort.Slice(boxes, func(i, j int) bool {
		return markerBox.DistanceTo(boxes[i]) < markerBox.DistanceTo(boxes[j])
	})

	for i := 1; i < len(boxes); i++ {
		fmt.Println(boxes[i], boxes[i-1], boxes[i].DistanceTo(boxes[i-1]))
	}
}
