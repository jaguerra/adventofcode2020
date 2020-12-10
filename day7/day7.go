package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	luggage := Luggage()
	fmt.Printf("%d\n", luggage.BagHolders("shiny gold"))
	fmt.Printf("%d\n", luggage.Part2("shiny gold"))
}

type luggage struct {
	bags map[string]bag
}

type bag struct {
	contains map[string]int
}

func (l *luggage) BagHolders(bagName string) int {
	return l.traverseBagHolders(bagName, 0, make(map[string]bool))
}

func (l *luggage) traverseBagHolders(containedBagName string, count int, visited map[string]bool) int {
	visited[containedBagName] = true
	for name, bag := range l.bags {
		if _, alreadyVisited := visited[name]; alreadyVisited == false {
			_, exists := bag.contains[containedBagName]
			if exists == true {
				count = l.traverseBagHolders(name, count+1, visited)
			}
		}
	}
	return count
}

func (l *luggage) Part2(bagName string) int {
	bag := l.bags[bagName]
	count := 0
	for name, capacity := range bag.contains {
		if capacity > 0 {
			count += capacity + capacity*l.Part2(name)
		}
	}
	return count
}

func Luggage() luggage {
	file, _ := os.Open("7.input")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	luggage := &luggage{make(map[string]bag)}

	for {
		success := scanner.Scan()
		if success == false {
			break
		}
		line := scanner.Text()
		luggage.ParseLine(line)
	}

	return *luggage
}

func (l *luggage) ParseLine(line string) {
	tokens := strings.Split(line, " bags contain")
	bagName := tokens[0]
	bag := &bag{make(map[string]int)}

	if tokens[1] != "no other bags" {
		t := regexp.MustCompile(` bag[s]{0,1}[.,]{0,1}`)
		contains := t.Split(tokens[1], -1)
		t = regexp.MustCompile(`([0-9]+) (.*)`)
		for _, s := range contains {
			matches := t.FindStringSubmatch(s)
			if len(matches) > 1 {
				count, _ := strconv.Atoi(matches[1])
				name := matches[2]
				bag.contains[name] = count
			}
		}
	}
	l.bags[bagName] = *bag
}
