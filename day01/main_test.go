package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	var sample = `3   4
	4   3
	2   5
	1   3
	3   9
	3   3`

	result := part1(sample)

	if result != 11 {
		t.Errorf("Expected 11, but got %d", result)
	}
}

func TestPart2(t *testing.T) {
	var sample = `3   4
4   3
2   5
1   3
3   9
3   3`

	result := part2(sample)

	if result != 31 {
		t.Errorf("Expected 31, but got %d", result)
	}
}
