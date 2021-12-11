package main

import (
	"fmt"
	"github.com/magicmonkey/adventofcode/2021/util"
	"sort"
	"strconv"
)

func main() {
	//lines := testInput()
	lines := util.ReadInputFile()
	fmt.Println("Part 1")
	part1(lines)
	fmt.Println("Part 2")
	part2(lines)
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
	heightmap, sizeX, sizeY := parseInput(lines)
	minima := findMinima(heightmap, sizeX, sizeY)

	var sizes []int
	for _, minimum := range minima {
		size := findBasinSize(minimum, heightmap, sizeX, sizeY)
		sizes = append(sizes, size)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(sizes)))

	fmt.Println(sizes[0] * sizes[1] * sizes[2])

}

func findBasinSize(coord tCoord, heightmap [][]int, sizeX int, sizeY int) int {
	// Allocate memory
	visited := make([][]bool, sizeX)
	for i := range visited {
		visited[i] = make([]bool, sizeY)
	}
	isInBasin := make([][]bool, sizeX)
	for i := range isInBasin {
		isInBasin[i] = make([]bool, sizeY)
	}

	visit(coord, heightmap, &visited, &isInBasin, sizeX, sizeY)

	size := countTruths(isInBasin)

	//renderBasin(heightmap, isInBasin)
	return size
}

func countTruths(matrix [][]bool) (retval int) {
	for _, row := range matrix {
		for _, el := range row {
			if el {
				retval++
			}
		}
	}
	return
}

func visit(coord tCoord, heightmap [][]int, visited *[][]bool, isInBasin *[][]bool, sizeX int, sizeY int) {

	if heightmap[coord.x][coord.y] == 9 {
		return
	}

	(*visited)[coord.x][coord.y] = true
	(*isInBasin)[coord.x][coord.y] = true

	// Look right
	if coord.x < (sizeX-1) && !(*visited)[coord.x+1][coord.y] && heightmap[coord.x][coord.y] < heightmap[coord.x+1][coord.y] {
		visit(tCoord{x: coord.x + 1, y: coord.y}, heightmap, visited, isInBasin, sizeX, sizeY)
	}

	// Look up
	if coord.y > 0 && !(*visited)[coord.x][coord.y-1] && heightmap[coord.x][coord.y] < heightmap[coord.x][coord.y-1] {
		visit(tCoord{x: coord.x, y: coord.y - 1}, heightmap, visited, isInBasin, sizeX, sizeY)
	}

	// Look left
	if coord.x > 0 && !(*visited)[coord.x-1][coord.y] && heightmap[coord.x][coord.y] < heightmap[coord.x-1][coord.y] {
		visit(tCoord{x: coord.x - 1, y: coord.y}, heightmap, visited, isInBasin, sizeX, sizeY)
	}

	// Look down
	if coord.y < (sizeY-1) && !(*visited)[coord.x][coord.y+1] && heightmap[coord.x][coord.y] < heightmap[coord.x][coord.y+1] {
		visit(tCoord{x: coord.x, y: coord.y + 1}, heightmap, visited, isInBasin, sizeX, sizeY)
	}
}

func renderBasin(heightmap [][]int, isInBasin [][]bool) {
	var size int
	for x, row := range heightmap {
		for y, height := range row {
			if isInBasin[x][y] {
				fmt.Printf("-%d-", height)
				size++
			} else {
				fmt.Printf(" %d ", height)
			}
		}
		fmt.Println("")
	}
	fmt.Println("===", size, "===")
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
