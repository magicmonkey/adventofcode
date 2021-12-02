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
	var horiz, depth int
	for _, line := range lines {
		command, amount := partsFromLine(line)
		switch command {
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
	fmt.Printf("Horiz: %d\nDepth: %d\n=== Result: %d ===\n", horiz, depth, (horiz * depth))
}

func part2(lines []string) {
	var horiz, depth, aim int
	for _, line := range lines {
		command, amount := partsFromLine(line)
		switch command {
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
	fmt.Printf("Horiz: %d\nDepth: %d\nAim: %d\n=== Result: %d ===\n", horiz, depth, aim, (horiz * depth))
}

func partsFromLine(line string) (command string, amount int) {
	lineParser := regexp.MustCompile(`^([a-z]*) (\d*)$`)
	parts := lineParser.FindStringSubmatch(line)
	command = parts[1]
	amountTmp, err := strconv.ParseInt(parts[2], 10, 32)
	if err != nil {
		panic(err)
	}
	amount = int(amountTmp) // int64 to int conversion
	return
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
