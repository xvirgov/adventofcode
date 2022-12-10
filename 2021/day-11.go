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

func flush(levels [][]int, i int, j int) [][]int {
	if i > 0 {
		levels[i-1][j]++
	}
	if j > 0 {
		levels[i][j-1]++
	}
	if i < len(levels)-1 {
		levels[i+1][j]++
	}
	if j < len(levels)-1 {
		levels[i][j+1]++
	}
	if i > 0 && j > 0 {
		levels[i-1][j-1]++
	}
	if i > 0 && j < len(levels)-1 {
		levels[i-1][j+1]++
	}
	if i < len(levels)-1 && j > 0 {
		levels[i+1][j-1]++
	}
	if i < len(levels)-1 && j < len(levels)-1 {
		levels[i+1][j+1]++
	}

	return levels
}

func allFlushed(flushed [][]bool) bool {
	for i := 0; i < len(flushed); i++ {
		for j := 0; j < len(flushed); j++ {
			if flushed[i][j] == false {
				return false
			}
		}
	}

	return true
}

func round(levels [][]int) ([][]int, int, bool) {
	flushed := make([][]bool, len(levels))
	for i := 0; i < len(levels); i++ {
		var flushedLine []bool
		for j := 0; j < len(levels); j++ {
			flushedLine = append(flushedLine, false)
		}
		flushed[i] = flushedLine
	}

	for i := 0; i < len(levels); i++ {
		for j := 0; j < len(levels); j++ {
			levels[i][j] = levels[i][j] + 1
		}
	}

	flushCount := 0

	// :(
	for x := 0; x < len(levels); x++ {
		for y := 0; y < len(levels); y++ {
			for i := 0; i < len(levels); i++ {
				for j := 0; j < len(levels); j++ {
					if levels[i][j] > 9 && !flushed[i][j] {
						levels = flush(levels, i, j)
						flushCount += 1
						flushed[i][j] = true
					}
				}
			}
		}
	}

	for i := 0; i < len(levels); i++ {
		for j := 0; j < len(levels); j++ {
			if levels[i][j] > 9 {
				levels[i][j] = 0
			}
		}
	}

	if allFlushed(flushed) {
		return levels, flushCount, true
	}

	return levels, flushCount, false
}

func printLevels(levels [][]int) {
	for i := 0; i < len(levels); i++ {
		for j := 0; j < len(levels); j++ {
			fmt.Printf("%d", levels[i][j])
		}
		fmt.Printf("\n")
	}
}

func main() {
	lines := LinesInFile("day-11.txt")

	levels := make([][]int, len(lines))

	for i := 0; i < len(lines); i++ {

		var lineLevels []int
		for j := 0; j < len(lines[0]); j++ {
			lineLevels = append(lineLevels, int(lines[i][j])-48)
		}

		levels[i] = lineLevels
	}

	nrounds := 1000
	totalFlushes := 0
	for i := 0; i < nrounds; i++ {

		flushes := 0
		allFlushed := false
		levels, flushes, allFlushed = round(levels)

		totalFlushes += flushes

		fmt.Printf("After %d-th round\n", i+1)
		printLevels(levels)

		if allFlushed {
			fmt.Printf("All flushed in %d-th round", i+1)
			break
		}
	}

	fmt.Println(totalFlushes)

}
