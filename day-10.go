package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func getValue(character uint8) int {
	switch character {
	case ')':
		return 3
	case ']':
		return 57
	case '}':
		return 1197
	case '>':
		return 25137
	}

	return 0
}

func syntaxCheck(line string) int {
	var array []uint8
	for i := 0; i < len(line); i++ {
		if line[i] == '(' || line[i] == '[' || line[i] == '{' || line[i] == '<' {
			array = append(array, line[i])
		} else if i == 0 {
			return getValue(line[i])
		} else {
			if (line[i] == ')' && array[len(array)-1] == '(') ||
				(line[i] == '}' && array[len(array)-1] == '{') ||
				(line[i] == ']' && array[len(array)-1] == '[') ||
				(line[i] == '>' && array[len(array)-1] == '<') {
				array = array[:len(array)-1]
			} else {
				return getValue(line[i])
			}
		}
	}

	return 0
}

func complete(line string) string {
	var array []uint8
	for i := 0; i < len(line); i++ {
		if line[i] == '(' || line[i] == '[' || line[i] == '{' || line[i] == '<' {
			array = append(array, line[i])
		} else {
			if (line[i] == ')' && array[len(array)-1] == '(') ||
				(line[i] == '}' && array[len(array)-1] == '{') ||
				(line[i] == ']' && array[len(array)-1] == '[') ||
				(line[i] == '>' && array[len(array)-1] == '<') {
				array = array[:len(array)-1]
			}
		}
	}

	var completeArry []uint8
	for i := len(array) - 1; i >= 0; i-- {
		if array[i] == '(' {
			completeArry = append(completeArry, ')')
		}
		if array[i] == '{' {
			completeArry = append(completeArry, '}')
		}
		if array[i] == '[' {
			completeArry = append(completeArry, ']')
		}
		if array[i] == '<' {
			completeArry = append(completeArry, '>')
		}
	}

	return string(completeArry)
}

func evaluate(completion string) int {
	accu := 0
	for i := 0; i < len(completion); i++ {
		accu *= 5

		switch completion[i] {
		case ')':
			accu += 1
		case ']':
			accu += 2
		case '}':
			accu += 3
		case '>':
			accu += 4
		}
	}

	return accu
}

func main() {
	file, err := os.Open("day-10.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var scores []int
	for scanner.Scan() {
		line := scanner.Text()

		if syntaxCheck(line) == 0 {
			completion := complete(line)
			scores = append(scores, evaluate(completion))
		}
	}

	sort.Ints(scores)

	fmt.Println(scores[len(scores)/2])

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
