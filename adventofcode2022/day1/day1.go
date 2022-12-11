package day1

import (
	"fmt"
	"main/common"
	"sort"
	"strconv"
)

func summarizeCalories(inputCalories []string) []int {
	elfIndex := 0
	caloriesPerElf := []int{}
	for i := 0; i < len(inputCalories); i++ {
		if inputCalories[i] == "" {
			elfIndex++
			continue
		}

		if len(caloriesPerElf) < elfIndex+1 {
			caloriesPerElf = append(caloriesPerElf, 0)
		}

		calories, _ := strconv.Atoi(inputCalories[i])
		caloriesPerElf[elfIndex] += calories
	}

	return caloriesPerElf
}

func partOneAndTwo() {

	input := common.LinesInFile("adventofcode2022/day1/day1.txt")

	caloriesPerElf := summarizeCalories(input)

	sort.Ints(caloriesPerElf)

	fmt.Println(caloriesPerElf[len(caloriesPerElf)-1])
	fmt.Println(caloriesPerElf[len(caloriesPerElf)-1] + caloriesPerElf[len(caloriesPerElf)-2] + caloriesPerElf[len(caloriesPerElf)-3])
}

func DayOne() {
	partOneAndTwo()
}
