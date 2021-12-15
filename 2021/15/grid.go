package main

import "fmt"

type tRisk int
type tGrid struct {
	val   map[int]map[int]tRisk
	sizeX int
	sizeY int
}

func (g *tGrid) set(x int, y int, val tRisk) {
	if x > g.sizeX {
		g.sizeX = x
	}
	if y > g.sizeY {
		g.sizeY = y
	}
	if _, ok := g.val[y]; !ok {
		g.val[y] = make(map[int]tRisk)
	}
	g.val[y][x] = val
}

func (g *tGrid) get(x int, y int) tRisk {
	return g.val[y][x]
}

func (g tGrid) print() {
	for x := 0; x <= g.sizeX; x++ {
		for y := 0; y <= g.sizeY; y++ {
			fmt.Printf("%d ", g.val[x][y])
		}
		fmt.Println("")
	}
	fmt.Println("")
}

func NewGrid() (retval tGrid) {
	retval = tGrid{}
	retval.val = make(map[int]map[int]tRisk)
	return
}
