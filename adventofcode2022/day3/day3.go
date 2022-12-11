package day3

import (
	"fmt"
	"main/common"
	"sort"
	"strings"
)

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func findIntersect(arr1, arr2 string) []rune {
	arr1Index := 0
	arr2Index := 0

	res := []rune{}

	for arr1Index < len(arr1) && arr2Index < len(arr2) {
		if arr1[arr1Index] < arr2[arr2Index] {
			arr1Index++
		} else if arr2[arr2Index] < arr1[arr1Index] {
			arr2Index++
		} else {
			res = append(res, rune(arr1[arr1Index]))
			arr1Index++
			arr2Index++
		}
	}

	return res
}

func getValue(in rune) int {
	res := 0
	if in >= 'a' && in <= 'z' {
		res = int(in-'a') + 1
	} else {
		res = int(in-'A') + 27
	}
	return res
}

func partOne() {
	contents := common.LinesInFile("adventofcode2022/day3/day3.txt")

	prioSum := 0

	for i := 0; i < len(contents); i++ {
		firstComp := contents[i][:len(contents[i])/2]
		secondComp := contents[i][len(contents[i])/2:]
		intersect := findIntersect(SortString(firstComp), SortString(secondComp))

		prioSum += getValue(intersect[0])
	}

	fmt.Println(prioSum)
}

func partTwo() {
	contents := common.LinesInFile("adventofcode2022/day3/day3.txt")
	prioSum := 0
	for i := 0; i < len(contents); i += 3 {
		rucksackOne := SortString(contents[i])
		rucksackTwo := SortString(contents[i+1])
		rucksackThree := SortString(contents[i+2])

		commonItems := findIntersect(rucksackOne, string(findIntersect(rucksackTwo, rucksackThree)))

		prioSum += getValue(commonItems[0])
	}
	fmt.Println(prioSum)
}

func DayThree() {
	partOne()
	partTwo()
}
