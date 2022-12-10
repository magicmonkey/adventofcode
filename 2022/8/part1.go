package main

import (
	"bufio"
	"fmt"
	"os"
)

var trees [][]byte

func CheckLeft(x, y int) bool {
	myHeight := trees[x][y]
	for i := 0; i < x; i++ {
		//fmt.Printf("L (%d,%d) [%d]...\n", i, y, trees[i][y])
		if myHeight <= trees[i][y] {
			//fmt.Printf("(%d,%d) [%d] is blocked from the left by (%d,%d) [%d]\n", x, y, myHeight, i, y, trees[i][y])
			return false
		}
	}
	return true
}

func CheckRight(x, y int) bool {
	myHeight := trees[x][y]
	for i := x + 1; i < len(trees); i++ {
		//fmt.Printf("R (%d,%d) [%d]...\n", i, y, trees[i][y])
		if myHeight <= trees[i][y] {
			//fmt.Printf("(%d,%d) [%d] is blocked from the right by (%d,%d) [%d]\n", x, y, myHeight, i, y, trees[i][y])
			return false
		}
	}
	return true
}

func CheckUp(x, y int) bool {
	myHeight := trees[x][y]
	for i := 0; i < y; i++ {
		//fmt.Printf("U (%d,%d) [%d]...\n", x, i, trees[x][i])
		if myHeight <= trees[x][i] {
			//fmt.Printf("(%d,%d) [%d] is blocked from the up by (%d,%d) [%d]\n", x, y, myHeight, x, i, trees[x][i])
			return false
		}
	}
	return true
}

func CheckDown(x, y int) bool {
	myHeight := trees[x][y]
	for i := y + 1; i < len(trees[0]); i++ {
		//fmt.Printf("D (%d,%d) [%d]...\n", x, i, trees[x][i])
		if myHeight <= trees[x][i] {
			//fmt.Printf("(%d,%d) [%d] is blocked from the down by (%d,%d) [%d]\n", x, y, myHeight, x, i, trees[x][i])
			return false
		}
	}
	return true
}

func CheckVisible(x, y int) bool {
	fmt.Printf("Checking (%d,%d) [%d]... ", x, y, trees[x][y])
	if x == 0 || y == 0 || x == len(trees[0])-1 || y == len(trees)-1 {
		fmt.Println("Visible from edge")
		return true
	}
	if CheckLeft(x, y) {
		fmt.Println("Visible from left")
		return true
	}
	if CheckRight(x, y) {
		fmt.Println("Visible from right")
		return true
	}
	if CheckUp(x, y) {
		fmt.Println("Visible from up")
		return true
	}
	if CheckDown(x, y) {
		fmt.Println("Visible from down")
		return true
	}
	fmt.Println("Not visible")
	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Bytes()
		var treeRow []byte
		for _, b := range line {
			treeRow = append(treeRow, b-'0')
		}
		trees = append(trees, treeRow)
	}

	numVisible := 0
	for x := 0; x < len(trees); x++ {
		for y := 0; y < len(trees[0]); y++ {
			if CheckVisible(x, y) {
				numVisible += 1
			}
		}
	}
	fmt.Println("Part 1", numVisible)
}
