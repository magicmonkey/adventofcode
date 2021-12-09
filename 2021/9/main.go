package main

import (
	"fmt"
	"strconv"
	//"github.com/magicmonkey/adventofcode/2021/util"
)

func main() {
	lines := testInput()
	//lines := util.ReadInputFile()
	part1(lines)
	//part2(lines)
}

func part1(lines []string) {
	heightmap, sizeX, sizeY := parseInput(lines)
	minima := findMinima(heightmap, sizeX, sizeY)

	// Analyse height map
	var risk int
	for _, coord := range minima {
		risk += heightmap[coord.x][coord.y] + 1
	}
	fmt.Println(risk)
}

func part2(lines []string) {
}

type tCoord struct {
	x int
	y int
}

func findMinima(heightmap [][]int, sizeX int, sizeY int) (retval []tCoord) {
	for x, col := range heightmap {
		for y, height := range col {
			//fmt.Printf("Checking %d,%d (%d)... ", x, y, height)
			if x == 0 {
				if height >= heightmap[x+1][y] {
					//fmt.Printf("No (1) because %d,%d (%d)\n", x+1, y, heightmap[x+1][y])
					continue
				}
			} else if x == sizeX-1 {
				if height >= heightmap[x-1][y] {
					//fmt.Printf("No (2) because %d,%d (%d)\n", x-1, y, heightmap[x-1][y])
					continue
				}
			} else {
				if height >= heightmap[x+1][y] {
					//fmt.Printf("No (3) because %d,%d (%d)\n", x+1, y, heightmap[x+1][y])
					continue
				}
				if height >= heightmap[x-1][y] {
					//fmt.Printf("No (4) because %d,%d (%d)\n", x-1, y, heightmap[x-1][y])
					continue
				}
			}
			if y == 0 {
				if height >= heightmap[x][y+1] {
					//fmt.Printf("No (5) because %d,%d (%d)\n", x, y+1, heightmap[x][y+1])
					continue
				}
			} else if y == sizeY-1 {
				if height >= heightmap[x][y-1] {
					//fmt.Printf("No (6) because %d,%d (%d)\n", x, y-1, heightmap[x][y-1])
					continue
				}
			} else {
				if height >= heightmap[x][y+1] {
					//fmt.Printf("No (7) because %d,%d (%d)\n", x, y+1, heightmap[x][y+1])
					continue
				}
				if height >= heightmap[x][y-1] {
					//fmt.Printf("No (8) because %d,%d (%d)\n", x, y-1, heightmap[x][y-1])
					continue
				}
			}
			retval = append(retval, tCoord{x: x, y: y})
		}
	}
	return
}

func parseInput(lines []string) (heightmap [][]int, sizeX int, sizeY int) {
	// Allocate memory
	sizeX = len(lines[0])
	sizeY = len(lines)
	heightmap = make([][]int, sizeX)
	for i := range heightmap {
		heightmap[i] = make([]int, sizeY)
	}

	// Read input
	for y, row := range lines {
		for x, heightStr := range row {
			heightInt, err := strconv.ParseInt(string(heightStr), 10, 32)
			if err != nil {
				panic(err)
			}
			heightmap[x][y] = int(heightInt)
		}
	}
	return
}

func testInput() []string {
	return []string{
		"2199943210",
		"3987894921",
		"9856789892",
		"8767896789",
		"9899965678",
	}
}
