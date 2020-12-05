package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	boardingPasses := BoardingPasses()
	minSeatID, maxSeatID := 128*8, 0
	occupiedSeats := make([]bool, 128*8)
	for _, boardingPass := range boardingPasses {
		if boardingPass.seatID > maxSeatID {
			maxSeatID = boardingPass.seatID
		}
		if boardingPass.seatID < minSeatID {
			minSeatID = boardingPass.seatID
		}
		occupiedSeats[boardingPass.seatID] = true
	}
	fmt.Println(maxSeatID)

	for i, occupied := range occupiedSeats[minSeatID:maxSeatID] {
		if occupied == false {
			fmt.Printf("%d\n", i+minSeatID)
		}
	}
}

type boardingPass struct {
	pass             string
	row, col, seatID int
}

func BoardingPasses() []boardingPass {
	file, _ := os.Open("5.input")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	boardingPasses := make([]boardingPass, 0)

	for {
		success := scanner.Scan()
		if success == false {
			break
		}
		boardingPasses = append(boardingPasses, BoardingPass(scanner.Text()))
	}

	return boardingPasses
}

func BoardingPass(pass string) boardingPass {
	row := BinarySearch(0, 127, string(pass[0:8]))
	col := BinarySearch(0, 7, string(pass[7:10]))
	boardingPass := &boardingPass{pass, row, col, row*8 + col}
	return *boardingPass
}

func BinarySearch(min, max int, pass string) int {
	m := (min + max) / 2

	switch string(pass[0]) {
	case "F", "L":
		// Lower half
		max = m
	case "B", "R":
		// Upper half
		min = m + 1
	}
	if min == max {
		return min
	} else {
		return BinarySearch(min, max, string(pass[1:]))
	}
}
