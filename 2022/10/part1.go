package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var x []int // Stores the value of X *after* the given cycle number

func processNoop() {
	// Add 1 cycle at the current X value
	x = append(x, x[len(x)-1])
}

func processAddx(val int) {
	// Add 1 cycle at the current X value
	x = append(x, x[len(x)-1])
	// Add 1 cycle at the new X value
	x = append(x, x[len(x)-1]+val)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	x = []int{1}

	r1 := regexp.MustCompile(`^addx ([-\d]+)$`)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "noop" {
			processNoop()
		}
		matches := r1.FindStringSubmatch(line)
		if len(matches) > 0 {
			num, err := strconv.Atoi(matches[1])
			if err != nil {
				panic(err)
			}
			processAddx(num)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	score := 0
	for i := 20; i < len(x); i += 40 {
		score += i * x[i-1]
	}
	fmt.Println("Part 1", score)

	// Now run the CRT simulation

	for row := 0; row < 6; row++ {
		for i := 0; i < 40; i++ {
			sprite := x[i+40*row]
			if sprite >= i-1 && sprite <= i+1 {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println("")
	}

}
