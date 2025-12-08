package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

type box struct {
	x, y, z   int
	id, best  int
	inCircuit bool
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

type circuit struct {
	first, second int
	boxes         []box
}

func part1(boxes []box) int {
	var result int

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

		boxes[i].best = boxes[best].id
	}

	var circuits []circuit
	for i := range boxes {
		if boxes[i].inCircuit {
			continue
		}

		if boxes[boxes[i].best].best == boxes[i].id {
			boxes[i].inCircuit = true
			boxes[boxes[i].best].inCircuit = true

			set := circuit{first: boxes[i].id, second: boxes[i].best, boxes: []box{boxes[i], boxes[boxes[i].best]}}
			circuits = append(circuits, set)
		}
	}

	for _, item := range boxes {
		if item.inCircuit {
			continue
		}

		for i := range circuits {
			if item.best == circuits[i].first || item.best == circuits[i].second {
				circuits[i].boxes = append(circuits[i].boxes, item)
				break
			}
		}
	}
	fmt.Println(circuits)

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
}
