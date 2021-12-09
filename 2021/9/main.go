package main

import (
	"fmt"
	"strconv"

	"github.com/magicmonkey/adventofcode/2021/util"
)

func main() {
	//lines := testInput()
	lines := util.ReadInputFile()
	part1(lines)
}

func part1(lines []string) {
	var heightmap [][]int

	// Allocate memory
	var sizeX int = len(lines[0])
	var sizeY int = len(lines)
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

	// Analyse height map
	var risk int
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
			//fmt.Println("Yes")
			risk += height + 1
		}
	}
	fmt.Println(risk)
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
