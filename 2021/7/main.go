package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

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
	posns, min, max := parsePosns(lines[0])

	var bestPos int
	var bestFuel int = math.MaxInt32
	for i := min; i <= max; i++ {
		f := totalFuelLinear(i, posns)
		if f < bestFuel {
			bestFuel = f
			bestPos = i
		}
	}
	fmt.Println(bestFuel, bestPos)
}

func part2(lines []string) {
	posns, min, max := parsePosns(lines[0])

	var bestPos int
	var bestFuel int = math.MaxInt32

	// Generate the lookup table (lut)
	var lut []int
	var acc int
	for i := 0; i <= max; i++ {
		acc += i
		lut = append(lut, acc)
	}

	for i := min; i <= max; i++ {
		f := totalFuelUsingLookup(i, posns, lut)
		if f < bestFuel {
			bestFuel = f
			bestPos = i
		}
	}

	fmt.Println(bestFuel, bestPos)
}

func totalFuelLinear(currPos int, posns []int) (retval int) {
	for _, posn := range posns {
		retval += absInt(posn - currPos)
	}
	return
}

func totalFuelUsingLookup(currPos int, posns []int, lut []int) (retval int) {
	for _, posn := range posns {
		retval += lut[absInt(posn-currPos)]
	}
	return
}

func parsePosns(line string) (retval []int, min int, max int) {
	posnsStr := strings.Split(line, ",")
	for _, posnStr := range posnsStr {
		posnInt, err := strconv.ParseInt(posnStr, 10, 32)
		if err != nil {
			panic(err)
		}
		retval = append(retval, int(posnInt))
	}

	max = -1
	min = math.MaxInt32
	for _, posn := range retval {
		if min > posn {
			min = posn
		}
		if max < posn {
			max = posn
		}
	}
	return
}

func absInt(i int) (retval int) {
	if i >= 0 {
		retval = i
	} else {
		retval = -1 * i
	}
	return
}

func testInput() []string {
	return []string{"16,1,2,0,4,2,7,1,2,14"}
}
