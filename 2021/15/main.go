package main

import (
	"fmt"
	"strconv"

	"github.com/magicmonkey/adventofcode/2021/util"
)

func main() {
	//lines := testInput()
	lines := util.ReadInputFile()
	fmt.Println("Part 1")
	part1(lines)
}

func part1(lines []string) {
	grid := parseInput(lines)
	//grid.print()

	risk := grid.takeStepsFrom(0, 0, 0)
	fmt.Println("Final risk:", risk)

}

func (g *tGrid) takeStepsFrom(posX int, posY int, risk tRisk) (finalRisk tRisk) {
	//fmt.Printf("Trying (%d,%d) == %d\n", posX, posY, g.get(posX, posY))

	var risk2 tRisk
	if posX == 0 && posY == 0 {
		risk2 = risk
	} else {
		risk2 = risk + g.get(posX, posY)
	}

	// if at the end, do nothing
	if posX == g.sizeX && posY == g.sizeY {
		//fmt.Println("At the end")
		finalRisk = risk2
		return
	}

	// if on the right wall, step down
	if posX == g.sizeX {
		//fmt.Println("Stepping down")
		finalRisk = g.takeStepsFrom(posX, posY+1, risk2)
		return
	}

	// if on the lower wall, step right
	if posY == g.sizeY {
		//fmt.Println("Stepping right")
		finalRisk = g.takeStepsFrom(posX+1, posY, risk2)
		return
	}

	// otherwise, try both ways
	//fmt.Println("Trying right")
	r1 := g.takeStepsFrom(posX+1, posY, risk2)

	//fmt.Println("Trying down")
	r2 := g.takeStepsFrom(posX, posY+1, risk2)

	if r1 < r2 {
		finalRisk = r1
	} else {
		finalRisk = r2
	}
	return
}

func parseInput(lines []string) (retval tGrid) {
	retval = NewGrid()
	for y, row := range lines {
		for x, valStr := range row {
			valInt, _ := strconv.ParseInt(string(valStr), 10, 32)
			retval.set(x, y, tRisk(valInt))
		}
	}
	return
}

func testInput() []string {
	return []string{
		"1163751742",
		"1381373672",
		"2136511328",
		"3694931569",
		"7463417111",
		"1319128137",
		"1359912421",
		"3125421639",
		"1293138521",
		"2311944581",
	}
}
