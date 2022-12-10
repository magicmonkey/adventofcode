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
	case 'X': // Need to lose
		switch r.col1 {
		case 'A': // Opponent chose rock, I choose scissors
			score += 3 + 0
		case 'B': // Opponent chose paper, I choose rock
			score += 1 + 0
		case 'C': // Opponent chose scissors, I choose paper
			score += 2 + 0
		}
	case 'Y': // Need to draw
		switch r.col1 {
		case 'A': // Opponent chose rock, I choose rock
			score += 1 + 3
		case 'B': // Opponent chose paper, I choose paper
			score += 2 + 3
		case 'C': // Opponent chose scissors, I choose scissors
			score += 3 + 3
		}
	case 'Z': // Need to win
		switch r.col1 {
		case 'A': // Opponent chose rock, I choose paper
			score += 2 + 6
		case 'B': // Opponent chose paper, I choose scissors
			score += 3 + 6
		case 'C': // Opponent chose scissors, I choose rock
			score += 1 + 6
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
		//fmt.Println(r.col1, r.col2, r.Score())
		total += r.Score()
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	fmt.Println("Part 2", total)

}
