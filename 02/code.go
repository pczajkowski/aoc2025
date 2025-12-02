package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func getRanges(data []byte) [][]int {
	var ranges [][]int
	ids := bytes.Split(data, []byte(","))
	for _, id := range ids {
		var start, end int
		n, err := fmt.Sscanf(string(id), "%d-%d", &start, &end)
		if err != nil || n != 2 {
			log.Fatalf("Invalid range: %s", id)
		}
		ranges = append(ranges, []int{start, end})
	}

	return ranges
}

func parts(ranges [][]int) (int, int) {
	invalid := make(map[int]int)
	invalidRest := make(map[int]int)
	for _, r := range ranges {
		for i := r[0]; i <= r[1]; i++ {
			s := strconv.Itoa(i)
			if len(s)%2 == 0 {
				half := len(s) / 2
				if s[:half] == s[half:] {
					invalid[i]++
					continue
				}
			}

			if len(s) > 1 {
				ds := s + s
				if strings.Contains(ds[1:len(ds)-2], s) {
					invalidRest[i]++
				}
			}
		}
	}

	var part1 int
	for key, value := range invalid {
		part1 += key * value
	}

	var part2 int
	for key, value := range invalidRest {
		part2 += key * value
	}

	return part1, part1 + part2
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("You need to specify a file!")
	}

	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	ranges := getRanges(data)
	part1, part2 := parts(ranges)
	fmt.Println("Part1:", part1)
	fmt.Println("Part2:", part2)
}
