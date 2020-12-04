package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	treeMap := TreeMap()
	// First half
	fmt.Println(treeMap.CalculateTreesInSlope(3, 1))

	// Second half
	results := make([]int, 0)
	results = append(results, treeMap.CalculateTreesInSlope(1, 1))
	results = append(results, treeMap.CalculateTreesInSlope(3, 1))
	results = append(results, treeMap.CalculateTreesInSlope(5, 1))
	results = append(results, treeMap.CalculateTreesInSlope(7, 1))
	results = append(results, treeMap.CalculateTreesInSlope(1, 2))
	result := 1
	for _, res := range results {
		result = result * res
	}
	fmt.Println(result)
}

type treeMap struct {
	trees  [][]bool
	height int
}

func (t treeMap) CalculateTreesInSlope(deltaX, deltaY int) int {
	x, y, treeCount := 0, 0, 0
	for {
		if t.hasTree(x, y) {
			treeCount += 1
		}
		x += deltaX
		y += deltaY
		if y >= t.height {
			break
		}
	}
	return treeCount
}

func (t treeMap) hasTree(x, y int) bool {
	width := len(t.trees[0])
	return t.trees[y][x%width]
}

func TreeMap() treeMap {
	file, _ := os.Open("3.input")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	treeMap := &treeMap{}

	for {
		success := scanner.Scan()
		if success == false {
			break
		}
		line := scanner.Text()

		mapRow := make([]bool, 0)
		for _, char := range line {
			if char == '#' {
				mapRow = append(mapRow, true)
			} else {
				mapRow = append(mapRow, false)
			}
		}
		treeMap.trees = append(treeMap.trees, mapRow)
	}
	treeMap.height = len(treeMap.trees)
	return *treeMap
}
