package main

import (
	"bytes"
	"fmt"
	"os"
)

func ParseFile(fname string) (chars [][]byte) {
	// Open the input file
	rawdata, err := os.ReadFile(fname)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	chars = bytes.Split(rawdata, []byte("\n"))

	return
}

func checkFrom(chars [][]byte, x, y int) (numMatch int) {
	// East
	if x < len(chars)-4 {
		if chars[y][x+1] == 'M' && chars[y][x+2] == 'A' && chars[y][x+3] == 'S' {
			numMatch++
		}
	}
	// South-east
	if x < len(chars)-4 && y < len(chars)-4 {
		if chars[y+1][x+1] == 'M' && chars[y+2][x+2] == 'A' && chars[y+3][x+3] == 'S' {
			numMatch++
		}
	}
	// South
	if y < len(chars)-4 {
		if chars[y+1][x] == 'M' && chars[y+2][x] == 'A' && chars[y+3][x] == 'S' {
			numMatch++
		}
	}
	// South-west
	if x > 2 && y < len(chars)-4 {
		if chars[y+1][x-1] == 'M' && chars[y+2][x-2] == 'A' && chars[y+3][x-3] == 'S' {
			numMatch++
		}
	}
	// West
	if x > 2 {
		if chars[y][x-1] == 'M' && chars[y][x-2] == 'A' && chars[y][x-3] == 'S' {
			numMatch++
		}
	}
	// North-west
	if x > 2 && y > 2 {
		if chars[y-1][x-1] == 'M' && chars[y-2][x-2] == 'A' && chars[y-3][x-3] == 'S' {
			numMatch++
		}
	}
	// North
	if y > 2 {
		if chars[y-1][x] == 'M' && chars[y-2][x] == 'A' && chars[y-3][x] == 'S' {
			numMatch++
		}
	}
	// North-east
	if x < len(chars)-4 && y > 2 {
		if chars[y-1][x+1] == 'M' && chars[y-2][x+2] == 'A' && chars[y-3][x+3] == 'S' {
			numMatch++
		}
	}
	return
}

func checkFrom2(chars [][]byte, x, y int) bool {
	if x == 0 || x >= len(chars[0])-1 || y == 0 || y >= len(chars)-2 {
		return false
	}
	// East
	if chars[y-1][x-1] == 'M' && chars[y+1][x-1] == 'M' && chars[y-1][x+1] == 'S' && chars[y+1][x+1] == 'S' {
		return true
	}
	// South
	if chars[y-1][x-1] == 'M' && chars[y-1][x+1] == 'M' && chars[y+1][x-1] == 'S' && chars[y+1][x+1] == 'S' {
		return true
	}
	// West
	if chars[y+1][x+1] == 'M' && chars[y-1][x+1] == 'M' && chars[y+1][x-1] == 'S' && chars[y-1][x-1] == 'S' {
		return true
	}
	// North
	if chars[y+1][x+1] == 'M' && chars[y+1][x-1] == 'M' && chars[y-1][x+1] == 'S' && chars[y-1][x-1] == 'S' {
		return true
	}
	return false
}

func part1(fname string) {
	chars := ParseFile(fname)

	total := 0
	for y, cy := range chars {
		for x, c := range cy {
			score := 0
			if c == 'X' {
				score = checkFrom(chars, x, y)
				total += score
			}
			//fmt.Printf("%d, %d = %v / %d\n", x, y, c, score)
		}
	}

	fmt.Println(total)
}

func part2(fname string) {
	chars := ParseFile(fname)

	total := 0
	for y, cy := range chars {
		for x, c := range cy {
			if c == 'A' {
				if checkFrom2(chars, x, y) {
					fmt.Printf("%d, %d\n", x, y)
					total += 1
				}
			}
		}
	}

	fmt.Println(total)
}

func main() {
	part2("input.txt")
}
