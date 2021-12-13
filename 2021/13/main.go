package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/magicmonkey/adventofcode/2021/util"
)

type tCoord struct {
	x int
	y int
}

type tGrid struct {
	coords []tCoord
}

type tFold struct {
	axis  string
	place int
}

func main() {
	//lines := testInput()
	lines := util.ReadInputFile()
	fmt.Println("Part 1")
	part1(lines)
	fmt.Println("Part 2")
	part2(lines)
}

func part1(lines []string) {
	grid, folds := parseInput(lines)
	grid.applyFold(folds[0])
	fmt.Println(len(grid.coords))
}

func part2(lines []string) {
	grid, folds := parseInput(lines)
	for _, f := range folds {
		grid.applyFold(f)
	}
	grid.print()
}

func (grid *tGrid) applyFold(f tFold) {
	if f.axis == "y" {
		grid.applyYFold(f.place)
	}
	if f.axis == "x" {
		grid.applyXFold(f.place)
	}
	grid.removeDupes()
}

func (grid *tGrid) print() {
	var maxX, maxY int
	fmt.Println("")
	for _, c := range grid.coords {
		if c.x > maxX {
			maxX = c.x
		}
		if c.y > maxY {
			maxY = c.y
		}
	}

	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			if grid.contains(tCoord{x: x, y: y}) {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println("")
	}

	fmt.Println(" === ", len(grid.coords), " === ")

}

func (grid *tGrid) contains(c1 tCoord) bool {
	for _, c := range grid.coords {
		if c.x == c1.x && c.y == c1.y {
			return true
		}
	}
	return false
}

func (grid *tGrid) removeDupes() {
	var newCoords []tCoord
outer:
	for i := 0; i < len(grid.coords); i++ {
		c1 := grid.coords[i]
		for j := i + 1; j < len(grid.coords); j++ {
			c2 := grid.coords[j]
			//fmt.Printf("Checking (%d,%d) %d against (%d,%d) %d\n", c1.x, c1.y, i, c2.x, c2.y, j)
			if c1.x == c2.x && c1.y == c2.y {
				continue outer
			}
		}
		newCoords = append(newCoords, c1)
	}
	grid.coords = newCoords
}

func (grid *tGrid) applyYFold(place int) {
	var newCoords []tCoord

	for _, c := range grid.coords {
		if c.y > place {
			newCoords = append(newCoords, tCoord{x: c.x, y: (place - (c.y - place))})
		} else {
			newCoords = append(newCoords, c)
		}
	}
	grid.coords = newCoords
}

func (grid *tGrid) applyXFold(place int) {
	var newCoords []tCoord

	for _, c := range grid.coords {
		if c.x > place {
			newCoords = append(newCoords, tCoord{x: (place - (c.x - place)), y: c.y})
		} else {
			newCoords = append(newCoords, c)
		}
	}
	grid.coords = newCoords
}

func (grid *tGrid) addCoord(c tCoord) {
	grid.coords = append(grid.coords, c)
}

func parseInput(lines []string) (grid *tGrid, folds []tFold) {
	grid = &tGrid{}
	foldRe := regexp.MustCompile(`^fold along ([xy])=(\d*)$`)
	coordRe := regexp.MustCompile(`^(\d*),(\d*)$`)
	for _, line := range lines {
		if line == "" {
			continue
		} else if line[0] == 'f' {
			// fold instruction
			parts := foldRe.FindStringSubmatch(line)
			placeInt, _ := strconv.ParseInt(parts[2], 10, 32)
			folds = append(folds, tFold{axis: parts[1], place: int(placeInt)})
		} else {
			// Coordinate
			parts := coordRe.FindStringSubmatch(line)
			xInt, _ := strconv.ParseInt(parts[1], 10, 32)
			yInt, _ := strconv.ParseInt(parts[2], 10, 32)
			grid.addCoord(tCoord{x: int(xInt), y: int(yInt)})
		}
	}
	return
}

func testInput() []string {
	return []string{
		"6,10",
		"0,14",
		"9,10",
		"0,3",
		"10,4",
		"4,11",
		"6,0",
		"6,12",
		"4,1",
		"0,13",
		"10,12",
		"3,4",
		"3,0",
		"8,4",
		"1,10",
		"2,14",
		"8,10",
		"9,0",
		"",
		"fold along y=7",
		"fold along x=5",
	}
}
