package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/magicmonkey/adventofcode/2021/util"
)

type tFish int64

var fishs []tFish

func main() {
	//lines := testInput()
	lines := util.ReadInputFile()
	fishs = parseInput(lines[0])

	fmt.Println("Part 1")
	part1(fishs, 80)

	fmt.Println("Part 2 - this ain't gonna complete")
	part1(fishs, 256)
}

func part1(fishs []tFish, numIterations int) {
	for i := 1; i <= numIterations; i++ {
		fishs = iterate(fishs)
	}
	fmt.Println(len(fishs))
}

func iterate(fishs []tFish) (retval []tFish) {
	var numSpawned int
	for _, fish := range fishs {
		fish--
		if fish < 0 {
			fish = 6
			numSpawned++
		}
		retval = append(retval, fish)
	}
	for i := 0; i < numSpawned; i++ {
		retval = append(retval, tFish(8))
	}
	return
}

func parseInput(line string) (retval []tFish) {
	fishsStr := strings.Split(line, ",")
	for _, fishStr := range fishsStr {
		fishInt, err := strconv.ParseInt(fishStr, 10, 32)
		if err != nil {
			panic(err)
		}
		retval = append(retval, tFish(fishInt))
	}
	return
}

func testInput() []string {
	return []string{
		"3,4,3,1,2",
	}
}
