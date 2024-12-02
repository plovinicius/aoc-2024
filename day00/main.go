package main

import (
	"fmt"

	"github.com/plovinicius/aoc-2024/utils"
)

func main() {
	data := utils.ReadFile("day00/input.txt")

	part1 := part1(data)
	fmt.Println("Part 1:", part1)

	part2 := part2(data)
	fmt.Println("Part 2:", part2)
}

func part1(data string) int {
	return 1
}

func part2(data string) int {
	return 2
}
