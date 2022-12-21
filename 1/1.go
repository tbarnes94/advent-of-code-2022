package main

import (
	"fmt"
	"io/ioutil"
	"sort"
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

	var calorieArr []int
	for _, stack := range arr {
		calories := tallyElfCalories(stack)
		calorieArr = append(calorieArr, calories)
	}
	sort.Ints(calorieArr)
	l := len(calorieArr)
	out := calorieArr[l-1] + calorieArr[l-2] + calorieArr[l-3]
	fmt.Println("Part 1:", calorieArr[l-1])
	fmt.Println("Part 2:", out)
}
