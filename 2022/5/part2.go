package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type tStacks map[int][]byte
type tInstruction struct {
	num  int
	from int
	to   int
}

func (stacks tStacks) Do(instr tInstruction) {
	stacks[instr.to] = append(stacks[instr.to], stacks[instr.from][len(stacks[instr.from])-instr.num:]...)
	stacks[instr.from] = stacks[instr.from][:len(stacks[instr.from])-instr.num]
}

var instructions []tInstruction

func main() {
	stacks := tStacks{}
	r1 := regexp.MustCompile(`^ 1`)
	var stacklines [][]byte
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Bytes()
		matches := r1.Match(line)
		if matches {
			break
		}
		stacklines = append(stacklines, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	numCols := (len(stacklines[len(stacklines)-1]) + 1) / 4
	numLines := len(stacklines)
	for i := numLines - 1; i >= 0; i-- {
		for c := 0; c < numCols; c++ {
			box := stacklines[i][(c*4)+1]
			if box != 32 {
				stacks[c] = append(stacks[c], box)
			}
		}
	}

	// Now we have the stacks of boxes

	//fmt.Println(stacks)

	// Parse the instructions

	r2 := regexp.MustCompile(`^move (\d+) from (\d+) to (\d+)$`)
	for scanner.Scan() {
		line := scanner.Text()
		matches := r2.FindStringSubmatch(line)
		if len(matches) > 0 {
			i := tInstruction{}
			val, err := strconv.Atoi(matches[1])
			if err != nil {
				panic(err)
			}
			i.num = val
			val, err = strconv.Atoi(matches[2])
			if err != nil {
				panic(err)
			}
			i.from = val - 1
			val, err = strconv.Atoi(matches[3])
			if err != nil {
				panic(err)
			}
			i.to = val - 1
			instructions = append(instructions, i)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	// Now run the process

	for _, instruction := range instructions {
		stacks.Do(instruction)
	}

	// Print the results

	for i := 0; i < len(stacks); i++ {
		fmt.Printf("%s", string(stacks[i][len(stacks[i])-1]))
	}
	fmt.Println("")
}
