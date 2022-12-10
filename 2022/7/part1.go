package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type tDir struct {
	parent   *tDir
	children map[string]*tDir
	files    map[string]int
}

var allDirs []*tDir

func NewDir() *tDir {
	d := &tDir{}
	d.children = make(map[string]*tDir)
	d.files = make(map[string]int)
	allDirs = append(allDirs, d)
	return d
}

func (d *tDir) MakeChild(name string) *tDir {
	c := NewDir()
	c.parent = d
	d.children[name] = c
	return c
}

func (d *tDir) TotalSize() (size int) {
	for _, s := range d.files {
		size += s
	}
	for _, c := range d.children {
		size += c.TotalSize()
	}
	return
}

func main() {
	root := NewDir()
	currDir := root
	scanner := bufio.NewScanner(os.Stdin)

	r1 := regexp.MustCompile(`^\$ cd /$`)
	r2 := regexp.MustCompile(`^\$ cd \.\.$`)
	r3 := regexp.MustCompile(`^\$ ls$`)
	r4 := regexp.MustCompile(`^dir ([a-z]+)$`)
	r5 := regexp.MustCompile(`^\$ cd ([a-z]+)$`)
	r6 := regexp.MustCompile(`^(\d+) ([a-z\.]+)$`)

	var line []byte
	for scanner.Scan() {
		line = scanner.Bytes()

		//fmt.Println(string(line))

		matches := r1.FindSubmatch(line)
		if len(matches) > 0 {
			currDir = root
			continue
		}

		matches = r2.FindSubmatch(line)
		if len(matches) > 0 {
			currDir = currDir.parent
			continue
		}

		matches = r3.FindSubmatch(line)
		if len(matches) > 0 {
			continue
		}

		matches = r4.FindSubmatch(line)
		if len(matches) > 0 {
			name := string(matches[1])
			_, ok := currDir.children[string(name)]
			if !ok {
				currDir.MakeChild(name)
			}
			continue
		}

		matches = r5.FindSubmatch(line)
		if len(matches) > 0 {
			currDir = currDir.children[string(matches[1])]
			continue
		}

		matches = r6.FindSubmatch(line)
		if len(matches) > 0 {
			size, err := strconv.Atoi(string(matches[1]))
			if err != nil {
				panic(err)
			}
			name := string(matches[2])
			currDir.files[name] = size
			continue
		}

	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	// We now have the whole tree, under root

	part1 := 0
	for _, d := range allDirs {
		s := d.TotalSize()
		if s <= 100000 {
			part1 += s
		}
	}
	fmt.Println("Part 1", part1)

	amountNeeded := 30000000 - (70000000 - root.TotalSize())

	candidate := 30000000
	for _, d := range allDirs {
		s := d.TotalSize()
		if s >= amountNeeded {
			if s < candidate {
				candidate = s
			}
		}
	}

	fmt.Println("Part 2", candidate)

}
