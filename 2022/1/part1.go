package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type tElf struct {
	food []int
}

type tElves []*tElf

var elves tElves

func NewElf() *tElf {
	e := &tElf{}
	return e
}

func (e *tElf) AddFood(calory int) {
	e.food = append(e.food, calory)
}

func (e *tElf) Total() (total int) {
	total = 0
	for _, v := range e.food {
		total = total + v
	}
	return
}

func (a tElves) Len() int           { return len(a) }
func (a tElves) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a tElves) Less(i, j int) bool { return a[i].Total() < a[j].Total() }

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var e *tElf
	e = NewElf()
	for scanner.Scan() {
		t := scanner.Text()
		if t == "" {
			e = NewElf()
			elves = append(elves, e)
		} else {
			i, err := strconv.Atoi(t)
			if err != nil {
				panic(err)
			}
			e.AddFood(i)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	max := 0
	for _, elf := range elves {
		if elf.Total() > max {
			max = elf.Total()
		}
	}
	fmt.Println("Part 1", max)

	sort.Sort(sort.Reverse(elves))

	amount := 0
	for _, elf := range elves[0:3] {
		amount = amount + elf.Total()
	}
	fmt.Println("Part 2", amount)

}
