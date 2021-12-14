package main

import (
	"fmt"
	"math"
	//"strings"
	"github.com/magicmonkey/adventofcode/2021/util"
	"github.com/mitchellh/copystructure"
)

func main() {
	lines := util.ReadInputFile()
	//lines := testInput()
	fmt.Println("Part 1")
	part1(lines)
	fmt.Println("Part 2")
	part2(lines)
}

func part2(lines []string) {
	var polymer string
	var rules map[string]byte
	var pairs map[byte]map[byte]int64 = make(map[byte]map[byte]int64)

	polymer, rules = parseInput(lines)

	for rule, _ := range rules {
		_, ok := pairs[rule[0]]
		if !ok {
			pairs[rule[0]] = make(map[byte]int64)
		}
		pairs[rule[0]][rule[1]] = 0
	}

	// Initialise the pairs counter with the starting polymer
	for i := 0; i < len(polymer)-1; i++ {
		pairs[polymer[i]][polymer[i+1]]++
	}

	for i := 0; i < 40; i++ {
		pairs = applyRules2(pairs, rules)
	}

	charCount, _ := countChars(pairs)

	var min int64 = math.MaxInt64
	var max int64 = 0

	for _, n := range charCount {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}

	fmt.Println(max - min)
}

func countChars(pairs map[byte]map[byte]int64) (map[byte]int64, int64) {
	// Count the chars
	var charCount map[byte]int64 = make(map[byte]int64)
	for _, t := range pairs {
		for e2, c := range t {
			charCount[e2] += c
		}
	}

	var length int64
	for _, c := range charCount {
		length += c
	}

	return charCount, length
}

func part1(lines []string) {
	var polymer string
	var rules map[string]byte
	polymer, rules = parseInput(lines)

	for i := 0; i < 10; i++ {
		polymer = applyRules(polymer, rules)
	}

	var charCount map[rune]int = make(map[rune]int)
	for _, c := range polymer {
		charCount[c]++
	}

	var min int = math.MaxInt
	var max int = 0

	for _, n := range charCount {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}

	fmt.Println(max - min)
}

func applyRules(polymer string, rules map[string]byte) (newPolymer string) {
	for i := 0; i < len(polymer)-1; i++ {
		newPolymer = newPolymer + polymer[i:i+1]
		newPolymer = newPolymer + string(rules[polymer[i:i+2]])
	}
	newPolymer = newPolymer + polymer[len(polymer)-1:len(polymer)-0]
	return
}

func applyRules2(pairs map[byte]map[byte]int64, rules map[string]byte) map[byte]map[byte]int64 {

	// Deep copy of pairs
	var origPairs map[byte]map[byte]int64
	t, err := copystructure.Copy(pairs)
	if err != nil {
		panic(err)
	}
	origPairs = t.(map[byte]map[byte]int64)

	for e1, t1 := range origPairs {
		for e2, c := range t1 {
			if c == 0 {
				continue
			}
			rule := rules[string(e1)+string(e2)]
			//fmt.Printf("%c%c (%d) becomes %c%c (%d) and %c%c (%d)\n", e1, e2, origPairs[e1][e2], e1, rule, origPairs[e1][rule], rule, e2, origPairs[rule][e2])
			pairs[e1][e2] -= c
			pairs[e1][rule] += c
			pairs[rule][e2] += c
		}
	}
	return pairs
}

func parseInput(lines []string) (polymer string, rules map[string]byte) {
	rules = make(map[string]byte)
	polymer = lines[0]
	for i := 2; i < len(lines); i++ {
		line := lines[i]
		rules[line[0:2]] = line[6]
	}
	return
}

func testInput() []string {
	return []string{
		"NNCB",
		"",
		"CH -> B",
		"HH -> N",
		"CB -> H",
		"NH -> C",
		"HB -> C",
		"HC -> B",
		"HN -> C",
		"NN -> C",
		"BH -> H",
		"NC -> B",
		"NB -> B",
		"BN -> B",
		"BB -> N",
		"BC -> B",
		"CC -> N",
		"CN -> C",
	}
}
