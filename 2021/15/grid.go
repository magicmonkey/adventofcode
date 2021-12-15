package main

import "fmt"

type tRisk int64
type tGrid struct {
	val   map[int]map[int]tRisk
	sizeX int
	sizeY int
}

func (g *tGrid) set(x int, y int, val tRisk) {
	if x >= g.sizeX {
		g.sizeX = x + 1
	}
	if y >= g.sizeY {
		g.sizeY = y + 1
	}
	if _, ok := g.val[y]; !ok {
		g.val[y] = make(map[int]tRisk)
	}
	g.val[y][x] = val
}

func (g tGrid) get(x int, y int) int64 {
	return int64(g.val[y][x])
}

func (g tGrid) print() {
	for x := 0; x < g.sizeX; x++ {
		for y := 0; y < g.sizeY; y++ {
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

func (g tGrid) replicate() (retval tGrid) {
	retval = NewGrid()

	for x := 0; x < g.sizeX; x++ {
		for y := 0; y < g.sizeY; y++ {
			//fmt.Println("Processing", x, y)
			for modX := 0; modX <= 4; modX++ {
				for modY := 0; modY <= 4; modY++ {
					newVal := (tRisk)(g.get(x, y) + int64(modX+modY))
					if newVal > 9 {
						newVal -= 9
					}
					newX := (modX * g.sizeX) + x
					newY := (modY * g.sizeY) + y
					//fmt.Printf("Setting (%d,%d) == %d\n", newX, newY, newVal)
					retval.set(newX, newY, newVal)
				}
			}
		}
	}
	return
}
