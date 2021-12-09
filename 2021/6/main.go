package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/magicmonkey/adventofcode/2021/util"
)

type tFish int64

func main() {
	//lines := testInput()
	lines := util.ReadInputFile()

	fmt.Println("Part 1")
	part1(lines)

	fmt.Println("Part 2")
	part2(lines)

}

// A count of fish at each age on a given day
type tDay [9]int64

func part2(lines []string) {
	var fishs []tFish

	var currDay tDay
	fishs = parseInput(lines[0])

	// Initial state
	for _, fish := range fishs {
		currDay[fish]++
	}

	for i := 1; i <= 256; i++ {
		currDay = iterateDay(currDay)
	}

	var numFish int64
	for _, fishAtAge := range currDay {
		numFish += fishAtAge
	}
	fmt.Println(numFish)
}

func iterateDay(day tDay) (retval tDay) {
	retval[0] = day[1]
	retval[1] = day[2]
	retval[2] = day[3]
	retval[3] = day[4]
	retval[4] = day[5]
	retval[5] = day[6]
	retval[6] = day[7] + day[0]
	retval[7] = day[8]
	retval[8] = day[0]
	return
}

func part1(lines []string) {
	var fishs []tFish
	fishs = parseInput(lines[0])
	for i := 1; i <= 80; i++ {
		fishs = iterateEachFish(fishs)
	}
	fmt.Println(len(fishs))
}

func iterateEachFish(fishs []tFish) (retval []tFish) {
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
