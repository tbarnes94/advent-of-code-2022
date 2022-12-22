package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func parseElfAssignments(e1, e2 string, mode int) int {
	arr1 := strings.Split(e1, "-")
	arr2 := strings.Split(e2, "-")
	start1, _ := strconv.Atoi(arr1[0])
	end1, _ := strconv.Atoi(arr1[1])
	start2, _ := strconv.Atoi(arr2[0])
	end2, _ := strconv.Atoi(arr2[1])

	Elf1TotallyWithinElf2 := start1 >= start2 && start1 <= end2 && end1 >= start2 && end1 <= end2
	Elf2TotallyWithinElf1 := start2 >= start1 && start2 <= end1 && end2 >= start1 && end2 <= end1

	Elf1SlightlyWithinElf2 := start1 >= start2 && start1 <= end2 || end1 >= start2 && end1 <= end2
	Elf2SlightlyWithinElf1 := start2 >= start1 && start2 <= end1 || end2 >= start1 && end2 <= end1

	if mode == 0 {
		if Elf1TotallyWithinElf2 || Elf2TotallyWithinElf1 {
			return 1
		} else {
			return 0
		}
	} else {
		if Elf1SlightlyWithinElf2 || Elf2SlightlyWithinElf1 {
			return 1
		} else {
			return 0
		}
	}
}

func parseAssignment(a string, mode int) int {
	arr := strings.Split(a, ",")
	elf1 := arr[0]
	elf2 := arr[1]
	return parseElfAssignments(elf1, elf2, mode)
}

func main() {
	content, _ := ioutil.ReadFile("input.txt")
	sc := string(content)
	arr := strings.Split(sc, "\n")

	sum := 0
	sum2 := 0
	for _, a := range arr {
		if a != "" {
			sum = sum + parseAssignment(a, 0)
			sum2 = sum2 + parseAssignment(a, 1)
		}
	}

	fmt.Println("Part 1:", sum)
	fmt.Println("Part 2:", sum2)
}
