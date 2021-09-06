package main

import (
	"fmt"
)

type Hand struct {
	Cards      []*Card
	IsComputer bool
	Name       string
	Color      string
}

func (h *Hand) Print() {

	print(h.Color)
	fmt.Printf("Player %s (%d cards)\n", h.Name, len(h.Cards))

	for i := 0; i < len(h.Cards); i++ {
		fmt.Printf("┌───┐ ")
	}
	fmt.Printf("\n")

	for _, v := range h.Cards {
		fmt.Printf("│ %d │ ", v.SideA)
	}
	fmt.Printf("\n")

	for i := 0; i < len(h.Cards); i++ {
		fmt.Printf("├───┤ ")
	}
	fmt.Printf("\n")

	for _, v := range h.Cards {
		fmt.Printf("│ %d │ ", v.SideB)
	}
	fmt.Printf("\n")

	for i := 0; i < len(h.Cards); i++ {
		fmt.Printf("└───┘ ")
	}
	fmt.Printf("\n")

	// fmt.Printf("  (len=%d)\n", len(h.Cards))
	print("\033[0m") // Reset
}
