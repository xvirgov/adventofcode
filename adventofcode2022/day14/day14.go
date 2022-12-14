package day14

import (
	"fmt"
	"log"
	"main/common"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

type Pair struct {
	x, y int
}

func getRockCoordinates() [][]Pair {
	rockSlices := common.LinesInFile("adventofcode2022/day14/day14.txt")

	rockSliceCoordinates := [][]Pair{}

	for _, rockSlice := range rockSlices {
		coordinates := []Pair{}

		points := strings.Split(rockSlice, " -> ")
		for _, point := range points {
			xy := strings.Split(point, ",")

			x, _ := strconv.Atoi(xy[0])
			y, _ := strconv.Atoi(xy[1])

			coordinates = append(coordinates, Pair{x: x, y: y})
		}
		rockSliceCoordinates = append(rockSliceCoordinates, coordinates)
	}
	return rockSliceCoordinates
}

func getMinMax(rockCoordinates [][]Pair) (Pair, Pair) {
	min := Pair{
		x: math.MaxInt,
		y: math.MaxInt,
	}

	max := Pair{
		x: math.MinInt,
		y: math.MinInt,
	}

	for _, coorCol := range rockCoordinates {
		for _, coorRow := range coorCol {
			if coorRow.x < min.x {
				min.x = coorRow.x
			}
			if coorRow.y < min.y {
				min.y = coorRow.y
			}
			if coorRow.x > max.x {
				max.x = coorRow.x
			}
			if coorRow.y > max.y {
				max.y = coorRow.y
			}
		}
	}

	return min, max
}

func generateRockMap(rockCoordinates [][]Pair) [][]rune {

	min, max := getMinMax(rockCoordinates)
	//fmt.Println(min)
	//fmt.Println(max)
	min = Pair{
		x: min.x,
		y: 0,
	}

	rockMap := [][]rune{}

	// initialize
	for i := min.y; i < max.y+1; i++ {
		rockMap = append(rockMap, []rune(strings.Repeat(".", max.x-min.x+1)))
	}

	// draw rocks
	for _, rockLine := range rockCoordinates {
		for i := 0; i < len(rockLine)-1; i++ {
			startPoint := Pair{
				x: rockLine[i].x - min.x,
				y: rockLine[i].y - min.y,
			}
			endPoint := Pair{
				x: rockLine[i+1].x - min.x,
				y: rockLine[i+1].y - min.y,
			}

			if startPoint.x == endPoint.x { // horizontal drawing
				if startPoint.y <= endPoint.y {
					for y := startPoint.y; y < endPoint.y; y++ {
						rockMap[y][startPoint.x] = '#'
					}
				} else {
					for y := endPoint.y; y <= startPoint.y; y++ {
						rockMap[y][startPoint.x] = '#'
					}
				}
			} else if startPoint.y == endPoint.y { // vertical drawing
				if startPoint.x < endPoint.x {
					for x := startPoint.x; x <= endPoint.x; x++ {
						rockMap[startPoint.y][x] = '#'
					}
				} else {
					for x := endPoint.x; x <= startPoint.x; x++ {
						rockMap[startPoint.y][x] = '#'
					}
				}
			} else { // error
				fmt.Println("This not right :(")
				os.Exit(1)
			}
		}
	}

	return rockMap
}

func generateRockMapV2(rockCoordinates [][]Pair) ([][]rune, int) {

	min, max := getMinMax(rockCoordinates)
	//fmt.Println(min)
	//fmt.Println(max)
	min = Pair{
		x: min.x,
		y: 0,
	}

	height := max.y - min.y
	//sandStartPoint := Pair{
	//	x: 500 - min.x,
	//	y: 0,
	//}
	//
	//min.x = sandStartPoint.x - height
	max.x = max.x + height

	rockMap := [][]rune{}

	// initialize
	for i := min.y; i < max.y+3; i++ {
		rockMap = append(rockMap, []rune(strings.Repeat(".", max.x-min.x+1+2*height)))
	}

	// draw rocks
	for _, rockLine := range rockCoordinates {
		for i := 0; i < len(rockLine)-1; i++ {

			startPoint := Pair{
				x: rockLine[i].x - min.x + height,
				y: rockLine[i].y - min.y,
			}
			endPoint := Pair{
				x: rockLine[i+1].x - min.x + height,
				y: rockLine[i+1].y - min.y,
			}

			if startPoint.x == endPoint.x { // horizontal drawing
				if startPoint.y <= endPoint.y {
					for y := startPoint.y; y < endPoint.y; y++ {
						rockMap[y][startPoint.x] = '#'
					}
				} else {
					for y := endPoint.y; y <= startPoint.y; y++ {
						rockMap[y][startPoint.x] = '#'
					}
				}
			} else if startPoint.y == endPoint.y { // vertical drawing
				if startPoint.x < endPoint.x {
					for x := startPoint.x; x <= endPoint.x; x++ {
						rockMap[startPoint.y][x] = '#'
					}
				} else {
					for x := endPoint.x; x <= startPoint.x; x++ {
						rockMap[startPoint.y][x] = '#'
					}
				}
			} else { // error
				fmt.Println("This not right :(")
				os.Exit(1)
			}
		}
	}

	for i := 0; i < len(rockMap[0]); i++ {
		rockMap[len(rockMap)-1][i] = '#'
	}

	return rockMap, height
}

func printRockMap(rockMap [][]rune) {
	for _, line := range rockMap {
		for _, char := range line {
			fmt.Print(string(char))
		}
		fmt.Println("")
	}
}

func addSand(rockMap [][]rune, startPoint Pair) ([][]rune, bool) {
	x := startPoint.x
	for y := startPoint.y; y < len(rockMap)-1; y++ {
		if rockMap[y+1][x] == '.' {
			continue
		} else if x-1 < 0 {
			return rockMap, false
		} else if rockMap[y+1][x-1] == '.' {
			x--
		} else if x+1 > len(rockMap[0]) {
			return rockMap, false
		} else if rockMap[y+1][x+1] == '.' {
			x++
		} else {
			rockMap[y][x] = 'o'
			return rockMap, true
		}
	}

	return rockMap, false
}

func addSandV2(rockMap [][]rune, startPoint Pair) ([][]rune, bool) {
	x := startPoint.x
	for y := startPoint.y; y < len(rockMap)-1; y++ {
		if rockMap[y][x] == 'o' {
			return rockMap, false
		} else if rockMap[y+1][x] == '.' {
			continue
		} else if x-1 < 0 {
			return rockMap, false
		} else if rockMap[y+1][x-1] == '.' {
			x--
		} else if x+1 > len(rockMap[0]) {
			return rockMap, false
		} else if rockMap[y+1][x+1] == '.' {
			x++
		} else {
			rockMap[y][x] = 'o'
			return rockMap, true
		}
	}

	return rockMap, false
}

func partOne() {
	rockCoordinates := getRockCoordinates()

	rockMap := generateRockMap(rockCoordinates)

	min, _ := getMinMax(rockCoordinates)

	startPoint := Pair{
		x: 500 - min.x,
		y: 0,
	}

	added := false
	counter := 0
	for true {
		rockMap, added = addSand(rockMap, startPoint)

		if !added {
			break
		}

		counter++
	}
	printRockMap(rockMap)
	fmt.Println(counter)
}

func partTwo() {
	rockCoordinates := getRockCoordinates()

	rockMap, height := generateRockMapV2(rockCoordinates)

	min, _ := getMinMax(rockCoordinates)

	startPoint := Pair{
		x: 500 - min.x + height,
		y: 0,
	}

	added := false
	counter := 0
	for true {
		rockMap, added = addSandV2(rockMap, startPoint)

		if !added {
			break
		}

		counter++
	}
	printRockMap(rockMap)
	fmt.Println(counter)
}

func DayFourteen() {
	start := time.Now()
	//partOne()
	//log.Printf("Part one took %s", time.Since(start))

	start = time.Now()
	partTwo()
	log.Printf("Part two took %s", time.Since(start))
}
