package main

import (
	"bufio"
	"fmt"
	"os"
)

func LinesInFile(fileName string) []string {
	f, _ := os.Open(fileName)
	// Create new Scanner.
	scanner := bufio.NewScanner(f)
	result := []string{}
	// Use Scan.
	for scanner.Scan() {
		line := scanner.Text()
		// Append line to result.
		result = append(result, line)
	}
	return result
}

func findLowPoints(points [][]int8) []int8 {

	var lowPoints []int8
	for i := 0; i < len(points); i++ {
		for j := 0; j < len(points[0]); j++ {
			lowPoint := true
			if lowPoint && i > 0 {
				lowPoint = points[i-1][j] > points[i][j]
			}
			if lowPoint && i < len(points)-1 {
				lowPoint = points[i+1][j] > points[i][j]
			}
			if lowPoint && j > 0 {
				lowPoint = points[i][j-1] > points[i][j]
			}
			if lowPoint && j < len(points[i])-1 {
				lowPoint = points[i][j+1] > points[i][j]
			}

			if lowPoint {
				lowPoints = append(lowPoints, points[i][j])
			}
		}
	}

	return lowPoints
}

func sumPlusOne(array []int8) int {
	result := 0
	for _, v := range array {
		result += int(v + 1)
	}
	return result
}

func main() {

	lines := LinesInFile("day-9.txt")

	//fmt.Println(lines)

	points := make([][]int8, len(lines))

	for i := 0; i < len(lines); i++ {

		var linePoints []int8
		for j := 0; j < len(lines[0]); j++ {
			linePoints = append(linePoints, int8(lines[i][j])-48)
		}

		points[i] = linePoints
	}

	lowPoints := findLowPoints(points)

	fmt.Println(sumPlusOne(lowPoints))

	//file, err := os.Open("day-9.txt")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer file.Close()
	//
	//scanner := bufio.NewScanner(file)
	//
	//for scanner.Scan() {
	//	line := scanner.Text()
	//	fmt.Println(line)
	//}
	//
	//if err := scanner.Err(); err != nil {
	//	log.Fatal(err)
	//}
}
