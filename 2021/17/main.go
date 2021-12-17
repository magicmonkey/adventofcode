package main

import (
	"fmt"
	//"github.com/magicmonkey/adventofcode/2021/util"
	"math"
	"regexp"
	"strconv"
)

func main() {
	//lines := util.ReadInputFile()
	lines := testInput()
	fmt.Println("Part 1")
	part1(lines)
	//fmt.Println("Part 2")
	//part2(lines)
}

var area struct {
	x1, y1, x2, y2 int
}

var currX, currY int

func part1(lines []string) {
	parseInput(lines)

	var maxHeight int = -1 * math.MaxInt

	minY := area.y1
	if area.y2 < minY {
		minY = area.y2
	}

	minX := area.x1
	if area.x2 < minX {
		minX = area.x2
	}

	for tryX := 1; tryX < minX; tryX++ {
		for tryY := 0; tryY < -1*minY; tryY++ {
			success, reached := try(tryX, tryY)
			if success {
				if reached > maxHeight {
					maxHeight = reached
				}
			}
		}
	}

	fmt.Println(maxHeight)

}

func try(velX int, velY int) (success bool, maxHeight int) {
	currX = 0
	currY = 0
	maxHeight = -1 * math.MaxInt
	var hit bool
	for !hit {
		hit, velX, velY = step(velX, velY)
		if currY > maxHeight {
			maxHeight = currY
		}

		// Break conditions
		if velX == 0 && currX < area.x1 {
			return false, 0
		}
		if currX > area.x2 {
			return false, 0
		}
		if currY < area.y1 {
			return false, 0
		}
	}

	return true, maxHeight
}

func step(startVelX int, startVelY int) (hit bool, velX int, velY int) {
	velX = startVelX
	velY = startVelY

	currX += velX
	currY += velY
	if velX > 0 {
		velX--
	}
	if velX < 0 {
		velX++
	}
	velY--

	if currX < area.x1 && currX < area.x2 {
		return false, velX, velY
	}
	if currX > area.x1 && currX > area.x2 {
		return false, velX, velY
	}
	if currY < area.y1 && currY < area.y2 {
		return false, velX, velY
	}
	if currY > area.y1 && currY > area.y2 {
		return false, velX, velY
	}
	return true, velX, velY
}

func part2(lines []string) {
}

func parseInput(lines []string) {
	re := regexp.MustCompile(`^target area: x=([-\d]*)\.\.([-\d]*), y=([-\d]*)\.\.([-\d]*)$`)
	parts := re.FindStringSubmatch(lines[0])
	var tempInt int64
	tempInt, _ = strconv.ParseInt(parts[1], 10, 32)
	area.x1 = int(tempInt)
	tempInt, _ = strconv.ParseInt(parts[2], 10, 32)
	area.x2 = int(tempInt)
	tempInt, _ = strconv.ParseInt(parts[3], 10, 32)
	area.y1 = int(tempInt)
	tempInt, _ = strconv.ParseInt(parts[4], 10, 32)
	area.y2 = int(tempInt)
}

func testInput() []string {
	return []string{"target area: x=20..30, y=-10..-5"}
}
