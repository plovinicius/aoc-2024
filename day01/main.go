package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/plovinicius/aoc-2024/utils"
)

func main() {
	data := utils.ReadFile("day01/input.txt")

	part1 := part1(data)
	fmt.Println("Part 1:", part1)

	part2 := part2(data)
	fmt.Println("Part 2:", part2)
}

func getLists(data string) ([]int, []int) {
	lines := strings.Split(data, "\n")
	var leftList []int
	var rightList []int

	for _, line := range lines {
		numbers := strings.Fields(line)
		leftNum, err := strconv.Atoi(numbers[0])
		if err != nil {
			fmt.Println("Error converting number:", err)
			continue
		}
		leftList = append(leftList, leftNum)

		rightNum, err := strconv.Atoi(numbers[1])
		if err != nil {
			fmt.Println("Error converting number:", err)
			continue
		}
		rightList = append(rightList, rightNum)
	}

	return leftList, rightList
}

func part1(data string) int {
	leftList, rightList := getLists(data)

	sort.Ints(sort.IntSlice(leftList))
	sort.Ints(sort.IntSlice(rightList))

	var result int

	for i := 0; i < len(leftList); i++ {
		if leftList[i] > rightList[i] {
			result += leftList[i] - rightList[i]
		} else {
			result += rightList[i] - leftList[i]
		}
	}

	return result
}

func part2(data string) int {
	leftList, rightList := getLists(data)

	var result int
	var nums = make(map[int]int)

	for i := 0; i < len(rightList); i++ {
		_, ok := nums[rightList[i]]

		if ok {
			nums[rightList[i]] += 1
		} else {
			nums[rightList[i]] = 1
		}
	}

	for i := 0; i < len(leftList); i++ {
		count, ok := nums[leftList[i]]

		if ok {
			result += leftList[i] * count
		}
	}

	return result
}
