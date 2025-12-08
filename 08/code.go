package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

type box struct {
	x, y, z     int
	id, circuit int
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

func part1(boxes []box) int {
	var result int

	circuit := 1

	for i := range boxes {
		bestDistance := math.MaxFloat64
		best := -1
		for j := range boxes {
			if j == i {
				continue
			}

			distance := boxes[j].DistanceTo(boxes[i])
			if distance < bestDistance {
				bestDistance = distance
				best = j
			}
		}

		if boxes[i].circuit > 0 {
			boxes[best].circuit = boxes[i].circuit
		} else if boxes[best].circuit > 0 {
			boxes[i].circuit = boxes[best].circuit
		} else {
			boxes[i].circuit = circuit
			boxes[best].circuit = circuit
			circuit++
		}
	}

	return result
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
	part1(boxes)
	fmt.Println(boxes)
}
