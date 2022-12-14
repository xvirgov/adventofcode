package day12

import (
	"fmt"
	"log"
	"main/common"
	"time"
)

type Position struct {
	row     int
	col     int
	val     rune
	visited bool
}

type Node struct {
	pos    *Position
	parent *Node
}

func canMove(from, to Position) bool {
	if to.visited {
		return false
	}
	if to.val == 'E' {
		to.val = 'z'
	}
	return from.val-to.val >= -1
}

func (n *Node) breadthFirstSearchElevation(positionsMap [][]Position) *Node {
	queue := []*Node{n}
	for len(queue) > 0 {
		//fmt.Println(len(queue))
		// dequeue
		current := queue[0]
		positionsMap[current.pos.row][current.pos.col].visited = true

		queue = queue[1:]

		// check if found the end
		if current.pos.val == 'E' {
			return current
		}

		currentRow := current.pos.row
		currentCol := current.pos.col
		//currentVal := current.pos.val

		// add children to the queue
		if currentRow > 0 {
			if canMove(*current.pos, positionsMap[currentRow-1][currentCol]) && !positionsMap[currentRow-1][currentCol].visited {
				positionsMap[currentRow-1][currentCol].visited = true
				child := Node{
					pos:    &positionsMap[currentRow-1][currentCol],
					parent: current,
				}
				queue = append(queue, &child)
			}
		}
		if currentRow < len(positionsMap)-1 {
			if canMove(*current.pos, positionsMap[currentRow+1][currentCol]) && !positionsMap[currentRow+1][currentCol].visited {
				positionsMap[currentRow+1][currentCol].visited = true
				child := Node{
					pos:    &positionsMap[currentRow+1][currentCol],
					parent: current,
				}
				queue = append(queue, &child)
			}
		}
		if currentCol > 0 {
			if canMove(*current.pos, positionsMap[currentRow][currentCol-1]) && !positionsMap[currentRow][currentCol-1].visited {
				positionsMap[currentRow][currentCol-1].visited = true
				child := Node{
					pos:    &positionsMap[currentRow][currentCol-1],
					parent: current,
				}
				queue = append(queue, &child)
			}
		}
		if currentCol < len(positionsMap[0])-1 {
			if canMove(*current.pos, positionsMap[currentRow][currentCol+1]) && !positionsMap[currentRow][currentCol+1].visited {
				positionsMap[currentRow][currentCol+1].visited = true
				child := Node{
					pos:    &positionsMap[currentRow][currentCol+1],
					parent: current,
				}
				queue = append(queue, &child)
			}
		}
	}
	return &Node{}
}

func convertMapToPositions(elevationMap []string) [][]Position {
	positions := make([][]Position, len(elevationMap))

	for i, elevationRow := range elevationMap {
		positions[i] = make([]Position, len(elevationRow))
		for j, elevation := range elevationMap[i] {
			position := Position{
				row:     i,
				col:     j,
				val:     elevation,
				visited: false,
			}
			positions[i][j] = position
		}
	}
	return positions
}

//func getPathLength

func getStartPosition(elevationMap []string) Position {
	position := Position{}

	for i, line := range elevationMap {
		for j, char := range line {
			if char == 'S' {
				position.row = i
				position.col = j
				position.val = 'S'
				return position
			}
		}
	}

	return position
}

func getPathLength(end *Node) int {
	node := end.parent
	count := 1
	for node != nil && node.parent != nil {
		count++
		node = node.parent
	}
	return count
}

func partOne() {
	elevationMap := common.LinesInFile("adventofcode2022/day12/day12.txt")

	startPosition := getStartPosition(elevationMap)
	startPosition.val = 'a'

	start := Node{
		pos: &startPosition,
	}

	positionsMap := convertMapToPositions(elevationMap)

	end := start.breadthFirstSearchElevation(positionsMap)

	fmt.Println(getPathLength(end))
}

func partTwo() {
	elevationMap := common.LinesInFile("adventofcode2022/day12/day12.txt")

	startPosition := getStartPosition(elevationMap)
	positionsMap := convertMapToPositions(elevationMap)
	positionsMap[startPosition.row][startPosition.col].val = 'a'

	length := len(positionsMap) * len(positionsMap[0])
	for _, positionRow := range positionsMap {
		for _, position := range positionRow {
			positionsMap = convertMapToPositions(elevationMap)
			positionsMap[startPosition.row][startPosition.col].val = 'a'
			if position.val == 'a' {
				start := Node{
					pos: &position,
				}
				end := start.breadthFirstSearchElevation(positionsMap)
				if length > getPathLength(end) && getPathLength(end) > 1 {
					length = getPathLength(end)
				}
				//fmt.Println(start.pos.row, " - ", start.pos.col, " ---> ", getPathLength(end))
			}
		}
	}
	fmt.Println(length)
}

func DayTwelve() {
	start := time.Now()
	partOne()
	log.Printf("Part one took %s", time.Since(start))

	start = time.Now()
	partTwo()
	log.Printf("Part two took %s", time.Since(start))
}
