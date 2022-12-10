package main

import (
	"bufio"
	"fmt"
	"os"
)

var trees [][]byte

func NumLeft(x, y int) (score int) {
	myHeight := trees[x][y]
	for i := x - 1; i >= 0; i-- {
		score += 1
		if myHeight <= trees[i][y] {
			return
		}
	}
	return
}

func NumRight(x, y int) (score int) {
	myHeight := trees[x][y]
	for i := x + 1; i < len(trees); i++ {
		score += 1
		if myHeight <= trees[i][y] {
			return
		}
	}
	return
}

func NumUp(x, y int) (score int) {
	myHeight := trees[x][y]
	for i := y - 1; i >= 0; i-- {
		score += 1
		if myHeight <= trees[x][i] {
			return
		}
	}
	return
}

func NumDown(x, y int) (score int) {
	myHeight := trees[x][y]
	for i := y + 1; i < len(trees[0]); i++ {
		score += 1
		if myHeight <= trees[x][i] {
			return
		}
	}
	return
}

func ScenicScore(x, y int) int {
	if x == 0 || y == 0 || x == len(trees[0])-1 || y == len(trees)-1 {
		return 0
	}
	return NumLeft(x, y) * NumRight(x, y) * NumUp(x, y) * NumDown(x, y)
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

	maxScore := 0
	for x := 0; x < len(trees); x++ {
		for y := 0; y < len(trees[0]); y++ {
			score := ScenicScore(x, y)
			if score > maxScore {
				maxScore = score
			}
		}
	}
	fmt.Println("Part 2", maxScore)
}
