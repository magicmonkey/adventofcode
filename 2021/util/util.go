package util

import (
	"bufio"
	"os"
	"strconv"
)

func ReadInputFile() []string {
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

func StringsToInts(lines []string) (linesNum []int64) {
	for _, line := range lines {
		lineNum, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			panic(err)
		}
		linesNum = append(linesNum, lineNum)
	}
	return
}
