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

	print("\033[0m") // Reset
}

func (h *Hand) CanPlay(b *Board) bool {
	for _, c := range h.Cards {
		if c.CanPlay(b) {
			return true
		}
	}
	return false
}

func (h *Hand) Play(b *Board, d *DrawPile) (error, *Card) {
	if h.IsComputer {
		return h.PlayAuto(b, d)
	} else {
		return h.PlayManual(b, d)
	}
}

func (h *Hand) PlayAuto(b *Board, d *DrawPile) (error, *Card) {
	for !h.CanPlay(b) {
		println("can't play, drawing...")
		if c := d.Pick(); c != nil {
			h.Cards = append(h.Cards, c)
			h.Print()
		} else {
			// @todo: Should can't play be an error?
			return nil, nil
		}
	}

	index, side, _ := h.determineBestPlay(b)
	return b.Play(h, index, side)
}

func (h *Hand) PlayManual(b *Board, d *DrawPile) (error, *Card) {
	for !h.CanPlay(b) {
		var u string
		println("You can't play, draw from pile (d)")
		fmt.Scanf("%s", &u)

		if c := d.Pick(); c != nil {
			h.Cards = append(h.Cards, c)
			h.Print()
		} else {
			return nil, nil
		}
	}

	index := 0
	for true {
		index = ChooseNumberInRange(0, len(h.Cards)-1)
		if h.Cards[index].CanPlay(b) {
			break
		}
		println("That card cannot be played")
	}

	return b.Play(h, index, AskMe)
}
