package main

import (
	"fmt"
	"os"
)

// from https://dev.to/toluwasethomas/simple-way-to-obtain-largest-number-in-an-array-or-slice-in-golang-2o06
func findLargestNumber(nums []int) int {
	if len(nums) == 0 {
		return 0 // handle empty slice case
	}
	largest := nums[0]               // Step 2
	for i := 1; i < len(nums); i++ { // Step 1
		if nums[i] > largest { // Step 3
			largest = nums[i] // Step 4
		}
	}
	return largest // Step 5
}

func ParseFile(fname string) (contents []int) {
	source, _ := os.ReadFile(fname)

	for _, b := range source {
		contents = append(contents, int(b-'0'))
	}
	return
}

func part1(fname string) {
	input := ParseFile(fname)

	expanded := make(map[int]int)
	compacted := make(map[int]int)

	id := 0
	place := 0

	for i := 0; i < len(input); i += 2 {
		// First lay down the file
		for j := 0; j < input[i]; j++ {
			expanded[place] = id
			place++
		}
		id++

		if len(input) <= i+1 {
			break
		}

		// Second lay down the gap
		for j := 0; j < input[i+1]; j++ {
			place++
		}
	}

	// Compact the data

	size := len(expanded)
	for i := 0; i < size; i++ {
		d, ok := expanded[i]
		if ok {
			compacted[i] = d
		} else {
			keys := make([]int, len(expanded))
			k := 0
			for m := range expanded {
				keys[k] = m
				k++
			}
			a := findLargestNumber(keys)
			compacted[i] = expanded[a]
			delete(expanded, a)
		}
	}

	// Checksum
	checksum := 0
	for k, v := range compacted {
		checksum += (k * v)
	}

	fmt.Println(checksum)
}

func main() {
	part1("input.txt")
}
