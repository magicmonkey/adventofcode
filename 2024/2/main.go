package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func MustInt(s string) int {
	numberInt, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return numberInt
}

func ParseFile(fname string) (list [][]int) {
	// Open the input file
	file, err := os.Open(fname)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	// Create the regex
	re := regexp.MustCompile(`(\d+)`)

	// Create scanner
	scanner := bufio.NewScanner(file)

	// Read and process lines
	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindAllStringSubmatch(line, -1)
		if matches != nil {
			var thingToInsert []int
			for _, v := range matches {
				thingToInsert = append(thingToInsert, MustInt(v[1]))
			}
			list = append(list, thingToInsert)
		}
	}

	// Check for scanning error
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	return
}

func Abs(num1, num2 int) (retval int) {
	retval = num1 - num2
	if retval < 0 {
		retval *= -1
	}
	return
}

func check(nums []int) bool {
	// Are we increasing or decreasing?
	increasing := false
	if nums[1] > nums[0] {
		increasing = true
	}

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return false
		}
		if increasing && nums[i+1] < nums[i] {
			return false
		}
		if !increasing && nums[i+1] > nums[i] {
			return false
		}
		diff := Abs(nums[i], nums[i+1])
		if diff > 3 {
			return false
		}
	}
	return true
}

func part1(fname string) {
	nums := ParseFile(fname)

	count := 0
	for _, v := range nums {
		if check(v) {
			count += 1
		}
	}
	fmt.Println(count)
}

func main() {
	part1("input.txt")
}
