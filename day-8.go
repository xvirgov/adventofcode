package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func getCountUnique(arr []string) int {
	uniqueCounts := []int{2, 4, 3, 7}
	count := 0
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(uniqueCounts); j++ {
			if len(arr[i]) == uniqueCounts[j] {
				count++
				break
			}
		}
	}

	return count
}

func getMapping(mapings map[string]int, number int) string {
	for k, v := range mapings {
		if v == number {
			return k
		}
	}

	return ""
}

func getCommonSubsetCount(str1 string, str2 string) int {}

func decodeSequences(input []string, output []string) {

	// get mappings for 1, 4, 7, 8

	mappings := make(map[string]int)

	// get the known sequences
	for i := 0; i < len(input); i++ {
		if len(input[i]) == 2 {
			mappings[input[i]] = 1
		} else if len(input[i]) == 4 {
			mappings[input[i]] = 4
		} else if len(input[i]) == 3 {
			mappings[input[i]] = 7
		} else if len(input[i]) == 7 {
			mappings[input[i]] = 8
		}
	}

	for i := 0; i < len(input); i++ {
		if len(input[i]) == 5 {
			// 3 -> length 5, but contains all from 1
			// 5 -> length 5 and contains 3 from 4
			// 2 -> length 5 else
			oneMapping := getMapping(mappings, 1)

			if strings.ContainsAny(input[i], oneMapping) {
				mappings[input[i]] = 3
			}

		} else if len(input[i]) == 6 {
			// 6 -> length 6 but doesnt contain all from 1
			// 0 -> length 6 but doesnt contain all from 4
			// 9 -> length 6, else

			oneMapping := getMapping(mappings, 1)
			fourMapping := getMapping(mappings, 4)

			if !strings.Contains(input[i], oneMapping) {
				mappings[input[i]] = 6
			} else if !strings.Contains(input[i], fourMapping) {
				mappings[input[i]] = 0
			} else {
				mappings[input[i]] = 9
			}
		}
	}

	fmt.Println(mappings)

}

func main() {
	file, err := os.Open("day-8.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	//countAll := 0
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		line := scanner.Text()
		arr := strings.Split(line, " ")
		input := arr[:10]
		output := arr[11:]
		//countAll += getCountUnique(output)
		decodeSequences(input, output)
	}

	//fmt.Println(countAll)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
