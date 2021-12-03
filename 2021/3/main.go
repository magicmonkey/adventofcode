package main

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/mitchellh/copystructure"
)

func main() {
	//lines := util.ReadInputFile()
	lines := testInput()
	fmt.Println("Part 1")
	part1(lines)
	fmt.Println("Part 2")
	part2(lines)
}

func part1(lines []string) {
	var gamma, epsilon int
	for i, _ := range lines[0] {
		var column []byte
		for _, line := range lines {
			column = append(column, line[i])
		}
		numZero := bytes.Count(column, []byte{'0'})
		gamma *= 2
		epsilon *= 2
		if numZero < (len(lines) / 2) {
			gamma++
		}
		if numZero > (len(lines) / 2) {
			epsilon++
		}
	}
	fmt.Println(gamma, epsilon, gamma*epsilon)
}

func part2(lines []string) {
	numLines := len(lines)
	inputLen := len(lines[0])

	var inputChars [][]byte
	var oxy, coo int

	inputChars = make([][]byte, numLines, numLines)

	for i, line := range lines {
		inputChars[i] = make([]byte, inputLen, inputLen)
		chars := strings.Split(line, "")
		for j, char := range chars {
			if char[0] == '0' {
				inputChars[i][j] = 0
			} else {
				inputChars[i][j] = 1
			}
		}
	}

	// Oxygen generator
	var tempInput [][]byte
	t, err := copystructure.Copy(inputChars)
	if err != nil {
		panic(err)
	}
	tempInput = t.([][]byte)

	for posCounter := 0; posCounter < inputLen; posCounter++ {
		var numZero, numOne int
		var nextInput [][]byte
		for line := 0; line < len(tempInput); line++ {
			if tempInput[line][posCounter] == 0 {
				numZero++
			} else if tempInput[line][posCounter] == 1 {
				numOne++
			}
		}

		if numZero > numOne {
			// Copy lines with zero at posCounter
			for i, _ := range tempInput {
				if tempInput[i][posCounter] == 0 {
					ti, err := copystructure.Copy(tempInput[i])
					if err != nil {
						panic(err)
					}
					nextInput = append(nextInput, ti.([]byte))
				}
			}
		} else {
			// Copy lines with one at posCounter
			for i, _ := range tempInput {
				if tempInput[i][posCounter] == 1 {
					ti, err := copystructure.Copy(tempInput[i])
					if err != nil {
						panic(err)
					}
					nextInput = append(nextInput, ti.([]byte))
				}
			}
		}

		if len(nextInput) == 1 {
			fmt.Println(nextInput)
			oxy = binaryToInt(nextInput[0])
			fmt.Println(oxy)
			break
		}

		// seed tempInput for the next time round the loop
		tempInput = nextInput
	}

	// CO2 scrubber
	t, err = copystructure.Copy(inputChars)
	if err != nil {
		panic(err)
	}
	tempInput = t.([][]byte)

	for posCounter := 0; posCounter < inputLen; posCounter++ {
		var numZero, numOne int
		var nextInput [][]byte
		for line := 0; line < len(tempInput); line++ {
			if tempInput[line][posCounter] == 0 {
				numZero++
			} else if tempInput[line][posCounter] == 1 {
				numOne++
			}
		}

		if numZero > numOne {
			// Copy lines with one at posCounter
			for i, _ := range tempInput {
				if tempInput[i][posCounter] == 1 {
					ti, err := copystructure.Copy(tempInput[i])
					if err != nil {
						panic(err)
					}
					nextInput = append(nextInput, ti.([]byte))
				}
			}
		} else {
			// Copy lines with zero at posCounter
			for i, _ := range tempInput {
				if tempInput[i][posCounter] == 0 {
					ti, err := copystructure.Copy(tempInput[i])
					if err != nil {
						panic(err)
					}
					nextInput = append(nextInput, ti.([]byte))
				}
			}
		}

		if len(nextInput) == 1 {
			fmt.Println(nextInput)
			coo = binaryToInt(nextInput[0])
			fmt.Println(coo)
			break
		}

		// seed tempInput for the next time round the loop
		tempInput = nextInput
	}

	fmt.Println(oxy * coo)

}

func binaryToInt(inp []byte) (retval int) {
	for _, in := range inp {
		retval *= 2
		retval += int(in)
	}
	return
}

func testInput() []string {
	return []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}
}
