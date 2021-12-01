package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"

	"github.com/magicmonkey/adventofcode/2021/util"
)

func main() {
	lines := util.ReadInputFile()

	fmt.Println("Part 1")
	part1(lines)
	fmt.Println("Part 2")
	part2(lines)
}

func part1(lines []string) {
	linesNum := util.StringsToInts(lines)
	var prevLine int64 = math.MaxInt64
	increaseCounter := 0
	for _, currLine := range linesNum {
		if currLine > prevLine {
			increaseCounter++
		}
		prevLine = currLine
	}
	fmt.Println(increaseCounter)
}

func part2(lines []string) {
	linesNum := stringsToInts(lines)
	var prevSum int64 = math.MaxInt64
	var currSum int64

	increaseCounter := 0

	for i, _ := range linesNum {
		if i >= (len(linesNum) - 2) {
			break
		}
		currSum = linesNum[i] + linesNum[i+1] + linesNum[i+2]
		if currSum > prevSum {
			increaseCounter++
		}
		prevSum = currSum
	}
	fmt.Println(increaseCounter)
}

func readInputFile() []string {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err = scanner.Err(); err != nil {
		panic(err)
	}
	return lines
}

func stringsToInts(lines []string) (linesNum []int64) {
	for _, line := range lines {
		lineNum, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			panic(err)
		}
		linesNum = append(linesNum, lineNum)
	}
	return
}
