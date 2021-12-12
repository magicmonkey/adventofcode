package main

import (
	"fmt"
	"github.com/mitchellh/copystructure"
	"strings"
)

func main() {
	input := testInput2()
	fmt.Println("Part 1")
	part1(input)
}

type tNode struct {
	next []string
}

var nodes map[string]*tNode
var numRoutes int

func part1(lines []string) {
	buildGraph(lines)

	visited := make(map[string]bool)
	nextStep("start", visited)

	fmt.Println(numRoutes)

}

func nextStep(nodeName string, visited map[string]bool) {
	n := nodes[nodeName]
	for _, next := range n.next {
		nextVisitedTmp, _ := copystructure.Copy(visited)
		nextVisited := nextVisitedTmp.(map[string]bool)
		fmt.Printf("%s - %s ... ", nodeName, next)
		if next == "start" {
			fmt.Println("No - start")
			continue
		} else if next == "end" {
			fmt.Println("end")
			numRoutes++
			return
		} else {
			if !nextVisited[next] {
				fmt.Println("Visiting...")
				if next[0] >= 97 { // lower case
					nextVisited[next] = true
				}
				nextStep(next, nextVisited)
			} else {
				fmt.Println("No (Already been)")
			}
		}
	}
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
