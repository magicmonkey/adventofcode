package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type tArea struct {
	min int
	max int
}

func NewArea(min, max int) (a *tArea) {
	a = &tArea{min: min, max: max}
	return
}

func (a1 *tArea) Contains(a2 *tArea) bool {
	if a1.min <= a2.min && a1.max >= a2.max {
		return true
	}
	return false
}

func (a1 *tArea) Overlaps(a2 *tArea) bool {
	if a1.max >= a2.min && a1.min <= a2.max {
		return true
	}
	return false
}

func main() {
	r := regexp.MustCompile(`^(\d+)-(\d+),(\d+)-(\d+)$`)
	scanner := bufio.NewScanner(os.Stdin)
	part1 := 0
	part2 := 0
	for scanner.Scan() {
		line := scanner.Text()
		matches := r.FindStringSubmatch(line)
		e1min, err := strconv.Atoi(matches[1])
		if err != nil {
			panic(err)
		}
		e1max, err := strconv.Atoi(matches[2])
		if err != nil {
			panic(err)
		}
		e2min, err := strconv.Atoi(matches[3])
		if err != nil {
			panic(err)
		}
		e2max, err := strconv.Atoi(matches[4])
		if err != nil {
			panic(err)
		}
		elf1 := NewArea(e1min, e1max)
		elf2 := NewArea(e2min, e2max)
		if elf1.Contains(elf2) || elf2.Contains(elf1) {
			part1 += 1
		}

		if elf1.Overlaps(elf2) {
			//fmt.Println(elf1, elf2)
			part2 += 1
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	fmt.Println("Part 1", part1)
	fmt.Println("Part 2", part2)

}
