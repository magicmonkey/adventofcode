package main

import (
	"bufio"
	"fmt"
	"os"
)

type round struct {
	col1 byte
	col2 byte
}

func NewRound(col1, col2 byte) *round {
	r := &round{
		col1: col1,
		col2: col2,
	}
	return r
}

func (r *round) Score() (score int) {
	switch r.col2 {
	case 'X':
		score += 1
		switch r.col1 {
		case 'A':
			score += 3
		case 'B':
			score += 0
		case 'C':
			score += 6
		}
	case 'Y':
		score += 2
		switch r.col1 {
		case 'A':
			score += 6
		case 'B':
			score += 3
		case 'C':
			score += 0
		}
	case 'Z':
		score += 3
		switch r.col1 {
		case 'A':
			score += 0
		case 'B':
			score += 6
		case 'C':
			score += 3
		}
	}
	return
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	total := 0
	for scanner.Scan() {
		t := scanner.Text()
		r := NewRound(t[0], t[2])
		total += r.Score()
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	fmt.Println("Part 1", total)

}
