package main

import (
	"fmt"
	"github.com/magicmonkey/adventofcode/2021/util"
	"strings"
)

func main() {
	//input := testInput3()
	input := util.ReadInputFile()
	fmt.Println("Part 1")
	part1(input)
	fmt.Println("Part 2")
	part2(input)
}

type tNode struct {
	next []string
}

var nodes map[string]*tNode
var numRoutes int

func part1(lines []string) {
	numRoutes = 0
	buildGraph(lines)

	var path []string
	nextStep("start", path)

	fmt.Println(numRoutes)

}

func part2(lines []string) {
	numRoutes = 0
	buildGraph(lines)

	var path []string
	nextStep2("start", path)

	fmt.Println(numRoutes)

}

func nextStep2(nodeName string, path []string) {
	path = append(path, nodeName)
	n := nodes[nodeName]
	for _, next := range n.next {
		var thisPath []string
		thisPath = append([]string{}, path...) // Deep copy
		if next == "start" {
			continue
		} else if next == "end" {
			thisPath = append(thisPath, "end")
			//fmt.Println(thisPath)
			numRoutes++
			continue
		} else {
			if canVisit(next, thisPath) {
				nextStep2(next, thisPath)
			}
		}
	}
}

func canVisit(next string, path []string) bool {
	if isSmallCave(next) {
		// Check if any lower case caves have been visited twice
		if smallCaveTwice(path) {
			if visited(next, path) {
				return false
			}
		}
	}
	return true
}

func smallCaveTwice(path []string) bool {
	for i, p := range path {
		if !isSmallCave(p) {
			continue
		}
		for j, q := range path {
			if p == q && i != j {
				return true
			}
		}
	}
	return false
}

func isSmallCave(name string) bool {
	if name[0] >= 97 { // lower case
		return true
	} else {
		return false
	}
}

func nextStep(nodeName string, path []string) {
	path = append(path, nodeName)
	n := nodes[nodeName]
	for _, next := range n.next {
		if next == "start" {
			continue
		} else if next == "end" {
			path = append(path, "end")
			//fmt.Println(path)
			numRoutes++
			continue
		} else {
			if isSmallCave(next) {
				if visited(next, path) {
					continue
				}
			}
			nextStep(next, path)
		}
	}
}

func visited(nodeName string, path []string) bool {
	for _, p := range path {
		if p == nodeName {
			return true
		}
	}
	return false
}

func buildGraph(lines []string) {
	nodes = make(map[string]*tNode)
	for _, line := range lines {
		var n1, n2 *tNode
		var ok bool
		parts := strings.Split(line, "-")
		n1, ok = nodes[parts[0]]
		if !ok {
			nodes[parts[0]] = &tNode{}
			n1 = nodes[parts[0]]
		}
		n2, ok = nodes[parts[1]]
		if !ok {
			nodes[parts[1]] = &tNode{}
			n2 = nodes[parts[1]]
		}
		appendIfNotExist(n1, parts[1])
		appendIfNotExist(n2, parts[0])
	}
}

func printGraph() {
	for src, node := range nodes {
		for _, dest := range node.next {
			fmt.Println(src, "-", dest)
		}
	}
}

func appendIfNotExist(n *tNode, next string) {
	for _, el := range n.next {
		if el == next {
			return
		}
	}
	n.next = append(n.next, next)
}

func testInput1() []string {
	return []string{
		"start-A",
		"start-b",
		"A-c",
		"A-b",
		"b-d",
		"A-end",
		"b-end",
	}
}
func testInput2() []string {
	return []string{
		"dc-end",
		"HN-start",
		"start-kj",
		"dc-start",
		"dc-HN",
		"LN-dc",
		"HN-end",
		"kj-sa",
		"kj-HN",
		"kj-dc",
	}
}
func testInput3() []string {
	return []string{
		"fs-end",
		"he-DX",
		"fs-he",
		"start-DX",
		"pj-DX",
		"end-zg",
		"zg-sl",
		"zg-pj",
		"pj-he",
		"RW-he",
		"fs-DX",
		"pj-RW",
		"zg-RW",
		"start-pj",
		"he-WI",
		"zg-he",
		"pj-fs",
		"start-RW",
	}
}
