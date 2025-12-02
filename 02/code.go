package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
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

func main() {
	if len(os.Args) < 2 {
		log.Fatal("You need to specify a file!")
	}

	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	ranges := getRanges(data)
	fmt.Println(ranges)
}
