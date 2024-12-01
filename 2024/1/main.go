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

func main() {
	fmt.Println("Starting...")

	// Open the input file
	fileName := "input.part1"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	// Create the regex
	re := regexp.MustCompile(`^(\d+)\s+(\d+)$`)

	// Create scanner
	scanner := bufio.NewScanner(file)

	var list1, list2 []int

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
