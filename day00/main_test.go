package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	expected := 11
	var sample = ``

	result := part1(sample)

	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestPart2(t *testing.T) {
	expected := 31
	var sample = ``

	result := part2(sample)

	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}
