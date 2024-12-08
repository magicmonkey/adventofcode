package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	x int
	y int
}

func ParseFile(fname string) (aerials map[string][]Point, height, width int) {
	aerials = make(map[string][]Point)

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
	var x, y int
	for scanner.Scan() {
		line := scanner.Text()
		x = -1
		for _, char := range line {
			x++
			if char == '.' {
				continue
			}
			p := Point{x: x, y: y}
			aerials[string(char)] = append(aerials[string(char)], p)
		}
		y++
	}

	height = y - 1
	width = x

	// Check for scanning error
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	return
}

func checkPair(aerial1, aerial2 Point, maxY, maxX int) (results []Point) {
	var antinode1, antinode2 Point

	xDiff := aerial1.x - aerial2.x
	if xDiff < 0 {
		xDiff *= -1
	}
	yDiff := aerial1.y - aerial2.y
	if yDiff < 0 {
		yDiff *= -1
	}

	if aerial1.x < aerial2.x {
		antinode1.x = aerial1.x - xDiff
		antinode2.x = aerial2.x + xDiff
	} else {
		antinode1.x = aerial1.x + xDiff
		antinode2.x = aerial2.x - xDiff
	}
	if aerial1.y < aerial2.y {
		antinode1.y = aerial1.y - yDiff
		antinode2.y = aerial2.y + yDiff
	} else {
		antinode1.y = aerial1.y + yDiff
		antinode2.y = aerial2.y - yDiff
	}

	// Check if antinode1 is off the map
	if antinode1.x < 0 || antinode1.x > maxX || antinode1.y < 0 || antinode1.y > maxY {
		// Off the map
	} else {
		results = append(results, antinode1)
	}

	// Check if antinode2 is off the map
	if antinode2.x < 0 || antinode2.x > maxX || antinode2.y < 0 || antinode2.y > maxY {
		// Off the map
	} else {
		results = append(results, antinode2)
	}

	return
}

func checkAerials(aerials []Point, maxY, maxX int) (results []Point) {
	var combinations [][2]Point
	for i := 0; i < len(aerials); i++ {
		for j := i + 1; j < len(aerials); j++ {
			combinations = append(combinations, [2]Point{aerials[i], aerials[j]})
		}
	}

	for _, combo := range combinations {
		res := checkPair(combo[0], combo[1], maxY, maxX)
		results = append(results, res...)
	}
	return
}

func part1(fname string) {
	allAerials, maxY, maxX := ParseFile(fname)

	locations := make(map[Point]bool)

	for _, aerials := range allAerials {
		antinodes := checkAerials(aerials, maxY, maxX)
		for _, antinode := range antinodes {
			if _, ok := locations[antinode]; !ok {
				locations[antinode] = true
			}
		}
	}
	fmt.Println(len(locations))

}

func main() {
	part1("input.txt")
}
