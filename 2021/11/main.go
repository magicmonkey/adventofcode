package main

import (
	"fmt"
	"strconv"
)

import "github.com/magicmonkey/adventofcode/2021/util"

func main() {
	//lines := testInput()
	lines := util.ReadInputFile()
	fmt.Println("Part 1")
	part1(lines)
	fmt.Println("Part 2")
	part2(lines)
}

type tMatrix [][]int

func part1(lines []string) {
	var numFlashes, n int
	matrix := parseInput(lines)
	for step := 0; step < 100; step++ {
		matrix, n = iterate(matrix)
		numFlashes += n
	}
	fmt.Println(numFlashes)
}

func part2(lines []string) {
	var n, step int
	matrix := parseInput(lines)
	for {
		step++
		matrix, n = iterate(matrix)
		if n == 100 {
			break
		}
	}
	fmt.Println(step)
}

func printMatrix(m tMatrix) {
	for _, row := range m {
		for _, cell := range row {
			fmt.Printf(" %d ", cell)
		}
		fmt.Println("")
	}
	fmt.Println("")
}

func iterate(src tMatrix) (dest tMatrix, numFlashes int) {

	// First, increment all cells by 1
	dest = incrementCells(src)

	// Check if any are exactly "10" and increment that and all around it
	var alreadyFlashed [10][10]bool
	for {
		numFlashesThisTime := checkFlash(&dest, &alreadyFlashed)
		if numFlashesThisTime == 0 {
			break
		}
	}

	// Count how many are 10 or over, and reset them to 0
	for y, row := range dest {
		for x, cell := range row {
			if cell >= 10 {
				numFlashes++
				dest[y][x] = 0
			}
		}
	}

	return
}

func checkFlash(src *tMatrix, alreadyFlashed *[10][10]bool) (numFlashes int) {
	size := 10
	for y, row := range *src {
		for x, cell := range row {
			if cell >= 10 && !(*alreadyFlashed)[y][x] {
				(*alreadyFlashed)[y][x] = true
				numFlashes++
				for modX := -1; modX <= 1; modX++ {
					for modY := -1; modY <= 1; modY++ {
						if x+modX >= 0 && x+modX < size && y+modY >= 0 && y+modY < size {
							(*src)[y+modY][x+modX]++
						}
					}
				}
			}
		}
	}
	return
}

func incrementCells(src tMatrix) (dest tMatrix) {
	// First, increment all cells by 1
	for _, row := range src {
		var energies []int
		for _, cell := range row {
			energies = append(energies, cell+1)
		}
		dest = append(dest, energies)
	}
	return
}

func parseInput(lines []string) (retval tMatrix) {
	for _, row := range lines {
		var energies []int
		for _, cell := range row {
			cellInt, _ := strconv.ParseInt(string(cell), 10, 32)
			energies = append(energies, int(cellInt))
		}
		retval = append(retval, energies)
	}
	return
}

func testInput() []string {
	return []string{
		"5483143223",
		"2745854711",
		"5264556173",
		"6141336146",
		"6357385478",
		"4167524645",
		"2176841721",
		"6882881134",
		"4846848554",
		"5283751526",
	}
}
