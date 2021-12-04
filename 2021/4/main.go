package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/magicmonkey/adventofcode/2021/util"
)

type tLine []int
type tCard []tLine // Data structure for a card

func main() {
	var cards []tCard
	var balls []int
	lines := util.ReadInputFile()
	//lines := testInput()

	balls = parseBalls(lines[0])
	cards = parseCards(lines[1:])

	fmt.Println("Part 1")
	part1(balls, cards)
	fmt.Println("Part 2")
	part2(balls, cards)
}

func part1(balls []int, cards []tCard) {
	winningCard := -1
	winner := false
	ballnum := -1

	for ballnum = 4; ballnum < len(balls); ballnum++ {
		winner, winningCard = checkCards(balls[0:ballnum], cards)
		if winner {
			break
		}
	}

	fmt.Println("First winning card is", cards[winningCard][0], "on ball number", ballnum)

	score := getScore(balls[0:ballnum], cards[winningCard])

	fmt.Println(score)
}

func part2(balls []int, cards []tCard) {
	winningCard := -1
	winner := false
	ballnum := -1
	var lastWinner tCard
	var lastBall int

outer:
	for ballnum = 4; ballnum <= len(balls); ballnum++ {
		winner = false
		for {
			winner, winningCard = checkCards(balls[0:ballnum], cards)
			if winner {
				lastWinner = cards[winningCard]
				lastBall = ballnum
				cards = removeCard(cards, winningCard)
				if len(cards) == 0 {
					break outer
				}
			} else {
				break
			}
		}
	}

	fmt.Println("Last winning card is", lastWinner[0], "on ball number", lastBall)
	score := getScore(balls[0:lastBall], lastWinner)
	fmt.Println(score)
}

func removeCard(cards []tCard, cardNum int) []tCard {
	return append(cards[:cardNum], cards[cardNum+1:]...)
}

func getScore(balls []int, card tCard) int {
	unmarked := getUnmarkedBalls(balls, card)
	sum := 0
	for _, num := range unmarked {
		sum += num
	}
	return sum * balls[len(balls)-1]
}

func getUnmarkedBalls(balls []int, card tCard) (retval []int) {
	var ball int
	for _, row := range card {
		for _, num := range row {
			var matched bool = false
			for _, ball = range balls {
				if ball == num {
					matched = true
					break
				}
			}
			if !matched {
				retval = append(retval, num)
			}
		}
	}
	return
}

func checkCards(balls []int, cards []tCard) (winner bool, winningCard int) {
	// Check rows
	for cardNum, card := range cards {
		for _, row := range card {
			if countNumInSet(balls, row) == 5 {
				return true, cardNum
			}
		}
	}
	// Check columns
	for cardNum, card := range cards {
		for colNum := 0; colNum < 5; colNum++ {
			var col []int = []int{
				card[0][colNum],
				card[1][colNum],
				card[2][colNum],
				card[3][colNum],
				card[4][colNum],
			}
			if countNumInSet(balls, col) == 5 {
				return true, cardNum
			}
		}
	}
	return false, -1
}

func countNumInSet(balls []int, row []int) (retval int) {
	for i := 0; i < len(balls); i++ {
		for j := 0; j < len(row); j++ {
			if balls[i] == row[j] {
				retval++
			}
		}
	}
	return
}

func parseCards(lines []string) (retval []tCard) {
	for i := 0; i < len(lines); i += 6 {
		var c tCard
		for j := 1; j <= 5; j++ {
			l := parseLine(lines[i+j])
			c = append(c, l)
		}
		retval = append(retval, c)
	}
	return
}

func parseLine(line string) (retval tLine) {
	line = strings.ReplaceAll(line, "  ", " ")
	line = strings.Trim(line, " ")
	numsString := strings.Split(line, " ")
	for _, numString := range numsString {
		numString = strings.Trim(numString, " ")
		numInt, err := strconv.ParseInt(numString, 10, 64)
		if err != nil {
			panic(err)
		}
		retval = append(retval, int(numInt))
	}
	return
}

func parseBalls(line string) (retval []int) {
	ballsString := strings.Split(line, ",")
	for _, ballString := range ballsString {
		ballInt, err := strconv.ParseInt(ballString, 10, 64)
		if err != nil {
			panic(err)
		}
		retval = append(retval, int(ballInt))
	}
	return
}

func testInput() []string {
	return []string{
		"7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1",
		"",
		"22 13 17 11  0",
		" 8  2 23  4 24",
		"21  9 14 16  7",
		" 6 10  3 18  5",
		" 1 12 20 15 19",
		"",
		" 3 15  0  2 22",
		" 9 18 13 17  5",
		"19  8  7 25 23",
		"20 11 10 24  4",
		"14 21 16 12  6",
		"",
		"14 21 17 24  4",
		"10 16 15  9 19",
		"18  8 23 26 20",
		"22 11 13  6  5",
		" 2  0 12  3  7",
	}
}
