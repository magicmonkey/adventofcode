package main

import (
	"fmt"
	"github.com/magicmonkey/adventofcode/2021/util"
	"regexp"
	"strconv"
)

type tVent struct {
	x1, y1, x2, y2 int
	xmod, ymod     int
	straight       bool
}

type tGrid [1000][1000]int

func main() {
	//lines := testInput()
	lines := util.ReadInputFile()
	fmt.Println("Part 1")
	part1(lines)
	fmt.Println("Part 2")
	part2(lines)
}

func part1(lines []string) {
	findNumIntersects(lines, true)
}

func part2(lines []string) {
	findNumIntersects(lines, false)
}

func findNumIntersects(lines []string, straightOnly bool) {
	var grid tGrid
	var vents []tVent
	vents = parseInput(lines)

	for _, vent := range vents {
		if straightOnly && !vent.straight {
			continue
		}
		var currx, curry int
		currx = vent.x1
		curry = vent.y1
		grid[currx][curry]++
		for currx != vent.x2 || curry != vent.y2 {
			currx += vent.xmod
			curry += vent.ymod
			grid[currx][curry]++
		}

	}

	answer := numIntersections(grid, 2)
	fmt.Println(answer)
}

func numIntersections(grid tGrid, threshold int) (retval int) {
	for _, row := range grid {
		for _, cell := range row {
			if cell >= threshold {
				retval++
			}
		}
	}
	return
}

func printGrid(grid tGrid) {
	for _, row := range grid {
		fmt.Println(row)
	}
	fmt.Println("---")
}

func parseInput(lines []string) (retval []tVent) {
	parser := regexp.MustCompile(`^(\d*),(\d*) -> (\d*),(\d*)$`)
	for _, line := range lines {
		coords := parser.FindStringSubmatch(line)
		x1, err := strconv.ParseInt(coords[1], 10, 32)
		if err != nil {
			panic(err)
		}
		y1, err := strconv.ParseInt(coords[2], 10, 32)
		if err != nil {
			panic(err)
		}
		x2, err := strconv.ParseInt(coords[3], 10, 32)
		if err != nil {
			panic(err)
		}
		y2, err := strconv.ParseInt(coords[4], 10, 32)
		if err != nil {
			panic(err)
		}
		var xmod, ymod int
		var straight bool
		if x2 > x1 {
			xmod = 1
		} else if x2 == x1 {
			xmod = 0
		} else {
			xmod = -1
		}
		if y2 > y1 {
			ymod = 1
		} else if y2 == y1 {
			ymod = 0
		} else {
			ymod = -1
		}
		if xmod == 0 || ymod == 0 {
			straight = true
		} else {
			straight = false
		}

		retval = append(retval, tVent{
			x1:       int(x1),
			y1:       int(y1),
			x2:       int(x2),
			y2:       int(y2),
			xmod:     xmod,
			ymod:     ymod,
			straight: straight,
		})
	}
	return
}

func testInput() []string {
	return []string{
		"0,9 -> 5,9",
		"8,0 -> 0,8",
		"9,4 -> 3,4",
		"2,2 -> 2,1",
		"7,0 -> 7,4",
		"6,4 -> 2,0",
		"0,9 -> 2,9",
		"3,4 -> 1,4",
		"0,0 -> 8,8",
		"5,5 -> 8,2",
	}
}
