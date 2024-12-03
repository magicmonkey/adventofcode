package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func MustInt(s string) int {
	numberInt, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return numberInt
}

func ParseFile(fname string) (muls [][]int) {
	// Open the input file
	file, err := os.Open(fname)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	// Create the regex
	re := regexp.MustCompile(`(mul\((\d+),(\d+)\)|do\(\)|don't\(\))`)

	// Create scanner
	scanner := bufio.NewScanner(file)

	ok := true

	// Read and process lines
	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindAllStringSubmatch(line, -1)
		if matches != nil {
			for _, v := range matches {
				switch v[0] {
				case `do()`:
					ok = true
				case `don't()`:
					ok = false
				default:
					if ok {
						var thingToInsert []int
						thingToInsert = append(thingToInsert, MustInt(v[2]))
						thingToInsert = append(thingToInsert, MustInt(v[3]))
						muls = append(muls, thingToInsert)
					}
				}
			}
		}
	}

	// Check for scanning error
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	return
}

func part1(fname string) {
	muls := ParseFile(fname)
	count := 0
	for _, v := range muls {
		count += v[0] * v[1]
	}
	fmt.Println(count)
}

func main() {
	part1("input.txt")
}
