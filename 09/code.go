package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

type Point struct {
	X, Y int
}

type PointList []Point

func (p PointList) Len() int {
	return len(p)
}

func (p PointList) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func Area2(a, b, c Point) int {
	return (b.X-a.X)*(c.Y-a.Y) - (c.X-a.X)*(b.Y-a.Y)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}

func (p PointList) Less(i, j int) bool {
	area := Area2(p[0], p[i], p[j])
	if area == 0 {
		x := abs(p[i].X-p[0].X) - abs(p[j].X-p[0].X)
		y := abs(p[i].Y-p[0].Y) - abs(p[j].Y-p[0].Y)
		if x < 0 || y < 0 {
			return true
		} else if x > 0 || y > 0 {
			return false
		} else {
			return false
		}
	}
	return area > 0
}

func (p PointList) FindLowestPoint() {
	m := 0
	for i := 1; i < len(p); i++ {
		//If lowest points are on the same line, take the rightmost point
		if (p[i].Y < p[m].Y) || ((p[i].Y == p[m].Y) && p[i].X > p[m].X) {
			m = i
		}
	}
	p[0], p[m] = p[m], p[0]
}

func isLeft(p0, p1, p2 Point) bool {
	return Area2(p0, p1, p2) > 0
}

func (points PointList) Compute() (PointList, bool) {
	if len(points) < 3 {
		return nil, false
	}

	stack := new(Stack)
	points.FindLowestPoint()
	sort.Sort(&points)

	stack.Push(points[0])
	stack.Push(points[1])

	i := 2
	for i < len(points) {
		pi := points[i]
		p1 := stack.top.next.value.(Point)
		p2 := stack.top.value.(Point)
		if isLeft(p1, p2, pi) {
			stack.Push(pi)
			i++
		} else {
			stack.Pop()
		}
	}

	//Copy the hull
	ret := make(PointList, stack.Len())
	top := stack.top
	count := 0
	for top != nil {
		ret[count] = top.value.(Point)
		top = top.next
		count++
	}
	return ret, true
}

func readInput(file *os.File) []Point {
	scanner := bufio.NewScanner(file)
	var tiles []Point

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		var x, y int
		n, err := fmt.Sscanf(line, "%d,%d", &x, &y)
		if n != 2 || err != nil {
			log.Fatalf("Bad input: %s", line)
		}

		tiles = append(tiles, Point{X: x, Y: y})
	}

	return tiles
}

func part1(tiles []Point) int {
	end := len(tiles)
	var maxArea int

	for i := range tiles {
		for j := i + 1; j < end; j++ {
			area := (abs(tiles[j].X-tiles[i].X) + 1) * (abs(tiles[j].Y-tiles[i].Y) + 1)
			if area > maxArea {
				maxArea = area
			}
		}
	}

	return maxArea
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

	tiles := readInput(file)
	fmt.Println("Part1:", part1(tiles))
}
