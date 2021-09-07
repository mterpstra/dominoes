package main

import (
	"fmt"
)

func ChooseLeftOrRight() string {
	side := ""
	for side != "L" && side != "R" {
		fmt.Printf("Enter Side of Board (L,R):")
		fmt.Scanf("%s", &side)
		if side == "l" {
			side = "L"
		}
		if side == "r" {
			side = "R"
		}
	}
	return side
}

func ChooseNumberInRange(min, max int) int {
	value := 0
	valid := false
	for !valid {
		fmt.Printf("Enter Card Index (%d,%d):", min, max)
		fmt.Scanf("%d", &value)
		if value >= min && value <= max {
			valid = true
		}
	}
	return value
}
