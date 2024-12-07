package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func MustInt(s string) int {
	numberInt, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return numberInt
}

type TLine struct {
	total    int
	operands []int
}

func NewTLine(txt string) TLine {

	re := regexp.MustCompile(`^(\d+): (.+)$`)
	matches := re.FindStringSubmatch(txt)

	if matches == nil {
		panic("unparseable")
	}

	fields := strings.Fields(matches[2])
	var operands []int
	for _, field := range fields {
		operands = append(operands, MustInt(field))
	}

	return TLine{
		total:    MustInt(matches[1]),
		operands: operands,
	}

}

func (line TLine) Total() int {
	return line.total
}

func (line TLine) Valid() bool {
	for i := 0; i < int(math.Pow(2, float64(len(line.operands)))); i++ {
		total := line.operands[0]
		for j := 1; j < len(line.operands); j++ {
			if (i & (1 << j)) > 0 {
				total += line.operands[j]
			} else {
				total *= line.operands[j]
			}
		}
		if total == line.total {
			return true
		}
	}
	return false
}

func ParseFile(fname string) (lines []TLine) {
	// Open the input file
	file, err := os.Open(fname)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	// Create scanner
	scanner := bufio.NewScanner(file)

	// Read and process lines
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, NewTLine(line))
	}

	// Check for scanning error
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	return
}

func part1(fname string) {
	lines := ParseFile(fname)
	total := 0
	for _, line := range lines {
		if line.Valid() {
			total += line.Total()
		}
	}
	fmt.Println(total)
}

func main() {
	part1("input.txt")
}
