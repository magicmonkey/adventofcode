package main

import (
	"bufio"
	"fmt"
	"os"
)

func ParseFile(fname string) (maze map[int]map[int]bool, pos [2]int) {

	maze = make(map[int]map[int]bool)

	// Open the input file
	file, err := os.Open(fname)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	// Create scanner
	scanner := bufio.NewScanner(file)

	// Read and process lines
	var y, x int
	for scanner.Scan() {
		line := scanner.Text()
		x = 0
		for _, b := range line {
			if _, ok := maze[x]; !ok {
				maze[x] = make(map[int]bool)
			}
			switch b {
			case '#':
				maze[x][y] = true
			case '.':
				maze[x][y] = false
			case '^':
				pos[0] = x
				pos[1] = y
				maze[x][y] = false
			default:
				fmt.Printf("Unknown value: %v\n", b)
			}
			x++
		}
		y++
	}

	// Check for scanning error
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	return
}

func takeStep(maze map[int]map[int]bool, pos [2]int, direction byte) (nextPos [2]int, nextDirection byte) {
	x := pos[0]
	y := pos[1]
	switch direction {
	case 'N':
		if maze[x][y-1] {
			nextDirection = 'E'
			nextPos[0] = x + 1
			nextPos[1] = y
		} else {
			nextDirection = 'N'
			nextPos[0] = x
			nextPos[1] = y - 1
		}
	case 'E':
		if maze[x+1][y] {
			nextDirection = 'S'
			nextPos[0] = x
			nextPos[1] = y + 1
		} else {
			nextDirection = 'E'
			nextPos[0] = x + 1
			nextPos[1] = y
		}
	case 'S':
		if maze[x][y+1] {
			nextDirection = 'W'
			nextPos[0] = x - 1
			nextPos[1] = y
		} else {
			nextDirection = 'S'
			nextPos[0] = x
			nextPos[1] = y + 1
		}
	case 'W':
		if maze[x-1][y] {
			nextDirection = 'N'
			nextPos[0] = x
			nextPos[1] = y - 1
		} else {
			nextDirection = 'W'
			nextPos[0] = x - 1
			nextPos[1] = y
		}
	}
	return
}

func part1(fname string) {
	maze, pos := ParseFile(fname)

	visited := make(map[int]map[int]bool)
	var direction byte = 'N'

	total := 0

	for {
		nextPos, nextDirection := takeStep(maze, pos, direction)
		if _, ok := maze[nextPos[0]][nextPos[1]]; !ok {
			break
		}
		if _, ok := visited[nextPos[0]]; !ok {
			visited[nextPos[0]] = make(map[int]bool)
		}
		if _, ok := visited[nextPos[0]][nextPos[1]]; !ok {
			total++
			visited[nextPos[0]][nextPos[1]] = true
		}
		pos = nextPos
		direction = nextDirection
	}

	// Count how many places have been visited
	fmt.Println(total)

}

func main() {
	part1("input.txt")
}
