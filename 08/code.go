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

type distance struct {
	first, second int
	value         float64
}

func part1(boxes []box) int {
	var result int
	var distances []distance

	for i := range boxes {
		for j := i + 1; j < len(boxes); j++ {
			howFar := boxes[j].DistanceTo(boxes[i])
			distances = append(distances, distance{first: boxes[i].id, second: boxes[j].id, value: howFar})
		}

	}

	sort.Slice(distances, func(i, j int) bool { return distances[i].value < distances[j].value })

	var circuits [][]int
	for i := range distances {
		first := boxes[distances[i].first]
		second := boxes[distances[i].second]
		if first.inCircuit && second.inCircuit {
			continue
		}

		if !first.inCircuit && !second.inCircuit {
			circuits = append(circuits, []int{first.id, second.id})
			boxes[distances[i].first].inCircuit = true
			boxes[distances[i].second].inCircuit = true
		} else if !first.inCircuit {
			var found bool
			for j := range circuits {
				for k := range circuits[j] {
					if circuits[j][k] == second.id {
						found = true
						circuits[j] = append(circuits[j], first.id)
						break
					}
				}
			}

			boxes[distances[i].first].inCircuit = true
			if !found {
				circuits = append(circuits, []int{first.id})
			}
		} else if !second.inCircuit {
			var found bool
			for j := range circuits {
				for k := range circuits[j] {
					if circuits[j][k] == first.id {
						found = true
						circuits[j] = append(circuits[j], second.id)
						break
					}
				}
			}

			boxes[distances[i].second].inCircuit = true
			if !found {
				circuits = append(circuits, []int{second.id})
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
