package main

import (
	"fmt"
	"strings"

	"github.com/magicmonkey/adventofcode/2021/util"
)

func main() {
	//lines := testInput()
	lines := util.ReadInputFile()

	fmt.Println("Part 1")
	part1(lines)
}

type tDisplay struct {
	inputs  [10][]byte
	outputs [4][]byte
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

func parseDisplays(lines []string) (retval []tDisplay) {
	for _, line := range lines {
		disp := tDisplay{}
		inpout := strings.Split(line, " | ")
		inps := strings.Split(inpout[0], " ")
		for i, inp := range inps {
			for _, b := range inp {
				disp.inputs[i] = append(disp.inputs[i], byte(b))
			}
		}
		outs := strings.Split(inpout[1], " ")
		for i, out := range outs {
			for _, b := range out {
				disp.outputs[i] = append(disp.outputs[i], byte(b))
			}
		}
		retval = append(retval, disp)
	}
	return
}

func testInput() []string {
	return []string{
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
