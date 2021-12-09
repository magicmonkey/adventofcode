package main

import (
	"fmt"
	"github.com/magicmonkey/adventofcode/2021/util"
	"strconv"
	"strings"
)

func main() {
	//lines := testInput()
	lines := util.ReadInputFile()

	fmt.Println("Part 1")
	part1(lines)

	fmt.Println("Part 2")
	part2(lines)
}

type tDisplay struct {
	inputs  [10][]rune
	outputs [4][]rune
}

func part1(lines []string) {
	disps := parseDisplays(lines)

	var count int
	for _, disp := range disps {
		for _, out := range disp.outputs {
			if len(out) == 2 || len(out) == 4 || len(out) == 3 || len(out) == 7 {
				count++
			}
		}
	}
	fmt.Println(count)
}

func part2(lines []string) {
	disps := parseDisplays(lines)
	var sum int
	for _, disp := range disps {
		digits := decodeDisp(disp)
		sum += digits
	}
	fmt.Println(sum)
}

/*

 AAAA
B    C
B    C
 DDDD
E    F
E    F
 GGGG
*/

func decodeDisp(disp tDisplay) int {
	wiring := decodeWiring(disp.inputs)
	//fmt.Println("Wiring", wiring)

	numbers := decodeNumbers(wiring, disp.inputs)
	//fmt.Println("Numbers", numbers)

	var digitsStr string
	for _, output := range disp.outputs {
		for i, input := range disp.inputs {
			if compareUnsorted(output, input) {
				for j, number := range numbers {
					if number == i {
						digitsStr = digitsStr + strconv.Itoa(j)
					}
				}
			}
		}
	}
	digits, _ := strconv.ParseInt(digitsStr, 10, 32)
	return int(digits)
}

func compareUnsorted(a []rune, b []rune) bool {
	l := len(a)
	if l != len(b) {
		return false
	}
	for _, c := range a {
		if !isInSlice(c, b) {
			return false
		}
	}
	return true
}

func decodeNumbers(wiring map[rune]rune, inputs [10][]rune) (numbers [10]int) {
	for i, input := range inputs {
		// 0
		if isInSlice(wiring['A'], input) && isInSlice(wiring['B'], input) && isInSlice(wiring['C'], input) && !isInSlice(wiring['D'], input) && isInSlice(wiring['E'], input) && isInSlice(wiring['F'], input) && isInSlice(wiring['G'], input) {
			numbers[0] = i
		}
		// 1
		if !isInSlice(wiring['A'], input) && !isInSlice(wiring['B'], input) && isInSlice(wiring['C'], input) && !isInSlice(wiring['D'], input) && !isInSlice(wiring['E'], input) && isInSlice(wiring['F'], input) && !isInSlice(wiring['G'], input) {
			numbers[1] = i
		}
		// 2
		if isInSlice(wiring['A'], input) && !isInSlice(wiring['B'], input) && isInSlice(wiring['C'], input) && isInSlice(wiring['D'], input) && isInSlice(wiring['E'], input) && !isInSlice(wiring['F'], input) && isInSlice(wiring['G'], input) {
			numbers[2] = i
		}
		// 3
		if isInSlice(wiring['A'], input) && !isInSlice(wiring['B'], input) && isInSlice(wiring['C'], input) && isInSlice(wiring['D'], input) && !isInSlice(wiring['E'], input) && isInSlice(wiring['F'], input) && isInSlice(wiring['G'], input) {
			numbers[3] = i
		}
		// 4
		if !isInSlice(wiring['A'], input) && isInSlice(wiring['B'], input) && isInSlice(wiring['C'], input) && isInSlice(wiring['D'], input) && !isInSlice(wiring['E'], input) && isInSlice(wiring['F'], input) && !isInSlice(wiring['G'], input) {
			numbers[4] = i
		}
		// 5
		if isInSlice(wiring['A'], input) && isInSlice(wiring['B'], input) && !isInSlice(wiring['C'], input) && isInSlice(wiring['D'], input) && !isInSlice(wiring['E'], input) && isInSlice(wiring['F'], input) && isInSlice(wiring['G'], input) {
			numbers[5] = i
		}
		// 6
		if isInSlice(wiring['A'], input) && isInSlice(wiring['B'], input) && !isInSlice(wiring['C'], input) && isInSlice(wiring['D'], input) && isInSlice(wiring['E'], input) && isInSlice(wiring['F'], input) && isInSlice(wiring['G'], input) {
			numbers[6] = i
		}
		// 7
		if isInSlice(wiring['A'], input) && !isInSlice(wiring['B'], input) && isInSlice(wiring['C'], input) && !isInSlice(wiring['D'], input) && !isInSlice(wiring['E'], input) && isInSlice(wiring['F'], input) && !isInSlice(wiring['G'], input) {
			numbers[7] = i
		}
		// 8
		if isInSlice(wiring['A'], input) && isInSlice(wiring['B'], input) && isInSlice(wiring['C'], input) && isInSlice(wiring['D'], input) && isInSlice(wiring['E'], input) && isInSlice(wiring['F'], input) && isInSlice(wiring['G'], input) {
			numbers[8] = i
		}
		// 9
		if isInSlice(wiring['A'], input) && isInSlice(wiring['B'], input) && isInSlice(wiring['C'], input) && isInSlice(wiring['D'], input) && !isInSlice(wiring['E'], input) && isInSlice(wiring['F'], input) && isInSlice(wiring['G'], input) {
			numbers[9] = i
		}
	}
	return
}

