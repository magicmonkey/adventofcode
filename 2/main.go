package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/magicmonkey/adventofcode/2021/util"
)

func main() {
	lines := util.ReadInputFile()
	//lines := testInput()

	fmt.Println("Part 1")
	part1(lines)
	fmt.Println("Part 2")
	part2(lines)
}

func part1(lines []string) {
	lineParser := regexp.MustCompile(`^([a-z]*) (\d*)$`)
	var horiz int64 = 0
	var depth int64 = 0
	for _, line := range lines {
		parts := lineParser.FindStringSubmatch(line)
		amount, err := strconv.ParseInt(parts[2], 10, 32)
		if err != nil {
			panic(err)
		}
		switch parts[1] {
		case "forward":
			horiz += amount
		case "down":
			depth += amount
		case "up":
			depth -= amount
		default:
			panic("Unknown command")
		}
	}
	fmt.Printf("Horiz: %d\nDepth: %d\nResult: %d\n", horiz, depth, (horiz * depth))
}

func part2(lines []string) {
	lineParser := regexp.MustCompile(`^([a-z]*) (\d*)$`)
	var horiz, depth, aim int64
	for _, line := range lines {
		parts := lineParser.FindStringSubmatch(line)
		amount, err := strconv.ParseInt(parts[2], 10, 32)
		if err != nil {
			panic(err)
		}
		switch parts[1] {
		case "forward":
			horiz += amount
			depth += (aim * amount)
		case "down":
			aim += amount
		case "up":
			aim -= amount
		default:
			panic("Unknown command")
		}
	}
	fmt.Printf("Horiz: %d\nDepth: %d\nResult: %d\n", horiz, depth, (horiz * depth))
}

func testInput() (lines []string) {
	return []string{
		"forward 5",
		"down 5",
		"forward 8",
		"up 3",
		"down 8",
		"forward 2",
	}
}
