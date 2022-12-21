package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func tallyElfCalories(stack string) int {
	arr := strings.Split(stack, "\n")
	sum := 0
	for _, value := range arr {
		intVal, _ := strconv.Atoi(value)
		sum = sum + intVal
	}
	return sum
}

func main() {
	content, _ := ioutil.ReadFile("input.txt")
	sc := string(content)
	arr := strings.Split(sc, "\n\n")

	maxCalories := 0
	for _, stack := range arr {
		calories := tallyElfCalories(stack)
		if maxCalories < calories {
			maxCalories = calories
		}
	}

	fmt.Println("Something to print:", maxCalories)
}
