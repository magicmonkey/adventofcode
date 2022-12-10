package main

import (
	"bufio"
	"fmt"
	"os"
)

type tRucksack struct {
	items map[byte]int
}

var rucksacks []*tRucksack

func NewRucksack(contents []byte) (r *tRucksack) {
	r = &tRucksack{}
	r.items = make(map[byte]int)
	for _, item := range contents {
		r.items[item] += 1
	}
	return
}

func WhatIsInAll(r1, r2, r3 *tRucksack) byte {
	for item, _ := range r1.items {
		_, ok := r2.items[item]
		if ok {
			_, ok := r3.items[item]
			if ok {
				return item
			}
		}
	}
	return 0
}

func score(item byte) int {
	s := 0
	if item <= 'Z' {
		s = int(item - 'A' + 27)
	} else {
		s = int(item - 'a' + 1)
	}
	return s
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	total := 0
	for scanner.Scan() {
		line := scanner.Bytes()
		r := NewRucksack(line)
		rucksacks = append(rucksacks, r)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	for i := 0; i < len(rucksacks); i += 3 {
		item := WhatIsInAll(rucksacks[i], rucksacks[i+1], rucksacks[i+2])
		//fmt.Println(string(item))
		total += score(item)
	}

	fmt.Println("Part 2", total)
}
