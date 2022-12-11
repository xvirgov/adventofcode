package day2

import (
	"fmt"
	"main/common"
)

func evaluateRPS(player, opponent rune) int {
	res := 0
	if (player == 'R' && opponent == 'S') || (player == 'P' && opponent == 'R') || (player == 'S' && opponent == 'P') {
		res += 6
	}
	if player == opponent {
		res += 3
	}

	switch player {
	case 'R':
		res += 1
	case 'P':
		res += 2
	case 'S':
		res += 3
	}

	return res
}

func getOpponentStrategy(oponent rune) rune {
	res := 'x'

	switch oponent {
	case 'A':
		res = 'R'
	case 'B':
		res = 'P'
	case 'C':
		res = 'S'
	}

	return res
}

func getPlayerStrategy(player rune) rune {
	res := 'x'

	switch player {
	case 'X':
		res = 'R'
	case 'Y':
		res = 'P'
	case 'Z':
		res = 'S'
	}

	return res
}

func partOne() {
	input := common.LinesInFile("adventofcode2022/day2/day2.txt")
	score := 0
	for i := 0; i < len(input); i++ {
		opponent := getOpponentStrategy(rune(input[i][0]))
		player := getPlayerStrategy(rune(input[i][2]))

		score += evaluateRPS(player, opponent)
	}

	fmt.Println(score)
}

func getPlayerStrategyV2(player, opponent rune) rune {
	res := 'x'
	switch player {
	case 'X':
		switch opponent {
		case 'R':
			res = 'S'
		case 'P':
			res = 'R'
		case 'S':
			res = 'P'
		}
	case 'Y':
		res = opponent
	case 'Z':
		switch opponent {
		case 'S':
			res = 'R'
		case 'R':
			res = 'P'
		case 'P':
			res = 'S'
		}
	}

	return res
}

func partTwo() {
	input := common.LinesInFile("adventofcode2022/day2.txt")
	score := 0
	for i := 0; i < len(input); i++ {
		opponent := getOpponentStrategy(rune(input[i][0]))
		player := getPlayerStrategyV2(rune(input[i][2]), opponent)

		score += evaluateRPS(player, opponent)
	}

	fmt.Println(score)
}

func DayTwo() {
	partOne()
	partTwo()
}
