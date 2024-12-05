package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
)

func MustInt(s string) int {
	numberInt, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return numberInt
}

func ParseFile(fname string) (rules map[int][]int, updates [][]int) {
	rules = make(map[int][]int)
	// Open the input file
	file, err := os.Open(fname)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	// Create the regexes
	re1 := regexp.MustCompile(`^(\d+)\|(\d+)$`)
	re2 := regexp.MustCompile(`(\d+)`)

	// Create scanner
	scanner := bufio.NewScanner(file)

	// Read and process lines
	for scanner.Scan() {
		line := scanner.Text()
		matches1 := re1.FindStringSubmatch(line)
		matches2 := re2.FindAllStringSubmatch(line, -1)
		if matches1 != nil {
			rule := [2]int{
				MustInt(matches1[1]),
				MustInt(matches1[2]),
			}
			if _, ok := rules[rule[0]]; !ok {
				rules[rule[0]] = []int{}
			}
			rules[rule[0]] = append(rules[rule[0]], rule[1])
		} else if line == "" {
			continue
		} else if matches2 != nil {
			var update []int
			for _, v := range matches2 {
				update = append(update, MustInt(v[1]))
			}
			updates = append(updates, update)
		}
	}

	// Check for scanning error
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	return
}

func checkRules(update []int, rules map[int][]int) bool {
	for i, thingToCheck := range update {
		for _, thingToCheckAgainst := range update[:i] {
			if slices.Contains(rules[thingToCheck], thingToCheckAgainst) {
				return false
			}
		}
	}
	return true
}

func part1(fname string) {
	rules, updates := ParseFile(fname)

	runningTotal := 0
	for _, update := range updates {
		if checkRules(update, rules) {
			fmt.Printf("Yes : %v\n", update)
			runningTotal += update[(len(update)-1)/2]
		}
	}
	fmt.Println(runningTotal)
}

func main() {
	part1("input.txt")
}
