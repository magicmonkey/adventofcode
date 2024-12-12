package main

import (
	"fmt"
	"strconv"
)

func MustInt(s string) int {
	numberInt, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return numberInt
}

func blink1(nums []int) (newNums []int) {
	for _, num := range nums {

		if num == 0 {
			newNums = append(newNums, 1)
		} else if strNum := strconv.Itoa(num); len(strNum)%2 == 0 {
			mid := len(strNum) / 2
			newNums = append(newNums, MustInt(strNum[:mid]))
			newNums = append(newNums, MustInt(strNum[mid:]))
		} else {
			newNums = append(newNums, num*2024)
		}
	}
	return
}

func blink2(nums map[int]int) (newNums map[int]int) {
	newNums = make(map[int]int)

	for num, count := range nums {

		if newNums[num] == 0 {
			delete(newNums, num)
		}
		if num == 0 {
			newNums[1] += count
		} else if strNum := strconv.Itoa(num); len(strNum)%2 == 0 {
			mid := len(strNum) / 2
			newNums[MustInt(strNum[:mid])] += count
			newNums[MustInt(strNum[mid:])] += count
		} else {
			newNums[num*2024] += count
		}
	}
	return
}

func part1(nums []int) {
	for i := 0; i < 25; i++ {
		nums = blink1(nums)
	}
	fmt.Println(len(nums))
}

func part2(nums map[int]int) {
	for i := 0; i < 75; i++ {
		nums = blink2(nums)
	}
	total := 0
	for _, count := range nums {
		total += count
	}
	fmt.Println(total)
}

func main() {
	//test := []int{125, 17}
	//input := []int{41078, 18, 7, 0, 4785508, 535256, 8154, 447}
	//part1(test)
	//part2(map[int]int{125: 1, 17: 1})
	part1([]int{41078, 18, 7, 0, 4785508, 535256, 8154, 447})
	part2(map[int]int{41078: 1, 18: 1, 7: 1, 0: 1, 4785508: 1, 535256: 1, 8154: 1, 447: 1})
}
