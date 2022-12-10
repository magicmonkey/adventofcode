package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var hx, hy, tx, ty int

func moveRight() {
	hx += 1
}

func moveLeft() {
	hx -= 1
}

func moveUp() {
	hy += 1
}

func moveDown() {
	hy -= 1
}

func updateTail() {
	/*
		. . . . .
		. . . . .
		. . H . .
		. . . . .
		. . . . .
	*/
	if tx == hx && ty == hy {
		return
	}
	/*
		. . . . .
		. . . . .
		. . H T .
		. . . . .
		. . . . .
	*/
	if tx-1 == hx && ty == hy {
		return
	}
	/*
		. . . . .
		. . . . .
		. . H . .
		. . . T .
		. . . . .
	*/
	if tx-1 == hx && ty+1 == hy {
		return
	}
	/*
		. . . . .
		. . . . .
		. . H . .
		. . T . .
		. . . . .
	*/
	if tx == hx && ty+1 == hy {
		return
	}
	/*
		. . . . .
		. . . . .
		. . H . .
		. T . . .
		. . . . .
	*/
	if tx+1 == hx && ty+1 == hy {
		return
	}
	/*
		. . . . .
		. . . . .
		. T H . .
		. . . . .
		. . . . .
	*/
	if tx+1 == hx && ty == hy {
		return
	}
	/*
		. . . . .
		. T . . .
		. . H . .
		. . . . .
		. . . . .
	*/
	if tx+1 == hx && ty-1 == hy {
		return
	}
	/*
		. . . . .
		. . T . .
		. . H . .
		. . . . .
		. . . . .
	*/
	if tx == hx && ty-1 == hy {
		return
	}
	/*
		. . . . .
		. . . T .
		. . H . .
		. . . . .
		. . . . .
	*/
	if tx-1 == hx && ty-1 == hy {
		return
	}
	//// Everything else needs the tail to move
	/*
		. T . . .
		. . . . .
		. . H . .
		. . . . .
		. . . . .
	*/
	if tx+1 == hx && ty-2 == hy {
		tx += 1
		ty -= 1
		return
	}
	/*
		. . T . .
		. . . . .
		. . H . .
		. . . . .
		. . . . .
	*/
	if tx == hx && ty-2 == hy {
		ty -= 1
		return
	}
	/*
		. . . T .
		. . . . .
		. . H . .
		. . . . .
		. . . . .
	*/
	if tx-1 == hx && ty-2 == hy {
		tx -= 1
		ty -= 1
		return
	}
	/*
		. . . . .
		. . . . T
		. . H . .
		. . . . .
		. . . . .
	*/
	if tx-2 == hx && ty-1 == hy {
		tx -= 1
		ty -= 1
		return
	}
	/*
		. . . . .
		. . . . .
		. . H . T
		. . . . .
		. . . . .
	*/
	if tx-2 == hx && ty == hy {
		tx -= 1
		return
	}
	/*
		. . . . .
		. . . . .
		. . H . .
		. . . . T
		. . . . .
	*/
	if tx-2 == hx && ty+1 == hy {
		tx -= 1
		ty += 1
		return
	}
	/*
		. . . . .
		. . . . .
		. . H . .
		. . . . .
		. . . T .
	*/
	if tx-1 == hx && ty+2 == hy {
		tx -= 1
		ty += 1
		return
	}
	/*
		. . . . .
		. . . . .
		. . H . .
		. . . . .
		. . T . .
	*/
	if tx == hx && ty+2 == hy {
		ty += 1
		return
	}
	/*
		. . . . .
		. . . . .
		. . H . .
		. . . . .
		. T . . .
	*/
	if tx+1 == hx && ty+2 == hy {
		tx += 1
		ty += 1
		return
	}
	/*
		. . . . .
		. . . . .
		. . H . .
		T . . . .
		. . . . .
	*/
	if tx+2 == hx && ty+1 == hy {
		tx += 1
		ty += 1
		return
	}
	/*
		. . . . .
		. . . . .
		T . H . .
		. . . . .
		. . . . .
	*/
	if tx+2 == hx && ty == hy {
		tx += 1
		return
	}
	/*
		. . . . .
		T . . . .
		. . H . .
		. . . . .
		. . . . .
	*/
	if tx+2 == hx && ty-1 == hy {
		tx += 1
		ty -= 1
		return
	}
}

var pos map[string]bool
var numPos int

func countTailPos() {
	//fmt.Printf("Head is at (%d,%d) - Tail is at (%d,%d)... ", hx, hy, tx, ty)
	coord := fmt.Sprintf("(%d,%d)", tx, ty)
	_, ok := pos[coord]
	if !ok {
		pos[coord] = true
		numPos += 1
		//fmt.Println("New")
	} else {
		//fmt.Println("Old")
	}
}

func main() {
	pos = make(map[string]bool)
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Bytes()
		direction := line[0]
		amount, err := strconv.Atoi(string(line[2:]))
		if err != nil {
			panic(err)
		}
		switch direction {
		case 'R':
			for i := 0; i < amount; i++ {
				moveRight()
				updateTail()
				countTailPos()
			}
		case 'L':
			for i := 0; i < amount; i++ {
				moveLeft()
				updateTail()
				countTailPos()
			}
		case 'U':
			for i := 0; i < amount; i++ {
				moveUp()
				updateTail()
				countTailPos()
			}
		case 'D':
			for i := 0; i < amount; i++ {
				moveDown()
				updateTail()
				countTailPos()
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	fmt.Println(numPos)

}
