package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var playbook = map[string]map[string]int{
	"A": {
		"X": 1 + 3,
		"Y": 2 + 0,
		"Z": 3 + 6,
	},
	"B": {
		"X": 1 + 6,
		"Y": 2 + 3,
		"Z": 3 + 0,
	},
	"C": {
		"X": 1 + 0,
		"Y": 2 + 6,
		"Z": 3 + 3,
	},
}

func scoreMatch(match string) int {
	arr := strings.Split(match, " ")
	p1 := arr[0]
	p2 := arr[1]
	return playbook[p1][p2]
}

func main() {
	content, _ := ioutil.ReadFile("input.txt")
	sc := string(content)
	arr := strings.Split(sc, "\n")

	sum := 0
	for _, match := range arr {
		if match != "" {
			sum = sum + scoreMatch(match)
		}
	}

	fmt.Println("Part 1:", sum)
}