func decodeWiring(inputs [10][]rune) (wiring map[rune]rune) {
	wiring = make(map[rune]rune)
	var numbers [10]int = [10]int{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1} // Mapping of numbers to entries in the inputs array
	var wires [7]rune = [7]rune{'a', 'b', 'c', 'd', 'e', 'f', 'g'}

	//fmt.Println("Inputs", inputs)

	// Find digit "1", which gives C,F
	var CorF []rune = make([]rune, 2, 2)
	for i, input := range inputs {
		if len(input) == 2 {
			numbers[1] = i
			CorF[0] = input[0]
			CorF[1] = input[1]
			break
		}
	}

	// Find digit "7", which gives A,C,F
	var AorCorF []rune = make([]rune, 3, 3)
	for i, input := range inputs {
		if len(input) == 3 {
			numbers[7] = i
			AorCorF[0] = input[0]
			AorCorF[1] = input[1]
			AorCorF[2] = input[2]
			break
		}
	}

	// Find the one that's in A,C,F and not in C,F which is A
	for _, c := range AorCorF {
		if !isInSlice(c, CorF) {
			// We've found A
			wiring['A'] = c
		}
	}

	// Find the one with 7 wires, which is 8
	for i, input := range inputs {
		if len(input) == 7 {
			numbers[8] = i
		}
	}

	// Find the one with 4 wires, which is 4
	for i, input := range inputs {
		if len(input) == 4 {
			numbers[4] = i
		}
	}

	// Find the wire which is not in 4 and is not A, which is E or G
	var EorG []rune
	for _, wire := range wires {
		if !isInSlice(wire, inputs[numbers[4]]) {
			if wiring['A'] != wire {
				EorG = append(EorG, wire)
			}
		}
	}

	// Find how many digits EorG is in, the max is G and the min is E
	var tempCount0, tempCount1 int
	for _, input := range inputs {
		if isInSlice(EorG[0], input) {
			tempCount0++
		}
		if isInSlice(EorG[1], input) {
			tempCount1++
		}
	}
	if tempCount0 == 4 {
		wiring['E'] = EorG[0]
		wiring['G'] = EorG[1]
	} else {
		wiring['E'] = EorG[1]
		wiring['G'] = EorG[0]
	}

	// Find a digit with everything wired except E, which is 9
	for i, input := range inputs {
		if len(input) == 6 {
			if !isInSlice(wiring['E'], input) {
				numbers[9] = i
			}
		}
	}

	// Find a digit with A,E,G and 2 extra wires, which is 2
	for i, input := range inputs {
		if len(input) == 5 {
			if isInSlice(wiring['A'], input) && isInSlice(wiring['E'], input) && isInSlice(wiring['G'], input) {
				numbers[2] = i
			}
		}
	}

	// Find a wire which is in all digits except 2, which is F
	for _, wire := range wires {
		var tempCount int
		for j, input := range inputs {
			if numbers[2] == j {
				continue
			}
			if isInSlice(wire, input) {
				tempCount++
			}
		}
		if tempCount == 9 {
			wiring['F'] = wire
		}
	}

	// Find a wire which is in 2 and is not A,E,G, which is C or D
	var CorD []rune
	for _, wire := range inputs[numbers[2]] {
		if wire == wiring['A'] || wire == wiring['E'] || wire == wiring['G'] {
			continue
		}
		CorD = append(CorD, wire)
	}

	// Find which of CorD is in 1, which is C
	if isInSlice(CorD[0], inputs[numbers[1]]) {
		wiring['C'] = CorD[0]
		wiring['D'] = CorD[1]
	} else {
		wiring['C'] = CorD[1]
		wiring['D'] = CorD[0]
	}

	// The remaining wire must be B
	for _, wire := range wires {
		if wire != wiring['A'] && wire != wiring['C'] && wire != wiring['D'] && wire != wiring['E'] && wire != wiring['F'] && wire != wiring['G'] {
			wiring['B'] = wire
			break
		}
	}

	return
}

func isInSlice(a rune, b []rune) bool {
	for _, d := range b {
		if a == d {
			return true
		}
	}
	return false
}

func parseDisplays(lines []string) (retval []tDisplay) {
	for _, line := range lines {
		disp := tDisplay{}
		inpout := strings.Split(line, " | ")
		inps := strings.Split(inpout[0], " ")
		for i, inp := range inps {
			for _, b := range inp {
				disp.inputs[i] = append(disp.inputs[i], b)
			}
		}
		outs := strings.Split(inpout[1], " ")
		for i, out := range outs {
			for _, b := range out {
				disp.outputs[i] = append(disp.outputs[i], b)
			}
		}
		retval = append(retval, disp)
	}
	return
}

func testInput() []string {
	return []string{
		//"acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbaf",
		"be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe",
		"edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc",
		"fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef | cg cg fdcagb cbg",
		"fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega | efabcd cedba gadfec cb",
		"aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga | gecf egdcabf bgf bfgea",
		"fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf | gebdcfa ecba ca fadegcb",
		"dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf | cefg dcbef fcge gbcadfe",
		"bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd | ed bcgafe cdgba cbgef",
		"egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg | gbdfcae bgc cg cgb",
		"gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc | fgae cfgab fg bagce",
	}
}
