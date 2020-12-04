package main

import "fmt"
import "os"
import "bufio"
import "strconv"

func main() {
	numbers := getExpenseReport()
	answer1 := getAnswer1(numbers)
	fmt.Println(answer1)
	answer2 := getAnswer2(numbers)
	fmt.Println(answer2)
}

func getAnswer1(numbers []int) int {
	answer := 0
	for x := 0; x < len(numbers); x++ {
		for y := 0; y < len(numbers); y++ {
			if x == y {
				continue
			}
			if numbers[x]+numbers[y] == 2020 {
				answer = numbers[x] * numbers[y]
				break
			}
		}
	}
	return answer
}

func getAnswer2(numbers []int) int {
	answer := 0
	for x := 0; x < len(numbers); x++ {
		for y := 0; y < len(numbers); y++ {
			for z := 0; z < len(numbers); z++ {
				if x == y || x == z || y == z {
					continue
				}
				if numbers[x]+numbers[y]+numbers[z] == 2020 {
					answer = numbers[x] * numbers[y] * numbers[z]
					break
				}
			}
		}
	}
	return answer
}

func getExpenseReport() []int {
	file, _ := os.Open("1.input")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	numbers := make([]int, 0)

	for {
		success := scanner.Scan()
		if success == false {
			break
		}
		number, _ := strconv.Atoi(scanner.Text())
		numbers = append(numbers, number)
	}
	return numbers
}
