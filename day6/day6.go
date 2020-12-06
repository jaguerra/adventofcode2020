package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	groups := Groups()
	sumAnyone, sumEveryone := 0, 0
	for _, group := range groups {
		sumAnyone = sumAnyone + group.CountAnswersByAnyone()
		sumEveryone = sumEveryone + group.CountAnswersByEveryone()
	}
	fmt.Println(sumAnyone)
	fmt.Println(sumEveryone)
}

type group struct {
	answerMap   map[string]int
	personCount int
}

func Group() group {
	return *&group{make(map[string]int), 0}
}

func (g *group) ParseAnswerLine(line string) {
	for _, char := range line {
		g.addAnswer(string(char))
	}
	g.personCount = g.personCount + 1
}

func (g *group) addAnswer(answer string) {
	count, exists := g.answerMap[answer]
	if exists == false {
		g.answerMap[answer] = 1
	} else {
		g.answerMap[answer] = count + 1
	}
}

func (g *group) CountAnswersByAnyone() int {
	return len(g.answerMap)
}

func (g *group) CountAnswersByEveryone() int {
	sum := 0
	for _, count := range g.answerMap {
		if count == g.personCount {
			sum++
		}
	}
	return sum
}

func Groups() []group {
	file, _ := os.Open("6.input")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	groups := make([]group, 0)
	group := Group()

	for {
		success := scanner.Scan()
		if success == false {
			break
		}
		line := scanner.Text()
		if len(line) == 0 {
			groups = append(groups, group)
			group = Group()
		} else {
			group.ParseAnswerLine(line)
		}
	}

	groups = append(groups, group)
	return groups
}
