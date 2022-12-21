package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var priorityMap = map[string]int{
	"a": 1, "b": 2, "c": 3,
	"d": 4, "e": 5, "f": 6,
	"g": 7, "h": 8, "i": 9,
	"j": 10, "k": 11, "l": 12,
	"m": 13, "n": 14, "o": 15,
	"p": 16, "q": 17, "r": 18,
	"s": 19, "t": 20, "u": 21,
	"v": 22, "w": 23, "x": 24,
	"y": 25, "z": 26, "A": 27,
	"B": 28, "C": 29, "D": 30,
	"E": 31, "F": 32, "G": 33,
	"H": 34, "I": 35, "J": 36,
	"K": 37, "L": 38, "M": 39,
	"N": 40, "O": 41, "P": 42,
	"Q": 43, "R": 44, "S": 45,
	"T": 46, "U": 47, "V": 48,
	"W": 49, "X": 50, "Y": 51,
	"Z": 52,
}

func checkPriority(item string) int {
	return priorityMap[item]
}

func computePriority(rucksack string) int {
	itemMap1 := make(map[string]int)
	itemMap2 := make(map[string]int)
	compartmentLength := len(rucksack)
	compartment1 := rucksack[0 : compartmentLength/2]
	compartment2 := rucksack[compartmentLength/2 : compartmentLength]

	// loop through each compartment for accounting
	for i := 0; i < compartmentLength/2; i++ {
		item1 := string(compartment1[i])
		if _, ok := itemMap1[item1]; ok {
			itemMap1[item1] = itemMap1[item1] + 1
		} else {
			itemMap1[item1] = 1
		}

		item2 := string(compartment2[i])
		if _, ok := itemMap2[item2]; ok {
			itemMap2[item2] = itemMap2[item2] + 1
		} else {
			itemMap2[item2] = 1
		}
	}

	sum := 0
	for item, num := range itemMap1 {
		if num > 0 {
			if val, ok := itemMap2[item]; ok && val > 0 {
				sum = sum + checkPriority(item)
			}
		}
	}
	return sum
}

func fillMap(rucksack string, m map[string]int) map[string]int {
	for _, item := range rucksack {
		i := string(item)
		if _, ok := m[i]; ok {
			m[i] = m[i] + 1
		} else {
			m[i] = 1
		}
	}
	return m
}

func computePriority2(rucksack1, rucksack2, rucksack3 string) int {
	itemMap1 := make(map[string]int)
	itemMap2 := make(map[string]int)
	itemMap3 := make(map[string]int)

	itemMap1 = fillMap(rucksack1, itemMap1)
	itemMap2 = fillMap(rucksack2, itemMap2)
	itemMap3 = fillMap(rucksack3, itemMap3)

	sum := 0
	for item, num := range itemMap1 {
		if num > 0 {
			if val1, ok := itemMap2[item]; ok && val1 > 0 {
				if val2, ok2 := itemMap3[item]; ok2 && val2 > 0 {
					sum = sum + checkPriority(item)
				}
			}
		}
	}

	return sum
}

func main() {
	content, _ := ioutil.ReadFile("input.txt")
	sc := string(content)
	arr := strings.Split(sc, "\n")

	sum := 0
	sum2 := 0
	for i, rucksack := range arr {
		if rucksack != "" {
			sum = sum + computePriority(rucksack)
			if i%3 == 0 {
				sum2 = sum2 + computePriority2(rucksack, arr[i+1], arr[i+2])
			}
		}
	}

	fmt.Println("Part 1:", sum)
	fmt.Println("Part 2:", sum2)
}
