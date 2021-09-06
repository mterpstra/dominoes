package main

import (
	"fmt"
)

type Board struct {
	Cards []*Card
}

func (b *Board) Left() int {
	if len(b.Cards) == 0 {
		return -1
	}
	return board.Cards[0].SideA
}

func (b *Board) Right() int {
	if len(b.Cards) == 0 {
		return -1
	}
	return board.Cards[len(board.Cards)-1].SideB
}

func (b *Board) Print() {
	print("\033[0;34m")
	defer func() {
		fmt.Printf("\n")
		print("\033[0m") // Reset
	}()

	if len(b.Cards) == 0 {
		fmt.Printf("Empty Board\n")
		return
	}

	for i := 0; i < len(b.Cards); i++ {
		fmt.Printf("┌───┬───┐")
	}
	fmt.Printf("\n")

	for _, v := range b.Cards {
		fmt.Printf("│ %d │ %d │", v.SideA, v.SideB)
	}
	fmt.Printf("\n")

	for i := 0; i < len(b.Cards); i++ {
		fmt.Printf("└───┴───┘")
	}
}
