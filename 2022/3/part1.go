package main

import (
	"bufio"
	"fmt"
	"os"
)

type tRucksack struct {
	c1 map[byte]int
	c2 map[byte]int
}

func NewRucksack(contents []byte) (r *tRucksack) {
	r = &tRucksack{}
	r.c1 = make(map[byte]int)
	r.c2 = make(map[byte]int)
	for _, item := range contents[0 : len(contents)/2] {
		r.c1[item] += 1
	}
	for _, item := range contents[len(contents)/2:] {
		r.c2[item] += 1
	}
	return
}

func (r *tRucksack) WhatIsInBoth() byte {
	for item, _ := range r.c1 {
		_, ok := r.c2[item]
		if ok {
			return item
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
		item := r.WhatIsInBoth()
		//fmt.Println(string(item), score(item))
		total += score(item)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	fmt.Println("Part 1", total)
}
