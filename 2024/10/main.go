package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func MustInt(s string) int {
	numberInt, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return numberInt
}

func ParseFile(fname string) (topo [][]int) {
	// Open the input file
	file, err := os.Open(fname)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	// Create scanner
	scanner := bufio.NewScanner(file)

	// Read and process lines
	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		topo = append(topo, []int{})
		for _, b := range line {
			topo[y] = append(topo[y], MustInt(string(b)))
		}
		y++
	}

	// Check for scanning error
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	return
}

func nextStep(topo [][]int, y, x int, height int, summits *[][2]int) (numTrails int) {
	fmt.Printf("%sChecking %d,%d for height %d\n", strings.Repeat(".", height), x, y, height)
	if height == 9 && topo[y][x] == 9 {
		if slices.Contains(*summits, [2]int{y, x}) {
			fmt.Printf("%sAlready been to %d,%d\n", strings.Repeat("x", height), x, y)
			return 0
		} else {
			*summits = append(*summits, [2]int{y, x})
			fmt.Printf("%sFound a summit ending at %d,%d\n", strings.Repeat("*", height), x, y)
			return 1
		}
	}
	total := 0
	// North
	if y > 0 && topo[y-1][x] == height+1 {
		total += nextStep(topo, y-1, x, height+1, summits)
	}
	// South
	if y < len(topo)-1 && topo[y+1][x] == height+1 {
		total += nextStep(topo, y+1, x, height+1, summits)
	}
	// West
	if x > 0 && topo[y][x-1] == height+1 {
		total += nextStep(topo, y, x-1, height+1, summits)
	}
	// East
	if x < len(topo[0])-1 && topo[y][x+1] == height+1 {
		total += nextStep(topo, y, x+1, height+1, summits)
	}
	fmt.Printf("%sGot %d for height %d\n", strings.Repeat("-", height), total, height)
	return total
}

func part1(fname string) {
	topo := ParseFile(fname)

	total := 0
	for y, row := range topo {
		for x, height := range row {
			if height == 0 {
				summits := &[][2]int{}
				val := nextStep(topo, y, x, height, summits)
				fmt.Printf("=== %d ===\n", val)
				total += val
			}
		}
	}
	fmt.Println(total)
}

// 1324
func part2(fname string) {
	topo := ParseFile(fname)
	summits := &[][2]int{}

	total := 0
	for y, row := range topo {
		for x, height := range row {
			if height == 0 {
				val := nextStep(topo, y, x, height, summits)
				fmt.Printf("=== %d ===\n", val)
				total += val
			}
		}
	}
	fmt.Println(total)
}

func main() {
	part1("input.txt")
}
