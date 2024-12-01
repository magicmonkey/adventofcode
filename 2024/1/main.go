package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
)

func MustInt(s string) int {
	numberInt, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return numberInt
}

func ParseFile(fname string) (list1, list2 []int) {
	// Open the input file
	file, err := os.Open(fname)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	// Create the regex
	re := regexp.MustCompile(`^(\d+)\s+(\d+)$`)

	// Create scanner
	scanner := bufio.NewScanner(file)

	// Read and process lines
	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindStringSubmatch(line)
		if matches != nil {
			list1 = append(list1, MustInt(matches[1]))
			list2 = append(list2, MustInt(matches[2]))
		}
	}

	// Check for scanning error
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	sort.Ints(list1)
	sort.Ints(list2)

	return
}

func part1() {
	list1, list2 := ParseFile("input.test")

	// Calculate the distances
	totalDist := 0
	for i, _ := range list1 {
		dist := list1[i] - list2[i]
		if dist < 0 {
			dist = -1 * dist
		}
		totalDist += dist
	}

	fmt.Println(totalDist)

}

func part2() {
	list1, list2 := ParseFile("input.test")
	weights := map[int]int{}
	for _, v := range list2 {
		val, exists := weights[v]
		if !exists {
			val = 0
		}
		weights[v] = val + 1
	}

	similarity := 0
	for _, v := range list1 {
		weight, exists := weights[v]
		if !exists {
			continue
		}
		similarity += weight * v
	}
	fmt.Println(similarity)
}

func main() {
	part2()
}
