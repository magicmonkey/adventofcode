package main

import (
	"fmt"
	"strconv"

	"github.com/RyanCarrier/dijkstra"
	"github.com/magicmonkey/adventofcode/2021/util"
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
	grid := parseInput(lines)
	grid.route()
}

func part2(lines []string) {
	startGrid := parseInput(lines)
	grid := startGrid.replicate()
	grid.route()
}
func (g *tGrid) route() {
	graph := dijkstra.NewGraph()

	// Add the vertices
	for x := 0; x < g.sizeX; x++ {
		for y := 0; y < g.sizeY; y++ {
			graph.AddVertex((y * g.sizeY) + x)
		}
	}

	// Add the weights/distances
	var srcID, destID int
	for x := 0; x < g.sizeX; x++ {
		for y := 0; y < g.sizeY; y++ {
			srcID = (y * g.sizeY) + x
			if x == 0 {
				destID = srcID + 1
				graph.AddArc(srcID, destID, g.get(x+1, y))
			} else if x == g.sizeX-1 {
				destID = srcID - 1
				graph.AddArc(srcID, destID, g.get(x-1, y))
			} else {
				destID = srcID + 1
				graph.AddArc(srcID, destID, g.get(x+1, y))
				destID = srcID - 1
				graph.AddArc(srcID, destID, g.get(x-1, y))
			}
			if y == 0 {
				destID = srcID + g.sizeY
				graph.AddArc(srcID, destID, g.get(x, y+1))
			} else if y == g.sizeY-1 {
				destID = srcID - g.sizeY
				graph.AddArc(srcID, destID, g.get(x, y-1))
			} else {
				destID = srcID + g.sizeY
				graph.AddArc(srcID, destID, g.get(x, y+1))
				destID = srcID - g.sizeY
				graph.AddArc(srcID, destID, g.get(x, y-1))
			}
		}
	}

	path, err := graph.Shortest(0, (g.sizeX*g.sizeY)-1)
	if err != nil {
		panic(err)
	}

	var score int64
	for _, p := range path.Path[1:] {
		x := int(p % g.sizeY)
		y := int(p / g.sizeY)
		score += g.get(x, y)
	}
	fmt.Println(score)
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
