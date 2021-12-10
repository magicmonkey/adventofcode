package main

import "fmt"
import "sort"
import "container/list"

import "github.com/magicmonkey/adventofcode/2021/util"

func main() {
	//lines := testInput()
	lines := util.ReadInputFile()
	fmt.Println("Part 1")
	part1(lines)
	fmt.Println("Part 2")
	part2(lines)
}

func part1(lines []string) {
	var score int
	for _, line := range lines {
		//fmt.Printf("%s ... ", line)
		corrupted, badChar := checkLineP1(line)
		if corrupted {
			switch badChar {
			case ')':
				score += 3
			case ']':
				score += 57
			case '}':
				score += 1197
			case '>':
				score += 25137
			}
		}
	}
	fmt.Println(score)
}

func part2(lines []string) {
	var scores []int
	for _, line := range lines {
		//fmt.Println(line)
		isCorrupted, remaining := checkLineP2(line)
		if !isCorrupted {
			var score int
			for _, c := range remaining {
				score *= 5
				switch c {
				case ')':
					score += 1
				case ']':
					score += 2
				case '}':
					score += 3
				case '>':
					score += 4
				}
			}
			scores = append(scores, score)
		}
	}
	sort.Sort(sort.IntSlice(scores))
	fmt.Println(scores[(len(scores)-1)/2])
}

func checkLineP2(line string) (isCorrupted bool, remaining []rune) {
	l := list.New()
	for _, c := range line {
		//fmt.Println("Checking", string(c), "...")
		switch c {
		case '[':
			l.PushBack(']')
		case '(':
			l.PushBack(')')
		case '{':
			l.PushBack('}')
		case '<':
			l.PushBack('>')
		case '>':
			fallthrough
		case '}':
			fallthrough
		case ')':
			fallthrough
		case ']':
			back := l.Back()
			expected := back.Value.(rune)
			if c != expected {
				//fmt.Println("Expected", string(expected), "but got", string(c))
				return true, []rune{}
			}
			l.Remove(back)
		default:
			panic("Unknown token")
		}
	}
	for e := l.Back(); e != nil; e = e.Prev() {
		remaining = append(remaining, e.Value.(rune))
	}
	return false, remaining
}

func checkLineP1(line string) (isCorrupted bool, badChar rune) {
	l := list.New()
	for _, c := range line {
		//fmt.Println("Checking", string(c), "...")
		switch c {
		case '[':
			l.PushBack(']')
		case '(':
			l.PushBack(')')
		case '{':
			l.PushBack('}')
		case '<':
			l.PushBack('>')
		case '>':
			fallthrough
		case '}':
			fallthrough
		case ')':
			fallthrough
		case ']':
			back := l.Back()
			expected := back.Value.(rune)
			if c != expected {
				//fmt.Println("Expected", string(expected), "but got", string(c))
				return true, c
			}
			l.Remove(back)
		default:
			panic("Unknown token")
		}
	}
	return false, '-'
}

func testInput() []string {
	return []string{
		"[({(<(())[]>[[{[]{<()<>>",
		"[(()[<>])]({[<{<<[]>>(",
		"{([(<{}[<>[]}>{[]{[(<()>",
		"(((({<>}<{<{<>}{[]{[]{}",
		"[[<[([]))<([[{}[[()]]]",
		"[{[{({}]{}}([{[{{{}}([]",
		"{<[[]]>}<{[{[{[]{()[[[]",
		"[<(<(<(<{}))><([]([]()",
		"<{([([[(<>()){}]>(<<{{",
		"<{([{{}}[<[[[<>{}]]]>[]]",
	}
}
