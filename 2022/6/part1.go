package main

import (
	"bufio"
	"fmt"
	"os"
)

func FindStart(line []byte, uniqueSize int) (pos int) {
	for i := 0; i < len(line); i++ {
		if i < uniqueSize-1 {
			continue
		}
		buffer := make(map[byte]bool)
		for j := 0; j < uniqueSize; j++ {
			ok, _ := buffer[line[i-j]]
			if !ok {
				buffer[line[i-j]] = true
			} else {
				goto collision
			}
		}
		pos = i + 1
		return
	collision:
	}
	return
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var line []byte
	for scanner.Scan() {
		line = scanner.Bytes()
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	fmt.Println("Part 1", FindStart(line, 4))
	fmt.Println("Part 2", FindStart(line, 14))
}
